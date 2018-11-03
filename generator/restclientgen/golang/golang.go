package golang

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/lestrrat-go/openapi/generator/restclientgen/compiler"
	codegen "github.com/lestrrat-go/openapi/internal/codegen/golang"
	"github.com/lestrrat-go/openapi/internal/stringutil"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

func goType(typ compiler.Type) string {
	if name := typ.Name(); name != "" {
		return name
	}

	switch typ := typ.(type) {
	case *compiler.Array:
		var prefix = "[]"
		if !isBuiltinType(typ.Elem()) {
			prefix = prefix + "*"
		}
		return prefix +typ.Elem()
	default:
		return "couldNotDeduce"
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

	client, err := compiler.Compile(spec, defaultServiceName)
	if err != nil {
		return errors.Wrap(err, `failed to compile spec`)
	}

	ctx := Context{
		dir:         dir,
		packageName: packageName,
		exportNew:   exportNew,
		client:      client,
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

	if err := writeTypesFile(&ctx); err != nil {
		return errors.Wrap(err, `failed to write options file`)
	}

	return nil
}

func writeTypesFile(ctx *Context) error {
	fn := filepath.Join(ctx.dir, "types_gen.go")
	log.Printf("Generating %s", fn)

	var buf bytes.Buffer
	var dst io.Writer = &buf
	codegen.WritePreamble(dst, ctx.packageName)

	var typDefs []compiler.TypeDefinition
	for _, typ := range ctx.client.Definitions() {
		typDefs = append(typDefs, typ)
	}
	sort.Slice(typDefs, func(i, j int) bool {
		return typDefs[i].Type.Name() < typDefs[j].Type.Name()
	})

	for _, typDef := range typDefs {
		typ := typDef.Type
		log.Printf("   * Generating definition for %s", typ.Name())
		switch t := typ.(type) {
		case *compiler.Array:
			fmt.Fprintf(dst, "\n\ntype %s []%s", t.Name(), t.Elem())
		case *compiler.Struct:
			fmt.Fprintf(dst, "\n\n// %s represents the data structure defined in %s", typ.Name(), typDef.Context)
			fmt.Fprintf(dst, "\ntype %s struct {", t.Name())
			for _, field := range t.Fields() {
				hints := field.Hints()
				fmt.Fprintf(dst, "\n%s %s `%s`", hints.GoName, goType(field.Type()), hints.GoTag)
			}
			fmt.Fprintf(dst, "\n}")
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
	fmt.Fprintf(dst, "\noptkeyJWT = `jwt`")
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

	fmt.Fprintf(dst, "\n\n// WithJWT is used to specify a signed token to be")
	fmt.Fprintf(dst, "\n// included in the Authorization header.")
	fmt.Fprintf(dst, "\nfunc WithJWT(s string) CallOption {")
	fmt.Fprintf(dst, "\nreturn newOption(optkeyJWT, s)")
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
	services := ctx.client.Services()
	for _, name := range ctx.client.ServiceNames() {
		// Remove the "service" from the filename
		fn := strings.TrimSuffix(name, "Service")
		fn = filepath.Join(ctx.dir, stringutil.Snake(fn)+"_service_gen.go")
		log.Printf("Generating %s", fn)

		var buf bytes.Buffer
		if err := formatService(ctx, &buf, services[name]); err != nil {
			return errors.Wrap(err, `failed to format service code`)
		}

		if err := codegen.WriteFormattedToFile(fn, buf.Bytes()); err != nil {
			codegen.DumpCode(os.Stdout, bytes.NewReader(buf.Bytes()))
			return errors.Wrapf(err, `failed to write to %s`, fn)
		}
	}
	return nil
}

func formatClient(ctx *Context, dst io.Writer, cl *compiler.ClientDefinition) error {
	codegen.WritePreamble(dst, ctx.packageName)
	codegen.WriteImports(dst, "bytes", "net/http", "github.com/pkg/errors")
	fmt.Fprintf(dst, "\n\n")

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
	for _, name := range ctx.client.ServiceNames() {
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
	for _, name := range ctx.client.ServiceNames() {
		fmt.Fprintf(dst, "\n%s: &%s{httpCl: cl, server: server},", codegen.UnexportedName(name), name)
	}
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	for _, name := range ctx.client.ServiceNames() {
		fmt.Fprintf(dst, "\nfunc (c *Client) %s() *%s {", name, name)
		fmt.Fprintf(dst, "\nreturn c.%s", codegen.UnexportedName(name))
		fmt.Fprintf(dst, "\n}")
	}

	return nil
}

func formatService(ctx *Context, dst io.Writer, svc *compiler.Service) error {
	log.Printf(" * Generating Service %s", svc.Name())
	codegen.WritePreamble(dst, ctx.packageName)

	// TODO: be smarter as to which libraries to include
	// for exmaple, oauth stuff
	codegen.WriteImports(dst, "context", "encoding/json", "fmt", "io", "mime", "net/http", "net/http/httputil", "net/url", "strings", "strconv", "github.com/pkg/errors", "github.com/lestrrat-go/urlenc", "github.com/lestrrat-go/jwx/jwt", "golang.org/x/oauth2")

	fmt.Fprintf(dst, "\n\ntype %s struct {", svc.Name())
	fmt.Fprintf(dst, "\nhttpCl *http.Client")
	fmt.Fprintf(dst, "\nserver string")
	fmt.Fprintf(dst, "\n}")

	for _, call := range svc.Calls() {
		if err := formatCall(dst, svc.Name(), call); err != nil {
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
func formatCall(dst io.Writer, svcName string, call *compiler.Call) error {
	log.Printf("   * Generating Call object %s", call.Name())
	fmt.Fprintf(dst, "\n\ntype %s struct {", call.Name())
	fmt.Fprintf(dst, "\nhttpCl *http.Client")
	fmt.Fprintf(dst, "\nserver string")
	fmt.Fprintf(dst, "\nmarshalers map[string]Marshaler")
	if call.Body() != nil {
		fmt.Fprintf(dst, "\nbody %s", call.Body().Name())
	}
	if call.Query() != nil {
		fmt.Fprintf(dst, "\nquery %s", call.Query().Name())
	}
	if call.Header() != nil {
		fmt.Fprintf(dst, "\nheader %s", call.Header().Name())
	}
	if call.Path() != nil {
		fmt.Fprintf(dst, "\npath %s", call.Path().Name())
	}

	fmt.Fprintf(dst, "\n}")

	log.Printf("      * Generating constructor")
	fmt.Fprintf(dst, "\n\nfunc (svc *%s) %s(", svcName, call.Method())
	for i, field := range call.Requireds() {
		fmt.Fprintf(dst, "%s %s", codegen.UnexportedName(field.Hints().GoName), goType(field.Type()))
		if i < len(call.Requireds())-1 {
			fmt.Fprintf(dst, ", ")
		}
	}
	fmt.Fprintf(dst, ") *%s {", call.Name())

	fmt.Fprintf(dst, "\nvar call %s", call.Name())
	fmt.Fprintf(dst, "\ncall.httpCl = svc.httpCl")
	fmt.Fprintf(dst, "\ncall.server = svc.server")
	fmt.Fprintf(dst, "\ncall.marshalers = map[string]Marshaler{")
	fmt.Fprintf(dst, "\n`application/json`: MarshalFunc(json.Marshal),")
	fmt.Fprintf(dst, "\n`application/x-www-form-urlencoded`: MarshalFunc(urlenc.Marshal),")
	fmt.Fprintf(dst, "\n}")
	for _, field := range call.Requireds() {
		fmt.Fprintf(dst, "\ncall.%s.%s = %s", field.ContainerName(), field.Hints().GoName, codegen.UnexportedName(field.Hints().GoName))
	}
	fmt.Fprintf(dst, "\nreturn &call")
	fmt.Fprintf(dst, "\n}")

	for _, optional := range call.Optionals() {
		log.Printf("      * Generating optional method for %s", codegen.ExportedName(optional.Name()))
		if gotyp := goType(optional.Type()); strings.HasPrefix(gotyp, "[]") {
			fmt.Fprintf(dst, "\n\nfunc (call *%s) %s(v ...%s) *%s {", call.Name(), codegen.ExportedName(optional.Name()), strings.TrimPrefix(gotyp, "[]"), call.Name())
			fmt.Fprintf(dst, "\ncall.%[1]s.%[2]s = append(call.%[1]s.%[2]s, v...)", optional.ContainerName(), codegen.ExportedName(optional.Name()))
			fmt.Fprintf(dst, "\nreturn call")
			fmt.Fprintf(dst, "\n}")
		} else {
			fmt.Fprintf(dst, "\n\nfunc (call *%s) %s(v %s) *%s {", call.Name(), stringutil.Camel(optional.Name()), optional.Type().Name(), call.Name())
			fmt.Fprintf(dst, "\ncall.%s.%s = v", optional.ContainerName(), codegen.ExportedName(optional.Name()))
			fmt.Fprintf(dst, "\nreturn call")
			fmt.Fprintf(dst, "\n}")
		}
	}

	/*
		fmt.Fprintf(dst, "\n\nfunc (call %s) AsMap() map[string]interface{} {", call.Name())
		fmt.Fprintf(dst, "\nm := make(map[string]interface{})")
		for _, param := range append(call.optionals, call.requireds...) {
			fmt.Fprintf(dst, "\nm[%#v] = call.%s", param.name, param.goName)
		}
		fmt.Fprintf(dst, "\nreturn m")
		fmt.Fprintf(dst, "\n}")
	*/

	fmt.Fprintf(dst, "\n\nfunc (call *%[1]s) Marshaler(mime string, m Marshaler) *%[1]s {", call.Name())
	fmt.Fprintf(dst, "\ncall.marshalers[mime] = m")
	fmt.Fprintf(dst, "\nreturn call")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (call *%s) Do(ctx context.Context, options ...CallOption) (Response, error) {", call.Name())

	var hasJWT bool
	if sslist := call.SecuritySettings(); len(sslist) > 0 {
		for _, settings := range sslist {
			switch settings.Definition().Type() {
			case "oauth2":
				// Require Authorization header
				hasJWT = true
			}
		}
	}

	fmt.Fprintf(dst, "\n\nvar debugOut io.Writer")
	if hasJWT {
		fmt.Fprintf(dst, "\nvar signedJWT string")
	}

	if call.Body() != nil {
		fmt.Fprintf(dst, "\ncontentType := %#v", call.DefaultConsumes())
	}

	fmt.Fprintf(dst, "\nfor _, option := range options {")
	fmt.Fprintf(dst, "\nswitch option.Name() {")
	fmt.Fprintf(dst, "\ncase optkeyDebugDump:")
	fmt.Fprintf(dst, "\ndebugOut = option.Value().(io.Writer)")
	if call.Body() != nil {
		fmt.Fprintf(dst, "\ncase optkeyRequestContentType:")
		fmt.Fprintf(dst, "\ncontentType = option.Value().(string)")
	}
	if hasJWT {
		fmt.Fprintf(dst, "\ncase optkeyJWT:")
		fmt.Fprintf(dst, "\nsignedJWT = option.Value().(string)")
	}
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\npath := %s", strconv.Quote(call.RequestPath()))
	if call.Path() != nil {
		for _, field := range call.Path().(*compiler.Struct).Fields() {
			fmt.Fprintf(dst, "\npath = strings.Replace(path, `{%s}`, ", field.Name())
			switch goType(field.Type()) {
			case "int64":
				fmt.Fprintf(dst, "strconv.FormatInt(call.path.%s, 10)", codegen.ExportedName(field.Name()))
			default:
				fmt.Fprintf(dst, "call.path.%s", codegen.ExportedName(field.Name()))
			}
			fmt.Fprintf(dst, ", -1)")
		}
	}

	if call.Query() != nil {
		fmt.Fprintf(dst, "\nv := url.Values{}")
		for _, query := range call.Query().(*compiler.Struct).Fields() {
			// XXX This needs to be more robust
			gotyp := goType(query.Type())
			if gotyp == "[]string" {
				fmt.Fprintf(dst, "\nfor _, param := range call.query.%s {", query.Hints().GoName)
				fmt.Fprintf(dst, "\nv.Add(%s, param)", strconv.Quote(query.Name()))
				fmt.Fprintf(dst, "\n}")
			} else {
				fmt.Fprintf(dst, "\nv.Add(%s, ", strconv.Quote(query.Name()))
				switch gotyp {
				case "int64":
					fmt.Fprintf(dst, "strconv.FormatInt(call.query.%s, 10)", codegen.ExportedName(query.Name()))
				case "bool":
					fmt.Fprintf(dst, "strconv.FormatBool(call.query.%s)", codegen.ExportedName(query.Name()))
				default:
					fmt.Fprintf(dst, "call.query.%s", query.Hints().GoName)
				}
				fmt.Fprintf(dst, ")")
			}
		}

		fmt.Fprintf(dst, "\npath = call.server + path + `?` + v.Encode()")
	}

	var body = "nil"
	if call.Body() != nil {
		body = "body"

		fmt.Fprintf(dst, "\nmtype, _, err := mime.ParseMediaType(contentType)")
		fmt.Fprintf(dst, "\nif err != nil {")
		fmt.Fprintf(dst, "\nreturn nil, errors.Wrapf(err, `failed to parse request content type %%s`, contentType)")
		fmt.Fprintf(dst, "\n}")

		if call.Body() != nil {
			fmt.Fprintf(dst, "\n\nbody, err := encodeCallPayload(call.marshalers, mtype, call.body)")
		} else {
			return errors.New(`can't proceed when call.Body() == nil`)
		}
		fmt.Fprintf(dst, "\nif err != nil {")
		fmt.Fprintf(dst, "\nreturn nil, errors.Wrapf(err, `failed to marshal request payload as %%s`, mtype)")
		fmt.Fprintf(dst, "\n}")
	}

	fmt.Fprintf(dst, "\n\nreq, err := http.NewRequest(%s, path, %s)", strconv.Quote(call.Verb()), body)
	fmt.Fprintf(dst, "\nif err != nil {")
	fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `failed to create request`)")
	fmt.Fprintf(dst, "\n}")

	if call.Body() != nil {
		fmt.Fprintf(dst, "\nreq.Header.Set(`Content-Type`, contentType)")
		fmt.Fprintf(dst, "\nreq.Header.Set(`Content-Length`, strconv.Itoa(body.Len()))")
	}

	if hasJWT {
		fmt.Fprintf(dst, "\nif len(signedJWT) > 0 {")
		fmt.Fprintf(dst, "\nreq.Header.Set(`Authorization`, `Bearer ` + signedJWT)")
		fmt.Fprintf(dst, "\nif debugOut != nil {")
		fmt.Fprintf(dst, "\ntoken, err := jwt.Parse(strings.NewReader(signedJWT))")
		fmt.Fprintf(dst, "\nif err != nil {")
		fmt.Fprintf(dst, "\nfmt.Fprintf(debugOut, `failed to decode JWT token: %%s`, err)")
		fmt.Fprintf(dst, "\n} else {")
		fmt.Fprintf(dst, "\nencoded, err := json.MarshalIndent(token, ``, `  `)")
		fmt.Fprintf(dst, "\nif err != nil {")
		fmt.Fprintf(dst, "\nfmt.Fprintf(debugOut, `failed to marshal token back into JSON: %%s`, err)")
		fmt.Fprintf(dst, "\n} else {")
		fmt.Fprintf(dst, "\nfmt.Fprintf(debugOut, \"=== ID Token ===\\n%%s\\n===============\", encoded)")
		fmt.Fprintf(dst, "\n}") // end if err != nil
		fmt.Fprintf(dst, "\n}") // end if err != nil
		fmt.Fprintf(dst, "\n}") // end if debugOut != nil
		fmt.Fprintf(dst, "\n}") // end if len(signedJWT) > 0
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
	for _, response := range call.Responses() {
		fmt.Fprintf(dst, "\ncase %s:", response.Code())
		// if typ == "" ignore response. we ain't expecting none.
		if response.Type() == "" {
			fmt.Fprintf(dst, "\n// no response body expected")
		} else {
			fmt.Fprintf(dst, "\nvar resdata %s", response.Type())
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
