package grpcgen

// In schema-like objects such as Schema, Parameter, and Items, we
// respect an extension called `x-proto-type`. This value should
// be a hash:
//
// "x-proto-type": {
//   "name": "MyMessage"
// }
//
// If the type requires importing from another library, you should
// specify the "import" field:
//
// "x-proto-type": {
//   "name": "google.protobuf.Timestamp",
//   "import": "google/protobuf/timestamp.proto"
// }

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/iancoleman/strcase"
	"github.com/lestrrat-go/openapi/internal/codegen/common"
	codegen "github.com/lestrrat-go/openapi/internal/codegen/grpc"
	"github.com/lestrrat-go/openapi/openapi2"
	"github.com/pkg/errors"
)

type SchemaLike interface {
	Type() openapi2.PrimitiveType
	Format() string
	Default() interface{}
	Maximum() float64
	ExclusiveMaximum() float64
	Minimum() float64
	ExclusiveMinimum() float64
	MaxLength() int
	MinLength() int
	Pattern() string
	MaxItems() int
	MinItems() int
	UniqueItems() bool
	Enum() *openapi2.InterfaceListIterator
	MultipleOf() float64
	Extension(string) (interface{}, bool)
	Extensions() *openapi2.ExtensionsIterator
	Reference() string
}

type globalOption struct {
	name  string
	value string
}

type nonFatal struct{}

func (e nonFatal) Error() string {
	return "non fatal resolution error"
}
func (e nonFatal) Fatal() bool {
	return false
}

type Resolver struct {
	resolver openapi2.Resolver
}

func (r *Resolver) Resolve(ref string) (interface{}, error) {
	// TOOD
	if ref == "google/protobuf/timestamp.proto#/google.protobuf.Timestamp" {
		return nil, nonFatal{}
	}

	return r.resolver.Resolve(ref)
}

func Generate(ctx context.Context, spec openapi2.OpenAPI, options ...Option) error {
	var dst io.Writer = os.Stdout
	var globalOptions []*globalOption
	var annotate bool
	var packageName string = "myapp"
	for _, o := range options {
		switch o.Name() {
		case optkeyPackageName:
			packageName = o.Value().(string)
		case optkeyDestination:
			dst = o.Value().(io.Writer)
		case optkeyGlobalOption:
			globalOptions = append(globalOptions, o.Value().(*globalOption))
		case optkeyAnnotation:
			annotate = o.Value().(bool)
		}
	}

	resolver := &Resolver{resolver: openapi2.NewResolver(spec)}
	/*
		if err := spec.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve references in openapi spec`)
		}
	*/

	c := &genCtx{
		Context:     ctx,
		annotate:    annotate,
		dst:         dst,
		isCompiling: map[string]struct{}{},
		isResolving: map[interface{}]struct{}{},
		resolver:    resolver,
		root:        spec,
		proto: &Protobuf{
			packageName:   packageName,
			globalOptions: globalOptions,
			imports:       make(map[string]struct{}),
		},

		// types that are defined at the top level through references.
		types: make(map[string]Type),
	}
	c.parent = c.proto
	if c.annotate {
		c.proto.AddImport("google/api/annotations.proto")
	}

	return generate(c)
}

func (ctx *genCtx) RegisterMessage(path string, typ Type) {
	ctx.log("* Registering type %s (%s)", path, typ.Name())
	ctx.types[path] = typ
}

func (ctx *genCtx) LookupType(path string) (Type, bool) {
	typ, ok := ctx.types[path]
	if ok {
		// If this is a pointer we need to clone this,
		// otherwise we have problems by consumers who
		// may want to modify this value
		rv := reflect.ValueOf(typ)
		switch rv.Kind() {
		case reflect.Ptr, reflect.Interface:
			copy := reflect.New(rv.Elem().Type())
			copy.Elem().Set(rv.Elem())
			typ = copy.Interface().(Type)
		}
	}
	return typ, ok
}

func (ctx *genCtx) IsResolving(v interface{}) bool {
	_, ok := ctx.isResolving[v]
	return ok
}

func (ctx *genCtx) MarkAsResolving(v interface{}) func() {
	ctx.isResolving[v] = struct{}{}
	return func() {
		delete(ctx.isResolving, v)
	}
}

func (ctx *genCtx) IsCompiling(path string) bool {
	_, ok := ctx.isCompiling[path]
	return ok
}

func (ctx *genCtx) MarkAsCompiling(path string) func() {
	ctx.isCompiling[path] = struct{}{}
	return func() {
		delete(ctx.isCompiling, path)
	}
}

func grpcMethodName(oper openapi2.Operation) string {
	if v, ok := oper.Extension(`x-rpc-name`); ok {
		if name, ok := v.(string); ok {
			return strcase.ToCamel(name)
		}
	}

	return grpcOperationID(oper)
}

func grpcOperationID(oper openapi2.Operation) string {
	if operID := oper.OperationID(); operID != "" {
		return strcase.ToCamel(operID)
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
	return strcase.ToCamel(verb + " " + strcase.ToCamel(oper.PathItem().Path()))
}

func generate(ctx *genCtx) error {
	if err := compile(ctx); err != nil {
		return errors.Wrap(err, `failed to compile from openapi spec`)
	}

	if err := format(ctx, ctx.dst, ctx.proto); err != nil {
		return errors.Wrap(err, `failed to format compile protobuf spec`)
	}

	return nil
}

func compile(ctx *genCtx) error {
	if err := compileGlobalDefinitions(ctx); err != nil {
		return errors.Wrap(err, `failed to compile definitions`)
	}

	if err := compileRPCs(ctx); err != nil {
		return errors.Wrap(err, `failed to compile RPC calls`)
	}

	return nil
}

func (ctx *genCtx) log(f string, args ...interface{}) {
	log.Printf(ctx.Indent(f, args...))
}

func (ctx *genCtx) Indent(f string, args ...interface{}) string {
	return ctx.indent + fmt.Sprintf(f, args...)
}

func (ctx *genCtx) Start(f string, args ...interface{}) func() {
	ctx.log(f, args...)
	ctx.indent = ctx.indent + "  "
	return func() {
		ctx.indent = ctx.indent[:len(ctx.indent)-2]
	}
}

func compileGlobalDefinitions(ctx *genCtx) error {
	done := ctx.Start("* Compiling Global Definitions")
	defer done()

	var messages []*Message
	var names []string
	for defiter := ctx.root.Definitions(); defiter.Next(); {
		name, thing := defiter.Item()

		var tmp openapi2.Schema
		if err := common.RoundTripDecode(&tmp, thing, openapi2.SchemaFromJSON); err != nil {
			return errors.Wrapf(err, `expected openapi2.Schema for #/defnitions/%s, got %T`, name, thing)
		}

		if err := tmp.Validate(true); err != nil {
			// If this doesn't look like a valid schema, it's OK. skip
			ctx.log(`* invalid schema object found in definition, skipping`)
			continue
		}

		ref := "#/definitions/" + name
		ctx.log("* Compiling Top Level definition %s", ref)
		cancel := ctx.MarkAsCompiling(ref)
		defer cancel()

		// If this is a message type (object), then we compile it as such.
		// Otherwise the user is simply using it as a way to reuse components,
		// so we ignore it
		if openapi2.GuessSchemaType(tmp) != openapi2.Object {
			ctx.log(`* referenced schema is not an object, not compiling as protobuf message`)
			typ, err := compileType(ctx, tmp)
			if err != nil {
				return errors.Wrapf(err, `failed to compile definition #/definitions/%s`, name)
			}

			ctx.RegisterMessage("#/definitions/"+name, typ)
		} else {
			typ, err := compileMessage(ctx, tmp)
			if err != nil {
				return errors.Wrapf(err, `failed to compile message #/definitions/%s`, name)
			}

			m, ok := typ.(*Message)
			if !ok {
				return errors.Errorf(`expected global definition to compile into a message, got %T`, typ)
			}

			m.name = codegen.MessageName(name)
			names = append(names, name)
			messages = append(messages, m)
			ctx.RegisterMessage("#/definitions/"+name, m)
		}
	}

	// Resolve incomplete types in the list
	for i, msg := range messages {
		ctx.log("Resolving incomplete types within #/definitions/%s", names[i])
		typ, err := msg.ResolveIncomplete(ctx)
		if err != nil {
			return errors.Wrapf(err, `failed to resolve incomplete types for #/definitions/%s`, names[i])
		}

		m, ok := typ.(*Message)
		if !ok {
			return errors.Errorf(`expected resolved definition to be a message, got %T`, typ)
		}

		ctx.proto.AddMessage(m)
	}
	return nil
}

func compileRPCs(ctx *genCtx) error {
	done := ctx.Start("* Compiling RPCs")
	defer done()

	for piter := ctx.root.Paths().Paths(); piter.Next(); {
		_, pi := piter.Item()
		for operiter := pi.Operations(); operiter.Next(); {
			oper := operiter.Item()
			rpc, err := compileRPC(ctx, oper)
			if err != nil {
				return errors.Wrap(err, `failed to compile rpc`)
			}
			var tag string
			if tagsiter := oper.Tags(); tagsiter.Next() {
				tag = tagsiter.Item()
			}
			if tag == "" {
				tag = ctx.proto.packageName
			}
			tag = strcase.ToCamel(tag)

			svc := ctx.proto.GetService(tag)
			svc.AddRPC(rpc)
		}
	}
	return nil
}

func compileBuiltin(ctx *genCtx, s openapi2.Schema) (Type, error) {
	switch s.Type() {
	case openapi2.Boolean:
		return Builtin("bool"), nil
	case openapi2.String:
		return Builtin(s.Type()), nil
	case openapi2.Number:
		switch f := s.Format(); f {
		case openapi2.Float, openapi2.Double:
			return Builtin(f), nil
		case "":
			return Builtin(openapi2.Float), nil
		default:
			return nil, errors.Errorf(`invalid number format %s`, f)
		}
	case openapi2.Integer:
		switch f := s.Format(); f {
		case openapi2.Int64, openapi2.Int32:
			return Builtin(f), nil
		case "":
			return Builtin(openapi2.Int32), nil
		default:
			return nil, errors.Errorf(`invalid integer format %s`, f)
		}
	default:
		return nil, errors.Errorf(`unknown builtin %s`, s.Type())
	}
}

// compileMessage takes a schema and compiles it into a Message.
// Messages normally need to be an object-like type, but in case
// we have an array, we will actually wrap the elements in an object.
func compileMessage(ctx *genCtx, schema openapi2.Schema) (Type, error) {
	done := ctx.Start("* Compiling Message")
	defer done()

	if ref := schema.Reference(); ref != "" {
		typ, ok := ctx.LookupType(ref)
		if ok {
			ctx.log("* Reference %s already compiled to %s", ref, typ.Name())
			return typ, nil
		}

		var tmp openapi2.Schema
		if err := common.RoundTripDecode(&tmp, typ, openapi2.SchemaFromJSON); err != nil {
			return nil, errors.Wrapf(err, `expected openapi2.Schema %s, got %T`, ref, typ)
		}
		schema = tmp
	}

	var allOfSchemas []openapi2.Schema
	for iter := schema.AllOf(); iter.Next(); {
		allOfSchemas = append(allOfSchemas, iter.Item())
	}

	if len(allOfSchemas) > 0 {
		ctx.log("* Merging %d schemas", len(allOfSchemas))
		merged := allOfSchemas[0]
		allOfSchemas = allOfSchemas[1:]
		for _, s := range allOfSchemas {
			m, err := openapi2.MergeSchemas(merged, s)
			if err != nil {
				return nil, errors.Wrap(err, `failed to merge schemas`)
			}
			merged = m
		}

		schema = merged
	}

	// it better be an object
	switch typ := openapi2.GuessSchemaType(schema); typ {
	case openapi2.Object:
	case openapi2.Array:
		// special case
		items, err := openapi2.NewSchema().
			Type(openapi2.Array).
			Items(schema.Items()).
			Build()
		if err != nil {
			return nil, errors.Wrap(err, `failed to create new schema for items`)
		}

		schema, err = openapi2.NewSchema().
			Type(openapi2.Object).
			Property("items", items).
			Build()
		if err != nil {
			return nil, errors.Wrap(err, `failed to create new schema wrapping array`)
		}
	default:
		return nil, errors.Errorf(`compileMessage: expected type "array" or "object", got %s`, typ)
	}

	var m Message
	oldparent := ctx.parent
	ctx.parent = &m
	defer func() {
		ctx.parent = oldparent
	}()

	for piter := schema.Properties(); piter.Next(); {
		name, prop := piter.Item()
		field, err := compileField(ctx, name, prop)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to compile property %s`, name)
		}
		m.fields = append(m.fields, field)
	}

	for i, f := range m.fields {
		f.id = i + 1
	}

	return &m, nil
}

func compileType(ctx *genCtx, src openapi2.SchemaConverter) (result Type, err error) {
	return compileTypeWithName(ctx, src, "")
}

// compileType takes something that looks like a schema, and creates an
// appropriate type for it.
func compileTypeWithName(ctx *genCtx, src openapi2.SchemaConverter, name string) (result Type, err error) {
	schema, err := src.ConvertToSchema()
	if err != nil {
		return nil, errors.Wrapf(err, `failed to extract schema out of %T`, src)
	}

	var registerGlobal bool
	if ref := schema.Reference(); ref != "" {
		done := ctx.Start("* Compiling Reference %s", ref)
		defer done()

		typ, ok := ctx.LookupType(ref)
		if ok {
			ctx.log("* Reference %s already compiled to %s", ref, typ.Name())
			return typ, nil
		}

		// If we had successfully compiled this reference, we would
		// have found it in the previous LookupType call. if we haven't
		// completed compiling it, we would see it in this next lookup
		if ctx.IsCompiling(ref) {
			ctx.log("* Circular dependency detected")
			return Incomplete(ref), nil
		}
		cancel := ctx.MarkAsCompiling(ref)
		defer cancel()

		thing, err := ctx.resolver.Resolve(ref)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to resolve reference %s`, ref)
		}

		var tmp openapi2.Schema
		if err := common.RoundTripDecode(&tmp, thing, openapi2.SchemaFromJSON); err != nil {
			return nil, errors.Errorf(`expected reference %s to resolve to a Schema`, ref)
		}
		registerGlobal = true
		schema = tmp
		if schema.Type() == openapi2.Object {
			if strings.HasPrefix(ref, "#/definitions/") {
				// If this refers to a global definition but the previous LookupType
				// failed, this can only mean that we encountered a global definition
				// that refers to another global definition.
				defer func() {
					if err != nil {
						return
					}
					if m, ok := result.(*Message); ok {
						m.name = codegen.MessageName(strings.TrimPrefix(ref, "#/definitions/"))
						m.reference = ref
						ctx.log("* Adding message %s (global)", m.Name())
						ctx.RegisterMessage(ref, m)
					}
				}()
			}
		}
	}

	merged, err := openapi2.MergedSchema(schema, ctx.resolver)
	if err != nil {
		return nil, errors.Wrap(err, `failed to merge schemas`)
	}
	schema = merged

	done := ctx.Start("* Compiling Type %s", schema.Type())
	defer done()

	switch openapi2.GuessSchemaType(schema) {
	case openapi2.String, openapi2.Number, openapi2.Boolean, openapi2.Integer:
		return compileBuiltin(ctx, schema)
	case openapi2.Object:
		typ, err := compileMessage(ctx, schema)
		if err != nil {
			return nil, errors.Wrap(err, `failed to compile message`)
		}
		ctx.log(" -----> %#v", typ)
		if !registerGlobal {
			if m, ok := typ.(*Message); ok {
				m.name = codegen.MessageName(name)
				ctx.log("* Adding message %s", m.Name())
				ctx.parent.AddMessage(m)
			}
		}
		return typ, nil
	case openapi2.Array:
		typ, err := compileType(ctx, schema.Items())
		if err != nil {
			return nil, errors.Wrap(err, `failed to compile array element`)
		}
		return &Array{element: typ}, nil
	}

	return nil, errors.Errorf(`compileType: unsupported schema.type = %s`, schema.Type())
}

func compileField(ctx *genCtx, name string, s openapi2.Schema) (*Field, error) {
	done := ctx.Start("* Compiling Field %s", name)
	defer done()

	typ, err := compileType(ctx, s)
	if err != nil {
		return nil, errors.Wrap(err, `failed to compile field`)
	}

	// name must be normalized to something snake_case
	return &Field{
		id:   1,
		name: codegen.FieldName(name),
		typ:  typ,
	}, nil
}

func compileRPCParameters(ctx *genCtx, name string, iter *openapi2.ParameterListIterator) (*Message, error) {
	done := ctx.Start("* Compiling RPC parameters for %s", name)
	defer done()

	var m Message

	oldparent := ctx.parent
	ctx.parent = &m
	defer func() {
		ctx.parent = oldparent
	}()

	var id = 1
	for iter.Next() {
		param := iter.Item()

		ctx.log("* Compiling parameter %s", param.Name())

		var typ Type
		var err error
		if param.In() == openapi2.InBody {
			typ, err = compileTypeWithName(ctx, param.Schema(), param.Name())
		} else {
			typ, err = compileType(ctx, param)
		}
		if err != nil {
			return nil, errors.Wrapf(err, `failed to deduce gRPC type for parameter %s`, strconv.Quote(param.Name()))
		}

		// if this type was not a builtin, we need to register it

		m.fields = append(m.fields, &Field{
			id:   id,
			name: codegen.FieldName(param.Name()),
			typ:  typ,
			body: param.In() == openapi2.InBody,
		})

		id++
	}

	m.name = name

	return &m, nil
}

func compileRPC(ctx *genCtx, oper openapi2.Operation) (*RPC, error) {
	var rpc RPC
	rpc.name = grpcMethodName(oper)
	rpc.in = Builtin("google.protobuf.Empty")
	rpc.out = Builtin("google.protobuf.Empty")
	if desc := oper.Description(); len(desc) > 0 {
		rpc.description = desc
	}

	done := ctx.Start("* Compiling RPC %s", rpc.name)
	defer done()

	operID := grpcOperationID(oper)

	paramiter := oper.Parameters()
	if paramiter.Next() {
		msg, err := compileRPCParameters(ctx, operID+"Request", paramiter)
		if err != nil {
			return nil, errors.Wrap(err, `failed to compile request parameters into message`)
		}
		rpc.in = msg

		if _, ok := ctx.proto.LookupMessage(msg.name); !ok {
			ctx.proto.AddMessage(msg)
		}
	}

	for resiter := oper.Responses().Responses(); resiter.Next(); {
		code, res := resiter.Item()

		switch code {
		case "200":
			s := res.Schema()
			if s == nil { // no content
				continue
			}

			var name string
			if s.Name() == "" {
				name = operID + "Response"
			}

			typ, err := compileMessage(ctx, s)
			if err != nil {
				return nil, errors.Wrapf(err, `failed to compile response message for %s (code = %s)`, oper.PathItem().Path(), code)
			}
			msg := typ.(*Message)
			if _, ok := ctx.proto.LookupMessage(name); !ok {
				msg.name = name
				ctx.proto.AddMessage(msg)
			}

			rpc.out = msg
			break
		}
	}
	rpc.path = oper.PathItem().Path()
	rpc.verb = oper.Verb()

	if rpc.in.Name() == "google.protobuf.Empty" || rpc.out.Name() == "google.protobuf.Empty" {
		ctx.proto.AddImport("google/protobuf/empty.proto")
	}

	return &rpc, nil
}

func format(ctx *genCtx, dst io.Writer, proto *Protobuf) error {
	fmt.Fprintf(dst, "syntax = \"proto3\";")
	fmt.Fprintf(dst, "\n\npackage %s;", proto.packageName)

	formatImports(dst, proto.imports)

	sort.Slice(proto.globalOptions, func(i, j int) bool {
		return proto.globalOptions[i].name < proto.globalOptions[j].name
	})

	if len(proto.globalOptions) > 0 {
		fmt.Fprintf(dst, "\n")
		for _, option := range proto.globalOptions {
			fmt.Fprintf(dst, "\noption %s = %s;", option.name, strconv.Quote(option.value))
		}
	}

	fmt.Fprintf(dst, "\n\n")
	formatMessages(ctx, dst, proto.messages)

	var serviceNames []string
	for name := range proto.services {
		serviceNames = append(serviceNames, name)
	}
	sort.Strings(serviceNames)

	for _, name := range serviceNames {
		service := proto.GetService(name)
		fmt.Fprintf(dst, "\n\nservice %s {", service.name)

		var rpcBuf bytes.Buffer
		formatRPCs(ctx, &rpcBuf, service.rpcs)
		copyIndent(dst, &rpcBuf)
		fmt.Fprintf(dst, "\n}")
	}

	return nil
}

func copyIndent(dst io.Writer, src io.Reader) {
	scanner := bufio.NewScanner(src)
	var n int
	for scanner.Scan() {
		n++
		txt := scanner.Text()
		if n == 1 && txt == "" {
			continue
		}

		fmt.Fprintf(dst, "\n  %s", txt)
	}
}

func formatFields(ctx *genCtx, dst io.Writer, fields []*Field) {
	for _, field := range fields {
		fmt.Fprintf(dst, "\n%s %s = %d;", field.typ.Name(), field.name, field.id)
	}
}

func formatRPCs(ctx *genCtx, dst io.Writer, rpcs []*RPC) {
	sort.Slice(rpcs, func(i, j int) bool {
		return rpcs[i].name < rpcs[j].name
	})

	var buf bytes.Buffer
	for _, rpc := range rpcs {
		if buf.Len() > 0 {
			fmt.Fprintf(&buf, "\n")
		}
		if desc := rpc.description; len(desc) > 0 {
			fmt.Fprintf(&buf, "\n// %s", desc)
		}
		fmt.Fprintf(&buf, "\nrpc %s(%s) returns (%s) {", rpc.name, rpc.in.Name(), rpc.out.Name())
		if ctx.annotate {
			var annotationBuf bytes.Buffer
			formatAnnotation(ctx, &annotationBuf, rpc)
			copyIndent(&buf, &annotationBuf)
		}
		fmt.Fprintf(&buf, "\n}")
	}

	buf.WriteTo(dst)
}

func formatAnnotation(ctx *genCtx, dst io.Writer, rpc *RPC) {
	fmt.Fprintf(dst, "\noption (google.api.http) = {")
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "\n%s: %s", strings.ToLower(rpc.verb), strconv.Quote(rpc.path))

	if m, ok := rpc.in.(*Message); ok {
		// There should be only one field with "in: body". This has been
		// checked with the parser's validation
		for _, f := range m.fields {
			if f.body {
				fmt.Fprintf(&buf, "\nbody: %s", strconv.Quote(f.name))
				break
			}
		}
	}
	copyIndent(dst, &buf)
	fmt.Fprintf(dst, "\n};")
}

func formatMessage(ctx *genCtx, dst io.Writer, msg *Message) {
	fmt.Fprintf(dst, "message %s {", msg.name)
	if len(msg.messages) > 0 {
		var buf bytes.Buffer

		var messageNames []string
		for name := range msg.messages {
			messageNames = append(messageNames, name)
		}
		sort.Strings(messageNames)

		for _, name := range messageNames {
			submsg := msg.messages[name]
			formatMessage(ctx, &buf, submsg)
		}
		copyIndent(dst, &buf)
	}

	if len(msg.fields) > 0 {
		var fieldsBuf bytes.Buffer
		formatFields(ctx, &fieldsBuf, msg.fields)
		copyIndent(dst, &fieldsBuf)
	}

	fmt.Fprintf(dst, "\n}")
}

func formatMessages(ctx *genCtx, dst io.Writer, messages map[string]*Message) {
	var messageNames []string
	for name := range messages {
		messageNames = append(messageNames, name)
	}
	sort.Strings(messageNames)

	for i, name := range messageNames {
		msg := messages[name]
		if i > 0 {
			fmt.Fprintf(dst, "\n\n")
		}
		formatMessage(ctx, dst, msg)
	}
}

func formatImports(dst io.Writer, imports map[string]struct{}) {
	var libs []string
	for lib := range imports {
		libs = append(libs, lib)
	}
	sort.Strings(libs)

	var buf bytes.Buffer
	for _, lib := range libs {
		if buf.Len() == 0 {
			fmt.Fprintf(&buf, "\n")
		}
		fmt.Fprintf(&buf, "\nimport %s;", strconv.Quote(lib))
	}
	buf.WriteTo(dst)
}
