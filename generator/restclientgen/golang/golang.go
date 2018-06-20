package golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	codegen "github.com/lestrrat-go/openapi/internal/codegen/golang"
	restclient "github.com/lestrrat-go/openapi/internal/codegen/restclient/golang"
	"github.com/lestrrat-go/openapi/internal/stringutil"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

func (f *Field) ContainerName() string {
	switch f.in {
	case openapi.InBody, openapi.InForm:
		return "body"
	case openapi.InQuery:
		return "query"
	case openapi.InHeader:
		return "header"
	case openapi.InPath:
		return "path"
	default:
		// No error case, as it should've been handled in Validate()
		return "(no container)"
	}
}

func setLocation(v interface{}, in openapi.Location) {
	switch v := v.(type) {
	case *Struct:
		for _, field := range v.fields {
			field.in = in
		}
	}
}

func extractFields(call *Call, v interface{}) {
	switch typ := v.(type) {
	case *Struct:
		for _, field := range typ.fields {
			if field.required {
				call.requireds = append(call.requireds, field)
			} else {
				call.optionals = append(call.optionals, field)
			}
		}
	}
}

func Generate(spec openapi.Swagger, options ...Option) error {
	var dir string
	var packageName string
	var defaultServiceName string
	exportNew := true
	for _, option := range options {
		switch option.Name() {
		case optkeyDefaultServiceName:
			defaultServiceName = option.Value().(string)
		case optkeyDirectory:
			dir = option.Value().(string)
		case optkeyPackageName:
			packageName = option.Value().(string)
		case optkeyExportNew:
			exportNew = option.Value().(bool)
		}
	}

	if dir == "" {
		dir = "restclient"
	}

	if packageName == "" {
		// Use the last component in the path
		i := strings.LastIndexByte(dir, os.PathSeparator)
		if i < 0 {
			packageName = dir
		} else {
			packageName = dir[i+1:]
		}
	}

	if defaultServiceName == "" {
		defaultServiceName = packageName
	}

	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.Wrapf(err, `failed to create directory %s`, dir)
		}
	}

	ctx := Context{
		compiling:          make(map[string]struct{}),
		dir:                dir,
		packageName:        packageName,
		defaultServiceName: defaultServiceName,
		exportNew:          exportNew,
		resolver:           openapi.NewResolver(spec),
		root:               spec,
		client:             &Client{services: make(map[string]*Service), types: make(map[string]Type)},
		types:              make(map[string]typeDefinition),
	}

	ctx.security = make(map[string]openapi.SecurityScheme)
	for iter := spec.SecurityDefinitions(); iter.Next(); {
		name, scheme := iter.Item()
		ctx.security[name] = scheme
	}

	consumes, err := canonicalConsumesList(spec.Consumes())
	if err != nil {
		return errors.Wrap(err, `failed to parse global "consumes" list`)
	}
	ctx.consumes = consumes

	if err := compileClient(&ctx); err != nil {
		return errors.Wrap(err, `failed to compile restclient`)
	}

	// declare types
	if err := writeTypesFile(&ctx); err != nil {
		return errors.Wrap(err, `failed to write options file`)
	}

	// define options
	if err := writeOptionsFile(&ctx); err != nil {
		return errors.Wrap(err, `failed to write options file`)
	}

	if err := writeClientFile(&ctx); err != nil {
		return errors.Wrap(err, `failed to generate client code`)
	}

	if err := writeServiceFiles(&ctx); err != nil {
		return errors.Wrap(err, `failed to generate service code`)
	}

	return nil
}

func writeTypesFile(ctx *Context) error {
	fn := filepath.Join(ctx.dir, "types_gen.go")
	log.Printf("Generating %s", fn)

	var buf bytes.Buffer
	var dst io.Writer = &buf
	codegen.WritePreamble(dst, ctx.packageName)

	var typDefs []typeDefinition
	for _, typ := range ctx.types {
		typDefs = append(typDefs, typ)
	}
	sort.Slice(typDefs, func(i, j int) bool {
		return typDefs[i].Type.Name() < typDefs[j].Type.Name()
	})

	for _, typDef := range typDefs {
		typ := typDef.Type
		log.Printf("   * Generating definition for %s", typ.Name())
		switch t := typ.(type) {
		case *Array:
			fmt.Fprintf(dst, "\n\ntype %s []%s", t.name, t.elem)
		case *Struct:
			fmt.Fprintf(dst, "\n\n// %s represents the data structure defined in %s", typ.Name(), typDef.Context)
			t.WriteCode(dst)
		}
	}

	if err := codegen.WriteFormattedToFile(fn, buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, bytes.NewReader(buf.Bytes()))
		return errors.Wrapf(err, `failed to write to %s`, fn)
	}
	return nil
}

func writeOptionsFile(ctx *Context) error {
	fn := filepath.Join(ctx.dir, "options_gen.go")
	log.Printf("Generating %s", fn)

	var buf bytes.Buffer
	var dst io.Writer = &buf
	codegen.WritePreamble(dst, ctx.packageName)
	codegen.WriteImports(dst, "io")

	fmt.Fprintf(dst, "\n\ntype Option interface {")
	fmt.Fprintf(dst, "\nName() string")
	fmt.Fprintf(dst, "\nValue() interface{}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n\ntype option struct {")
	fmt.Fprintf(dst, "\nname  string")
	fmt.Fprintf(dst, "\nvalue interface{}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n\nfunc newOption(name string, value interface{}) Option {")
	fmt.Fprintf(dst, "\nreturn &option{")
	fmt.Fprintf(dst, "\nname:  name,")
	fmt.Fprintf(dst, "\nvalue: value,")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n")
	fmt.Fprintf(dst, "\n\nfunc (o *option) Name() string {")
	fmt.Fprintf(dst, "\nreturn o.name")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n\nfunc (o *option) Value() interface{} {")
	fmt.Fprintf(dst, "\nreturn o.value")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\ntype ClientOption = Option")
	fmt.Fprintf(dst, "\n\nconst (")
	fmt.Fprintf(dst, "\noptkeyServer = `server`")
	fmt.Fprintf(dst, "\n)")

	fmt.Fprintf(dst, "\n\ntype CallOption = Option")

	fmt.Fprintf(dst, "\n\nconst (")
	fmt.Fprintf(dst, "\noptkeyAccessToken = `accessToken`")
	fmt.Fprintf(dst, "\noptkeyDebugDump = `debugDump`")
	fmt.Fprintf(dst, "\noptkeyRequestContentType = `requestContentType`")
	fmt.Fprintf(dst, "\n)")

	fmt.Fprintf(dst, "\n\n// WithServer specifies the server that the client will access.")
	fmt.Fprintf(dst, "\nfunc WithServer(s string) ClientOption {")
	fmt.Fprintf(dst, "\nreturn newOption(optkeyServer, s)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// WithContentType is used to specify the content type  that we will")
	fmt.Fprintf(dst, "\n// be using to send the payload. This is useful when you have")
	fmt.Fprintf(dst, "\nfunc WithContentType(s string) CallOption {")
	fmt.Fprintf(dst, "\nreturn newOption(optkeyRequestContentType, s)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// WithAccessToken is used to specify access token used to perform")
	fmt.Fprintf(dst, "\n// authorization on the requested endpoint")
	fmt.Fprintf(dst, "\nfunc WithAccessToken(s string) CallOption {")
	fmt.Fprintf(dst, "\nreturn newOption(optkeyAccessToken, s)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// WithDebugDump is used to dump request and response to")
	fmt.Fprintf(dst, "\n// the sepcified io.Writer")
	fmt.Fprintf(dst, "\nfunc WithDebugDump(dst io.Writer) CallOption {")
	fmt.Fprintf(dst, "\nreturn newOption(optkeyDebugDump, dst)")
	fmt.Fprintf(dst, "\n}")

	if err := codegen.WriteFormattedToFile(fn, buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, bytes.NewReader(buf.Bytes()))
		return errors.Wrapf(err, `failed to write to %s`, fn)
	}
	return nil
}

func writeClientFile(ctx *Context) error {
	fn := filepath.Join(ctx.dir, "client_gen.go")
	log.Printf("Generating %s", fn)

	var buf bytes.Buffer
	if err := formatClient(ctx, &buf, ctx.client); err != nil {
		return errors.Wrap(err, `failed to format client code`)
	}

	if err := codegen.WriteFormattedToFile(fn, buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, bytes.NewReader(buf.Bytes()))
		return errors.Wrapf(err, `failed to write to %s`, fn)
	}
	return nil
}

func writeServiceFiles(ctx *Context) error {
	var serviceNames []string
	for name := range ctx.client.services {
		serviceNames = append(serviceNames, name)
	}
	sort.Strings(serviceNames)

	for _, name := range serviceNames {
		// Remove the "service" from the filename
		fn := strings.TrimSuffix(name, "Service")
		fn = filepath.Join(ctx.dir, stringutil.Snake(fn)+"_service_gen.go")
		log.Printf("Generating %s", fn)

		var buf bytes.Buffer
		if err := formatService(ctx, &buf, ctx.client.services[name]); err != nil {
			return errors.Wrap(err, `failed to format service code`)
		}

		if err := codegen.WriteFormattedToFile(fn, buf.Bytes()); err != nil {
			codegen.DumpCode(os.Stdout, bytes.NewReader(buf.Bytes()))
			return errors.Wrapf(err, `failed to write to %s`, fn)
		}
	}
	return nil
}

func callMethodName(oper openapi.Operation) string {
	rawMethodName, ok := oper.Extension(`x-call-method-name`)
	if !ok {
		return ""
	}

	if s, ok := rawMethodName.(string); ok {
		return s
	}
	return ""
}

func compileClient(ctx *Context) error {
	for piter := ctx.root.Paths().Paths(); piter.Next(); {
		_, pi := piter.Item()
		for operiter := pi.Operations(); operiter.Next(); {
			if err := compileCall(ctx, operiter.Item()); err != nil {
				return errors.Wrapf(err, `failed to compile call object for path %s`, pi.Path())
			}
		}
	}

	return nil
}

func registerType(ctx *Context, path string, t Type, where string) {
	if t.Name() == "" {
		panic("anonymous type")
	}

	if _, ok := ctx.client.types[t.Name()]; ok {
		return
	}

	log.Printf(" * Registering type %s (%s)", t.Name(), path)
	ctx.types[path] = typeDefinition{
		Path:    path,
		Type:    t,
		Context: where,
	}
	ctx.client.types[t.Name()] = t
}

func compileNumeric(s string) Type {
	switch s {
	case "double":
		return Builtin("double")
	case "int64":
		return Builtin("int64")
	default:
		return Builtin("float32")
	}
}

func compileStruct(ctx *Context, schema openapi.Schema) (Type, error) {
	var obj Struct

	for piter := schema.Properties(); piter.Next(); {
		name, prop := piter.Item()

		fieldMsg, err := compileSchema(ctx, prop)
		if err != nil {
			return nil, errors.Wrap(err, `failed to compile schema for object property`)
		}

		obj.fields = append(obj.fields, &Field{
			name:     name,
			goName:   codegen.ExportedName(name),
			tag:      fmt.Sprintf(`json:"%s"`, name),
			typ:      fieldMsg.Name(),
			required: schema.IsRequiredProperty(name),
		})
	}
	return &obj, nil
}

// To solve json references properly, we need to check if the refered
// object matches what we are expecting. to do this, we need to
// marshal/unmarshal and see if it's successful
func encodeDecodeJSON(src interface{}, decodeFunc func([]byte) error) error {
	var encoded bytes.Buffer
	if err := json.NewEncoder(&encoded).Encode(src); err != nil {
		return errors.Wrap(err, `failed to encode temporary structure to JSON`)
	}

	if err := decodeFunc(encoded.Bytes()); err != nil {
		return errors.Wrap(err, `failed to decode temporary structure from JSON`)
	}
	return nil
}

func resolveReference(ctx *Context, ref string, decodeFunc func([]byte) error) error {
	if _, ok := ctx.compiling[ref]; ok {
		return errors.Errorf(`circular dep on %s`, ref)
	}

	ctx.compiling[ref] = struct{}{}
	defer func() {
		delete(ctx.compiling, ref)
	}()

	var thing interface{}
	if cached, ok := ctx.types[ref]; ok {
		thing = cached
	} else {

		// this better be resolvable via Definitions
		resolved, err := ctx.resolver.Resolve(ref)
		if err != nil {
			return errors.Wrapf(err, `failed to resolve %s`, ref)
		}
		thing = resolved
	}

	// The only way to truly make sure that this resolved thingy
	// is a "Schema", is by encoding it to JSON, and decoding
	// it back

	if err := encodeDecodeJSON(thing, decodeFunc); err != nil {
		return errors.Wrapf(err, `failed to extract schema out of %s`, ref)
	}
	return nil
}

func compileItems(ctx *Context, items openapi.Items) (t Type, err error) {
	return compileSchemaLike(ctx, items)
}

func compileSchema(ctx *Context, schema openapi.Schema) (t Type, err error) {
	if schema == nil {
		return nil, errors.New(`nil schema`)
	}

	if ref := schema.Reference(); ref != "" {
		var news openapi.Schema
		fun := func(buf []byte) error {
			return openapi.SchemaFromJSON(buf, &news)
		}

		if err := resolveReference(ctx, ref, fun); err != nil {
			return nil, errors.Wrapf(err, `failed to resolve reference %s`, ref)
		}
		schema = news
		defer func() {
			if strings.HasPrefix(ref, "#/definitions/") {
				n := codegen.ExportedName(strings.TrimPrefix(ref, "#/definitions/"))
				t.SetName(n)
			}
			registerType(ctx, ref, t, ref)
		}()

	}

	return compileSchemaLike(ctx, schema)
}

type openapiItemser interface {
	Items() openapi.Items
}

func compileSchemaLike(ctx *Context, schema openapiTypeFormater) (Type, error) {
	switch schema.Type() {
	case openapi.String, openapi.Integer, openapi.Boolean:
		return compileBuiltin(ctx, schema)
	case openapi.Array:
		return compileArray(ctx, schema)
	default:
		// In order for this to work, schema must be a full-blown openapi.Schema,
		// not a openapi.Items
		fullSchema, ok := schema.(openapi.Schema)
		if !ok {
			return nil, errors.Errorf(`target must be an openapi.Schema (was %T)`, fullSchema)
		}
		object, err := compileStruct(ctx, fullSchema)
		if err != nil {
			return nil, errors.Wrap(err, `failed to compile object schema`)
		}
		return object, nil
	}

	return nil, errors.New(`unreachable`)
}

type openapiTyper interface {
	Type() openapi.PrimitiveType
}

type openapiFormater interface {
	Format() string
}

type openapiTypeFormater interface {
	openapiTyper
	openapiFormater
}

func compileBuiltin(ctx *Context, schema openapiTypeFormater) (Type, error) {
	switch schema.Type() {
	case openapi.Boolean:
		return Builtin("bool"), nil
	case openapi.String:
		return Builtin(schema.Type()), nil
	case openapi.Integer:
		switch schema.Format() {
		case "int64":
			return Builtin("int64"), nil
		default:
			return Builtin("int"), nil
		}
	default:
		return nil, errors.Errorf(`unknown builtin %s`, schema.Type())
	}
}

func compileResponseType(ctx *Context, response openapi.Response) (string, error) {
	if ref := response.Reference(); ref != "" {
		// this better be resolvable via Definitions
		thing, err := ctx.resolver.Resolve(ref)
		if err != nil {
			return "", errors.Wrapf(err, `failed to resolve %s`, ref)
		}

		// The only way to truly make sure that this resolved thingy
		// is a "Parameter", is by encoding it to JSON, and decoding
		// it back
		var encoded bytes.Buffer
		if err := json.NewEncoder(&encoded).Encode(thing); err != nil {
			return "", errors.Wrap(err, `failed to encode temporary structure to JSON`)
		}

		var newr openapi.Response
		if err := openapi.ResponseFromJSON(encoded.Bytes(), &newr); err != nil {
			return "", errors.Wrap(err, `failed to decode temporary structure from JSON`)
		}
		response = newr
	}

	schema := response.Schema()
	if schema == nil {
		// empty schema means we ain't expecting a response
		return "", nil
	}
	// If this is an array type, we create a []T  instead of type T struct { something []X }
	switch schema.Type() {
	case openapi.Array:
		typ, err := compileSchema(ctx, schema.Items())
		if err != nil {
			return "", errors.Wrap(err, `failed to compile array response`)
		}

		return "[]*" + typ.Name(), nil
	case "", openapi.Object:
		typ, err := compileSchema(ctx, schema)
		if err != nil {
			return "", errors.Wrap(err, `failed to compile object response`)
		}
		if typ.Name() == "" {
			typ.SetName(codegen.ExportedName(ctx.currentCall.name + "_Response"))
			registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, ctx.currentCall.name+" response")
		}
		return typ.Name(), nil
	default:
		return "", errors.Errorf(`unimplemented %s`, schema.Type())
	}
}

func compileParameterType(ctx *Context, param openapi.Parameter) (Type, error) {
	if ref := param.Reference(); ref != "" {
		// this better be resolvable via Definitions
		thing, err := ctx.resolver.Resolve(ref)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to resolve %s`, ref)
		}

		// The only way to truly make sure that this resolved thingy
		// is a "Parameter", is by encoding it to JSON, and decoding
		// it back
		var encoded bytes.Buffer
		if err := json.NewEncoder(&encoded).Encode(thing); err != nil {
			return nil, errors.Wrap(err, `failed to encode temporary structure to JSON`)
		}

		var newp openapi.Parameter
		if err := openapi.ParameterFromJSON(encoded.Bytes(), &newp); err != nil {
			return nil, errors.Wrap(err, `failed to decode temporary structure from JSON`)
		}
		param = newp
	}

	if param.In() == openapi.InBody {
		schema := param.Schema() // presence of this element should be guaranteed by calling validate
		// If this is an array type, we create a []T  instead of type T struct { something []X }
		switch schema.Type() {
		case openapi.Array:
			typ, err := compileSchema(ctx, schema.Items())
			if err != nil {
				return nil, errors.Wrap(err, `failed to compile array parameter`)
			}

			return &Array{elem: typ.Name()}, nil
		case openapi.Object:
			typ, err := compileSchema(ctx, schema)
			if err != nil {
				return nil, errors.Wrap(err, `failed to compile object parameter`)
			}

			if typ.Name() == "" {
				typ.SetName(codegen.ExportedName(ctx.currentCall.name + "_" + param.Name()))
			}

			registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, ctx.currentCall.name+" parameter")

			return typ, nil
		case openapi.String:
			return compileBuiltin(ctx, schema)
		default:
			return nil, errors.Errorf(`unhandled parameter type %s`, strconv.Quote(string(schema.Type())))
		}
	}

	switch param.Type() {
	case openapi.Number:
		return compileNumeric(param.Format()), nil
	case openapi.Array:
		return compileArray(ctx, param)
	}

	return Builtin(param.Type()), nil
}

func compileArray(ctx *Context, schema interface{}) (Type, error) {
	var subtyp Type
	var err error
	if s, ok := schema.(openapi.Schema); ok {
		subtyp, err = compileSchema(ctx, s.Items())
	} else if i, ok := schema.(openapi.Parameter); ok {
		subtyp, err = compileItems(ctx, i.Items())
	} else {
		return nil, errors.Wrapf(err, `cannot compile array element %T`, schema)
	}

	if err != nil {
		return nil, errors.Wrap(err, `failed to compile array schema`)
	}
	return &Array{elem: subtyp.Name()}, nil
}

func compileParameterToProperty(parentBuilder *openapi.SchemaBuilder, param openapi.Parameter) error {
	prop, err := param.ConvertToSchema()
	if err != nil {
		return errors.Wrap(err, `failed to convert parameter to property (openapi.Schema)`)
	}
	parentBuilder.Property(param.Name(), prop)
	return nil
}

func compileStructFromBuilder(ctx *Context, builder interface {
	Do(...openapi.Option) (openapi.Schema, error)
}, location openapi.Location) (Type, error) {
	schema, err := builder.Do()
	if err != nil {
		return nil, errors.Wrap(err, `failed to build schema`)
	}

	typ, err := compileStruct(ctx, schema)
	if err != nil {
		return nil, errors.Wrap(err, `failed to compile schema`)
	}
	setLocation(typ, location)

	return typ, nil
}

func compileCall(ctx *Context, oper openapi.Operation) error {
	// x-service dictates the service name. If not present,
	// the default service, which is named after the package
	// is used.

	callName := restclient.CallObjectName(oper)
	methodName := restclient.CallMethodName(oper)
	if methodName == "" {
		methodName = strings.TrimSuffix(callName, "Call")
	}

	call := &Call{
		name:        callName,
		method:      methodName,
		requestPath: oper.PathItem().Path(),
		verb:        oper.Verb(),
	}
	ctx.currentCall = call

	// The OpenAPI spec allows you to specify multiple "consumes"
	// clause, but we only support JSON or appliation/x-www-form-urlencoded
	// by default

	consumesList, err := canonicalConsumesList(oper.Consumes())
	if err != nil {
		return errors.Wrapf(err, `failed to parse consumes list for %s:%s`, oper.PathItem().Path(), oper.Verb())
	}

	if len(consumesList) == 0 {
		// Use the default consumes list if not provided
		consumesList = append(consumesList, ctx.consumes...)
	}
	call.consumes = consumesList

	for respiter := oper.Responses().Responses(); respiter.Next(); {
		_, resp := respiter.Item()
		typ, err := compileResponseType(ctx, resp)
		if err != nil {
			return errors.Wrap(err, `failed to compile response type`)
		}

		call.responses = append(call.responses, &Response{
			code: resp.StatusCode(),
			typ:  typ,
		})
	}

	var queryBuilder *openapi.SchemaBuilder
	var pathBuilder *openapi.SchemaBuilder
	var headerBuilder *openapi.SchemaBuilder
	var formBuilder *openapi.SchemaBuilder
	for piter := oper.Parameters(); piter.Next(); {
		param := piter.Item()

		switch param.In() {
		case openapi.InQuery:
			if queryBuilder == nil {
				queryBuilder = openapi.NewSchema()
			}
			if err := compileParameterToProperty(queryBuilder, param); err != nil {
				return errors.Wrap(err, `failed to compile parameter into openapi.Schema`)
			}
		case openapi.InPath:
			if pathBuilder == nil {
				pathBuilder = openapi.NewSchema()
			}
			if err := compileParameterToProperty(pathBuilder, param); err != nil {
				return errors.Wrap(err, `failed to compile parameter into openapi.Schema`)
			}
		case openapi.InHeader:
			if headerBuilder == nil {
				headerBuilder = openapi.NewSchema()
			}
			if err := compileParameterToProperty(headerBuilder, param); err != nil {
				return errors.Wrap(err, `failed to compile parameter into openapi.Schema`)
			}
		case openapi.InForm:
			if formBuilder == nil {
				formBuilder = openapi.NewSchema()
			}
			if err := compileParameterToProperty(formBuilder, param); err != nil {
				return errors.Wrap(err, `failed to compile parameter into openapi.Schema`)
			}
		case openapi.InBody:
			// sanity check (although this should have already been taken care
			// of in Validate())
			if call.body != nil {
				return errors.New(`multiple body elements found in parameters`)
			}

			// compile this field into type, and enqueue its fields as
			// if they are part of the Call object
			// (the body parameter is no longer visible to the user,
			// but we want the user to populate its fields)
			typ, err := compileParameterType(ctx, param)
			if err != nil {
				// XXX use param.Name, not field.name because we might have
				// transformed it
				return errors.Wrapf(err, `failed to compile parameter %s`, param.Name())
			}

			setLocation(typ, openapi.InBody)
			extractFields(call, typ)
			call.body = typ
		default:
			return errors.Errorf(`invalid location in parmaeter: %s`, param.In())
			/*
				typ, err := compileParameterType(ctx, param)
				if err != nil {
					// XXX use param.Name, not field.name because we might have
					// transformed it
					return errors.Wrapf(err, `failed to compile parameter %s`, param.Name())
				}
				var field Field
				field.name = param.Name()
				field.goName = stringutil.LowerCamel(param.Name())
				field.typ = typ.Name()
				if param.Required() {
					call.requireds = append(call.requireds, &field)
				} else {
					call.optionals = append(call.optionals, &field)
				}
			*/
		}
	}

	if queryBuilder != nil {
		typ, err := compileStructFromBuilder(ctx, queryBuilder, openapi.InQuery)
		if err != nil {
			return errors.Wrap(err, `failed to compile schema for query fields`)
		}
		typ.SetName(codegen.ExportedName(call.name + "_Query"))
		registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, call.name+" query")
		extractFields(call, typ)
		call.query = typ
	}

	if headerBuilder != nil {
		typ, err := compileStructFromBuilder(ctx, headerBuilder, openapi.InHeader)
		if err != nil {
			return errors.Wrap(err, `failed to compile schema for header fields`)
		}
		typ.SetName(codegen.ExportedName(call.name + "_Header"))
		registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, call.name+" header")
		extractFields(call, typ)
		call.header = typ
	}

	if pathBuilder != nil {
		typ, err := compileStructFromBuilder(ctx, pathBuilder, openapi.InPath)
		if err != nil {
			return errors.Wrap(err, `failed to compile schema for path fields`)
		}
		typ.SetName(codegen.ExportedName(call.name + "_Path"))
		registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, call.name+" path")
		extractFields(call, typ)
		call.path = typ
	}

	if formBuilder != nil {
		typ, err := compileStructFromBuilder(ctx, formBuilder, openapi.InPath)
		if err != nil {
			return errors.Wrap(err, `failed to compile schema for form fields`)
		}
		typ.SetName(codegen.ExportedName(call.name + "_Form"))
		registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, call.name+" form")
		extractFields(call, typ)
		call.body = typ

		if len(call.consumes) == 0 {
			call.consumes = append(call.consumes, "application/x-www-form-urlencoded")
		}
	}

	// grab the service.
	svcName := ctx.defaultServiceName

	// x-oagen-service, then x-service
	for _, key := range []string{"x-oagen-service", "x-service"} {
		rawSvcName, ok := oper.Extension(key)
		if ok {
			svcName, ok = rawSvcName.(string)
			if !ok {
				return errors.Errorf(`expected x-service to be a string`)
			}
			break
		}
	}

	// Check if we have security
	for iter := oper.Security(); iter.Next(); {
		requirement := iter.Item()
		for siter := requirement.Scopes(); siter.Next(); {
			name, scopes := siter.Item()
			security, ok := ctx.security[name]
			if !ok {
				return errors.Errorf(`invalid security definition %s (not found)`, name)
			}
			call.securitySettings = append(call.securitySettings, &SecuritySettings{
				definition: security,
				scopes:     scopes,
			})
		}
	}

	svcName = strcase.ToCamel(svcName)
	svc := ctx.client.getServiceFor(svcName)
	svc.addCall(call)
	return nil
}

func formatClient(ctx *Context, dst io.Writer, cl *Client) error {
	codegen.WritePreamble(dst, ctx.packageName)
	codegen.WriteImports(dst, "bytes", "net/http", "github.com/pkg/errors")
	fmt.Fprintf(dst, "\n\n")
	var serviceNames []string
	for name := range cl.services {
		serviceNames = append(serviceNames, name)
	}
	sort.Strings(serviceNames)

	fmt.Fprintf(dst, "\n\ntype Marshaler interface {")
	fmt.Fprintf(dst, "\nMarshal(interface{}) ([]byte, error)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\ntype MarshalFunc func(interface{}) ([]byte, error)")
	fmt.Fprintf(dst, "\nfunc (f MarshalFunc) Marshal(v interface{}) ([]byte, error) {")
	fmt.Fprintf(dst, "\nreturn f(v)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Response is the interface to wrap all possible")
	fmt.Fprintf(dst, "\n// responses. The actual data returned from the server can")
	fmt.Fprintf(dst, "\n// be accessed through the Data() method.")
	fmt.Fprintf(dst, "\ntype Response interface {")
	fmt.Fprintf(dst, "\nStatusCode() int")
	fmt.Fprintf(dst, "\nData() interface{}")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\ntype response struct {")
	fmt.Fprintf(dst, "\ncode int")
	fmt.Fprintf(dst, "\ndata interface{}")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (r *response) StatusCode() int {")
	fmt.Fprintf(dst, "\nreturn r.code")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (r *response) Data() interface{} {")
	fmt.Fprintf(dst, "\nreturn r.data")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc encodeCallPayload(marshalers map[string]Marshaler, mtype string, payload interface{}) (*bytes.Buffer, error) {")
	fmt.Fprintf(dst, "\nmarshaler, ok := marshalers[mtype]")
	fmt.Fprintf(dst, "\nif !ok {")
	fmt.Fprintf(dst, "\nreturn nil, errors.Errorf(`missing marshaler for request content type %%s`, mtype)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n\nencoded, err := marshaler.Marshal(payload)")
	fmt.Fprintf(dst, "\nif err != nil {")
	fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `failed to marshal payload`)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nreturn bytes.NewBuffer(encoded), nil")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\ntype Client struct {")
	for _, name := range serviceNames {
		fmt.Fprintf(dst, "\n%s *%s", codegen.UnexportedName(name), name)
	}
	fmt.Fprintf(dst, "\n}")

	// There are cases where we don't want to export New(): e.g.
	// you want to put more custom logic around the creation of
	// a new client object.
	// In such cases, use the WithExportNew option
	newFuncName := "New"
	if !ctx.exportNew {
		newFuncName = "newClient"
	}

	fmt.Fprintf(dst, "\n\n// %s creates a new client. For example, if your API require additional", newFuncName)
	fmt.Fprintf(dst, "\n// OAuth authentication, JWT authorization, etc, pass an http.Client with")
	fmt.Fprintf(dst, "\n// a custom Transport to handle it")
	fmt.Fprintf(dst, "\nfunc %s(cl *http.Client, options ...ClientOption) *Client {", newFuncName)
	fmt.Fprintf(dst, "\nvar server string")
	fmt.Fprintf(dst, "\nfor _, option := range options {")
	fmt.Fprintf(dst, "\nswitch option.Name() {")
	fmt.Fprintf(dst, "\ncase optkeyServer:")
	fmt.Fprintf(dst, "\nserver = option.Value().(string)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nif cl == nil {")
	fmt.Fprintf(dst, "\ncl = &http.Client{}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nreturn &Client{")
	for _, name := range serviceNames {
		fmt.Fprintf(dst, "\n%s: &%s{httpCl: cl, server: server},", codegen.UnexportedName(name), name)
	}
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	for _, name := range serviceNames {
		fmt.Fprintf(dst, "\nfunc (c *Client) %s() *%s {", name, name)
		fmt.Fprintf(dst, "\nreturn c.%s", codegen.UnexportedName(name))
		fmt.Fprintf(dst, "\n}")
	}

	return nil
}

func formatService(ctx *Context, dst io.Writer, svc *Service) error {
	log.Printf(" * Generating Service %s", svc.name)
	codegen.WritePreamble(dst, ctx.packageName)
	codegen.WriteImports(dst, "context", "encoding/json", "fmt", "io", "mime", "net/http", "net/http/httputil", "net/url", "strings", "strconv", "github.com/pkg/errors", "github.com/lestrrat-go/urlenc")

	fmt.Fprintf(dst, "\n\ntype %s struct {", svc.name)
	fmt.Fprintf(dst, "\nhttpCl *http.Client")
	fmt.Fprintf(dst, "\nserver string")
	fmt.Fprintf(dst, "\n}")

	sort.Slice(svc.calls, func(i, j int) bool {
		return svc.calls[i].name < svc.calls[j].name
	})

	for _, call := range svc.calls {
		if err := formatCall(dst, svc.name, call); err != nil {
			return errors.Wrap(err, `failed to format call`)
		}
	}
	return nil
}

// res, err := restclient.Cats().
// 	GetCatsprotojson().
//		Protojson("json").
//		Cats(message.Cat{...}, ...).
//		Do(ctx)
func formatCall(dst io.Writer, svcName string, call *Call) error {
	log.Printf("   * Generating Call object %s", call.name)
	var allFields []*Field
	allFields = append(append(allFields, call.requireds...), call.optionals...)
	sort.Slice(allFields, func(i, j int) bool {
		return allFields[i].name < allFields[j].name
	})

	fmt.Fprintf(dst, "\n\ntype %s struct {", call.name)
	fmt.Fprintf(dst, "\nhttpCl *http.Client")
	fmt.Fprintf(dst, "\nserver string")
	fmt.Fprintf(dst, "\nmarshalers map[string]Marshaler")
	if call.body != nil {
		fmt.Fprintf(dst, "\nbody %s", call.body.Name())
	}
	if call.query != nil {
		fmt.Fprintf(dst, "\nquery %s", call.query.Name())
	}
	if call.header != nil {
		fmt.Fprintf(dst, "\nheader %s", call.header.Name())
	}
	if call.path != nil {
		fmt.Fprintf(dst, "\npath %s", call.path.Name())
	}

	fmt.Fprintf(dst, "\n}")

	sort.Slice(call.requireds, func(i, j int) bool {
		return call.requireds[i].name < call.requireds[j].name
	})

	log.Printf("      * Generating constructor")
	fmt.Fprintf(dst, "\n\nfunc (svc *%s) %s(", svcName, call.method)
	for i, field := range call.requireds {
		fmt.Fprintf(dst, "%s %s", codegen.UnexportedName(field.goName), field.typ)
		if i < len(call.requireds)-1 {
			fmt.Fprintf(dst, ", ")
		}
	}
	fmt.Fprintf(dst, ") *%s {", call.name)

	fmt.Fprintf(dst, "\nvar call %s", call.name)
	fmt.Fprintf(dst, "\ncall.httpCl = svc.httpCl")
	fmt.Fprintf(dst, "\ncall.server = svc.server")
	fmt.Fprintf(dst, "\ncall.marshalers = map[string]Marshaler{")
	fmt.Fprintf(dst, "\n`application/json`: MarshalFunc(json.Marshal),")
	fmt.Fprintf(dst, "\n`application/x-www-form-urlencoded`: MarshalFunc(urlenc.Marshal),")
	fmt.Fprintf(dst, "\n}")
	for _, field := range call.requireds {
		fmt.Fprintf(dst, "\ncall.%s.%s = %s", field.ContainerName(), field.goName, codegen.UnexportedName(field.goName))
	}
	fmt.Fprintf(dst, "\nreturn &call")
	fmt.Fprintf(dst, "\n}")

	sort.Slice(call.optionals, func(i, j int) bool {
		return call.optionals[i].name < call.optionals[i].name
	})
	for _, optional := range call.optionals {
		log.Printf("      * Generating optional method for %s", codegen.ExportedName(optional.name))
		if strings.HasPrefix(optional.typ, "[]") {
			fmt.Fprintf(dst, "\n\nfunc (call *%s) %s(v ...%s) *%s {", call.name, codegen.ExportedName(optional.name), strings.TrimPrefix(optional.typ, "[]"), call.name)
			fmt.Fprintf(dst, "\ncall.%[1]s.%[2]s = append(call.%[1]s.%[2]s, v...)", optional.ContainerName(), codegen.ExportedName(optional.name))
			fmt.Fprintf(dst, "\nreturn call")
			fmt.Fprintf(dst, "\n}")
		} else {
			fmt.Fprintf(dst, "\n\nfunc (call *%s) %s(v %s) *%s {", call.name, stringutil.Camel(optional.name), optional.typ, call.name)
			fmt.Fprintf(dst, "\ncall.%s.%s = v", optional.ContainerName(), codegen.ExportedName(optional.name))
			fmt.Fprintf(dst, "\nreturn call")
			fmt.Fprintf(dst, "\n}")
		}
	}

	/*
		fmt.Fprintf(dst, "\n\nfunc (call %s) AsMap() map[string]interface{} {", call.name)
		fmt.Fprintf(dst, "\nm := make(map[string]interface{})")
		for _, param := range append(call.optionals, call.requireds...) {
			fmt.Fprintf(dst, "\nm[%#v] = call.%s", param.name, param.goName)
		}
		fmt.Fprintf(dst, "\nreturn m")
		fmt.Fprintf(dst, "\n}")
	*/

	fmt.Fprintf(dst, "\n\nfunc (call *%[1]s) Marshaler(mime string, m Marshaler) *%[1]s {", call.name)
	fmt.Fprintf(dst, "\ncall.marshalers[mime] = m")
	fmt.Fprintf(dst, "\nreturn call")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (call *%s) Do(ctx context.Context, options ...CallOption) (Response, error) {", call.name)

	var hasOAuth2 bool
	if len(call.securitySettings) > 0 {
		for _, settings := range call.securitySettings {
			switch settings.definition.Type() {
			case "oauth2":
				// Require Authorization header
				hasOAuth2 = true
			}
		}
	}

	fmt.Fprintf(dst, "\n\nvar debugOut io.Writer")
	if hasOAuth2 {
		fmt.Fprintf(dst, "\nvar accessToken string")
	}

	if call.body != nil {
		fmt.Fprintf(dst, "\ncontentType := %#v", call.DefaultConsumes())
	}

		fmt.Fprintf(dst, "\nfor _, option := range options {")
		fmt.Fprintf(dst, "\nswitch option.Name() {")
		fmt.Fprintf(dst, "\ncase optkeyDebugDump:")
		fmt.Fprintf(dst, "\ndebugOut = option.Value().(io.Writer)")
		if call.body != nil {
			fmt.Fprintf(dst, "\ncase optkeyRequestContentType:")
			fmt.Fprintf(dst, "\ncontentType = option.Value().(string)")
		}
		if hasOAuth2 {
			fmt.Fprintf(dst, "\ncase optkeyAccessToken:")
			fmt.Fprintf(dst, "\naccessToken = option.Value().(string)")
		}
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\npath := %s", strconv.Quote(call.requestPath))
	if call.path != nil {
		for _, field := range call.path.(*Struct).fields {
			fmt.Fprintf(dst, "\npath = strings.Replace(path, `{%s}`, ", field.name)
			switch field.typ {
			case "int64":
				fmt.Fprintf(dst, "strconv.FormatInt(call.path.%s, 10)", codegen.ExportedName(field.name))
			default:
				fmt.Fprintf(dst, "call.path.%s", codegen.ExportedName(field.name))
			}
			fmt.Fprintf(dst, ", -1)")
		}
	}

	if call.query != nil {
		fmt.Fprintf(dst, "\nv := url.Values{}")
		for _, query := range call.query.(*Struct).fields {
			// XXX This needs to be more robust
			if query.typ == "[]string" {
				fmt.Fprintf(dst, "\nfor _, param := range call.query.%s {", query.goName)
				fmt.Fprintf(dst, "\nv.Add(%s, param)", strconv.Quote(query.name))
				fmt.Fprintf(dst, "\n}")
			} else {
				fmt.Fprintf(dst, "\nv.Add(%s, ", strconv.Quote(query.name))
				switch query.typ {
				case "int64":
					fmt.Fprintf(dst, "strconv.FormatInt(call.query.%s, 10)", query.name)
				default:
					fmt.Fprintf(dst, "call.query.%s", query.goName)
				}
				fmt.Fprintf(dst, ")")
			}
		}

		fmt.Fprintf(dst, "\npath = call.server + path + `?` + v.Encode()")
	}

	var body = "nil"
	if call.body != nil {
		body = "body"

		fmt.Fprintf(dst, "\nmtype, _, err := mime.ParseMediaType(contentType)")
		fmt.Fprintf(dst, "\nif err != nil {")
		fmt.Fprintf(dst, "\nreturn nil, errors.Wrapf(err, `failed to parse request content type %%s`, contentType)")
		fmt.Fprintf(dst, "\n}")

		if call.body != nil {
			fmt.Fprintf(dst, "\n\nbody, err := encodeCallPayload(call.marshalers, mtype, call.body)")
		} else {
			return errors.New(`can't proceed when call.body == nil`)
		}
		fmt.Fprintf(dst, "\nif err != nil {")
		fmt.Fprintf(dst, "\nreturn nil, errors.Wrapf(err, `failed to marshal request payload as %%s`, mtype)")
		fmt.Fprintf(dst, "\n}")
	}

	fmt.Fprintf(dst, "\n\nreq, err := http.NewRequest(%s, path, %s)", strconv.Quote(call.verb), body)
	fmt.Fprintf(dst, "\nif err != nil {")
	fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `failed to create request`)")
	fmt.Fprintf(dst, "\n}")

	if call.body != nil {
		fmt.Fprintf(dst, "\nreq.Header.Set(`Content-Type`, contentType)")
		fmt.Fprintf(dst, "\nreq.Header.Set(`Content-Length`, strconv.Itoa(body.Len()))")
	}

	if hasOAuth2 {
		fmt.Fprintf(dst, "\nif len(accessToken) > 0 {")
		fmt.Fprintf(dst, "\nreq.Header.Set(`Authorization`, `Bearer ` + accessToken)")
		fmt.Fprintf(dst, "\n}")
	}

	fmt.Fprintf(dst, "\n\nif debugOut != nil {")
	fmt.Fprintf(dst, "\ndump, _ := httputil.DumpRequest(req, true)")
	fmt.Fprintf(dst, "\nfmt.Fprintf(debugOut, \"=== REQUEST ===\\n%%s\\n===============\\n\", dump)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nres, err := call.httpCl.Do(req)")
	fmt.Fprintf(dst, "\nif err != nil {")
	fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `failed to make HTTP request`)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nif debugOut != nil {")
	fmt.Fprintf(dst, "\ndump, _ := httputil.DumpResponse(res, true)")
	fmt.Fprintf(dst, "\nfmt.Fprintf(debugOut, \"=== RESPONSE ===\\n%%s\\n===============\\n\", dump)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nvar apires response")
	fmt.Fprintf(dst, "\napires.code = res.StatusCode")

	fmt.Fprintf(dst, "\n\nswitch res.StatusCode {")
	for _, response := range call.responses {
		fmt.Fprintf(dst, "\ncase %s:", response.code)
		// if typ == "" ignore response. we ain't expecting none.
		if response.typ == "" {
			fmt.Fprintf(dst, "\n// no response body expected")
		} else {
			fmt.Fprintf(dst, "\nvar resdata %s", response.typ)
			fmt.Fprintf(dst, "\nswitch ct := strings.ToLower(res.Header.Get(`Content-Type`)); {")
			fmt.Fprintf(dst, "\ncase strings.HasPrefix(ct, `application/json`):")
			fmt.Fprintf(dst, "\nif err := json.NewDecoder(res.Body).Decode(&resdata); err != nil {")
			fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `failed to decode JSON content`)")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\napires.data = resdata")
		}
	}
	fmt.Fprintf(dst, "\ndefault:")
	fmt.Fprintf(dst, "\nreturn nil, errors.Errorf(`invalid response code %%d`, res.StatusCode)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nreturn &apires, nil")
	fmt.Fprintf(dst, "\n}")
	return nil
}

func compilePrimitiveType(typ openapi.PrimitiveType, format string) (Type, error) {
	switch typ {
	case openapi.Number:
		switch format {
		case "", "float":
			return Builtin("float32"), nil
		case "double":
			return Builtin("float64"), nil
		default:
			return nil, errors.Errorf(`unknown "number" format: %s`, format)
		}
	case openapi.Integer:
		switch format {
		case "":
			return Builtin("int"), nil
		case "int32":
			return Builtin("int32"), nil
		case "int64":
			return Builtin("int64"), nil
		default:
			return nil, errors.Errorf(`unknown "integer" format: %s`, format)
		}
	case openapi.String:
		switch format {
		case "":
			return Builtin("string"), nil
		case "byte":
			return Builtin("[]byte"), nil // Array{}?
		case "binary":
			return Builtin("[]byte"), nil // differentiate between byte and binary?
		case "date", "date-time":
			return Builtin("time.Time"), nil // I think this will barf
		default:
			return nil, errors.Errorf(`unknown "string" format: %s`, format)
		}
	case openapi.Boolean:
		return Builtin("bool"), nil
	default:
		return nil, errors.Errorf(`primitive type %s should not have gone through CompilePrimitive`, typ)
	}
}
