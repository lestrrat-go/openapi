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
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

type SchemaLike interface {
	Type() openapi.PrimitiveType
	Format() string
	DefaultValue() interface{}
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

func New() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(ctx context.Context, spec openapi.OpenAPI, options ...Option) error {
	var dst io.Writer = os.Stdout
	for _, o := range options {
		switch o.Name() {
		case optkeyDestination:
			dst = o.Value().(io.Writer)
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
		messages: make(map[string]*Message),
	}

	return generate(c)
}

func registerMessage(ctx *genCtx, name string, message *Message) {
	log.Printf(" * Registering message %s (%s)", name, message.name)
	ctx.messages[name] = message
}

func lookupMessage(ctx *genCtx, name string) (*Message, bool) {
	message, ok := ctx.messages[name]
	return message, ok
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
	ctx.proto = &Protobuf{
		packageName: "myapp",
		imports:     make(map[string]struct{}),
	}

	if err := compileGlobalDefinitions(ctx); err != nil {
		return errors.Wrap(err, `failed to compile definitions`)
	}

	if err := compileMethods(ctx); err != nil {
		return errors.Wrap(err, `failed to compile methods`)
	}

	return nil
}

func compileGlobalDefinitions(ctx *genCtx) error {
	for defiter := ctx.root.Definitions(); defiter.Next(); {
		name, schema := defiter.Item()

		log.Printf(" * Compiling #/definitions/%s", name)
		m, err := compileMessage(ctx, name, schema)
		if err != nil {
			return errors.Wrapf(err, `failed to compile message #/definitions/%s`, name)
		}

		ctx.proto.AddMessage(m)
		registerMessage(ctx, "#/definitions/"+name, m)
	}
	return nil
}

func compileMethods(ctx *genCtx) error {
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

func compileMessage(ctx *genCtx, name string, s openapi.Schema) (message *Message, err error) {
	if ref := s.Reference(); ref != "" {
		// if it's a reference, the chances are it's already registered in our type registry
		if m, ok := lookupMessage(ctx, ref); ok {
			return m, nil
		}
		log.Printf(" * Compiling %s", ref)
		defer func() {
			if err != nil {
				return
			}
			registerMessage(ctx, ref, message)
		}()

		v, err := ctx.resolver.Resolve(ref)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to resolve reference %s`, ref)
		}
		asserted, ok := v.(openapi.Schema)
		if !ok {
			return nil, errors.Errorf(`expected resolved object to be an openapi.Schema, but got %T`, v)
		}
		s = asserted
	}
	var m Message

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

func compileArrayElement(ctx *genCtx, s SchemaLike) (*Message, error) {
	if ref := s.Reference(); ref != "" {
		// if ref looks like it's from #/definitions, cheat
		if strings.HasPrefix(ref, "#/definitions/") {
			return &Message{name: strings.TrimPrefix(ref, "#/definitions/")}, nil
		}
	}

	return nil, errors.New(`unimplemented`)
}

func grpcType(ctx *genCtx, s SchemaLike) (string, bool, error) {
	var typ string
	var repeated bool
	// If this is a reference, resolve it
	if ref := s.Reference(); ref != "" {
		v, err := ctx.resolver.Resolve(ref)
		if err != nil {
			return "", false, errors.Wrap(err, `failed to resolve referece`)
		}

		tmp, ok := v.(openapi.Schema)
		if !ok {
			return "", false, errors.Errorf(`expected reference %s to resolve to a Schema, got %T`, ref, v)
		}
		s = tmp
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
		case openapi.Array:
			var items SchemaLike
			if tmp, ok := s.(openapi.Schema); ok {
				items = tmp.Items()
			} else if tmp, ok := s.(openapi.Items); ok {
				items = tmp.Items()
			}

			msg, err := compileArrayElement(ctx, items)
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
		default:
			return "", false, errors.Errorf(`unsupported field type %s`, s.Type())
		}
	}

	return typ, repeated, nil
}

func compileField(ctx *genCtx, name string, s openapi.Schema) (*Field, error) {
	typ, repeated, err := grpcType(ctx, s)
	if err != nil {
		return nil, errors.Wrap(err, `failed to deduce gRPC type`)
	}

	return &Field{
		id:       1,
		name:     name,
		typ:      typ,
		repeated: repeated,
	}, nil
}

func compileRPCParameters(ctx *genCtx, name string, iter *openapi.ParameterListIterator) (*Message, error) {
	log.Printf(" * Compiling RPC parameters for %s", name)
	var m Message
	var id = 1
	for iter.Next() {
		param := iter.Item()

		var typ string
		var repeated bool
		var err error
		if param.In() == openapi.InBody {
			typ, repeated, err = grpcType(ctx, param.Schema())
		} else {
			typ, repeated, err = grpcType(ctx, param)
		}
		if err != nil {
			return nil, errors.Wrap(err, `failed to deduce gRPC type`)
		}
		m.fields = append(m.fields, &Field{
			id:       id,
			name:     param.Name(),
			typ:      typ,
			repeated: repeated,
		})

		id++
	}
	m.name = name

	return &m, nil
}

func compileRPC(ctx *genCtx, oper openapi.Operation) (*RPC, error) {
	var rpc RPC
	rpc.name = grpcMethodName(oper)
	rpc.in = "google.protobuf.Empty"
	rpc.out = "google.protobuf.Empty"

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
			var name string
			s := res.Schema()
			if s.Name() == "" {
				name = rpc.name + "Response"
			}
			msg, err := compileMessage(ctx, name, s)
			if err != nil {
				return nil, errors.Wrapf(err, `failed to compile response message for %s (code = %s)`, oper.PathItem().Path(), code)
			}
			ctx.proto.AddMessage(msg)

			rpc.out = msg.name
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
		if field.repeated {
			fmt.Fprintf(dst, "repeated ")
		}
		fmt.Fprintf(dst, "%s %s = %d;", field.typ, field.name, field.id)
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
