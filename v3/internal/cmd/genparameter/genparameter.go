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
	"strings"
	"unicode"
	"unicode/utf8"

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

	return strings.Replace(string(out), "open_api", "openapi", -1)
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
	default:
		return n
	}
}

func typname(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Ptr:
		return "*" + typname(t.Elem())
	case reflect.Interface:
		return "interface{}"
	case reflect.Slice:
		return "[]" + typname(t.Elem())
	case reflect.Map:
		return "map[" + typname(t.Key()) + "]" + typname(t.Elem())
	default:
		return t.Name()
	}
}

type field struct {
	Name string
	Type string
	Tag  string
}

func _main() error {
	commonFields := []field{
		{"Description", "string", `json:"description,omitempty"`},
		{"Deprecated", "bool", `json:"deprecated,omitempty"`},
		{"AllowEmptyValue", "bool", `json:"allowEmptyValue,omitempty"`},
		{"Explode", "bool", `json:"explode,omitempty"`},
		{"AllowReserved", "bool", `json:"allowReserved,omitempty"`},
		{"Schema", "*Schema", `json:"schema,omitempty"`},
		{"Example", "interface{}", `json:"example,omitempty"`},
		{"Examples", "map[string]*Example", `json:"examples,omitempty"`},
		{"Content", "map[string]*MediaType", `json:"content,omitempty"`},
	}

	parameterFields := []field{
		{"Name", "string", `json:"name,omitempty" builder:"required"`},
		{"In", "Location", `json:"in" builder:"required"`},
		{"Required", "bool", `json:"required,omitempty" default:"defaultParameterRequiredFromLocation(in)"`},
	}

	headerFields := []field{
		{"In", "Location", `json:"-" builder:"required" default:"entity.InHeader"`},
		{"Required", "bool", `json:"required,omitempty"`},
	}

	generateEntity("Parameter", parameterFields, commonFields)
	generateEntity("Header", headerFields, commonFields)
	return nil
}

func generateEntity(name string, extras, commons []field) error {
	var buf bytes.Buffer
	var dst io.Writer = &buf

	fmt.Fprintf(dst, "package entity")
	fmt.Fprintf(dst, "\n\ntype %s struct {", name)
	for _, extra := range extras {
		fmt.Fprintf(dst, "\n%s %s `%s`", extra.Name, extra.Type, extra.Tag)
	}
	for _, common := range commons {
		fmt.Fprintf(dst, "\n%s %s `%s`", common.Name, common.Type, common.Tag)
	}
	fmt.Fprintf(dst, "\n}")

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.String())
		return errors.Wrap(err, `failed to format source`)
	}

	filename := filepath.Join("entity", fmt.Sprintf("%s_gen.go", snakeCase(name)))
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrapf(err, `failed to open file %s`, filename)
	}
	defer f.Close()

	f.Write(formatted)

	return nil
}
