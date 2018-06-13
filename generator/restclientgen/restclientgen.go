package restclientgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/iancoleman/strcase"
	"github.com/lestrrat-go/openapi/internal/codegen"
	"github.com/lestrrat-go/openapi/internal/stringutil"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

type typeDefinition struct {
	Path    string
	Context string
	Type    Type
}

type genCtx struct {
	client             *Client
	compiling          map[string]struct{}
	currentCall        *Call
	dir                string
	packageName        string
	defaultServiceName string
	resolver           openapi.Resolver
	root               openapi.Swagger
	types              map[string]typeDefinition
	exportNew          bool
	consumes           []string
	produces           []string
}

type Client struct {
	services map[string]*Service
	types    map[string]Type
}

type Type interface {
	Name() string
	SetName(string)
}

type Builtin string

func (b Builtin) Name() string {
	return string(b)
}

func (b Builtin) SetName(s string) {
	panic("oops?")
}

type Array struct {
	name string
	elem string
}

type Struct struct {
	name   string
	fields []*Field
}

func (v *Struct) WriteCode(dst io.Writer) error {
	fmt.Fprintf(dst, "\ntype %s struct {", v.name)
	for _, field := range v.fields {
		fmt.Fprintf(dst, "\n%s %s `%s`", field.goName, field.typ, field.tag)
	}
	fmt.Fprintf(dst, "\n}")
	return nil
}

func (a *Array) SetName(s string) {
	a.name = s
}

func isBuiltinType(s string) bool {
	switch s {
	case "string",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "byte", "rune", "bool":
		return true
	default:
		return false
	}
}

func (a *Array) Name() string {
	if a.name == "" {
		if isBuiltinType(a.elem) {
			return "[]" + a.elem
		} else {
			return "[]*" + a.elem
		}
	}
	return a.name
}

func (o *Struct) SetName(s string) {
	o.name = s
}

func (o *Struct) Name() string {
	return o.name
}

func (c *Client) getServiceFor(name string) *Service {
	name = name + "Service"
	svc, ok := c.services[name]
	if !ok {
		svc = &Service{name: name}
		c.services[name] = svc
	}
	return svc
}

type Service struct {
	name  string
	calls []*Call
}

func (s *Service) addCall(call *Call) {
	s.calls = append(s.calls, call)
}

type Call struct {
	name        string
	path        string
	verb        string
	consumes    []string
	produces    []string
	requireds   []*Field
	optionals   []*Field
	pathparams  []*Field
	queryparams []*Field
	formType    Type
	bodyType    Type
	responses   []*Response
}

func (call *Call) DefaultConsumes() string {
	// default to "application/x-www-form-urlencoded"
	if len(call.consumes) == 0 {
		return "application/x-www-form-urlencoded"
	}
	return call.consumes[0]
}

func (call *Call) Consumes() []string {
	return call.consumes
}

type Response struct {
	code string
	typ  string
}

type Field struct {
	name     string // raw name
	goName   string // camelCase name
	typ      string
	tag      string
	required bool
	inBody   bool
}

func canonicalConsumesList(iter *openapi.MIMETypeListIterator) ([]string, error) {
	consumesSeen := map[string]struct{}{}

	var consumesList []string
	for iter.Next() {
		mt := iter.Item()
		if _, ok := consumesSeen[mt]; ok {
			continue
		}
		consumesList = append(consumesList, mt)
		consumesSeen[mt] = struct{}{}
	}

	// Make sure the consumes list is valid
	for i, v := range consumesList {
		mt, _, err := mime.ParseMediaType(v)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to parse "consumes" value %s`, v)
		}
		// Use the canonical mime type (the parsed one)
		consumesList[i] = mt
	}
	return consumesList, nil
}

func Generate(spec openapi.Swagger, options ...Option) error {
	if err := spec.Validate(true); err != nil {
		return errors.Wrap(err, `failed to validate spec`)
	}

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

	ctx := genCtx{
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

func writeTypesFile(ctx *genCtx) error {
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

func writeOptionsFile(ctx *genCtx) error {
	fn := filepath.Join(ctx.dir, "options_gen.go")
	log.Printf("Generating %s", fn)

	var buf bytes.Buffer
	var dst io.Writer = &buf
	codegen.WritePreamble(dst, ctx.packageName)

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

	if err := codegen.WriteFormattedToFile(fn, buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, bytes.NewReader(buf.Bytes()))
		return errors.Wrapf(err, `failed to write to %s`, fn)
	}
	return nil
}

func writeClientFile(ctx *genCtx) error {
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

func writeServiceFiles(ctx *genCtx) error {
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

func callName(oper openapi.Operation) string {
	if operID := oper.OperationID(); operID != "" {
		return strcase.ToCamel(operID) + "Call"
	}

	pi := oper.PathItem()
	if pi == nil {
		buf, err := yaml.Marshal(oper)
		if err != nil {
			fmt.Fprintf(os.Stdout, err.Error())
		} else {
			os.Stdout.Write(buf)
		}
		panic("PathItem for operation is nil")
	}

	verb := strings.ToLower(oper.Verb())
	return strcase.ToCamel(verb+" "+strcase.ToCamel(oper.PathItem().Path())) + "Call"
}

func compileClient(ctx *genCtx) error {
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

func registerType(ctx *genCtx, path string, t Type, where string) {
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

func compileStruct(ctx *genCtx, schema openapi.Schema) (Type, error) {
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

func resolveReference(ctx *genCtx, ref string, decodeFunc func([]byte) error) error {
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

func compileItems(ctx *genCtx, items openapi.Items) (t Type, err error) {
	return compileSchemaLike(ctx, items)
}

func compileSchema(ctx *genCtx, schema openapi.Schema) (t Type, err error) {
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

func compileSchemaLike(ctx *genCtx, schema openapiTypeFormater) (Type, error) {
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

func compileBuiltin(ctx *genCtx, schema openapiTypeFormater) (Type, error) {
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

func compileResponseType(ctx *genCtx, response openapi.Response) (string, error) {
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

func compileParameterType(ctx *genCtx, param openapi.Parameter) (Type, error) {
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

func compileArray(ctx *genCtx, schema interface{}) (Type, error) {
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

func compileCall(ctx *genCtx, oper openapi.Operation) error {
	// x-service dictates the service name. If not present,
	// the default service, which is named after the package
	// is used.

	call := &Call{
		name: callName(oper),
		path: oper.PathItem().Path(),
		verb: oper.Verb(),
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

	var formBuilder *openapi.SchemaBuilder
	for piter := oper.Parameters(); piter.Next(); {
		param := piter.Item()

		switch param.In() {
		case openapi.InForm:
			// Use the values from the parameter to construct a schema
			if formBuilder == nil {
				formBuilder = openapi.NewSchema()
			}

			collectionFormat := param.CollectionFormat()
			if collectionFormat == "" {
				collectionFormat = openapi.CSV
			}
			b := openapi.NewSchema().
				Type(param.Type()).
				Format(param.Format()).
				Pattern(param.Pattern()).
				UniqueItems(param.UniqueItems()).
				Enum(param.Enum()).
				Default(param.Default())
			if param.HasMaximum() {
				b.Maximum(param.Maximum())
			}
			if param.HasMinimum() {
				b.Minimum(param.Minimum())
			}
			if param.HasExclusiveMaximum() {
				b.ExclusiveMaximum(param.ExclusiveMaximum())
			}
			if param.HasExclusiveMinimum() {
				b.ExclusiveMinimum(param.ExclusiveMinimum())
			}
			if param.HasMaxLength() {
				b.MaxLength(param.MaxLength())
			}
			if param.HasMinLength() {
				b.MinLength(param.MinLength())
			}
			if param.HasMaxItems() {
				b.MaxItems(param.MaxItems())
			}
			if param.HasMinItems() {
				b.MinItems(param.MinItems())
			}
			if param.HasMultipleOf() {
				b.MultipleOf(param.MultipleOf())
			}

			prop, err := b.Do()
			if err != nil {
				return errors.Wrapf(err, `failed to create schema for form parameter %s`, param.Name())
			}

			formBuilder.Property(param.Name(), prop)
		case openapi.InBody:
			// sanity check (although this should have already been taken care
			// of in Validate())
			if call.bodyType != nil {
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

			switch typ := typ.(type) {
			case *Struct:
				call.bodyType = typ
				for _, field := range typ.fields {
					field.inBody = true
					if field.required {
						call.requireds = append(call.requireds, field)
					} else {
						call.optionals = append(call.optionals, field)
					}
				}
			default:
				return errors.Errorf("body parameter handling for %T is not implemented", typ)
			}
		default:
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

			switch param.In() {
			case openapi.InPath:
				call.pathparams = append(call.pathparams, &field)
			case openapi.InQuery:
				call.queryparams = append(call.queryparams, &field)
			}
		}
	}

	// if we have form fields, compile that into a struct
	if formBuilder != nil {
		formSchema, err := formBuilder.Do()
		if err != nil {
			return errors.Wrap(err, `failed to build schema for formData fields`)
		}
		formType, err := compileStruct(ctx, formSchema)
		if err != nil {
			return errors.Wrap(err, `failed to compile schema for formData fields`)
		}
		if formType.Name() == "" {
			formType.SetName(codegen.ExportedName(call.name + "_Form"))
			registerType(ctx, fmt.Sprintf("#/generated/%s", formType.Name()), formType, call.name+" body form")
		}
		call.formType = formType

		if len(call.consumes) == 0 {
			call.consumes = append(call.consumes, "application/x-www-form-urlencoded")
		}
	}

	// grab the service.
	svcName := ctx.defaultServiceName
	rawSvcName, ok := oper.Extension("x-service")
	if ok {
		svcName, ok = rawSvcName.(string)
		if !ok {
			return errors.Errorf(`expected x-service to be a string`)
		}
	}

	svcName = strcase.ToCamel(svcName)
	svc := ctx.client.getServiceFor(svcName)
	svc.addCall(call)
	return nil
}

func formatClient(ctx *genCtx, dst io.Writer, cl *Client) error {
	codegen.WritePreamble(dst, ctx.packageName)
	codegen.WriteImports(dst, "bytes", "io", "net/http", "github.com/pkg/errors")
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

	fmt.Fprintf(dst, "\n\nfunc encodeCallPayload(marshalers map[string]Marshaler, mtype string, payload interface{}) (io.Reader, error) {")
	fmt.Fprintf(dst, "\nmarshaler, ok := marshalers[mtype]")
	fmt.Fprintf(dst, "\nif !ok {")
	fmt.Fprintf(dst, "\nreturn nil, errors.Errorf(`missing marshaler for request content type %s`, mtype)")
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

func formatService(ctx *genCtx, dst io.Writer, svc *Service) error {
	log.Printf(" * Generating Service %s", svc.name)
	codegen.WritePreamble(dst, ctx.packageName)
	codegen.WriteImports(dst, "context", "encoding/json", "mime", "net/http", "net/url", "strings", "strconv", "github.com/pkg/errors", "github.com/lestrrat-go/urlenc")

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
	if call.bodyType != nil {
		fmt.Fprintf(dst, "\nbody %s", call.bodyType.Name())
	} else if call.formType != nil {
		fmt.Fprintf(dst, "\nform %s", call.formType.Name())
	}

	for _, field := range allFields {
		if field.inBody {
			continue
		}

		fmt.Fprintf(dst, "\n%s %s", field.goName, field.typ)
	}
	fmt.Fprintf(dst, "\n}")

	sort.Slice(call.requireds, func(i, j int) bool {
		return call.requireds[i].name < call.requireds[j].name
	})

	log.Printf("      * Generating constructor")
	fmt.Fprintf(dst, "\n\nfunc (svc *%s) %s(", svcName, strings.TrimSuffix(call.name, "Call"))
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
		if field.inBody {
			fmt.Fprintf(dst, "\ncall.body.%s = %s", field.goName, codegen.UnexportedName(field.goName))
		} else {
			fmt.Fprintf(dst, "\ncall.%s = %s", field.goName, codegen.UnexportedName(field.goName))
		}
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
			if optional.inBody {
				fmt.Fprintf(dst, "\ncall.body.%[1]s = append(call.body.%[1]s, v...)", codegen.ExportedName(optional.name))
			} else {
				fmt.Fprintf(dst, "\ncall.%[1]s = append(call.%[1]s, v...)", optional.name)
			}
			fmt.Fprintf(dst, "\nreturn call")
			fmt.Fprintf(dst, "\n}")
		} else {
			fmt.Fprintf(dst, "\n\nfunc (call *%s) %s(v %s) *%s {", call.name, stringutil.Camel(optional.name), optional.typ, call.name)
			if optional.inBody {
				fmt.Fprintf(dst, "\ncall.body.%s = v", codegen.ExportedName(optional.name))
			} else {
				fmt.Fprintf(dst, "\ncall.%s = v", optional.name)
			}
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
	fmt.Fprintf(dst, "\nconst basepath = %s", strconv.Quote(call.path))

	if call.bodyType != nil || call.formType != nil {
		fmt.Fprintf(dst, "\n\ncontentType := %#v", call.DefaultConsumes())
		fmt.Fprintf(dst, "\nfor _, option := range options {")
		fmt.Fprintf(dst, "\nswitch option.Name() {")
		fmt.Fprintf(dst, "\ncase optkeyRequestContentType:")
		fmt.Fprintf(dst, "\ncontentType = option.Value().(string)")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\n}")
	}

	fmt.Fprintf(dst, "\npath := basepath")
	for _, pathparam := range call.pathparams {
		fmt.Fprintf(dst, "\npath = strings.Replace(path, `{%s}`, ", pathparam.name)
		switch pathparam.typ {
		case "int64":
			fmt.Fprintf(dst, "strconv.FormatInt(call.%s, 10)", pathparam.name)
		default:
			fmt.Fprintf(dst, "call.%s", pathparam.name)
		}
		fmt.Fprintf(dst, ", -1)")
	}

	if len(call.queryparams) > 0 {
		fmt.Fprintf(dst, "\nv := url.Values{}")
		for _, queryparam := range call.queryparams {
			// XXX This needs to be more robust
			if queryparam.typ == "[]string" {
				fmt.Fprintf(dst, "\nfor _, param := range call.%s {", queryparam.goName)
				fmt.Fprintf(dst, "\nv.Add(%s, param)", strconv.Quote(queryparam.name))
				fmt.Fprintf(dst, "\n}")
			} else {
				fmt.Fprintf(dst, "\nv.Add(%s, ", strconv.Quote(queryparam.name))
				switch queryparam.typ {
				case "int64":
					fmt.Fprintf(dst, "strconv.FormatInt(call.%s, 10)", queryparam.name)
				default:
					fmt.Fprintf(dst, "call.%s", queryparam.goName)
				}
				fmt.Fprintf(dst, ")")
			}
		}

		fmt.Fprintf(dst, "\npath = call.server + path + `?` + v.Encode()")
	}

	var body = "nil"
	if call.bodyType != nil || call.formType != nil {
		body = "body"

		fmt.Fprintf(dst, "\nmtype, _, err := mime.ParseMediaType(contentType)")
		fmt.Fprintf(dst, "\nif err != nil {")
		fmt.Fprintf(dst, "\nreturn nil, errors.Wrapf(err, `failed to parse request content type %%s`, contentType)")
		fmt.Fprintf(dst, "\n}")

		if call.bodyType != nil {
			fmt.Fprintf(dst, "\n\nbody, err := encodeCallPayload(call.marshalers, mtype, call.body)")
		} else if call.formType != nil {
			fmt.Fprintf(dst, "\n\nbody, err := encodeCallPayload(call.marshalers, mtype, call.form)")
		} else {
			return errors.New(`can't proceed when call.bodyType == nil and call.formType == nil`)
		}
		fmt.Fprintf(dst, "\nif err != nil {")
		fmt.Fprintf(dst, "\nreturn nil, errors.Wrapf(err, `failed to marshal request payload as %%s`, mtype)")
		fmt.Fprintf(dst, "\n}")
	}

	fmt.Fprintf(dst, "\n\nreq, err := http.NewRequest(%s, path, %s)", strconv.Quote(call.verb), body)
	fmt.Fprintf(dst, "\nif err != nil {")
	fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `failed to create request`)")
	fmt.Fprintf(dst, "\n}")

	if call.bodyType != nil || call.formType != nil {
		fmt.Fprintf(dst, "\nreq.Header.Set(`Content-Type`, contentType)")
	}

	fmt.Fprintf(dst, "\n\nres, err := call.httpCl.Do(req)")
	fmt.Fprintf(dst, "\nif err != nil {")
	fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `failed to make HTTP request`)")
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
