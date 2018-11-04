package es6flow

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/lestrrat-go/openapi/generator/restclientgen/compiler"
	codegen "github.com/lestrrat-go/openapi/internal/codegen/es6"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

func esType(typ compiler.Type) string {
	if name := typ.Name(); name != "" {
		return name
	}

	switch typ := typ.(type) {
	case *compiler.Array:
		return "Array<" + typ.Elem() + ">"
	default:
		return "couldNotDeduce"
	}
}

func Generate(spec openapi.Swagger, options ...Option) error {
	var clientName string
	var dir string
	var defaultServiceName string
	for _, option := range options {
		switch option.Name() {
		case optkeyClientName:
			clientName = option.Value().(string)
		case optkeyDefaultServiceName:
			defaultServiceName = option.Value().(string)
		case optkeyDirectory:
			dir = option.Value().(string)
		}
	}

	if dir == "" {
		dir = "restclient"
	}

	if defaultServiceName == "" {
		defaultServiceName = "restclient"
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
		compiling:          make(map[string]struct{}),
		clientName:         clientName,
		dir:                dir,
		defaultServiceName: defaultServiceName,
		resolver:           openapi.NewResolver(spec),
		root:               spec,
		client:             client,
		types:              make(map[string]compiler.TypeDefinition),
	}

	// declare types
	if err := writeResponseFile(&ctx); err != nil {
		return errors.Wrap(err, `failed to write response file`)
	}

	if err := writeTypesFile(&ctx); err != nil {
		return errors.Wrap(err, `failed to write options file`)
	}

	if err := writeTypesClientOptions(&ctx); err != nil {
		return errors.Wrap(err, `failed to write client options file`)
	}

	if err := writeServiceFiles(&ctx); err != nil {
		return errors.Wrapf(err, `failed to generate services`)
	}

	if err := writeClientFile(&ctx); err != nil {
		return errors.Wrap(err, `failed to generate client code`)
	}

	if err := runFlow(&ctx); err != nil {
		return errors.Wrap(err, `failed to run "npm run flow"`)
	}

	return nil
}

func runFlow(ctx *Context) error {
	cmd := exec.Command("npm", "run", "flow")
	cmd.Dir = ctx.dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, `failed to execute "flow" tool`)
	}

	return nil
}

func writeResponseFile(ctx *Context) error {
	fn := filepath.Join(ctx.dir, "types", "response.js")
	log.Printf("Generating %s", fn)

	var buf bytes.Buffer
	var dst io.Writer = &buf
	codegen.WritePreamble(dst, ctx.packageName)

	fmt.Fprintf(dst, "\n\nexport default class RESTResponse {")
	fmt.Fprintf(dst, "\ncode :number")
	fmt.Fprintf(dst, "\ndata :any")
	fmt.Fprintf(dst, "\nconstructor(code: number, data: any) {")
	fmt.Fprintf(dst, "\nthis.code = code;")
	fmt.Fprintf(dst, "\nthis.data = data;")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	if err := codegen.WriteFormattedToFile(fn, buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, bytes.NewReader(buf.Bytes()))
		return errors.Wrapf(err, `failed to write to %s`, fn)
	}
	return nil
}

func writeTypesClientOptions(ctx *Context) error {
	fn := filepath.Join(ctx.dir, "types", "client_options.js")
	log.Printf("Generating %s", fn)

	var buf bytes.Buffer
	var dst io.Writer = &buf
	codegen.WritePreamble(dst, ctx.packageName)

	fmt.Fprintf(dst, "\n\nexport type ClientOptions = {")
	fmt.Fprintf(dst, "\nmodifyOptions: (RequestOptions) => RequestOptions;")
	fmt.Fprintf(dst, "\nheaders?: {string: [string]};")
	fmt.Fprintf(dst, "\n}")

	if err := codegen.WriteFormattedToFile(fn, buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, bytes.NewReader(buf.Bytes()))
		return errors.Wrapf(err, `failed to write to %s`, fn)
	}
	return nil
}

func writeTypesFile(ctx *Context) error {
	fn := filepath.Join(ctx.dir, "types.js")
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

	log.Printf("  * %d types to define", len(typDefs))

	for _, typDef := range typDefs {
		typ := typDef.Type
		log.Printf("   * Generating definition for %s", typ.Name())
		switch t := typ.(type) {
		//		case *Array:
		//			fmt.Fprintf(dst, "\n\ntype %s []%s", t.name, t.elem)
		case *compiler.Struct:
			fmt.Fprintf(dst, "\n\n// %s represents the data structure defined in %s", typ.Name(), typDef.Context)
			fmt.Fprintf(dst, "\nexport type %s = {", t.Name())
			fields := t.Fields()
			for i, field := range fields {
				optional := ""
				if field.Required() {
					optional = "?"
				}
				fmt.Fprintf(dst, "\n%s: %s%s", field.Hints().JsName, optional, field.Type().Name())
				if i != len(fields) {
					fmt.Fprintf(dst, ",")
				}
			}

			fmt.Fprintf(dst, "\npayload(): string")
			fmt.Fprintf(dst, "\n}")
		}
	}

	if err := codegen.WriteFormattedToFile(fn, buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, bytes.NewReader(buf.Bytes()))
		return errors.Wrapf(err, `failed to write to %s`, fn)
	}
	return nil
}

func writeClientFile(ctx *Context) error {
	fn := filepath.Join(ctx.dir, "client.js")
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
		fn = filepath.Join(ctx.dir, "services", codegen.FileName(fn)+".js")
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
	fmt.Fprintf(dst, "\nimport type { ClientOptions } from './types/client_options.js';")
	for _, name := range ctx.client.ServiceNames() {
		fmt.Fprintf(dst, "\nimport %s from './services/%s'", codegen.ClassName(name), codegen.FileName(strings.TrimSuffix(name, "Service")))
	}

	fmt.Fprintf(dst, "\n\nexport default class %s {", ctx.clientName)
	for _, name := range ctx.client.ServiceNames() {
		fmt.Fprintf(dst, "\n%s: %s;", codegen.FieldName(name), codegen.ClassName(name))
	}

	fmt.Fprintf(dst, "\n\nconstructor(server: string, options: ?ClientOptions) {")
	for _, name := range ctx.client.ServiceNames() {
		fmt.Fprintf(dst, "\nthis.%s = new %s(server, options);", codegen.FieldName(name), codegen.ClassName(name))
	}
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	return nil
}

func formatService(ctx *Context, dst io.Writer, svc *compiler.Service) error {
	log.Printf(" * Generating Service %s", svc.Name())
	codegen.WritePreamble(dst, ctx.packageName)

	fmt.Fprintf(dst, "\n\nimport RESTResponse from '../types/response';")
	fmt.Fprintf(dst, "\n\nimport type { ClientOptions } from '../types/client_options';")

	for _, call := range svc.Calls() {
		if err := formatCall(ctx, dst, call); err != nil {
			return errors.Wrapf(err, `failed to format call %s`, call.Name())
		}
	}

	fmt.Fprintf(dst, "\n\nexport default class %s {", svc.Name())
	fmt.Fprintf(dst, "\nserver: string;")
	fmt.Fprintf(dst, "\noptions: ?ClientOptions;")
	fmt.Fprintf(dst, "\n\nconstructor(server: string, options: ?ClientOptions) {")
	fmt.Fprintf(dst, "\nthis.server = server;")
	fmt.Fprintf(dst, "\nthis.options = options;")
	fmt.Fprintf(dst, "\n}")
	for _, call := range svc.Calls() {
		fmt.Fprintf(dst, "\n\n%s (", codegen.MethodName(call.Method()))

		requireds := call.Requireds()
		for i, field := range requireds {
			fmt.Fprintf(dst, "%s: %s", codegen.FieldName(field.Name()), esType(field.Type()))
			if i < len(requireds)-1 {
				fmt.Fprintf(dst, ", ")
			}
		}

		fmt.Fprintf(dst, ") :%s {", codegen.ClassName(call.Name()))
		fmt.Fprintf(dst, "\nlet call = new %s(this.server, this.options", codegen.ClassName(call.Name()))
		for _, field := range requireds {
			fmt.Fprintf(dst, ", %s", codegen.FieldName(field.Name()))
		}
		fmt.Fprintf(dst, ");")
		fmt.Fprintf(dst, "\nreturn call;")
		fmt.Fprintf(dst, "\n}")
	}
	fmt.Fprintf(dst, "\n}")
	return nil
}

func jsType(s string) string {
	if strings.HasPrefix(s, "[]") {
		return "Array<" + s[2:] + ">"
	}
	return s
}

func formatCall(ctx *Context, dst io.Writer, call *compiler.Call) error {
	log.Printf("formatCall for %s", call.Name())

	for _, typ := range []compiler.Type{call.Body(), call.Path(), call.Query(), call.Header()} {
		if typ == nil {
			continue
		}

		if err := formatCallPayload(ctx, dst, call, typ); err != nil {
			return errors.Wrap(err, `failed to declare call body`)
		}
	}

	fmt.Fprintf(dst, "\n\nclass %s {", codegen.ClassName(call.Name()))
	fmt.Fprintf(dst, "\n_contentType :string;")
	fmt.Fprintf(dst, "\n_options :?ClientOptions;")
	fmt.Fprintf(dst, "\n_server :string;")
	if bodyType := call.Body(); bodyType != nil {
		fmt.Fprintf(dst, "\n_body :%s;", jsType(bodyType.Name()))
	}
	if queryType := call.Query(); queryType != nil {
		fmt.Fprintf(dst, "\n_query :%s", queryType.Name())
	}
	if pathType := call.Path(); pathType != nil {
		fmt.Fprintf(dst, "\n_path :%s", pathType.Name())
	}
	if headerType := call.Header(); headerType != nil {
		fmt.Fprintf(dst, "\n_header :%s", headerType.Name())
	}

	fmt.Fprintf(dst, "\n\nconstructor(server :string, options: ?ClientOptions")
	if requireds := call.Requireds(); len(requireds) > 0 {
		for _, field := range requireds {
			fmt.Fprintf(dst, ", %s: %s", codegen.FieldName(field.Name()), esType(field.Type()))
		}
	}
	fmt.Fprintf(dst, ") {")
	fmt.Fprintf(dst, "\nthis._options = options;")
	fmt.Fprintf(dst, "\nthis._server = server;")
	if requireds := call.Requireds(); len(requireds) > 0 {
		for _, field := range requireds {
			fmt.Fprintf(dst, "\nthis._%[1]s.%[2]s = %[2]s", field.ContainerName(), codegen.FieldName(field.Name()))
		}
	}
	fmt.Fprintf(dst, "\n}")

	if optionals := call.Optionals(); len(optionals) > 0 {
		for _, field := range optionals {
			fmt.Fprintf(dst, "\n\n%s(v: %s) :%s{", codegen.MethodName(field.Name()), esType(field.Type()), codegen.ClassName(call.Name()))
			fmt.Fprintf(dst, "\nthis._%s.%s = v;", field.ContainerName(), codegen.FieldName(field.Name()))
			fmt.Fprintf(dst, "\nreturn this;")
			fmt.Fprintf(dst, "\n}")
		}
	}

	fmt.Fprintf(dst, "\n\nasync do(): Promise<RESTResponse> {")

	if call.Path() == nil {
		fmt.Fprintf(dst, "\nconst path :string = %s;", strconv.Quote(call.RequestPath()))
	} else {
		fmt.Fprintf(dst, "\nlet path :string = %s;", strconv.Quote(call.RequestPath()))
		if pp, ok := call.Path().(*compiler.Struct); ok {
			for _, param := range pp.Fields() {
				fmt.Fprintf(dst, "\npath = path.replace(\"{%s}\", this._path.%s);", param.Name(), codegen.FieldName(param.Name()))
			}
		}
	}

	fmt.Fprintf(dst, "\n\nlet url = this._server + path;")
	if call.Query() != nil {
		fmt.Fprintf(dst, "\nlet query :Array<string> = [];")
		// XXX currently this code doesn't handle complex query params
		if pp, ok := call.Path().(*compiler.Struct); ok {
			for _, param := range pp.Fields() {
				fmt.Fprintf(dst, "\nif (this.%[1]s !== null && this.%[1]s !== undefined) {", codegen.FieldName(param.Name()))
				if strings.HasPrefix(esType(param.Type()), "Array<") {
					fmt.Fprintf(dst, "query.push(this.%s.map(v => '%s=' + encodeURIComponent(v)).join('&'))", codegen.FieldName(param.Name()), param.Name())
				} else {
					fmt.Fprintf(dst, "query.push('%s=' + encodeURIComponent(this.%s))", param.Name(), codegen.FieldName(param.Name()))
				}
				fmt.Fprintf(dst, "\n}")
			}
		}
		fmt.Fprintf(dst, "\nif (query.length > 0) {")
		fmt.Fprintf(dst, "\nurl = url + '?' + query.join('&');")
		fmt.Fprintf(dst, "\n}")
	}
	fmt.Fprintf(dst, ";")

	if call.Body() != nil {
		fmt.Fprintf(dst, "\n\nlet contentType = this._contentType;")
		fmt.Fprintf(dst, "\nif (contentType == '') {")
		fmt.Fprintf(dst, "\ncontentType = %s;", strconv.Quote(call.Consumes()[0]))
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\nlet mime = contentType;")
		fmt.Fprintf(dst, "\nlet seploc = mime.indexOf(';');")
		fmt.Fprintf(dst, "\nif (seploc > -1) {")
		fmt.Fprintf(dst, "\nmime = mime.substr(seploc);")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\nlet body :string | FormData")
		fmt.Fprintf(dst, "\nswitch (mime) {")
		fmt.Fprintf(dst, "\ncase 'application/json':")
		fmt.Fprintf(dst, "\nbody = this._body._json();")
		fmt.Fprintf(dst, "\nbreak;")
		fmt.Fprintf(dst, "\ncase 'application/x-www-form-urlencoded':")
		fmt.Fprintf(dst, "\nbody = this._body._form();")
		fmt.Fprintf(dst, "\nbreak;")
		fmt.Fprintf(dst, "\ndefault:")
		fmt.Fprintf(dst, "\nreturn Promise.resolve(new RESTResponse(500, {error: 'unsupported content-type ' + mime}))")
		fmt.Fprintf(dst, "\n}")
	}

	fmt.Fprintf(dst, "\nlet options = {")
	fmt.Fprintf(dst, "\nmethod: %s,", strconv.Quote(call.Verb()))
	if call.Body() != nil {
		fmt.Fprintf(dst, "\nheaders: {")
		fmt.Fprintf(dst, "\n'Content-Type': contentType")
		fmt.Fprintf(dst, "\n},")
		fmt.Fprintf(dst, "\nbody:  body")
	}
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nif (this._options !== null && this._options !== undefined) {")
	fmt.Fprintf(dst, "\nif (this._options.modifyOptions !== null && this._options.modifyOptions !== undefined) {")
	fmt.Fprintf(dst, "\noptions = this._options.modifyOptions(options);")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\nlet promise = fetch(url, options).")
	fmt.Fprintf(dst, "\nthen(response => {")
	if len(call.Responses()) == 0 {
		fmt.Fprintf(dst, "\nreturn Promise.resolve(new RESTResponse(response.status))")
	} else {
		fmt.Fprintf(dst, "\nif (response.headers.get('Content-Type').match(/^application\\/json/i)) {")
		fmt.Fprintf(dst, "\nreturn Promise.resolve(new RESTResponse(response.status, response.json()));")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nreturn Promise.resolve(new RESTResponse(500, \"Unhandled response type \" + response.headers.get('Content-Type')));")
	}

	fmt.Fprintf(dst, "\n});")
	fmt.Fprintf(dst, "\nreturn promise;")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	return nil
}

func formatCallPayload(ctx *Context, dst io.Writer, call *compiler.Call, typ compiler.Type) error {
	fmt.Fprintf(dst, "\n\nclass %s {", codegen.ClassName(typ.Name()))
	switch typ := typ.(type) {
	case *compiler.Array:
		fmt.Fprintf(dst, "\nelements: %s;", typ.Elem())
		fmt.Fprintf(dst, "\n\n_json(): %s {", typ.Elem())
		fmt.Fprintf(dst, "\nreturn JSON.stringify(this.elements);")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\n_form(): FormData {")
		fmt.Fprintf(dst, "\nlet form = new FormData();")
		fmt.Fprintf(dst, "\nthis.elements.forEach(v => form.append(%s, String(v)));", codegen.FieldName(typ.Name()))
		fmt.Fprintf(dst, "\n}")
	case *compiler.Struct:
		fields := typ.Fields()
		for _, field := range fields {
			fmt.Fprintf(dst, "\n%s: %s;", codegen.FieldName(field.Name()), esType(field.Type()))
		}

		fmt.Fprintf(dst, "\n\n_json(): string {")
		fmt.Fprintf(dst, "\nlet object = {")
		for i, field := range fields {
			fmt.Fprintf(dst, "\n%s: this.%s", strconv.Quote(field.Name()), codegen.FieldName(field.Name()))
			if i < len(fields)-1 {
				fmt.Fprintf(dst, ",")
			}
		}
		fmt.Fprintf(dst, "\n};")
		fmt.Fprintf(dst, "\nreturn JSON.stringify(object);")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\n_form(): FormData {")
		fmt.Fprintf(dst, "\nlet form = new FormData();")
		for _, field := range fields {
			if strings.HasPrefix(esType(field.Type()), "Array<") {
				fmt.Fprintf(dst, "\nthis.%s.forEach(v => form.append(%s, String(v)));", codegen.FieldName(field.Name()), strconv.Quote(field.Name()))
			} else {
				fmt.Fprintf(dst, "\nform.append(%s, String(this.%s));", strconv.Quote(field.Name()), codegen.FieldName(field.Name()))
			}
		}
		fmt.Fprintf(dst, "\nreturn form;")
		fmt.Fprintf(dst, "\n}")
	}
	fmt.Fprintf(dst, "\n}")
	return nil
}

/*

		if call.Body() != nil {
			var bodyFields []*Field
			for _, field := range call.AllFields() {
				if !field.inBody {
					continue
				}
				bodyFields = append(bodyFields, field)
			}

			fmt.Fprintf(dst, "\n\nclass %s {", codegen.ClassName(call.Body().Name()))
			for _, field := range bodyFields {
				fmt.Fprintf(dst, "\n%s: %s;", codegen.FieldName(field.name), field.typ)
			}

			fmt.Fprintf(dst, "\n\n_json(): string {")
			fmt.Fprintf(dst, "\nlet object = {")
			for i, field := range bodyFields {
				fmt.Fprintf(dst, "\n%s: this.%s", field.name, codegen.FieldName(field.name))
				if i < len(bodyFields)-1 {
					fmt.Fprintf(dst, ",")
				}
			}
			fmt.Fprintf(dst, "\n};")
			fmt.Fprintf(dst, "\nreturn JSON.stringify(object);")
			fmt.Fprintf(dst, "\n}")

			fmt.Fprintf(dst, "\n\n_form(): FormData {")
			fmt.Fprintf(dst, "\nlet form = new FormData();")
			for _, field := range bodyFields {
				if strings.HasPrefix(field.typ, "Array<") {
					fmt.Fprintf(dst, "\nthis.%s.forEach(v => form.append(%s, String(v)));", codegen.FieldName(field.name), strconv.Quote(field.name))
				} else {
					fmt.Fprintf(dst, "\nform.append(%s, String(this.%s));", strconv.Quote(field.name), codegen.FieldName(field.name))
				}
			}
			fmt.Fprintf(dst, "\nreturn form;")
			fmt.Fprintf(dst, "\n}")

			fmt.Fprintf(dst, "\n}")
		}

	return nil
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

func compileCall(ctx *Context, oper openapi.Operation) error {
	// x-service dictates the service name. If not present,
	// the default service, which is named after the package
	// is used.

	callName := codegen.CallObjectName(oper)
	methodName := codegen.CallMethodName(oper)
	if methodName == "" {
		methodName = codegen.MethodName(strings.TrimSuffix(callName, "Call"))
	}

	call := &Call{
		name:   callName,
		method: methodName,
		path:   oper.PathItem().Path(),
		verb:   oper.Verb(),
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
		log.Printf("resp.StatusCode = %s", resp.StatusCode())
		if resp.StatusCode() != "200" {
			log.Printf("---> skipping response for %s", call.path)
			continue
		}
		json.NewEncoder(os.Stdout).Encode(resp)
		typ, err := compileResponseType(ctx, resp)
		if err != nil {
			return errors.Wrap(err, `failed to compile response type`)
		}
		if typ == "" {
			log.Printf("---> skipping response for %s (no type)", call.path)
			continue
		}
		log.Printf("---> appending response for %s (%s)", call.path, typ)
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

			prop, err := b.Build()
			if err != nil {
				return errors.Wrapf(err, `failed to create schema for form parameter %s`, param.Name())
			}

			formBuilder.Property(param.Name(), prop)
		case openapi.InBody:
			// sanity check (although this should have already been taken care
			// of in Validate())
			if call.Body() != nil {
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
			field.jsName = codegen.FieldName(param.Name())
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
		formSchema, err := formBuilder.Build()
		if err != nil {
			return errors.Wrap(err, `failed to build schema for formData fields`)
		}
		formType, err := compileStruct(ctx, formSchema)
		if err != nil {
			return errors.Wrap(err, `failed to compile schema for formData fields`)
		}
		if formType.Name() == "" {
			formType.SetName(codegen.ClassName(call.name + "_Form"))
			registerType(ctx, fmt.Sprintf("#/generated/%s", formType.Name()), formType, call.name+" body form")
		}
		switch typ := formType.(type) {
		case *Struct:
			for _, field := range typ.fields {
				field.inBody = true
				if field.required {
					call.requireds = append(call.requireds, field)
				} else {
					call.optionals = append(call.optionals, field)
				}
			}
		}
		call.bodyType = formType

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

	svcName = strcase.ToCamel(svcName)
	svc := ctx.client.getServiceFor(svcName)
	svc.addCall(call)
	return nil
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
			typ.SetName(codegen.ClassName(ctx.currentCall.name + "_Response"))
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
				typ.SetName(codegen.ClassName(ctx.currentCall.name + "_" + param.Name()))
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
	case openapi.Array:
		return compileArray(ctx, param)
	default:
		return compilePrimitiveType(param.Type(), param.Format())
	}

	return Builtin(param.Type()), nil
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
			jsName:   codegen.FieldName(name),
			tag:      fmt.Sprintf(`json:"%s"`, name),
			typ:      fieldMsg.Name(),
			required: schema.IsRequiredProperty(name),
		})
	}
	return &obj, nil
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

func compileItems(ctx *Context, items openapi.Items) (t Type, err error) {
	return compileSchemaLike(ctx, items)
}

func compileSchema(ctx *Context, schema openapi.Schema) (t Type, err error) {
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
				n := codegen.ClassName(strings.TrimPrefix(ref, "#/definitions/"))
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

func compilePrimitiveType(typ openapi.PrimitiveType, format string) (Type, error) {
	switch typ {
	case openapi.Number, openapi.Integer:
		return Builtin("number"), nil
	case openapi.String:
		return Builtin("string"), nil
	case openapi.Boolean:
		return Builtin("boolean"), nil
	default:
		return nil, errors.Errorf(`primitive type %s should not have gone through CompilePrimitive`, typ)
	}
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
*/
