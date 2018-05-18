package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/lestrrat-go/openapi/v3/entity"
	"github.com/pkg/errors"
)

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

func snakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	switch s := string(out); {
	case strings.Contains(s, "open_api"):
		return strings.Replace(s, "open_api", "openapi", -1)
	case strings.Contains(s, "o_auth"):
		return strings.Replace(s, "o_auth", "oauth", -1)
	default:
		return s
	}
}

func lcfirst(s string) string {
	if len(s) <= 0 {
		return s
	}

	switch s {
	case "URL":
		return "url"
	}

	r, w := utf8.DecodeRuneInString(s)
	var buf bytes.Buffer
	buf.WriteRune(unicode.ToLower(r))
	buf.WriteString(s[w:])
	return buf.String()
}

func fieldName(s string) string {
	n := lcfirst(s)
	switch n {
	case "default":
		return "defaultValue"
	case "type":
		return "typ"
	default:
		return n
	}
}

func typname(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Interface:
		return "interface{}"
	case reflect.Ptr:
		return "*" + typname(t.Elem())
	case reflect.Slice:
		return "[]" + typname(t.Elem())
	case reflect.Map:
		return "map[" + typname(t.Key()) + "]" + typname(t.Elem())
	default:
		return t.Name()
	}
}

func _main() error {
	entities := []interface{}{
		entity.Callback{},
		entity.Components{},
		entity.Contact{},
		entity.Discriminator{},
		entity.Encoding{},
		entity.Example{},
		entity.ExternalDocumentation{},
		entity.Header{},
		entity.Info{},
		entity.License{},
		entity.Link{},
		entity.MediaType{},
		entity.OAuthFlow{},
		entity.OAuthFlows{},
		entity.OpenAPI{},
		entity.Operation{},
		entity.Parameter{},
		entity.PathItem{},
		entity.Paths{},
		entity.RequestBody{},
		entity.Response{},
		entity.Responses{},
		entity.Schema{},
		entity.SecurityRequirement{},
		entity.SecurityScheme{},
		entity.Server{},
		entity.ServerVariable{},
		entity.Tag{},
	}

	for _, e := range entities {
		if err := generateFromEntity(e); err != nil {
			return errors.Wrap(err, `failed to encode entity`)
		}
	}

	return nil
}

func writeImports(dst io.Writer, pkgs []string) {
	fmt.Fprintf(dst, "\n\nimport (")
	for _, pkg := range pkgs {
		fmt.Fprintf(dst, "\n%s", strconv.Quote(pkg))
	}
	fmt.Fprintf(dst, "\n)")
}

func generateFromEntity(e interface{}) error {
	rv := reflect.ValueOf(e)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	fmt.Fprintf(dst, "package builder")

	writeImports(dst, []string{"github.com/lestrrat-go/openapi/v3/entity"})

	fmt.Fprintf(dst, "\n\ntype %s struct {", rv.Type().Name())
	fmt.Fprintf(dst, "\ntarget %s", reflect.PtrTo(rv.Type()))
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (b *%s) Build() %s {", rv.Type().Name(), reflect.PtrTo(rv.Type()))
	fmt.Fprintf(dst, "\nreturn b.target")
	fmt.Fprintf(dst, "\n}")

	// Iterate through the fields, check if they are required / explicitly
	// excluded from certain operations
	var defaults []reflect.StructField
	var requireds []reflect.StructField
	var optionals []reflect.StructField
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Type().Field(i)
		hasDefault := len(fv.Tag.Get("default")) > 0
		if hasDefault {
			defaults = append(defaults, fv)
		}
		switch fv.Tag.Get("builder") {
		case "-":
			continue // ignore this
		case "required":
			if hasDefault {
				optionals = append(optionals, fv)
			} else {
				requireds = append(requireds, fv)
			}
		default:
			optionals = append(optionals, fv)
		}
	}

	fmt.Fprintf(dst, "\n\nfunc (b *Builder) New%s(", rv.Type().Name())
	for i, fv := range requireds {
		fmt.Fprintf(dst, "%s %s", fieldName(fv.Name), fv.Type)
		if i < len(requireds)-1 {
			fmt.Fprintf(dst, ", ")
		}
	}
	fmt.Fprintf(dst, ") *%s {", rv.Type().Name())
	fmt.Fprintf(dst, "\nreturn &%s{", rv.Type().Name())
	fmt.Fprintf(dst, "\ntarget: &entity.%s{", rv.Type().Name())
	for _, fv := range defaults {
		fmt.Fprintf(dst, "\n%s: %s,", fv.Name, fv.Tag.Get("default"))
	}
	for _, fv := range requireds {
		fmt.Fprintf(dst, "\n%s: %s,", fv.Name, fieldName(fv.Name))
	}
	fmt.Fprintf(dst, "\n},")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	for _, fv := range optionals {
		fmt.Fprintf(dst, "\n\nfunc (b *%s) %s(v %s) *%s {", rv.Type().Name(), fv.Name, fv.Type, rv.Type().Name())
		fmt.Fprintf(dst, "\nb.target.%s = v", fv.Name)
		fmt.Fprintf(dst, "\nreturn b")
		fmt.Fprintf(dst, "\n}")
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.String())
		return errors.Wrap(err, `failed to format source`)
	}

	filename := filepath.Join("builder", fmt.Sprintf("%s_gen.go", snakeCase(rv.Type().Name())))
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrapf(err, `failed to open file %s`, filename)
	}
	defer f.Close()

	f.Write(formatted)

	return nil
}
