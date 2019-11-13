package types

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	codegen "github.com/lestrrat-go/openapi/internal/codegen/golang"
	"github.com/lestrrat-go/openapi/internal/stringutil"
	"github.com/pkg/errors"
)

var containers = []interface{}{
	ExampleMap{},
	HeaderMap{},
	InterfaceList{},
	InterfaceMap{},
	MIMETypeList{},
	ParameterList{},
	ParameterMap{},
	PathItemMap{},
	ResponseMap{},
	SchemaList{},
	SchemaMap{},
	SchemeList{},
	ScopesMap{},
	SecurityRequirementList{},
	SecuritySchemeMap{},
	StringList{},
	StringMap{},
	TagList{},
}
var entities = []interface{}{
	swagger{},
	info{},
	contact{},
	license{},
	paths{},
	pathItem{},
	operation{},
	externalDocumentation{},
	parameter{},
	items{},
	responses{},
	response{},
	header{},
	schema{},
	xml{},
	securityScheme{},
	securityRequirement{},
	tag{},
}
var postUnmarshalJSONHooks = map[string]struct{}{}
var entityTypes = make(map[string]interface{})
var containerTypes = make(map[string]interface{})
var validators = make(map[string]struct{})
var packageName = "openapi2"

func GenerateCode() error {
	for _, e := range entities {
		name := reflect.TypeOf(e).Name()
		entityTypes[name] = e
		switch name {
		default:
			validators[name] = struct{}{}
		}
	}

	for _, c := range containers {
		containerTypes[reflect.TypeOf(c).Name()] = c
	}

	if err := copyInterface(); err != nil {
		return errors.Wrap(err, `failed to copy interface`)
	}

	for _, e := range entities {
		if err := generateJSONHandlersFromEntity(e); err != nil {
			return errors.Wrap(err, `failed to generate JSON handlers from entity`)
		}

		if err := generateAccessorsFromEntity(e); err != nil {
			return errors.Wrap(err, `failed to generate accessors from entity`)
		}

		if err := generateBuilderFromEntity(e); err != nil {
			return errors.Wrap(err, `failed to generate builder from entity`)
		}

		if err := generateMutatorFromEntity(e); err != nil {
			return errors.Wrap(err, `failed to generate mutator from entity`)
		}

		if err := generateVisitorsFromEntity(e); err != nil {
			return errors.Wrap(err, `failed to generate visitors from entity`)
		}
	}

	for _, c := range containers {
		if err := generateContainer(c); err != nil {
			return errors.Wrap(err, `failed to generate container`)
		}
	}

	if err := generateClonersFromEntity(entities); err != nil {
		return errors.Wrap(err, `failed to generate clone methods from entity list`)
	}

	if err := generateIteratorsFromEntity(containers); err != nil {
		return errors.Wrap(err, `failed to generate iterators from entity list`)
	}

	if err := generateVisitor(entities); err != nil {
		return errors.Wrap(err, `failed to generate visitor from entity list`)
	}

	return nil
}

func isStockValidator(s string) bool {
	_, ok := validators[s]
	return ok
}

func isContainer(s string) bool {
	_, ok := containerTypes[s]
	return ok
}

func isEntity(s string) bool {
	_, ok := entityTypes[s]
	return ok
}

func isMap(s string) bool {
	return strings.HasSuffix(s, "Map")
}

func isList(s string) bool {
	return strings.HasSuffix(s, "List")
}

func typname(t reflect.Type) string {
	if t.Name() != "" {
		return t.Name()
	}

	switch t.Kind() {
	case reflect.Interface:
		if t.Name() == "" {
			return "interface{}"
		}
		return t.Name()
	case reflect.Ptr:
		return "*" + typname(t.Elem())
	case reflect.Slice:
		return "[]" + typname(t.Elem())
	case reflect.Map:
		return "map[" + typname(t.Key()) + "]" + typname(t.Elem())
	default:
		name := t.Name()
		if name == "" {
			panic("typname empty: " + t.String())
		}
		return name
	}
}

func iteratorName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Map, reflect.Slice:
	default:
		panic("unimplemented " + t.String())
	}

	// This should only happen for those types that we
	// have declared wrapper containers for. Use the
	// naming schemes
	var name string = t.Name()
	if !isContainer(name) { // && !isEntity(t.Elem().Name()) {
		name = typname(t.Elem())
		if strings.HasPrefix(name, "[]") {
			name = strings.TrimPrefix(name, "[]") + "List"
		}
		if strings.HasPrefix(name, "map[") {
			name = strings.TrimPrefix(name, "[]") + "Map"
		}
		name = strcase.ToCamel(name)
	}

	if name == "" {
		panic("empty name " + t.String())
	}

	return strcase.ToCamel(name + "Iterator")
}

func completeInterface(dst io.Writer, ifacename string) {
	log.Printf(" * Completing interface for %s", ifacename)

	e, ok := entityTypes[codegen.UnexportedName(ifacename)]
	if !ok {
		panic(fmt.Sprintf("Could not find value for %s", strcase.ToLowerCamel(ifacename)))
	}

	rv := reflect.ValueOf(e)

	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Type().Field(i)
		if fv.Tag.Get("accessor") == "-" {
			continue
		}

		// If it's a container type, we need to return iterators
		fieldType := fv.Type.Name()
		exportedField := codegen.ExportedName(fv.Name)
		log.Printf("   - %s", exportedField)
		if _, ok := containerTypes[fieldType]; ok {
			fmt.Fprintf(dst, "\n%s() *%s", exportedField, iteratorName(fv.Type))
		} else {
			if fv.Tag.Get("accessor") == "indirect" {
				fmt.Fprintf(dst, "\nHas%s() bool", exportedField)
				fmt.Fprintf(dst, "\n%s() %s", exportedField, typname(fv.Type.Elem()))
			} else {
				fmt.Fprintf(dst, "\n%s() %s", exportedField, typname(fv.Type))
			}
		}
	}

	fmt.Fprintf(dst, "\nExtension(string) (interface{}, bool)")
	fmt.Fprintf(dst, "\nExtensions() *ExtensionsIterator")
	fmt.Fprintf(dst, "\nClone() %s", ifacename)
	fmt.Fprintf(dst, "\nIsUnresolved() bool")
	fmt.Fprintf(dst, "\nMarshalJSON() ([]byte, error)")
	fmt.Fprintf(dst, "\nReference() string")
	fmt.Fprintf(dst, "\nValidator")
}

var importDummies = map[string]string{
	"encoding/json":         "json.Unmarshal",
	"fmt":                   "fmt.Fprintf",
	"github.com/pkg/errors": "errors.Cause",
	"log":                   "log.Printf",
	"sort":                  "sort.Strings",
	"strconv":               "strconv.ParseInt",
}

func writeImports(dst io.Writer, pkgs []string) {
	fmt.Fprintf(dst, "\n\nimport (")
	for _, pkg := range pkgs {
		fmt.Fprintf(dst, "\n%s", strconv.Quote(pkg))
	}
	fmt.Fprintf(dst, "\n)")

	// check to see if we need dummies
	var buf bytes.Buffer
	for _, pkg := range pkgs {
		if v, ok := importDummies[pkg]; ok {
			fmt.Fprintf(&buf, "\nvar _ = %s", v)
		}
	}

	if buf.Len() > 0 {
		fmt.Fprintf(dst, "\n")
		buf.WriteTo(dst)
	}
}

func copyInterface() error {
	log.Printf("Generating interface.go")

	var buf bytes.Buffer
	var dst io.Writer = &buf

	src, err := os.Open(filepath.Join("internal", "types", "interface.go"))
	if err != nil {
		return errors.Wrapf(err, `failed to open source file`)
	}

	scanner := bufio.NewScanner(src)

	// throw away first N lines
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "package types" {
			break
		}
	}

	codegen.WritePreamble(dst, packageName)

	for scanner.Scan() {
		txt := scanner.Text()

		txt = strings.Replace(txt, "//gen:lazy ", "", -1)

		fmt.Fprintf(dst, "\n%s", txt)
		if strings.HasPrefix(txt, "type ") && strings.HasSuffix(txt, " interface {") {
			i := strings.Index(txt[5:], " ")
			switch txt[5 : 5+i] {
			case "ResolveError", "Resolver", "Validator", "SchemaConverter":
			default:
				completeInterface(dst, txt[5:5+i])
			}
		} else if strings.HasPrefix(txt, "type ") && strings.HasSuffix(txt, " struct {") {
			i := strings.Index(txt[5:], " ")
			if isEntity(txt[5 : 5+i]) {
				fmt.Fprintf(dst, "\nreference string `json:\"$ref,omitempty\"`")
				fmt.Fprintf(dst, "\nresolved bool `json:\"-\"`")
				fmt.Fprintf(dst, "\nextensions Extensions `json:\"-\"`")
			}
		}
	}
	if err := codegen.WriteFormattedToFile("interface_gen.go", buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, &buf)
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func generateIteratorsFromEntity(entities []interface{}) error {
	filename := "iterator_gen.go"
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, packageName)
	writeImports(dst, []string{"sync"})

	fmt.Fprintf(dst, "\n\ntype mapIteratorItem struct {")
	fmt.Fprintf(dst, "\nitem interface{}")
	fmt.Fprintf(dst, "\nkey  interface{}")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\ntype mapIterator struct {")
	fmt.Fprintf(dst, "\nlist listIterator")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (iter *mapIterator) Next() bool{")
	fmt.Fprintf(dst, "\nreturn iter.list.Next()")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (iter *mapIterator) Size() int{")
	fmt.Fprintf(dst, "\nreturn iter.list.Size()")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (iter *mapIterator) Item() *mapIteratorItem {")
	fmt.Fprintf(dst, "\nv := iter.list.Item()")
	fmt.Fprintf(dst, "\nif v == nil {")
	fmt.Fprintf(dst, "\nreturn nil")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nreturn v.(*mapIteratorItem)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\ntype listIterator struct {")
	fmt.Fprintf(dst, "\nmu sync.RWMutex")
	fmt.Fprintf(dst, "\nitems []interface{}")
	fmt.Fprintf(dst, "\nsize int")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Size returns the size of the iterator. This size")
	fmt.Fprintf(dst, "\n// is fixed at creation time. It does not represent")
	fmt.Fprintf(dst, "\n// the remaining number of items")
	fmt.Fprintf(dst, "\nfunc (iter *listIterator) Size() int {")
	fmt.Fprintf(dst, "\nreturn iter.size")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Item returns the next item in this iterator")
	fmt.Fprintf(dst, "\nfunc (iter *listIterator) Item() interface{} {")
	fmt.Fprintf(dst, "\niter.mu.Lock()")
	fmt.Fprintf(dst, "\ndefer iter.mu.Unlock()")
	fmt.Fprintf(dst, "\n\nif !iter.nextNoLock() {")
	fmt.Fprintf(dst, "\nreturn nil")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n\nitem := iter.items[0]")
	fmt.Fprintf(dst, "\niter.items = iter.items[1:]")
	fmt.Fprintf(dst, "\nreturn item")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (iter *listIterator) nextNoLock() bool {")
	fmt.Fprintf(dst, "\nreturn len(iter.items) > 0")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Next returns true if there are more elements in this iterator")
	fmt.Fprintf(dst, "\nfunc (iter *listIterator) Next() bool {")
	fmt.Fprintf(dst, "\niter.mu.RLock()")
	fmt.Fprintf(dst, "\ndefer iter.mu.RUnlock()")
	fmt.Fprintf(dst, "\nreturn iter.nextNoLock()")
	fmt.Fprintf(dst, "\n}")

	writeListIterator := func(dst io.Writer, typ, item string) {
		fmt.Fprintf(dst, "\n\ntype %sListIterator struct {", typ)
		fmt.Fprintf(dst, "\nlistIterator")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\n// Item returns the next item in this iterator. Make sure to call Next()")
		fmt.Fprintf(dst, "\n// before hand to check if the iterator has more items")
		fmt.Fprintf(dst, "\nfunc (iter *%sListIterator) Item() %s {", typ, item)
		fmt.Fprintf(dst, "\nreturn iter.listIterator.Item().(%s)", item)
		fmt.Fprintf(dst, "\n}")
	}

	writeMapIterator := func(dst io.Writer, t reflect.Type) {
		name := t.Name()
		keyType := t.Key().Name()
		elemType := t.Elem().Name()
		if elemType == "" {
			elemType = typname(t.Elem())
		}

		fmt.Fprintf(dst, "\n\ntype %sIterator struct {", name)
		fmt.Fprintf(dst, "\nmapIterator")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\n// Item returns the next item in this iterator. Make sure to call Next()")
		fmt.Fprintf(dst, "\n// before hand to check if the iterator has more items")
		fmt.Fprintf(dst, "\nfunc (iter *%sIterator) Item() (%s, %s) {", name, keyType, elemType)
		fmt.Fprintf(dst, "\nitem := iter.mapIterator.Item()")
		fmt.Fprintf(dst, "\nif item == nil {")
		keyZero := strings.TrimPrefix(fmt.Sprintf("%#v", reflect.Zero(t.Key())), "types.")
		elemZero := strings.TrimPrefix(fmt.Sprintf("%#v", reflect.Zero(t.Elem())), "types.")
		fmt.Fprintf(dst, "\nreturn %s, %s", keyZero, elemZero)
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nreturn item.key.(%s), item.item.(%s)", keyType, elemType)
		fmt.Fprintf(dst, "\n}")
	}

	writeListIterator(dst, "Interface", "interface{}")
	writeListIterator(dst, "String", "string")
	writeListIterator(dst, "StringList", "[]string")
	writeListIterator(dst, "Operation", "Operation")
	writeListIterator(dst, "PathItem", "PathItem")
	writeListIterator(dst, "MediaType", "MediaType")
	writeMapIterator(dst, reflect.TypeOf(Extensions{}))

	for _, e := range entities {
		rv := reflect.ValueOf(e)
		switch rv.Type().Name() {
		case "InterfaceList", "StringList":
			continue
		}

		if rv.Kind() == reflect.Map {
			writeMapIterator(dst, rv.Type())
		} else {
			ifacename := strings.TrimSuffix(codegen.ExportedName(rv.Type().Name()), "List")
			writeListIterator(dst, ifacename, ifacename)
		}
	}

	return codegen.WriteFormattedToFile(filename, buf.Bytes())
}

func generateBuilderFromEntity(e interface{}) error {
	rv := reflect.ValueOf(e)
	filename := fmt.Sprintf("%s_builder_gen.go", strcase.ToSnake(rv.Type().Name()))
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, packageName)
	writeImports(dst, []string{"sync", "github.com/pkg/errors"})

	ifacename := codegen.ExportedName(rv.Type().Name())
	structname := rv.Type().Name()

	fmt.Fprintf(dst, "\n\n// %sBuilder is used to build an instance of %s. The user must", ifacename, ifacename)
	fmt.Fprintf(dst, "\n// call `Build()` after providing all the necessary information to")
	fmt.Fprintf(dst, "\n// build an instance of %s.", ifacename)
	fmt.Fprintf(dst, "\n// Builders may NOT be reused. It must be created for every instance")
	fmt.Fprintf(dst, "\n// of %s that you want to create", ifacename)
	fmt.Fprintf(dst, "\ntype %sBuilder struct {", ifacename)
	fmt.Fprintf(dst, "\nlock sync.Locker")
	fmt.Fprintf(dst, "\ntarget *%s", structname)
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// MustBuild is a convenience function for those time when you know that")
	fmt.Fprintf(dst, "\n// the result of the builder must be successful")
	fmt.Fprintf(dst, "\nfunc (b *%[1]sBuilder) MustBuild(options ...Option) %[1]s {", ifacename)
	fmt.Fprintf(dst, "\nv, err := b.Build(options...)")
	fmt.Fprintf(dst, "\nif err != nil {")
	fmt.Fprintf(dst, "\npanic(err)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nreturn v")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Build finalizes the building process for %s and returns the result", ifacename)
	fmt.Fprintf(dst, "\n// By default, Build() will validate if the given structure is valid")
	fmt.Fprintf(dst, "\nfunc (b *%[1]sBuilder) Build(options ...Option) (%[1]s, error) {", ifacename)
	fmt.Fprintf(dst, "\nb.lock.Lock()")
	fmt.Fprintf(dst, "\ndefer b.lock.Unlock()")
	fmt.Fprintf(dst, "\nif b.target == nil {")
	fmt.Fprintf(dst, "\nreturn nil, errors.New(`builder has already been used`)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nvalidate := true")
	fmt.Fprintf(dst, "\nfor _, option := range options {")
	fmt.Fprintf(dst, "\nswitch option.Name() {")
	fmt.Fprintf(dst, "\ncase optkeyValidate:")
	fmt.Fprintf(dst, "\nvalidate = option.Value().(bool)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nif validate {")
	fmt.Fprintf(dst, "\nif err := b.target.Validate(false); err != nil {")
	fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `validation failed`)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\ndefer func() { b.target = nil }()")
	fmt.Fprintf(dst, "\nreturn b.target, nil")
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

	fmt.Fprintf(dst, "\n\n// New%[1]s creates a new builder object for %[1]s", ifacename)
	fmt.Fprintf(dst, "\nfunc New%s(", ifacename)
	for _, fv := range requireds {
		fmt.Fprintf(dst, "%s %s", codegen.UnexportedName(fv.Name), typname(fv.Type))
		fmt.Fprintf(dst, ", ")
	}
	fmt.Fprintf(dst, "options ...Option) *%sBuilder {", ifacename)
	fmt.Fprintf(dst, "\nvar lock sync.Locker = &sync.Mutex{}")
	fmt.Fprintf(dst, "\nfor _, option := range options {")
	fmt.Fprintf(dst, "\nswitch option.Name() {")
	fmt.Fprintf(dst, "\ncase optkeyLocker:")
	fmt.Fprintf(dst, "\nlock = option.Value().(sync.Locker)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nvar b %sBuilder", ifacename)
	fmt.Fprintf(dst, "\nif lock == nil {")
	fmt.Fprintf(dst, "\nlock = nilLock{}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nb.lock = lock")
	fmt.Fprintf(dst, "\nb.target = &%s{", structname)
	hasDefault := make(map[string]struct{})
	for _, fv := range defaults {
		hasDefault[fv.Name] = struct{}{}
		fmt.Fprintf(dst, "\n%s: %s,", fv.Name, fv.Tag.Get("default"))
	}
	for _, fv := range requireds {
		fmt.Fprintf(dst, "\n%s: %s,", fv.Name, codegen.UnexportedName(fv.Name))
	}
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nreturn &b")
	fmt.Fprintf(dst, "\n}")

	for _, fv := range optionals {
		exportedFieldName := codegen.ExportedName(fv.Name)
		fmt.Fprintf(dst, "\n\n// %s sets the %s field for object %s.", exportedFieldName, fv.Name, ifacename)
		if _, ok := hasDefault[fv.Name]; ok {
			fmt.Fprintf(dst, " If this is not called,\n// a default value (%s) is assigned to this field", fv.Tag.Get("default"))
		}

		// if the inferred argument type is a list, we should allow a
		// variadic expression in the argument
		var argType string
		indirect := fv.Tag.Get("builder") == "indirect"
		if indirect {
			argType = typname(fv.Type.Elem())
		} else {
			argType = typname(fv.Type)
		}

		if strings.HasPrefix(argType, "[]") {
			argType = "..." + strings.TrimPrefix(argType, "[]")
		} else if isList(argType) {
			argType = strings.TrimSuffix(argType, "List")
			switch argType {
			case "String":
				argType = stringutil.LcFirst(argType)
			case "Interface":
				argType = "interface{}"
			}
			argType = "..." + argType
		}
		fmt.Fprintf(dst, "\nfunc (b *%sBuilder) %s(v %s) *%sBuilder {", ifacename, exportedFieldName, argType, ifacename)
		fmt.Fprintf(dst, "\nb.lock.Lock()")
		fmt.Fprintf(dst, "\ndefer b.lock.Unlock()")
		fmt.Fprintf(dst, "\nif b.target == nil {")
		fmt.Fprintf(dst, "\nreturn b")
		fmt.Fprintf(dst, "\n}")
		if indirect {
			fmt.Fprintf(dst, "\nb.target.%s = &v", fv.Name)
		} else {
			fmt.Fprintf(dst, "\nb.target.%s = v", fv.Name)
		}
		fmt.Fprintf(dst, "\nreturn b")
		fmt.Fprintf(dst, "\n}")
	}

	fmt.Fprintf(dst, "\n\n// Reference sets the $ref (reference) field for object %s.", ifacename)
	fmt.Fprintf(dst, "\nfunc (b *%sBuilder) Reference(v string) *%sBuilder {", ifacename, ifacename)
	fmt.Fprintf(dst, "\nb.lock.Lock()")
	fmt.Fprintf(dst, "\ndefer b.lock.Unlock()")
	fmt.Fprintf(dst, "\nif b.target == nil {")
	fmt.Fprintf(dst, "\nreturn b")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nb.target.reference = v")
	fmt.Fprintf(dst, "\nreturn b")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Extension sets an arbitrary element (an extension) to the")
	fmt.Fprintf(dst, "\n// object %s. The extension name should start with a \"x-\"", ifacename)
	fmt.Fprintf(dst, "\nfunc (b *%[1]sBuilder) Extension(name string, value interface{}) *%[1]sBuilder {", ifacename)
	fmt.Fprintf(dst, "\nb.lock.Lock()")
	fmt.Fprintf(dst, "\ndefer b.lock.Unlock()")
	fmt.Fprintf(dst, "\nif b.target == nil {")
	fmt.Fprintf(dst, "\nreturn b")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nb.target.extensions[name] = value")
	fmt.Fprintf(dst, "\nreturn b")
	fmt.Fprintf(dst, "\n}")

	if err := codegen.WriteFormattedToFile(filename, buf.Bytes()); err != nil {
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func generateClonersFromEntity(entities []interface{}) error {
	filename := "cloner_gen.go"
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, packageName)

	for _, e := range entities {
		rv := reflect.ValueOf(e)
		fmt.Fprintf(dst, "\n\nfunc (v *%s) Clone() %s {", rv.Type().Name(), codegen.ExportedName(rv.Type().Name()))
		fmt.Fprintf(dst, "\nvar dst %s", rv.Type().Name())
		fmt.Fprintf(dst, "\ndst = *v")
		fmt.Fprintf(dst, "\nreturn &dst")
		fmt.Fprintf(dst, "\n}")
	}

	return codegen.WriteFormattedToFile(filename, buf.Bytes())
}

func generateAccessorsFromEntity(e interface{}) error {
	rv := reflect.ValueOf(e)
	filename := fmt.Sprintf("%s_accessors_gen.go", strcase.ToSnake(rv.Type().Name()))
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, packageName)
	codegen.WriteImports(dst, "context", "sort", "github.com/pkg/errors")

	structname := rv.Type().Name()

	fmt.Fprintf(dst, "\n\nfunc (v *%s) IsValid() bool {", structname)
	fmt.Fprintf(dst, "\nreturn v != nil")
	fmt.Fprintf(dst, "\n}")

	var entityFields []reflect.StructField

	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Type().Field(i)
		if fv.Tag.Get("accessor") == "-" {
			log.Printf(" - Skipping field %s", fv.Name)
			continue
		}
		log.Printf(" * Checking field %s", fv.Name)

		exportedName := codegen.ExportedName(fv.Name)
		unexportedName := codegen.UnexportedName(fv.Name)
		fieldType := fv.Type.Name()

		// keep track of all fields whose type is one of our entity types
		if fv.Tag.Get("json") != "-" {
			if isEntity(codegen.UnexportedName(fieldType)) || isContainer(fieldType) {
				entityFields = append(entityFields, fv)
			}
		}

		switch {
		case isMap(fieldType):
			iteratorName := iteratorName(fv.Type)
			if iteratorName == "InterfaceIterator" {
				panic("fuck")
			}
			fmt.Fprintf(dst, "\n\nfunc (v *%s) %s() *%s {", structname, exportedName, iteratorName)
			fmt.Fprintf(dst, "\nvar keys []string")
			fmt.Fprintf(dst, "\nfor key := range v.%s {", unexportedName)
			fmt.Fprintf(dst, "\nkeys = append(keys, key)")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\nsort.Strings(keys)")
			fmt.Fprintf(dst, "\nvar items []interface{}")
			fmt.Fprintf(dst, "\nfor _, key := range keys {")
			fmt.Fprintf(dst, "\nitem := v.%s[key]", unexportedName)
			fmt.Fprintf(dst, "\nitems = append(items, &mapIteratorItem{key: key, item: item})")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\nvar iter %s", iteratorName)
			fmt.Fprintf(dst, "\niter.list.size = len(items)")
			fmt.Fprintf(dst, "\niter.list.items = items")
			fmt.Fprintf(dst, "\nreturn &iter")
			fmt.Fprintf(dst, "\n}")
		case isList(fieldType):
			iteratorName := iteratorName(fv.Type)
			fmt.Fprintf(dst, "\n\nfunc (v *%s) %s() *%s {", structname, exportedName, iteratorName)
			fmt.Fprintf(dst, "\nvar items []interface{}")
			fmt.Fprintf(dst, "\nfor _, item := range v.%s {", unexportedName)
			fmt.Fprintf(dst, "\nitems = append(items, item)")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\nvar iter %s", iteratorName)
			fmt.Fprintf(dst, "\niter.size = len(items)")
			fmt.Fprintf(dst, "\niter.items = items")
			fmt.Fprintf(dst, "\nreturn &iter")
			fmt.Fprintf(dst, "\n}")
		default:
			if fv.Tag.Get("accessor") == "indirect" {
				fmt.Fprintf(dst, "\n\n// Has%s returns true if the value for %s has been", exportedName, unexportedName)
				fmt.Fprintf(dst, "\n// explicitly specified")
				fmt.Fprintf(dst, "\nfunc (v *%s) Has%s() bool {", structname, exportedName)
				fmt.Fprintf(dst, "\nreturn v.%s != nil", unexportedName)
				fmt.Fprintf(dst, "\n}")

				fmt.Fprintf(dst, "\n\n// %s returns the value of %s. If the value has not", exportedName, unexportedName)
				fmt.Fprintf(dst, "\n// been explicitly, set, the zero value will be returned")
				fmt.Fprintf(dst, "\nfunc (v *%s) %s() %s {", structname, exportedName, typname(fv.Type.Elem()))
				fmt.Fprintf(dst, "\nif !v.Has%s() {", exportedName)
				fmt.Fprintf(dst, "\nreturn %v", reflect.Zero(fv.Type.Elem()).Interface())
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\nreturn *v.%s", unexportedName)
				fmt.Fprintf(dst, "\n}")
			} else {
				fmt.Fprintf(dst, "\n\nfunc (v *%s) %s() %s {", structname, exportedName, typname(fv.Type))
				fmt.Fprintf(dst, "\nreturn v.%s", unexportedName)
				fmt.Fprintf(dst, "\n}")
			}
		}
	}

	fmt.Fprintf(dst, "\n\n// Reference returns the value of `$ref` field")
	fmt.Fprintf(dst, "\nfunc (v *%s) Reference() string {", structname)
	fmt.Fprintf(dst, "\nreturn v.reference")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n\nfunc (v *%s) IsUnresolved() bool {", structname)
	fmt.Fprintf(dst, "\nreturn v.reference != \"\" && !v.resolved")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Extension returns the value of an arbitrary extension")
	fmt.Fprintf(dst, "\nfunc (v *%s) Extension(key string) (interface{}, bool) {", structname)
	fmt.Fprintf(dst, "\ne, ok := v.extensions[key]")
	fmt.Fprintf(dst, "\nreturn e, ok")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Extensions return an iterator to iterate over all extensions")
	fmt.Fprintf(dst, "\nfunc (v *%s) Extensions() *ExtensionsIterator {", structname)
	fmt.Fprintf(dst, "\nvar items []interface{}")
	fmt.Fprintf(dst, "\nfor key, item := range v.extensions {")
	fmt.Fprintf(dst, "\nitems = append(items, &mapIteratorItem{key: key, item: item})")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nvar iter ExtensionsIterator")
	fmt.Fprintf(dst, "\niter.list.size = len(items)")
	fmt.Fprintf(dst, "\niter.list.items = items")
	fmt.Fprintf(dst, "\nreturn &iter")
	fmt.Fprintf(dst, "\n}")

	if isStockValidator(structname) {
		fmt.Fprintf(dst, "\n\nfunc (v *%s) Validate(recurse bool) error {", structname)
		fmt.Fprintf(dst, "\nreturn newValidator(recurse).Validate(context.Background(), v)")
		fmt.Fprintf(dst, "\n}")
	}

	if err := codegen.WriteFormattedToFile(filename, buf.Bytes()); err != nil {
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func generateJSONHandlersFromEntity(e interface{}) error {
	rv := reflect.ValueOf(e)

	filename := fmt.Sprintf("%s_json_gen.go", strcase.ToSnake(rv.Type().Name()))
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf
	codegen.WritePreamble(dst, packageName)

	ifacename := codegen.ExportedName(rv.Type().Name())

	writeImports(dst, []string{"encoding/json", "fmt", "log", "strings", "strconv", "github.com/pkg/errors"})
	switch rv.Type().Name() {
	case "paths", "responses", "securityRequirement":
	default:
		mpname := rv.Type().Name() + "MarshalProxy"

		fmt.Fprintf(dst, "\n\ntype %s struct {", mpname)
		fmt.Fprintf(dst, "\nReference string `json:\"$ref,omitempty\"`")
		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Type().Field(i)
			if fv.Tag.Get("json") == "-" {
				continue
			}

			fieldType := fv.Type.Name()
			if fieldType == "" {
				fieldType = typname(fv.Type)
			}

			fmt.Fprintf(dst, "\n%s %s `json:\"%s\"`", codegen.ExportedName(fv.Name), fieldType, fv.Tag.Get("json"))
		}
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\nfunc (v *%s) MarshalJSON() ([]byte, error) {", rv.Type().Name())
		fmt.Fprintf(dst, "\nvar proxy %s", mpname)

		fmt.Fprintf(dst, "\nif s := v.reference; len(s) > 0 {")
		fmt.Fprintf(dst, "\nreturn []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil")
		fmt.Fprintf(dst, "\n}")

		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Type().Field(i)
			if fv.Name == "reference" || fv.Tag.Get("json") == "-" {
				continue
			}
			fmt.Fprintf(dst, "\nproxy.%s = v.%s", codegen.ExportedName(fv.Name), codegen.UnexportedName(fv.Name))
		}
		fmt.Fprintf(dst, "\nbuf, err := json.Marshal(proxy)")
		fmt.Fprintf(dst, "\nif err != nil {")
		fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `failed to marshal struct`)")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nif len(v.extensions) > 0 {")
		fmt.Fprintf(dst, "\nextBuf, err := json.Marshal(v.extensions)")
		fmt.Fprintf(dst, "\nif err != nil || len(extBuf) <= 2 {")
		fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `failed to marshal struct (extensions)`)")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nbuf = append(append(buf[:len(buf)-1], ','), extBuf[1:]...)")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nreturn buf, nil")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\n// UnmarshalJSON defines how %s is deserialized from JSON", rv.Type().Name())
		fmt.Fprintf(dst, "\nfunc (v *%s) UnmarshalJSON(data []byte) error {", rv.Type().Name())
		fmt.Fprintf(dst, "\nvar proxy map[string]json.RawMessage")
		fmt.Fprintf(dst, "\nif err := json.Unmarshal(data, &proxy); err != nil {")
		fmt.Fprintf(dst, "\nreturn err")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\nif raw, ok := proxy[\"$ref\"]; ok {")
		fmt.Fprintf(dst, "\nif err := json.Unmarshal(raw, &v.reference); err != nil {")
		fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to unmarshal $ref`)")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\nmutator := Mutate%s(v)", ifacename)
		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Type().Field(i)
			if fv.Name == "reference" || fv.Tag.Get("json") == "-" {
				continue
			}

			// unexportedFieldName := codegen.UnexportedName(fv.Name)
			exportedFieldName := codegen.ExportedName(fv.Name)

			// XXX assume that we always have a json field name specified
			jsonName := fv.Tag.Get("json")
			if i := strings.IndexByte(jsonName, ','); i > -1 {
				jsonName = jsonName[:i]
			}

			mapKey := codegen.UnexportedName(fv.Name) + "MapKey"
			fmt.Fprintf(dst, "\n\nconst %s = %s", mapKey, strconv.Quote(jsonName))
			if isList(codegen.ExportedName(fv.Type.Name())) {
				fmt.Fprintf(dst, "\nif raw, ok := proxy[%s]; ok {", mapKey)
				fmt.Fprintf(dst, "\nvar decoded %s", typname(fv.Type))
				fmt.Fprintf(dst, "\nif err := json.Unmarshal(raw, &decoded); err != nil {")
				fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to unmarshal field %s`)", codegen.ExportedName(fv.Name))
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\nfor _, elem := range decoded {")
				fmt.Fprintf(dst, "\nmutator.%s(elem)", inflection.Singular(exportedFieldName))
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\ndelete(proxy, %s)", mapKey)
				fmt.Fprintf(dst, "\n}")
			} else if isMap(codegen.ExportedName(fv.Type.Name())) {
				fmt.Fprintf(dst, "\nif raw, ok := proxy[%s]; ok {", mapKey)
				fmt.Fprintf(dst, "\nvar decoded %s", typname(fv.Type))
				fmt.Fprintf(dst, "\nif err := json.Unmarshal(raw, &decoded); err != nil {")
				fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to unmarshal field %s`)", codegen.ExportedName(fv.Name))
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\nfor key, elem := range decoded {")
				fmt.Fprintf(dst, "\nmutator.%s(key, elem)", inflection.Singular(exportedFieldName))
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\ndelete(proxy, %s)", mapKey)
				fmt.Fprintf(dst, "\n}")
			} else if isEntity(codegen.UnexportedName(fv.Type.Name())) {
				fmt.Fprintf(dst, "\nif raw, ok := proxy[%s]; ok {", mapKey)
				fmt.Fprintf(dst, "\nvar decoded %s", codegen.UnexportedName(typname(fv.Type)))
				fmt.Fprintf(dst, "\nif err := json.Unmarshal(raw, &decoded); err != nil {")
				fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to unmarshal field %s`)", codegen.ExportedName(fv.Name))
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\n\nmutator.%s(&decoded)", exportedFieldName)
				fmt.Fprintf(dst, "\ndelete(proxy, %s)", mapKey)
				fmt.Fprintf(dst, "\n}")
			} else {
				fmt.Fprintf(dst, "\nif raw, ok := proxy[%s]; ok {", mapKey)
				if fv.Tag.Get("mutator") == "indirect" {
					fmt.Fprintf(dst, "\nvar decoded %s", typname(fv.Type.Elem()))
				} else {
					fmt.Fprintf(dst, "\nvar decoded %s", typname(fv.Type))
				}
				fmt.Fprintf(dst, "\nif err := json.Unmarshal(raw, &decoded); err != nil {")
				fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to unmarshal field %s`)", jsonName)
				fmt.Fprintf(dst, "\n}")

				fmt.Fprintf(dst, "\nmutator.%s(decoded)", exportedFieldName)
				fmt.Fprintf(dst, "\ndelete(proxy, %s)", mapKey)
				fmt.Fprintf(dst, "\n}")
			}
		}

		// iterate through the proxy to look for any element that starts
		// with a ^x-.
		fmt.Fprintf(dst, "\n\nfor name, raw := range proxy {")
		fmt.Fprintf(dst, "\nif strings.HasPrefix(name, `x-`) {")
		fmt.Fprintf(dst, "\nvar ext interface{}")
		fmt.Fprintf(dst, "\nif err := json.Unmarshal(raw, &ext); err != nil {")
		fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to unmarshal field %%s`, name)")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nmutator.Extension(name, ext)")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\nif err := mutator.Apply(); err != nil {")
		fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to  unmarshal JSON`)")
		fmt.Fprintf(dst, "\n}")

		if _, ok := postUnmarshalJSONHooks[rv.Type().Name()]; ok {
			fmt.Fprintf(dst, "\n\nv.postUnmarshalJSON()")
		}
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")
	}

	/*
		fmt.Fprintf(dst, "\n\nfunc (v *%s) Resolve(resolver Resolver) error {", rv.Type().Name())
		fmt.Fprintf(dst, "\nif v.IsUnresolved() {")
		fmt.Fprintf(dst, "\n\nresolved, err := resolver.Resolve(v.Reference())")
		fmt.Fprintf(dst, "\nif err != nil {")
		fmt.Fprintf(dst, "\nif re, ok := err.(ResolveError); !ok || ok && re.Fatal() {")
		fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to resolve reference %%s`, v.Reference())")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nif resolved != nil { // can happen if !re.Fatal()")
		fmt.Fprintf(dst, "\nasserted, ok := resolved.(*%s)", rv.Type().Name())
		fmt.Fprintf(dst, "\nif !ok {")
		fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `expected resolved reference to be of type %s, but got %%T`, resolved)", codegen.ExportedName(rv.Type().Name()))
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nmutator := Mutate%s(v)", codegen.ExportedName(rv.Type().Name()))
		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Type().Field(i)
			if fv.Tag.Get("resolve") == "-" {
				continue
			}

			// If this is a container type, it has a corresponding iterator.
			// Use the iterator to assign new values
			exported := codegen.ExportedName(fv.Name)
			fieldType := fv.Type.Name()
			switch {
			default:
				fmt.Fprintf(dst, "\nmutator.%s(asserted.%s())", exported, exported)
			case isContainer(fieldType):
				switch {
				case isMap(fieldType):
					fmt.Fprintf(dst, "\nfor iter := asserted.%s(); iter.Next(); {", exported)
					fmt.Fprintf(dst, "\nkey, item := iter.Item()")
					fmt.Fprintf(dst, "\nmutator.%s(key, item)", inflection.Singular(exported))
					fmt.Fprintf(dst, "\n}")
				case isList(fieldType):
					fmt.Fprintf(dst, "\nfor iter := asserted.%s(); iter.Next(); {", exported)
					fmt.Fprintf(dst, "\nitem := iter.Item()")
					fmt.Fprintf(dst, "\nmutator.%s(item)", inflection.Singular(exported))
					fmt.Fprintf(dst, "\n}")
				default:
					return errors.Errorf(`unknown cotainer %s`, exported)
				}
			}
		}
		fmt.Fprintf(dst, "\nif err := mutator.Apply(); err != nil {")
		fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to mutate`)")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nv.resolved = true")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\n}")

		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Type().Field(i)
			if fv.Name == "reference" || fv.Tag.Get("resolve") == "-" {
				continue
			}

			// if it's an entity, or a container for entity, resolve
			var resolve bool
			if _, ok := entityTypes[codegen.UnexportedName(fv.Type.Name())]; ok {
				resolve = true
			} else if _, ok := containerTypes[fv.Type.Name()]; ok {
				resolve = true
			}

			if resolve {
				fmt.Fprintf(dst, "\nif v.%s != nil {", codegen.UnexportedName(fv.Name))
				fmt.Fprintf(dst, "\nif err := v.%s.Resolve(resolver); err != nil {", codegen.UnexportedName(fv.Name))
				fmt.Fprintf(dst, "\nif re, ok := err.(ResolveError); !ok || ok && re.Fatal() {")
				fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to resolve %s`)", codegen.ExportedName(fv.Name))
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\n}")
			}
		}
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")
	*/

	fmt.Fprintf(dst, "\n\nfunc (v *%s) QueryJSON(path string) (ret interface{}, ok bool)  {", rv.Type().Name())
	fmt.Fprintf(dst, "\npath = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })")
	fmt.Fprintf(dst, "\nif path == \"\" {")
	fmt.Fprintf(dst, "\nreturn v, true")
	fmt.Fprintf(dst, "\n}")

	var fragFields []reflect.StructField
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Type().Field(i)
		if fv.Name == "reference" {
			continue
		}
		jsname := fv.Tag.Get("json")
		if jsname == "-" {
			continue
		}
		fragFields = append(fragFields, fv)
	}

	if len(fragFields) > 0 {
		fmt.Fprintf(dst, "\n\nvar frag string")
		fmt.Fprintf(dst, "\nif i := strings.Index(path, \"/\"); i > -1 {")
		fmt.Fprintf(dst, "\nfrag = path[:i]")
		fmt.Fprintf(dst, "\npath = path[i+1:]")
		fmt.Fprintf(dst, "\n} else {")
		fmt.Fprintf(dst, "\nfrag = path")
		fmt.Fprintf(dst, "\npath = \"\"")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\n\nvar target interface{}")
		fmt.Fprintf(dst, "\n\nswitch frag {")
		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Type().Field(i)
			if fv.Name == "reference" {
				continue
			}
			jsname := fv.Tag.Get("json")
			if jsname == "-" {
				continue
			}
			if i := strings.Index(jsname, ","); i > -1 {
				jsname = jsname[:i]
			}

			if jsname == "" {
				jsname = fv.Name
			}

			fmt.Fprintf(dst, "\ncase %s:", strconv.Quote(jsname))
			fmt.Fprintf(dst, "\ntarget = v.%s", fv.Name)
		}
		fmt.Fprintf(dst, "\ndefault:")
		fmt.Fprintf(dst, "\nreturn nil, false")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\nif qj, ok := target.(QueryJSONer); ok {")
		fmt.Fprintf(dst, "\nreturn qj.QueryJSON(path)")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nif path == \"\" {")
		fmt.Fprintf(dst, "\nreturn target, true")
		fmt.Fprintf(dst, "\n}")
	}
	fmt.Fprintf(dst, "\nreturn nil, false")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// %[1]sFromJSON constructs a %[1]s from JSON buffer. `dst` must", ifacename)
	fmt.Fprintf(dst, "\n// be a pointer to `%s`", ifacename)
	fmt.Fprintf(dst, "\nfunc %[1]sFromJSON(buf []byte, dst interface{}) error {", ifacename)
	fmt.Fprintf(dst, "\nv, ok := dst.(*%s)", ifacename)
	fmt.Fprintf(dst, "\nif !ok {")
	fmt.Fprintf(dst, "\nreturn errors.Errorf(`dst needs to be a pointer to %s, but got %%T`, dst)", ifacename)
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nvar tmp %s", codegen.UnexportedName(rv.Type().Name()))
	fmt.Fprintf(dst, "\nif err := json.Unmarshal(buf, &tmp); err != nil {")
	fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to unmarshal %s`)", ifacename)
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n*v = &tmp")
	fmt.Fprintf(dst, "\nreturn nil")
	fmt.Fprintf(dst, "\n}")

	if err := codegen.WriteFormattedToFile(filename, buf.Bytes()); err != nil {
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func generateMutatorFromEntity(e interface{}) error {
	rv := reflect.ValueOf(e)
	filename := fmt.Sprintf("%s_mutator_gen.go", strcase.ToSnake(rv.Type().Name()))
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, packageName)
	writeImports(dst, []string{"sync"})

	ifacename := codegen.ExportedName(rv.Type().Name())
	structname := rv.Type().Name()

	fmt.Fprintf(dst, "\n\n// %sMutator is used to build an instance of %s. The user must", ifacename, ifacename)
	fmt.Fprintf(dst, "\n// call `Apply()` after providing all the necessary information to")
	fmt.Fprintf(dst, "\n// the new instance of %s with new values", ifacename)
	fmt.Fprintf(dst, "\ntype %sMutator struct {", ifacename)
	fmt.Fprintf(dst, "\nlock sync.Locker")
	fmt.Fprintf(dst, "\nproxy *%s", structname)
	fmt.Fprintf(dst, "\ntarget *%s", structname)
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Apply finalizes the matuation process for %s and returns the result", ifacename)
	fmt.Fprintf(dst, "\nfunc (m *%sMutator) Apply() error {", ifacename)
	// TODO: validation
	fmt.Fprintf(dst, "\nm.lock.Lock()")
	fmt.Fprintf(dst, "\ndefer m.lock.Unlock()")
	fmt.Fprintf(dst, "\n*m.target = *m.proxy")
	fmt.Fprintf(dst, "\nreturn nil")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Mutate%s creates a new mutator object for %s", ifacename, ifacename)
	fmt.Fprintf(dst, "\n// Operations on the mutator are safe to be used concurrently, except for")
	fmt.Fprintf(dst, "\n// when calling `Apply()`, where the user is responsible for restricting access")
	fmt.Fprintf(dst, "\n// to the target object to be mutated")
	fmt.Fprintf(dst, "\nfunc Mutate%s(v %s, options ...Option) *%sMutator {", ifacename, ifacename, ifacename)
	fmt.Fprintf(dst, "\nvar lock sync.Locker = &sync.Mutex{}")
	fmt.Fprintf(dst, "\nfor _, option := range options {")
	fmt.Fprintf(dst, "\nswitch option.Name() {")
	fmt.Fprintf(dst, "\ncase optkeyLocker:")
	fmt.Fprintf(dst, "\nlock = option.Value().(sync.Locker)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nif lock == nil {")
	fmt.Fprintf(dst, "\nlock = nilLock{}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nreturn &%sMutator{", ifacename)
	fmt.Fprintf(dst, "\nlock: lock,")
	fmt.Fprintf(dst, "\ntarget: v.(*%s),", structname)
	fmt.Fprintf(dst, "\nproxy: v.Clone().(*%s),", structname)
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Type().Field(i)
		if fv.Tag.Get("mutator") == "-" {
			continue
		}

		exportedName := codegen.ExportedName(fv.Name)
		unexportedName := codegen.UnexportedName(fv.Name)
		fieldType := fv.Type.Name()
		switch {
		case isMap(fieldType):
			log.Printf(" * Generating map element mutator for %s", fieldType)
			fmt.Fprintf(dst, "\n\n// Clear%s removes all values in %s field", exportedName, unexportedName)
			fmt.Fprintf(dst, "\nfunc (m *%[1]sMutator) Clear%[2]s() *%[1]sMutator {", ifacename, exportedName)
			fmt.Fprintf(dst, "\nm.lock.Lock()")
			fmt.Fprintf(dst, "\ndefer m.lock.Unlock()")
			fmt.Fprintf(dst, "\n_ = m.proxy.%s.Clear()", unexportedName)
			fmt.Fprintf(dst, "\nreturn m")
			fmt.Fprintf(dst, "\n}")

			fmt.Fprintf(dst, "\n\n// %s sets the value of %s", inflection.Singular(exportedName), unexportedName)
			fmt.Fprintf(dst, "\nfunc (m *%sMutator) %s(key %sKey, value %s) *%sMutator {", ifacename, inflection.Singular(exportedName), fieldType, typname(fv.Type.Elem()), ifacename)
			fmt.Fprintf(dst, "\nm.lock.Lock()")
			fmt.Fprintf(dst, "\ndefer m.lock.Unlock()")
			fmt.Fprintf(dst, "\nif m.proxy.%s == nil {", unexportedName)
			fmt.Fprintf(dst, "\nm.proxy.%s = %s{}", unexportedName, fieldType)
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n\nm.proxy.%s[key] = value", unexportedName)
			if isEntity(fv.Type.Elem().Name()) {
				fmt.Fprintf(dst, ".Clone()")
			}
			fmt.Fprintf(dst, "\nreturn m")
			fmt.Fprintf(dst, "\n}")
		case isList(fieldType):
			log.Printf(" * Generating list element mutator for %s", fieldType)
			fmt.Fprintf(dst, "\n\n// Clear%s clears all elements in %s", exportedName, unexportedName)
			fmt.Fprintf(dst, "\nfunc (m *%sMutator) Clear%s() *%sMutator {", ifacename, exportedName, ifacename)
			fmt.Fprintf(dst, "\nm.lock.Lock()")
			fmt.Fprintf(dst, "\ndefer m.lock.Unlock()")
			fmt.Fprintf(dst, "\n_ = m.proxy.%s.Clear()", unexportedName)
			fmt.Fprintf(dst, "\nreturn m")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n\n// %s appends a value to %s", inflection.Singular(exportedName), unexportedName)
			fmt.Fprintf(dst, "\nfunc (m *%sMutator) %s(value %s) *%sMutator {", ifacename, inflection.Singular(exportedName), typname(fv.Type.Elem()), ifacename)
			fmt.Fprintf(dst, "\nm.lock.Lock()")
			fmt.Fprintf(dst, "\ndefer m.lock.Unlock()")
			fmt.Fprintf(dst, "\nm.proxy.%s = append(m.proxy.%s, value)", unexportedName, unexportedName)
			fmt.Fprintf(dst, "\nreturn m")
			fmt.Fprintf(dst, "\n}")
		default:
			if fv.Tag.Get("mutator") == "indirect" {
				fmt.Fprintf(dst, "\n\n// Clear%s clears the %s field", exportedName, unexportedName)
				fmt.Fprintf(dst, "\nfunc (m *%sMutator) Clear%s() *%sMutator {", ifacename, exportedName, ifacename)
				fmt.Fprintf(dst, "\nm.lock.Lock()")
				fmt.Fprintf(dst, "\ndefer m.lock.Unlock()")
				fmt.Fprintf(dst, "\nm.proxy.%s = nil", unexportedName)
				fmt.Fprintf(dst, "\nreturn m")
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\n\n// %s sets the %s field.", exportedName, unexportedName)
				fmt.Fprintf(dst, "\nfunc (m *%sMutator) %s(v %s) *%sMutator {", ifacename, exportedName, typname(fv.Type.Elem()), ifacename)
				fmt.Fprintf(dst, "\nm.proxy.%s = &v", unexportedName)
				fmt.Fprintf(dst, "\nreturn m")
				fmt.Fprintf(dst, "\n}")
			} else {

				fmt.Fprintf(dst, "\n\n// %s sets the %s field for object %s.", exportedName, exportedName, ifacename)
				fmt.Fprintf(dst, "\nfunc (m *%sMutator) %s(v %s) *%sMutator {", ifacename, exportedName, typname(fv.Type), ifacename)
				fmt.Fprintf(dst, "\nm.lock.Lock()")
				fmt.Fprintf(dst, "\ndefer m.lock.Unlock()")
				fmt.Fprintf(dst, "\nm.proxy.%s = v", unexportedName)
				fmt.Fprintf(dst, "\nreturn m")
				fmt.Fprintf(dst, "\n}")
			}
		}
	}

	fmt.Fprintf(dst, "\n\n// Extension sets an arbitrary extension field in %s", ifacename)
	fmt.Fprintf(dst, "\nfunc (m *%sMutator) Extension(name string, value interface{}) *%sMutator {", ifacename, ifacename)
	fmt.Fprintf(dst, "\nif m.proxy.extensions == nil {")
	fmt.Fprintf(dst, "\nm.proxy.extensions = Extensions{}")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nm.proxy.extensions[name] = value")
	fmt.Fprintf(dst, "\nreturn m")
	fmt.Fprintf(dst, "\n}")
	if err := codegen.WriteFormattedToFile(filename, buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, &buf)
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func generateContainer(c interface{}) error {
	rv := reflect.ValueOf(c)

	filename := fmt.Sprintf("%s_gen.go", strcase.ToSnake(rv.Type().Name()))
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, packageName)
	codegen.WriteImports(dst, "context", "encoding/json", "github.com/pkg/errors")

	typeName := rv.Type().Name()
	switch {
	case isList(typeName):
		fmt.Fprintf(dst, "\n\n// Clear removes all values from %s", typeName)
		fmt.Fprintf(dst, "\nfunc (v *%s) Clear() error {", typeName)
		fmt.Fprintf(dst, "\n*v = %s(nil)", rv.Type().Name())
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\n// Validate checks for the values for correctness. If `recurse`")
		fmt.Fprintf(dst, "\n// is specified, child elements are also validated")
		fmt.Fprintf(dst, "\nfunc (v *%s) Validate(recurse bool) error {", typeName)
		elemType := rv.Type().Elem()
		if elemType.Kind() == reflect.Interface {
			fmt.Fprintf(dst, "\nfor i, elem := range *v {")
			fmt.Fprintf(dst, "\nif validator, ok := elem.(Validator); ok {")
			fmt.Fprintf(dst, "\nif err := validator.Validate(recurse); err != nil {")
			fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to validate element %%d`, i)")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n}")
		} else if isEntity(codegen.UnexportedName(elemType.Name())) || isContainer(elemType.Name()) {
			fmt.Fprintf(dst, "\nfor i, elem := range *v {")
			fmt.Fprintf(dst, "\nif err := elem.Validate(recurse); err != nil {")
			fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to validate element %%d`, i)")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n}")
		}
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")

		if rv.Type().Elem().Name() != "" && rv.Type().Elem().Kind() == reflect.Interface {
			fmt.Fprintf(dst, "\n\n// UnmarshalJSON defines how %s is deserialized from JSON", typeName)
			fmt.Fprintf(dst, "\nfunc (v *%s) UnmarshalJSON(data []byte) error {", typeName)
			fmt.Fprintf(dst, "\nvar proxy []*%s", codegen.UnexportedName(typname(rv.Type().Elem())))
			fmt.Fprintf(dst, "\nif err := json.Unmarshal(data, &proxy); err != nil {")
			fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to unmarshal`)")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n\nif len(proxy) == 0 {")
			fmt.Fprintf(dst, "\n*v = %s(nil)", typeName)
			fmt.Fprintf(dst, "\nreturn nil")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n\ntmp := make(%s, len(proxy))", typeName)
			fmt.Fprintf(dst, "\nfor i, value := range proxy {")
			fmt.Fprintf(dst, "\ntmp[i] = value")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n*v = tmp")
			fmt.Fprintf(dst, "\nreturn nil")
			fmt.Fprintf(dst, "\n}")
		}
	case isMap(typeName):
		fmt.Fprintf(dst, "\n\n// Clear removes all elements from %s", typeName)
		fmt.Fprintf(dst, "\nfunc (v *%s) Clear() error {", typeName)
		fmt.Fprintf(dst, "\n*v = make(%s)", typeName)
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\n// Validate checks the correctness of values in %s", typeName)
		fmt.Fprintf(dst, "\nfunc (v *%s) Validate(recurse bool) error {", typeName)
		fmt.Fprintf(dst, "\nreturn Visit(context.Background(), newValidator(recurse), v)")
		fmt.Fprintf(dst, "\n}")
		/*
			elemType := rv.Type().Elem()
			if elemType.Kind() == reflect.Interface {
				fmt.Fprintf(dst, "\nfor name, elem := range *v {")
				fmt.Fprintf(dst, "\nif validator, ok := elem.(Validator); ok {")
				fmt.Fprintf(dst, "\nif err := validator.Validate(recurse); err != nil {")
				fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to validate element %%v`, name)")
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\n}")
			} else if isEntity(codegen.UnexportedName(elemType.Name())) || isContainer(elemType.Name()) {
				fmt.Fprintf(dst, "\nfor name, elem := range *v {")
				fmt.Fprintf(dst, "\nif err := elem.Validate(recurse); err != nil {")
				fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to validate element %%v`, name)")
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\n}")
			}
			fmt.Fprintf(dst, "\nreturn nil")
			fmt.Fprintf(dst, "\n}")
			/*
				fmt.Fprintf(dst, "\n\nfunc (v %s) Resolve(resolver Resolver) error {", rv.Type().Name())
				if _, ok := entityTypes[codegen.UnexportedName(rv.Type().Elem().Name())]; ok {
					fmt.Fprintf(dst, "\nif len(v) > 0 {")
					fmt.Fprintf(dst, "\nfor name, elem := range v {")
					fmt.Fprintf(dst, "\nif err := elem.Resolve(resolver); err != nil {")
					fmt.Fprintf(dst, "\nif re, ok := err.(ResolveError); !ok || ok && re.Fatal() {")
					fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to resolve %s (key = %%s)`, name)", rv.Type().Name())
					fmt.Fprintf(dst, "\n}")
					fmt.Fprintf(dst, "\n}")
					fmt.Fprintf(dst, "\n}")
					fmt.Fprintf(dst, "\n}")
				}
				fmt.Fprintf(dst, "\nreturn nil")
				fmt.Fprintf(dst, "\n}")
		*/

		fmt.Fprintf(dst, "\n\n// QueryJSON is used to query an element within the document")
		fmt.Fprintf(dst, "\n// Using jsonref")
		fmt.Fprintf(dst, "\nfunc (v %s) QueryJSON(path string) (ret interface{}, ok bool) {", rv.Type().Name())
		fmt.Fprintf(dst, "\nif path == `` {")
		fmt.Fprintf(dst, "\nreturn v, true")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\nvar frag string")
		fmt.Fprintf(dst, "\nfrag, path = extractFragFromPath(path)")
		fmt.Fprintf(dst, "\ntarget, ok := v[frag]")
		fmt.Fprintf(dst, "\nif !ok {")
		fmt.Fprintf(dst, "\nreturn nil, false")
		fmt.Fprintf(dst, "\n}")

		if rv.Type().Elem().Kind() == reflect.Interface {
			fmt.Fprintf(dst, "\n\nif qj, ok := target.(QueryJSONer); ok {")
			fmt.Fprintf(dst, "\nreturn qj.QueryJSON(path)")
			fmt.Fprintf(dst, "\n}")
		}

		fmt.Fprintf(dst, "\n\nif path == `` {")
		fmt.Fprintf(dst, "\nreturn target, true")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nreturn nil, false")
		fmt.Fprintf(dst, "\n}")

		if rv.Type().Elem().Name() != "" && rv.Type().Elem().Kind() == reflect.Interface {
			fmt.Fprintf(dst, "\n\n// UnmarshalJSON takes a JSON buffer and properly populates `v`")
			fmt.Fprintf(dst, "\nfunc (v *%s) UnmarshalJSON(data []byte) error {", typeName)
			fmt.Fprintf(dst, "\nvar proxy map[%s]*%s", typname(rv.Type().Key()), codegen.UnexportedName(typname(rv.Type().Elem())))
			fmt.Fprintf(dst, "\nif err := json.Unmarshal(data, &proxy); err != nil {")
			fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to unmarshal`)")
			fmt.Fprintf(dst, "\n}")

			fmt.Fprintf(dst, "\ntmp := make(map[%s]%s)", typname(rv.Type().Key()), typname(rv.Type().Elem()))
			fmt.Fprintf(dst, "\nfor name, value := range proxy {")
			switch typeName {
			case "ParameterMap", "SecuritySchemeMap":
			default:
				// assume they all have setName(string)
				if isEntity(codegen.UnexportedName(rv.Type().Elem().Name())) {
					fmt.Fprintf(dst, "\nvalue.setName(name)")
				}
			}
			fmt.Fprintf(dst, "\ntmp[name] = value")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n*v = tmp")
			fmt.Fprintf(dst, "\nreturn nil")
			fmt.Fprintf(dst, "\n}")
		}
	}

	return codegen.WriteFormattedToFile(filename, buf.Bytes())
}

func generateVisitorsFromEntity(e interface{}) error {
	tv := reflect.TypeOf(e)

	filename := fmt.Sprintf("%s_visitor_gen.go", stringutil.Snake(tv.Name()))
	log.Printf("Generating %s", filename)

	ifacename := codegen.ExportedName(tv.Name())
	structname := codegen.UnexportedName(tv.Name())
	ctxKey := structname + "VisitorCtxKey"

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, packageName)
	codegen.WriteImports(dst, "context", "github.com/pkg/errors")

	fmt.Fprintf(dst, "\n\n// %sVisitor is an interface for objects that knows", ifacename)
	fmt.Fprintf(dst, "\n// how to process %s elements while traversing the OpenAPI structure", ifacename)
	fmt.Fprintf(dst, "\ntype %sVisitor interface {", ifacename)
	fmt.Fprintf(dst, "\nVisit%[1]s(context.Context, %[1]s) error", ifacename)
	fmt.Fprintf(dst, "\n}")

	type entityField struct {
		Name string
		Type reflect.Type
	}
	var entityFields []entityField

	// If it's a struct, find the element types from the StructField
	switch tv.Kind() {
	case reflect.Struct:
		for i := 0; i < tv.NumField(); i++ {
			fv := tv.Field(i)
			fieldType := fv.Type.Name()

			// keep track of all fields whose type is one of our entity types
			if fv.Tag.Get("visit") == "true" || fv.Tag.Get("json") != "-" {
				if isEntity(codegen.UnexportedName(fieldType)) || isContainer(fieldType) {
					entityFields = append(entityFields, entityField{Type: fv.Type, Name: fv.Name})
				}
			}
		}
	default:
		entityFields = append(entityFields, entityField{Type: tv, Name: strings.TrimSuffix(tv.Name(), "List") + "s"})
	}

	switch ifacename {
	case "Paths":
	default:
		fmt.Fprintf(dst, "\n\nfunc visit%[1]s(ctx context.Context, elem %[1]s) error {", ifacename)
		fmt.Fprintf(dst, "\nif checker, ok := elem.(interface { IsValid() bool }); ok {")
		fmt.Fprintf(dst, "\nif !checker.IsValid() {")
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nselect {")
		fmt.Fprintf(dst, "\ncase <-ctx.Done():")
		fmt.Fprintf(dst, "\nreturn ctx.Err()")
		fmt.Fprintf(dst, "\ndefault:")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\nif v, ok := ctx.Value(%s{}).(%sVisitor); ok {", ctxKey, ifacename)
		fmt.Fprintf(dst, "\nif err := v.Visit%s(ctx, elem); err != nil {", ifacename)
		fmt.Fprintf(dst, "\nif err == ErrVisitAbort {")
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to visit %s element`)", ifacename)
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\n}")
		for _, f := range entityFields {
			if isMap(f.Type.Name()) {
				// skip things like map[string]string
				if isEntity(codegen.UnexportedName(f.Type.Elem().Name())) {
					fmt.Fprintf(dst, "\n\nfor iter := elem.%s(); iter.Next(); {", codegen.ExportedName(f.Name))
					fmt.Fprintf(dst, "\nkey, value := iter.Item()")
					fmt.Fprintf(dst, "\nif err := visit%s(context.WithValue(ctx, %sKeyVisitorCtxKey{}, key), value); err != nil {", codegen.ExportedName(f.Type.Elem().Name()), codegen.UnexportedName(f.Type.Name()))
					fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to visit %s element for %s`)", codegen.ExportedName(f.Name), ifacename)
					fmt.Fprintf(dst, "\n}")
					fmt.Fprintf(dst, "\n}")
				}
			} else if isList(f.Type.Name()) {
				// skip things like []string
				if isEntity(codegen.UnexportedName(f.Type.Elem().Name())) {
					fmt.Fprintf(dst, "\n\nfor i, iter := 0, elem.%s(); iter.Next(); {", codegen.ExportedName(f.Name))
					fmt.Fprintf(dst, "\nif err := visit%s(ctx, iter.Item()); err != nil {", codegen.ExportedName(f.Type.Elem().Name()))
					fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to visit element %%d for %s`, i)", ifacename)
					fmt.Fprintf(dst, "\n}")
					fmt.Fprintf(dst, "\ni++")
					fmt.Fprintf(dst, "\n}")
				}
			} else {
				fmt.Fprintf(dst, "\n\nif child := elem.%s(); child != nil {", codegen.ExportedName(f.Name))
				fmt.Fprintf(dst, "\nif err := visit%s(ctx, child); err != nil {", codegen.ExportedName(f.Type.Name()))
				fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to visit %s element for %s`)", codegen.ExportedName(f.Name), ifacename)
				fmt.Fprintf(dst, "\n}")
				fmt.Fprintf(dst, "\n}")
			}
		}
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")
	}

	if err := codegen.WriteFormattedToFile(filename, buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, &buf)
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func generateVisitor(entities []interface{}) error {
	filename := fmt.Sprintf("visitor_gen.go")
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, packageName)
	codegen.WriteImports(dst, "context", "github.com/pkg/errors")

	fmt.Fprintf(dst, "\n\nvar ErrVisitAbort = errors.New(`visit aborted (non-error)`)")

	for _, c := range containers {
		tv := reflect.TypeOf(c)
		if isMap(tv.Name()) {
			structname := codegen.UnexportedName(tv.Name())
			fmt.Fprintf(dst, "\n\ntype %sKeyVisitorCtxKey struct{}", structname)
		}
	}

	for _, e := range entities {
		tv := reflect.TypeOf(e)
		structname := codegen.UnexportedName(tv.Name())
		fmt.Fprintf(dst, "\n\ntype %sVisitorCtxKey struct{}", structname)
	}

	fmt.Fprintf(dst, "\n\n// Visit allows you to traverse through the OpenAPI structure")
	fmt.Fprintf(dst, "\nfunc Visit(ctx context.Context, handler, elem interface{}) error {")

	for i, e := range entities {
		tv := reflect.TypeOf(e)
		ifacename := codegen.ExportedName(tv.Name())
		structname := codegen.UnexportedName(tv.Name())
		if i > 0 {
			fmt.Fprintf(dst, "\n")
		}
		fmt.Fprintf(dst, "\nif v, ok := handler.(%sVisitor); ok {", ifacename)
		fmt.Fprintf(dst, "\nctx = context.WithValue(ctx, %sVisitorCtxKey{}, v)", structname)
		fmt.Fprintf(dst, "\n}")
	}
	fmt.Fprintf(dst, "\n\nreturn visit(ctx, elem)")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc visit(ctx context.Context, elem interface{}) error {")
	fmt.Fprintf(dst, "\nswitch elem := elem.(type) {")
	for _, e := range entities {
		tv := reflect.TypeOf(e)
		ifacename := codegen.ExportedName(tv.Name())
		fmt.Fprintf(dst, "\ncase %s:", ifacename)
		fmt.Fprintf(dst, "\nreturn visit%s(ctx, elem)", ifacename)
	}
	fmt.Fprintf(dst, "\ndefault:")
	fmt.Fprintf(dst, "\nreturn errors.Errorf(`unknown element %%T`, elem)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	if err := codegen.WriteFormattedToFile(filename, buf.Bytes()); err != nil {
		codegen.DumpCode(os.Stdout, &buf)
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}
