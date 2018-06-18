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
	"sort"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/iancoleman/strcase"
	"github.com/lestrrat-go/openapi/internal/codegen/common"
	codegen "github.com/lestrrat-go/openapi/internal/codegen/golang"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

type SchemaLike interface {
	Type() openapi.PrimitiveType
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
	Enum() *openapi.InterfaceListIterator
	MultipleOf() float64
	Extension(string) (interface{}, bool)
	Extensions() *openapi.ExtensionsIterator
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
	resolver openapi.Resolver
}

func (r *Resolver) Resolve(ref string) (interface{}, error) {
	// TOOD
	if ref == "google/protobuf/timestamp.proto#/google.protobuf.Timestamp" {
		return nil, nonFatal{}
	}

	return r.resolver.Resolve(ref)
}

func Generate(ctx context.Context, spec openapi.OpenAPI, options ...Option) error {
	var dst io.Writer = os.Stdout
	var globalOptions []*globalOption
	for _, o := range options {
		switch o.Name() {
		case optkeyDestination:
			dst = o.Value().(io.Writer)
		case optkeyGlobalOption:
			globalOptions = append(globalOptions, o.Value().(*globalOption))
		}
	}

	resolver := &Resolver{resolver: openapi.NewResolver(spec)}
	/*
		if err := spec.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve references in openapi spec`)
		}
	*/

	c := &genCtx{
		Context:  ctx,
		resolver: resolver,
		root:     spec,
		dst:      dst,
		proto: &Protobuf{
			packageName:   "myapp",
			globalOptions: globalOptions,
			imports:       make(map[string]struct{}),
		},

		// types that are defined at the top level through references.
		types: make(map[string]Type),
	}

	return generate(c)
}

func (ctx *genCtx) RegisterMessage(path string, typ Type) {
	ctx.log("* Registering type %s (%s)", path, typ.Name())
	ctx.types[path] = typ
}

func (ctx *genCtx) LookupType(path string) (Type, bool) {
	typ, ok := ctx.types[path]
	return typ, ok
}

func grpcMethodName(oper openapi.Operation) string {
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

	if err := format(ctx.dst, ctx.proto); err != nil {
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

	for defiter := ctx.root.Definitions(); defiter.Next(); {
		name, thing := defiter.Item()

		var tmp openapi.Schema
		if err := common.RoundTripDecode(&tmp, thing, openapi.SchemaFromJSON); err != nil {
			return errors.Wrapf(err, `expected openapi.Schema for #/defnitions/%s, got %T`, name, thing)
		}

		ctx.log("* Compiling #/definitions/%s", name)
		typ, err := compileMessage(ctx, tmp)
		if err != nil {
			return errors.Wrapf(err, `failed to compile message #/definitions/%s`, name)
		}

		m, ok := typ.(*Message)
		if !ok {
			return errors.Errorf(`expected global definition to compile into a message, got %T`, typ)
		}

		m.name = codegen.ExportedName(name)
		ctx.proto.AddMessage(m)
		ctx.RegisterMessage("#/definitions/"+name, m)
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

/*
// compileMessage takes a schema an compiles it into a Message object.
// if the value is a built-in, then it's an error
func compileMessage(ctx *genCtx, name string, s openapi.Schema) (message *Message, err error) {
	if ref := s.Reference(); ref != "" {
		// if it's a reference, the chances are it's already registered in our type registry
		if m, ok := ctx.LookupType(ref); ok {
			if m, ok := m.(*Message); ok {
				return m, nil
			}
			return nil, errors.Errorf(`expected Message object, but %s resolved to %T`, ref, m)
		}

		log.Printf(" * Compiling %s", ref)
		defer func() {
			if err != nil {
				return
			}
			ctx.RegisterMessage(ref, message)
		}()

		v, err := ctx.resolver.Resolve(ref)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to resolve reference %s`, ref)
		}

		var tmp openapi.Schema
		if err := common.RoundTripDecode(&tmp, v, openapi.SchemaFromJSON); err != nil {
			return nil, errors.Errorf(`expected resolved object to be an openapi.Schema, but got %T`, v)
		}
		s = tmp
	}
	var m Message

	log.Printf("%#v", s)

	m.name = name

	// objects only, please
	if s.Type() != openapi.Object {
		return nil, errors.Errorf(`expected openapi.Object, got %s`, s.Type())
	}

	for piter := s.Properties(); piter.Next(); {
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
*/

func compileBuiltin(ctx *genCtx, s openapi.Schema) (Type, error) {
	switch s.Type() {
	case openapi.Boolean:
		return Builtin("bool"), nil
	case openapi.String:
		return Builtin(s.Type()), nil
	case openapi.Integer:
		switch s.Format() {
		case "int64":
			return Builtin("int64"), nil
		default:
			return Builtin("int"), nil
		}
	default:
		return nil, errors.Errorf(`unknown builtin %s`, s.Type())
	}
}

// compileMessage takes a schema and compiles it into a Message.
// Messages normally need to be an object-like type, but in case
// we have an array, we will actually wrap the elements in an object.
func compileMessage(ctx *genCtx, schema openapi.Schema) (Type, error) {
	done := ctx.Start("* Compiling Message")
	defer done()

	// it better be an object
	switch schema.Type() {
	case openapi.Object:
	case openapi.Array:
		// special case
		items, err := openapi.NewSchema().
			Type(openapi.Array).
			Items(schema.Items()).
			Do()
		if err != nil {
			return nil, errors.Wrap(err, `failed to create new schema for items`)
		}

		schema, err = openapi.NewSchema().
			Type(openapi.Object).
			Property("items", items).
			Do()
		if err != nil {
			return nil, errors.Wrap(err, `failed to create new schema wrapping array`)
		}
	default:
		return nil, errors.Errorf(`compileMessage: expected type "object", got %s`, schema.Type())
	}

	var m Message
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

// compileType takes something that looks like a schema, and creates an
// appropriate type for it.
func compileType(ctx *genCtx, src openapi.SchemaConverter) (result Type, err error) {
	schema, err := src.ConvertToSchema()
	if err != nil {
		return nil, errors.Wrapf(err, `failed to extract schema out of %T`, src)
	}

	done := ctx.Start("* Compiling Type %s", schema.Type())
	defer done()

	if ref := schema.Reference(); ref != "" {
		typ, ok := ctx.LookupType(ref)
		if ok {
			ctx.log("* Reference %s already compileted to %s", ref, typ.Name())
			return typ, nil
		}

		thing, err := ctx.resolver.Resolve(ref)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to resolve reference %s`, ref)
		}

		// Are we sure this should always resolve to a schema?
		var tmp openapi.Schema
		if err := common.RoundTripDecode(&tmp, thing, openapi.SchemaFromJSON); err != nil {
			return nil, errors.Errorf(`expected reference %s to resolve to a Schema`, ref)
		}

		schema = tmp

		if strings.HasPrefix(ref, "#/definitions/") {
			// If this refers to a global definition but the previous LookupType
			// failed, this can only mean that we encountered a global definition
			// that refers to another global definition.
			// In this case we need to register this message
			defer func() {
				if err != nil {
					return
				}

				if m, ok := result.(*Message); ok {
					m.name = codegen.ExportedName(strings.TrimPrefix(ref, "#/definitions/"))
					ctx.proto.AddMessage(m)
					ctx.RegisterMessage(ref, m)
				}
			}()
		}
	}

	switch schema.Type() {
	case openapi.String, openapi.Number, openapi.Boolean:
		return compileBuiltin(ctx, schema)
	case openapi.Object:
		return compileMessage(ctx, schema)
	case openapi.Array:
		ctx.log("* Compiling Array Element")
		typ, err := compileType(ctx, schema.Items())
		if err != nil {
			return nil, errors.Wrap(err, `failed to compile array element`)
		}
		return &Array{element: typ}, nil
	default:
		return nil, errors.Errorf(`compileType: unsupported schema.type = %s`, schema.Type())
	}
}

/*

func compileArrayElement(ctx *genCtx, s SchemaLike) (Type, error) {
	if ref := s.Reference(); ref != "" {
		// if ref looks like it's from #/definitions, cheat
		if strings.HasPrefix(ref, "#/definitions/") {
			return &Message{name: strings.TrimPrefix(ref, "#/definitions/")}, nil
		}
	}

	typ, err := compileType(ctx, s)


	return grpcType(ctx, s, "dummy")

//	return nil, errors.Errorf(`unimplemented %T`, s)
}
*/

/*
func grpcType(ctx *genCtx, s SchemaLike, name string) (string, bool, error) {
	var typ string
	var repeated bool
	// If this is a reference, resolve it
	if ref := s.Reference(); ref != "" {
		// if it's a reference, the chances are it's already registered in our type registry
		if m, ok := lookupMessage(ctx, ref); ok {
			return m.name, false, nil
		}

		v, err := ctx.resolver.Resolve(ref)
		if err != nil {
			return "", false, errors.Wrap(err, `failed to resolve referece`)
		}

		var tmp openapi.Schema
		if err := common.RoundTripDecode(&tmp, v, openapi.SchemaFromJSON); err != nil {
			return "", false, errors.Errorf(`expected reference %s to resolve to a Schema`, ref)
		}

		msg, err := compileMessage(ctx, codegen.ExportedName(strings.TrimPrefix(ref, "#/definitions/")), tmp)
		if err != nil {
			return "", false, errors.Wrapf(err, `failed to compile message for %s`, ref)
		}
		ctx.proto.AddMessage(msg)

		return msg.name, false, nil
	}

	if raw, ok := s.Extension("x-proto-type"); ok {
		// better be a map of strings (disguised as interface{} and map[string]interface{})
		proto, ok := raw.(map[string]interface{})
		if !ok {
			return "", false, errors.Errorf(`expected x-proto-type to be a map`)
		}
		rawName := proto["name"]
		name, ok := rawName.(string)
		if !ok {
			return "", false, errors.Errorf(`expected x-proto-type.name to be a string`)
		}
		rawImport, ok := proto["import"]
		lib, ok := rawImport.(string)
		if !ok {
			return "", false, errors.Errorf(`expected x-proto-type.import to be a string`)
		}

		typ = name
		ctx.proto.AddImport(lib)
	}

	if typ == "" {
		switch s.Type() {
		case openapi.Object:
			sc, ok := s.(openapi.SchemaConverter)
			if !ok {
				return "", false, errors.Errorf(`%T does not implement openapi.SchemaConverter`, s)
			}
			converted, err := sc.ConvertToSchema()
			if err != nil {
				return "", false, errors.Wrapf(err, `failed to convert %T to schema`, s)
			}
			msg, err := compileMessage(ctx, name, converted)
			if err != nil {
				return "", false, errors.Wrap(err, `failed to compile message`)
			}
			return msg.name, false, nil
		case openapi.Array:
			var items SchemaLike
			if tmp, ok := s.(openapi.Schema); ok {
				items = tmp.Items()
			} else if tmp, ok := s.(openapi.Items); ok {
				items = tmp.Items()
			}

			msg, err := com(ctx, items)
			if err != nil {
				return "", false, errors.Wrap(err, `failed to compile array items`)
			}
			typ = msg.name
			repeated = true
		case openapi.String:
			typ = "string"
		case openapi.Integer:
			typ = "int32" // TODO consider format
		case openapi.Number:
			typ = "float32"
		case openapi.Boolean:
			typ = "boolean"
		default:
			return "", false, errors.Errorf(`unsupported field type %s`, s.Type())
		}
	}

	return typ, repeated, nil
}
*/

func compileField(ctx *genCtx, name string, s openapi.Schema) (*Field, error) {
	done := ctx.Start("* Compiling Field %s", name)
	defer done()

	typ, err := compileType(ctx, s)
	if err != nil {
		return nil, errors.Wrap(err, `failed to compile field`)
	}

	return &Field{
		id:   1,
		name: name,
		typ:  typ,
	}, nil
}

func compileRPCParameters(ctx *genCtx, name string, iter *openapi.ParameterListIterator) (*Message, error) {
	done := ctx.Start("* Compiling RPC parameters for %s", name)
	defer done()

	var m Message
	var id = 1
	var inBody bool
	for iter.Next() {
		param := iter.Item()

		ctx.log("* Compiling parameter %s", param.Name())

		var typ Type
		var err error
		if param.In() == openapi.InBody {
			typ, err = compileType(ctx, param.Schema())
			inBody = true
		} else {
			typ, err = compileType(ctx, param)
		}
		if err != nil {
			return nil, errors.Wrap(err, `failed to deduce gRPC type`)
		}

		m.fields = append(m.fields, &Field{
			id:   id,
			name: param.Name(),
			typ:  typ,
		})

		id++
	}

	if len(m.fields) == 1 && inBody {
		ctx.log("* Parameter single body parameter only. Replacing entire message with body message type")
		// this element *MUST* be a message
		body := m.fields[0].typ.(*Message)
		m = *body
	}
	m.name = name

	return &m, nil
}

func compileRPC(ctx *genCtx, oper openapi.Operation) (*RPC, error) {
	var rpc RPC
	rpc.name = grpcMethodName(oper)
	rpc.in = "google.protobuf.Empty"
	rpc.out = "google.protobuf.Empty"
	if desc := oper.Description(); len(desc) > 0 {
		rpc.description = desc
	}

	done := ctx.Start("* Compiling %s", rpc.name)
	defer done()

	paramiter := oper.Parameters()
	if paramiter.Next() {
		msg, err := compileRPCParameters(ctx, rpc.name+"Request", paramiter)
		if err != nil {
			return nil, errors.Wrap(err, `failed to compile request parameters into message`)
		}
		rpc.in = msg.name
		ctx.proto.AddMessage(msg)
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
				name = rpc.name + "Response"
			}

			typ, err := compileMessage(ctx, s)
			if err != nil {
				return nil, errors.Wrapf(err, `failed to compile response message for %s (code = %s)`, oper.PathItem().Path(), code)
			}
			msg := typ.(*Message)
			msg.name = name
			ctx.proto.AddMessage(msg)

			rpc.out = msg.Name()
			break
		}
	}
	rpc.path = oper.PathItem().Path()

	if rpc.in == "google.protobuf.Empty" || rpc.out == "google.protobuf.Empty" {
		ctx.proto.AddImport("google/protobuf/empty.proto")
	}

	return &rpc, nil
}

func format(dst io.Writer, proto *Protobuf) error {
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

	formatMessages(dst, proto.messages)

	var serviceNames []string
	for name := range proto.services {
		serviceNames = append(serviceNames, name)
	}
	sort.Strings(serviceNames)

	for _, name := range serviceNames {
		service := proto.GetService(name)
		fmt.Fprintf(dst, "\n\nservice %s {", service.name)

		var rpcBuf bytes.Buffer
		formatRPCs(&rpcBuf, service.rpcs)
		copyIndent(dst, &rpcBuf)
		fmt.Fprintf(dst, "\n}")
	}

	return nil
}

func copyIndent(dst io.Writer, src io.Reader) {
	scanner := bufio.NewScanner(src)
	var n int
	for scanner.Scan() {
		if n > 0 {
			fmt.Fprintf(dst, "\n")
		}
		fmt.Fprintf(dst, "    %s", scanner.Text())
		n++
	}
}

func formatFields(dst io.Writer, fields []*Field) {
	for _, field := range fields {
		fmt.Fprintf(dst, "\n")
		fmt.Fprintf(dst, "%s %s = %d;", field.typ.Name(), field.name, field.id)
	}
}

func formatRPCs(dst io.Writer, rpcs []*RPC) {
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
		fmt.Fprintf(&buf, "\nrpc %s(%s) returns (%s) {", rpc.name, rpc.in, rpc.out)
		fmt.Fprintf(&buf, "\n}")
	}

	buf.WriteTo(dst)
}

func formatMessages(dst io.Writer, messages map[string]*Message) {
	var messageNames []string
	for name := range messages {
		messageNames = append(messageNames, name)
	}
	sort.Strings(messageNames)

	for _, name := range messageNames {
		msg := messages[name]

		fmt.Fprintf(dst, "\n\nmessage %s {", msg.name)

		var fieldsBuf bytes.Buffer
		formatFields(&fieldsBuf, msg.fields)
		copyIndent(dst, &fieldsBuf)

		fmt.Fprintf(dst, "\n}")
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
		fmt.Fprintf(&buf, "\nimport %s", strconv.Quote(lib))
	}
	buf.WriteTo(dst)
}
