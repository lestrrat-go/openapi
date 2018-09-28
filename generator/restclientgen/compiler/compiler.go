package compiler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"sort"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/lestrrat-go/openapi/internal/codegen/golang"
	restclient "github.com/lestrrat-go/openapi/internal/codegen/restclient/golang"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

// Takes an OpenAPI v2 structure, and compiles it into a form that all REST
// clients share

func Compile(spec openapi.OpenAPI, defaultServiceName string) (*ClientDefinition, error) {
	ctx := &compileCtx{
		compiling:          make(map[string]struct{}),
		defaultServiceName: defaultServiceName,
		isCompiling:        make(map[interface{}]struct{}),
		isResolving:        make(map[interface{}]struct{}),
		resolver:           openapi.NewResolver(spec),
		root:               spec,
		client: &ClientDefinition{
			services:    make(map[string]*Service),
			definitions: make(map[string]TypeDefinition),
			types:       make(map[string]Type),
		},
	}

	ctx.security = make(map[string]openapi.SecurityScheme)
	for iter := spec.SecurityDefinitions(); iter.Next(); {
		name, scheme := iter.Item()
		ctx.security[name] = scheme
	}

	// declare types
	if err := compileGlobalDefaults(ctx); err != nil {
		return nil, errors.Wrap(err, `failed to compile global defaults`)
	}

	if err := compileClient(ctx); err != nil {
		return nil, errors.Wrap(err, `failed to compile client`)
	}

	if err := ctx.resolveIncomplete(ctx.client); err != nil {
		return nil, errors.Wrap(err, `failed to resolve incomplete types`)
	}

	for _, svc := range ctx.client.services {
		sort.Slice(svc.calls, func(i, j int) bool {
			return svc.calls[i].name < svc.calls[j].name
		})

		for _, call := range svc.calls {
			sort.Slice(call.optionals, func(i, j int) bool {
				return call.optionals[i].name < call.optionals[i].name
			})
			sort.Slice(call.requireds, func(i, j int) bool {
				return call.requireds[i].name < call.requireds[j].name
			})
			sort.Slice(call.responses, func(i, j int) bool {
				iv, _ := strconv.Atoi(call.responses[i].code)
				jv, _ := strconv.Atoi(call.responses[j].code)
				return iv < jv
			})

			call.allFields = append(append([]*Field(nil), call.requireds...), call.optionals...)
			sort.Slice(call.allFields, func(i, j int) bool {
				return call.allFields[i].Name() < call.allFields[j].Name()
			})
		}
	}

	return ctx.client, nil
}

func (ctx *compileCtx) IsCompiling(v interface{}) bool {
	_, ok := ctx.isCompiling[v]
	return ok
}

func (ctx *compileCtx) MarkAsCompiling(v interface{}) func() {
	ctx.isCompiling[v] = struct{}{}
	return func() {
		delete(ctx.isCompiling, v)
	}
}

func (ctx *compileCtx) IsResolving(v interface{}) bool {
	_, ok := ctx.isResolving[v]
	return ok
}

func (ctx *compileCtx) MarkAsResolving(v interface{}) func() {
	ctx.isResolving[v] = struct{}{}
	return func() {
		delete(ctx.isResolving, v)
	}
}

func (ctx *compileCtx) resolveIncomplete(c *ClientDefinition) error {
	for _, def := range c.definitions {
		typ, err := def.Type.ResolveIncomplete(ctx)
		if err != nil {
			return errors.Wrap(err, `failed to resolve incomplete type in TypeDefinition`)
		}
		def.Type = typ
	}

	for name, typ := range c.types {
		typ, err := typ.ResolveIncomplete(ctx)
		if err != nil {
			return errors.Wrap(err, `failed to resolve incomplete type in types`)
		}
		c.types[name] = typ
	}
	return nil
}

func compileGlobalDefaults(ctx *compileCtx) error {
	consumes, err := canonicalConsumesList(ctx.root.Consumes())
	if err != nil {
		return errors.Wrap(err, `failed to parse global "consumes" list`)
	}
	ctx.consumes = consumes
	return nil
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

func compileClient(ctx *compileCtx) error {
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

func createCall(ctx *compileCtx, oper openapi.Operation) (*Call, error) {
	callName := restclient.CallObjectName(oper)
	methodName := restclient.CallMethodName(oper)
	if methodName == "" {
		methodName = golang.ExportedName(strings.TrimSuffix(callName, "Call"))
	}

	return &Call{
		name:        callName,
		method:      methodName,
		requestPath: oper.PathItem().Path(),
		verb:        oper.Verb(),
	}, nil
}

func compileCallConsumesList(ctx *compileCtx, oper openapi.Operation) ([]string, error) {
	consumesList, err := canonicalConsumesList(oper.Consumes())
	if err != nil {
		return nil, errors.Wrapf(err, `failed to parse consumes list for %s:%s`, oper.PathItem().Path(), oper.Verb())
	}

	if len(consumesList) == 0 {
		// Use the default consumes list if not provided
		consumesList = append(consumesList, ctx.consumes...)
	}
	return consumesList, nil
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

func resolveReference(ctx *compileCtx, ref string, decodeFunc func([]byte) error) error {
	if _, ok := ctx.compiling[ref]; ok {
		return errors.Errorf(`circular dep on %s`, ref)
	}

	ctx.compiling[ref] = struct{}{}
	defer func() {
		delete(ctx.compiling, ref)
	}()

	var thing interface{}
	if cached, ok := ctx.client.definitions[ref]; ok {
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

func lookupReferencedType(ctx *compileCtx, path string) (Type, bool) {
	typdef, ok := ctx.client.definitions[path]
	if !ok {
		return nil, false
	}

	return typdef.Type, true
}

func registerType(ctx *compileCtx, path string, t Type, where string) {
	if t.Name() == "" {
		panic("anonymous type")
	}

	if _, ok := ctx.client.types[t.Name()]; ok {
		return
	}

	log.Printf(" * Registering type %s (%s)", t.Name(), path)
	ctx.client.definitions[path] = TypeDefinition{
		Path:    path,
		Type:    t,
		Context: where,
	}
	ctx.client.types[t.Name()] = t
}

func compileBuiltin(ctx *compileCtx, schema openapiTypeFormater) (Type, error) {
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

func compileItems(ctx *compileCtx, items openapi.Items) (t Type, err error) {
	return compileSchemaLike(ctx, items)
}

func compileArray(ctx *compileCtx, schema interface{}) (Type, error) {
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
	return &Array{elem: subtyp}, nil
}

func compileParameterToProperty(parentBuilder *openapi.SchemaBuilder, param openapi.Parameter) error {
	prop, err := param.ConvertToSchema()
	if err != nil {
		return errors.Wrap(err, `failed to convert parameter to property (openapi.Schema)`)
	}
	parentBuilder.Property(param.Name(), prop)
	return nil
}

func compileStruct(ctx *compileCtx, schema openapi.Schema) (Type, error) {
	var obj Struct

	for piter := schema.Properties(); piter.Next(); {
		name, prop := piter.Item()

		fieldMsg, err := compileSchema(ctx, prop)
		if err != nil {
			return nil, errors.Wrap(err, `failed to compile schema for object property`)
		}

		obj.fields = append(obj.fields, &Field{
			name: name,
			hints: Hints{
				GoName: golang.ExportedName(name),
				GoTag:  fmt.Sprintf(`json:"%s"`, name),
			},
			typ:      fieldMsg,
			required: schema.IsRequiredProperty(name),
		})
	}
	return &obj, nil
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

func compileSchemaLike(ctx *compileCtx, schema openapiTypeFormater) (Type, error) {
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

func compileSchema(ctx *compileCtx, schema openapi.Schema) (t Type, err error) {
	if schema == nil {
		return nil, errors.New(`nil schema`)
	}

	if ref := schema.Reference(); ref != "" {
		if typ, ok := lookupReferencedType(ctx, ref); ok {
			return typ, nil
		}

		if ctx.IsCompiling(ref) {
			return Incomplete(ref), nil
		}

		cancel := ctx.MarkAsCompiling(ref)
		defer cancel()

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
				n := golang.ExportedName(strings.TrimPrefix(ref, "#/definitions/"))
				t.SetName(n)
			}
			registerType(ctx, ref, t, ref)
		}()

	}

	return compileSchemaLike(ctx, schema)
}

func compileResponseType(ctx *compileCtx, response openapi.Response) (string, error) {
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
			typ.SetName(golang.ExportedName(ctx.currentCall.name + "_Response"))
			registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, ctx.currentCall.name+" response")
		}
		return typ.Name(), nil
	default:
		return "", errors.Errorf(`unimplemented %s`, schema.Type())
	}
}

func compileResponse(ctx *compileCtx, res openapi.Response) (*Response, error) {
	typ, err := compileResponseType(ctx, res)
	if err != nil {
		return nil, errors.Wrap(err, `failed to compile response type`)
	}

	return &Response{
		code: res.StatusCode(),
		typ:  typ,
	}, nil
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

func compileParameterType(ctx *compileCtx, param openapi.Parameter) (Type, error) {
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

			return &Array{elem: typ}, nil
		case openapi.Object:
			typ, err := compileSchema(ctx, schema)
			if err != nil {
				return nil, errors.Wrap(err, `failed to compile object parameter`)
			}

			if typ.Name() == "" {
				typ.SetName(golang.ExportedName(ctx.currentCall.name + "_" + param.Name()))
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

func compileStructFromBuilder(ctx *compileCtx, builder interface {
	Build(...openapi.Option) (openapi.Schema, error)
}, location openapi.Location) (Type, error) {
	schema, err := builder.Build()
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

func compileCallParameters(ctx *compileCtx, oper openapi.Operation, call *Call) error {
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
		}
	}

	if queryBuilder != nil {
		typ, err := compileStructFromBuilder(ctx, queryBuilder, openapi.InQuery)
		if err != nil {
			return errors.Wrap(err, `failed to compile schema for query fields`)
		}
		typ.SetName(golang.ExportedName(call.name + "_Query"))
		registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, call.name+" query")
		extractFields(call, typ)
		call.query = typ
	}

	if headerBuilder != nil {
		typ, err := compileStructFromBuilder(ctx, headerBuilder, openapi.InHeader)
		if err != nil {
			return errors.Wrap(err, `failed to compile schema for header fields`)
		}
		typ.SetName(golang.ExportedName(call.name + "_Header"))
		registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, call.name+" header")
		extractFields(call, typ)
		call.header = typ
	}

	if pathBuilder != nil {
		typ, err := compileStructFromBuilder(ctx, pathBuilder, openapi.InPath)
		if err != nil {
			return errors.Wrap(err, `failed to compile schema for path fields`)
		}
		typ.SetName(golang.ExportedName(call.name + "_Path"))
		registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, call.name+" path")
		extractFields(call, typ)
		call.path = typ
	}

	if formBuilder != nil {
		typ, err := compileStructFromBuilder(ctx, formBuilder, openapi.InPath)
		if err != nil {
			return errors.Wrap(err, `failed to compile schema for form fields`)
		}
		typ.SetName(golang.ExportedName(call.name + "_Form"))
		registerType(ctx, fmt.Sprintf("#/generated/%s", typ.Name()), typ, call.name+" form")
		extractFields(call, typ)
		call.body = typ

		if len(call.consumes) == 0 {
			call.consumes = append(call.consumes, "application/x-www-form-urlencoded")
		}
	}

	return nil
}

func compileServiceName(ctx *compileCtx, oper openapi.Operation) (string, error) {
	svcName := ctx.defaultServiceName

	// x-oagen-service, then x-service
	for _, key := range []string{"x-oagen-service", "x-service"} {
		rawSvcName, ok := oper.Extension(key)
		if ok {
			svcName, ok = rawSvcName.(string)
			if !ok {
				return "", errors.Errorf(`expected x-service to be a string`)
			}
			break
		}
	}

	svcName = strcase.ToCamel(svcName)
	return svcName, nil
}

func compileSecuritySettings(ctx *compileCtx, requirement openapi.SecurityRequirement) ([]*SecuritySettings, error) {
	var list []*SecuritySettings
	for siter := requirement.Scopes(); siter.Next(); {
		name, scopes := siter.Item()
		security, ok := ctx.security[name]
		if !ok {
			return nil, errors.Errorf(`invalid security definition %s (not found)`, name)
		}
		list = append(list, &SecuritySettings{
			definition: security,
			scopes:     scopes,
		})
	}
	return list, nil
}

func compileCall(ctx *compileCtx, oper openapi.Operation) error {
	// x-service dictates the service name. If not present,
	// the default service, which is named after the package
	// is used.

	call, err := createCall(ctx, oper)
	if err != nil {
		return errors.Wrap(err, `failed to create Call object`)
	}
	ctx.currentCall = call

	// The OpenAPI spec allows you to specify multiple "consumes"
	// clause, but we only support JSON or appliation/x-www-form-urlencoded
	// by default
	consumesList, err := compileCallConsumesList(ctx, oper)
	if err != nil {
		return errors.Wrapf(err, `failed to parse consumes list for %s:%s`, oper.PathItem().Path(), oper.Verb())
	}
	call.consumes = consumesList

	for respiter := oper.Responses().Responses(); respiter.Next(); {
		_, resp := respiter.Item()
		res, err := compileResponse(ctx, resp)
		if err != nil {
			return errors.Wrapf(err, `failed to compile response for %s:%s`, oper.PathItem().Path(), oper.Verb())
		}
		call.responses = append(call.responses, res)

	}

	if err := compileCallParameters(ctx, oper, call); err != nil {
		return errors.Wrapf(err, `failed to compile operation parameters for %s:%s`, oper.PathItem().Path(), oper.Verb())
	}

	// grab the service.
	svcName, err := compileServiceName(ctx, oper)
	if err != nil {
		return errors.Wrapf(err, `failed to compile service name for %s:%s`, oper.PathItem().Path(), oper.Verb())
	}

	// Check if we have security
	for iter := oper.Security(); iter.Next(); {
		requirement := iter.Item()
		ss, err := compileSecuritySettings(ctx, requirement)
		if err != nil {
			return errors.Wrapf(err, `failed to compile security settings for %s:%s`, oper.PathItem().Path(), oper.Verb())
		}
		call.securitySettings = ss
	}

	svc := ctx.client.getServiceFor(svcName)
	svc.addCall(call)
	return nil
}
