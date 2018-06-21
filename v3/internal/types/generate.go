package types

import (
	"bufio"
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

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	codegen "github.com/lestrrat-go/openapi/internal/codegen/golang"
	"github.com/lestrrat-go/openapi/internal/stringutil"
	"github.com/pkg/errors"
)

var entityTypes = map[string]interface{}{}
var containerTypes = map[string]interface{}{}
var postUnmarshalJSONHooks = map[string]struct{}{
	"SchemaMap":   struct{}{},
	"pathItem":    struct{}{},
	"requestBody": struct{}{},
}

func GenerateCode() error {
	entities := []interface{}{
		callback{},
		components{},
		contact{},
		discriminator{},
		encoding{},
		example{},
		externalDocumentation{},
		header{},
		info{},
		license{},
		link{},
		mediaType{},
		oauthFlow{},
		oauthFlows{},
		openAPI{},
		operation{},
		parameter{},
		pathItem{},
		paths{},
		requestBody{},
		response{},
		responses{},
		schema{},
		securityRequirement{},
		securityScheme{},
		server{},
		serverVariable{},
		tag{},
	}

	containers := []interface{}{
		CallbackMap{},
		EncodingMap{},
		ExampleMap{},
		HeaderMap{},
		InterfaceList{},
		InterfaceMap{},
		LinkMap{},
		MediaTypeMap{},
		PathItemMap{},
		ParameterList{},
		ParameterMap{},
		RequestBodyMap{},
		ResponseMap{},
		SchemaList{},
		SchemaMap{},
		ScopeMap{},
		SecurityRequirementList{},
		SecuritySchemeMap{},
		ServerList{},
		ServerVariableMap{},
		StringList{},
		StringMap{},
		StringListMap{},
		TagList{},
	}

	for _, e := range entities {
		name := reflect.ValueOf(e).Type().Name()
		entityTypes[name] = e
	}

	for _, c := range containers {
		name := reflect.ValueOf(c).Type().Name()
		containerTypes[name] = c
	}

	// Copy the interfaces file after swapping the package name
	copyInterface()

	for _, e := range entities {
		if err := generateJSONHandlersFromEntity(e); err != nil {
			return errors.Wrap(err, `failed to generate JSON handlers from entity`)
		}

		if err := generateAccessorsFromEntity(e); err != nil {
			return errors.Wrap(err, `failed to generate accessors from entity`)
		}

		if err := generateBuildersFromEntity(e); err != nil {
			return errors.Wrap(err, `failed to generate builders from entity`)
		}

		if err := generateMutatorsFromEntity(e); err != nil {
			return errors.Wrap(err, `failed to generate mutators from entity`)
		}
	}

	for _, c := range containers {
		if err := generateContainers(c); err != nil {
			return errors.Wrap(err, `failed to generate containers`)
		}
	}

	if err := generateClonersFromEntity(entities); err != nil {
		return errors.Wrap(err, `failed to generate cloners from entity list`)
	}

	if err := generateIteratorsFromEntity(containers); err != nil {
		return errors.Wrap(err, `failed to generate iterators from entity list`)
	}

	return nil
}

func typname(t reflect.Type) string {
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

	codegen.WritePreamble(dst, "openapi")

	for scanner.Scan() {
		txt := scanner.Text()

		txt = strings.Replace(txt, "//gen:lazy ", "", -1)

		fmt.Fprintf(dst, "\n%s", txt)
		if strings.HasPrefix(txt, "type ") && strings.HasSuffix(txt, " interface {") {
			i := strings.Index(txt[5:], " ")
			completeInterface(dst, txt[5:5+i])
		} else if strings.HasPrefix(txt, "type ") && strings.HasSuffix(txt, " struct {") {
			i := strings.Index(txt[5:], " ")
			if _, ok := entityTypes[txt[5:5+i]]; ok {
				fmt.Fprintf(dst, "\nreference string `json:\"$ref,omitempty\"`")
				fmt.Fprintf(dst, "\nresolved bool `json:\"-\"`")
			}
		}
	}
	if err := writeFormattedSource(&buf, "interface_gen.go"); err != nil {
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func completeInterface(dst io.Writer, ifacename string) {
	log.Printf("Completing interface for %s", ifacename)

	e, ok := entityTypes[codegen.UnexportedName(ifacename)]
	if !ok {
		panic(fmt.Sprintf("Could not find value for %s", codegen.UnexportedName(ifacename)))
	}

	rv := reflect.ValueOf(e)

	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Type().Field(i)
		if fv.Tag.Get("accessor") == "-" {
			continue
		}

		// If it's a container type, we need to return iterators
		fieldType := fv.Type.Name()
		exportedFieldName := codegen.ExportedName(fv.Name)
		if _, ok := containerTypes[fieldType]; ok {
			fmt.Fprintf(dst, "\n%s() *%s", exportedFieldName, iteratorName(fv.Type))
		} else {
			fmt.Fprintf(dst, "\n%s() %s", exportedFieldName, typname(fv.Type))
		}
	}

	fmt.Fprintf(dst, "\nMarshalJSON() ([]byte, error)")
	fmt.Fprintf(dst, "\nIsUnresolved() bool")
	fmt.Fprintf(dst, "\nResolve(*Resolver) error")
	fmt.Fprintf(dst, "\nClone() %s", ifacename)
}

func generateJSONHandlersFromEntity(e interface{}) error {
	rv := reflect.ValueOf(e)

	filename := fmt.Sprintf("%s_json_gen.go", stringutil.Snake(rv.Type().Name()))
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf
	codegen.WritePreamble(dst, "openapi")

	switch rv.Type().Name() {
	case "paths", "responses":
		codegen.WriteImports(dst, "log", "strings", "github.com/pkg/errors")
	default:
		codegen.WriteImports(dst, "log", "encoding/json", "strings", "github.com/pkg/errors")

		mpname := rv.Type().Name() + "MarshalProxy"
		upname := rv.Type().Name() + "UnmarshalProxy"

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

			exportedFieldName := codegen.ExportedName(fv.Name)
			fmt.Fprintf(dst, "\n%s %s `json:\"%s\"`", exportedFieldName, fieldType, fv.Tag.Get("json"))
		}
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\ntype %s struct {", upname)
		fmt.Fprintf(dst, "\nReference string `json:\"$ref,omitempty\"`")
		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Type().Field(i)
			if fv.Tag.Get("json") == "-" {
				continue
			}

			exportedFieldName := codegen.ExportedName(fv.Name)

			switch {
			case isContainer(fv.Type.Name()):
				fmt.Fprintf(dst, "\n%s %s `json:\"%s\"`", exportedFieldName, fv.Type.Name(), fv.Tag.Get("json"))
			case fv.Type.Kind() == reflect.Slice:
				if _, ok := entityTypes[codegen.UnexportedName(typname(fv.Type.Elem()))]; ok {
					fmt.Fprintf(dst, "\n%s []json.RawMessage `json:\"%s\"`", exportedFieldName, fv.Tag.Get("json"))
				}
			case fv.Type.Kind() == reflect.Map:
				if _, ok := entityTypes[codegen.UnexportedName(typname(fv.Type.Elem()))]; ok {
					fmt.Fprintf(dst, "\n%s map[string]json.RawMessage `json:\"%s\"`", exportedFieldName, fv.Tag.Get("json"))
				}
			default:
				if _, ok := entityTypes[codegen.UnexportedName(typname(fv.Type))]; ok {
					fmt.Fprintf(dst, "\n%s json.RawMessage `json:\"%s\"`", exportedFieldName, fv.Tag.Get("json"))
				} else {
					fmt.Fprintf(dst, "\n%s %s `json:\"%s\"`", exportedFieldName, typname(fv.Type), fv.Tag.Get("json"))
				}
			}
		}
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\nfunc (v *%s) MarshalJSON() ([]byte, error) {", rv.Type().Name())
		fmt.Fprintf(dst, "\nvar proxy %s", mpname)

		fmt.Fprintf(dst, "\nif len(v.reference) > 0 {")
		fmt.Fprintf(dst, "\nproxy.Reference = v.reference")
		fmt.Fprintf(dst, "\nreturn json.Marshal(proxy)")
		fmt.Fprintf(dst, "\n}")

		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Type().Field(i)
			if fv.Name == "reference" || fv.Tag.Get("json") == "-" {
				continue
			}
			fmt.Fprintf(dst, "\nproxy.%s = v.%s", codegen.ExportedName(fv.Name), codegen.UnexportedName(fv.Name))
		}
		fmt.Fprintf(dst, "\nreturn json.Marshal(proxy)")
		fmt.Fprintf(dst, "\n}")

		// Unmarshaling interfaces is tricky. We need to construct a concrete
		// type that fulfills the interface, and unmarshal using that.

		fmt.Fprintf(dst, "\n\nfunc (v *%s) UnmarshalJSON(data []byte) error {", rv.Type().Name())
		fmt.Fprintf(dst, "\nvar proxy %s", upname)
		fmt.Fprintf(dst, "\nif err := json.Unmarshal(data, &proxy); err != nil {")
		fmt.Fprintf(dst, "\nreturn err")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\nif len(proxy.Reference) > 0 {")
		fmt.Fprintf(dst, "\nv.reference = proxy.Reference")
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")
		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Type().Field(i)
			if fv.Name == "reference" || fv.Tag.Get("json") == "-" {
				continue
			}

			exportedFieldName := codegen.ExportedName(fv.Name)
			unexportedFieldName := codegen.UnexportedName(fv.Name)
			switch {
			case isContainer(fv.Type.Name()):
				fmt.Fprintf(dst, "\nv.%s = proxy.%s", unexportedFieldName, exportedFieldName)
			case fv.Type.Kind() == reflect.Slice:
				elemType := codegen.UnexportedName(fv.Type.Elem().Name())
				elemIface := codegen.ExportedName(fv.Type.Elem().Name())
				if _, ok := entityTypes[elemType]; ok {
					fmt.Fprintf(dst, "\n\nif len(proxy.%s) > 0 {", exportedFieldName)
					fmt.Fprintf(dst, "\nvar list []%s", elemIface)
					fmt.Fprintf(dst, "\nfor i, pv := range proxy.%s {", exportedFieldName)
					fmt.Fprintf(dst, "\nvar decoded %s", elemType)
					fmt.Fprintf(dst, "\nif err := json.Unmarshal(pv, &decoded); err != nil {")
					fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to unmasrhal element %%d of field %s`, i)", exportedFieldName)
					fmt.Fprintf(dst, "\n}")
					fmt.Fprintf(dst, "\nlist = append(list, &decoded)")
					fmt.Fprintf(dst, "\n}")
					fmt.Fprintf(dst, "\nv.%s = list", unexportedFieldName)
					fmt.Fprintf(dst, "\n}")
				}
			case fv.Type.Kind() == reflect.Map:
				elemType := codegen.UnexportedName(fv.Type.Elem().Name())
				elemIface := codegen.ExportedName(fv.Type.Elem().Name())
				if _, ok := entityTypes[elemType]; ok {
					fmt.Fprintf(dst, "\n\nif len(proxy.%s) > 0 {", exportedFieldName)
					fmt.Fprintf(dst, "\nm := make(map[string]%s)", elemIface)
					fmt.Fprintf(dst, "\nfor key, pv := range proxy.%s {", exportedFieldName)
					fmt.Fprintf(dst, "\nvar decoded %s", elemType)
					fmt.Fprintf(dst, "\nif err := json.Unmarshal(pv, &decoded); err != nil {")
					fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to unmasrhal element %%s of field %s`, key)", exportedFieldName)
					fmt.Fprintf(dst, "\n}")
					fmt.Fprintf(dst, "\nm[key] = &decoded")
					fmt.Fprintf(dst, "\n}")
					fmt.Fprintf(dst, "\nv.%s = m", unexportedFieldName)
					fmt.Fprintf(dst, "\n}")
				}
			default:
				elemType := codegen.UnexportedName(fv.Type.Name())
				if _, ok := entityTypes[elemType]; ok {
					fmt.Fprintf(dst, "\n\nif len(proxy.%s) > 0 {", exportedFieldName)
					fmt.Fprintf(dst, "\nvar decoded %s", elemType)
					fmt.Fprintf(dst, "\nif err := json.Unmarshal(proxy.%s, &decoded); err != nil {", exportedFieldName)
					fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to unmarshal field %s`)", exportedFieldName)
					fmt.Fprintf(dst, "\n}")
					fmt.Fprintf(dst, "\n\nv.%s = &decoded", unexportedFieldName)
					fmt.Fprintf(dst, "\n}")
				} else {
					fmt.Fprintf(dst, "\nv.%s = proxy.%s", unexportedFieldName, exportedFieldName)
				}
			}
		}

		if _, ok := postUnmarshalJSONHooks[rv.Type().Name()]; ok {
			fmt.Fprintf(dst, "\n\nv.postUnmarshalJSON()")
		}
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")
	}

	fmt.Fprintf(dst, "\n\nfunc (v *%s) Resolve(resolver *Resolver) error {", rv.Type().Name())
	fmt.Fprintf(dst, "\nif v.IsUnresolved() {")
	fmt.Fprintf(dst, "\n\nresolved, err := resolver.Resolve(v.Reference())")
	fmt.Fprintf(dst, "\nif err != nil {")
	fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to resolve reference %%s`, v.Reference())")
	fmt.Fprintf(dst, "\n}")
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
	fmt.Fprintf(dst, "\nif err := mutator.Do(); err != nil {")
	fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to mutate`)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nv.resolved = true")
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
			fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to resolve %s`)", codegen.ExportedName(fv.Name))
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n}")
		}
	}
	fmt.Fprintf(dst, "\nreturn nil")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (v *%s) QueryJSON(path string) (ret interface{}, ok bool)  {", rv.Type().Name())
	fmt.Fprintf(dst, "\npath = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })")
	fmt.Fprintf(dst, "\nif path == \"\" {")
	fmt.Fprintf(dst, "\nreturn v, true")
	fmt.Fprintf(dst, "\n}")

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
	fmt.Fprintf(dst, "\nreturn nil, false")
	fmt.Fprintf(dst, "\n}")

	if err := writeFormattedSource(&buf, filename); err != nil {
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func generateAccessorsFromEntity(e interface{}) error {
	rv := reflect.ValueOf(e)
	filename := fmt.Sprintf("%s_accessors_gen.go", stringutil.Snake(rv.Type().Name()))
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, "openapi")

	structname := rv.Type().Name()

	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Type().Field(i)
		if fv.Tag.Get("accessor") == "-" {
			continue
		}

		exportedName := codegen.ExportedName(fv.Name)
		unexportedName := codegen.UnexportedName(fv.Name)
		fieldType := fv.Type.Name()
		switch {
		case isMap(fieldType):
			iteratorName := iteratorName(fv.Type)
			fmt.Fprintf(dst, "\n\nfunc (v *%s) %s() *%s {", structname, exportedName, iteratorName)
			fmt.Fprintf(dst, "\nvar items []interface{}")
			fmt.Fprintf(dst, "\nfor key, item := range v.%s {", unexportedName)
			fmt.Fprintf(dst, "\nitems = append(items, &mapIteratorItem{key: key, item: item})")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\nvar iter %s", iteratorName)
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
			fmt.Fprintf(dst, "\niter.items = items")
			fmt.Fprintf(dst, "\nreturn &iter")
			fmt.Fprintf(dst, "\n}")
		default:
			fmt.Fprintf(dst, "\n\nfunc (v *%s) %s() %s {", structname, exportedName, typname(fv.Type))
			fmt.Fprintf(dst, "\nreturn v.%s", unexportedName)
			fmt.Fprintf(dst, "\n}")
		}
	}

	fmt.Fprintf(dst, "\n\nfunc (v *%s) Reference() string {", structname)
	fmt.Fprintf(dst, "\nreturn v.reference")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n\nfunc (v *%s) IsUnresolved() bool {", structname)
	fmt.Fprintf(dst, "\nreturn v.reference != \"\" && !v.resolved")
	fmt.Fprintf(dst, "\n}")

	if err := writeFormattedSource(&buf, filename); err != nil {
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func generateBuildersFromEntity(e interface{}) error {
	rv := reflect.ValueOf(e)
	filename := fmt.Sprintf("%s_builder_gen.go", stringutil.Snake(rv.Type().Name()))
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, "openapi")

	ifacename := codegen.ExportedName(rv.Type().Name())
	structname := rv.Type().Name()

	fmt.Fprintf(dst, "\n\n// %sBuilder is used to build an instance of %s. The user must", ifacename, ifacename)
	fmt.Fprintf(dst, "\n// call `Do()` after providing all the necessary information to")
	fmt.Fprintf(dst, "\n// build an instance of %s", ifacename)
	fmt.Fprintf(dst, "\ntype %sBuilder struct {", ifacename)
	fmt.Fprintf(dst, "\ntarget *%s", structname)
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Do finalizes the building process for %s and returns the result", ifacename)
	fmt.Fprintf(dst, "\nfunc (b *%sBuilder) Do() %s {", ifacename, ifacename)
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

	fmt.Fprintf(dst, "\n\n// New%s creates a new builder object for %s", ifacename, ifacename)
	fmt.Fprintf(dst, "\nfunc New%s(", ifacename)
	for i, fv := range requireds {
		fmt.Fprintf(dst, "%s %s", codegen.UnexportedName(fv.Name), typname(fv.Type))
		if i < len(requireds)-1 {
			fmt.Fprintf(dst, ", ")
		}
	}
	fmt.Fprintf(dst, ") *%sBuilder {", ifacename)
	fmt.Fprintf(dst, "\nreturn &%sBuilder{", ifacename)
	fmt.Fprintf(dst, "\ntarget: &%s{", structname)

	hasDefault := make(map[string]struct{})
	for _, fv := range defaults {
		hasDefault[fv.Name] = struct{}{}
		fmt.Fprintf(dst, "\n%s: %s,", fv.Name, fv.Tag.Get("default"))
	}
	for _, fv := range requireds {
		fmt.Fprintf(dst, "\n%s: %s,", fv.Name, codegen.UnexportedName(fv.Name))
	}
	fmt.Fprintf(dst, "\n},")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	for _, fv := range optionals {
		fmt.Fprintf(dst, "\n\n// %[1]s sets the %[1]s field for object %[2]s.", codegen.ExportedName(fv.Name), ifacename)
		if _, ok := hasDefault[fv.Name]; ok {
			fmt.Fprintf(dst, " If this is not called,\n// a default value (%s) is assigned to this field", fv.Tag.Get("default"))
		}
		fmt.Fprintf(dst, "\nfunc (b *%[1]sBuilder) %[2]s(v %[3]s) *%[1]sBuilder {", ifacename, codegen.ExportedName(fv.Name), typname(fv.Type))
		fmt.Fprintf(dst, "\nb.target.%s = v", fv.Name)
		fmt.Fprintf(dst, "\nreturn b")
		fmt.Fprintf(dst, "\n}")
	}

	fmt.Fprintf(dst, "\n\n// Reference sets the $ref (reference) field for object %s.", ifacename)
	fmt.Fprintf(dst, "\nfunc (b *%sBuilder) Reference(v string) *%sBuilder {", ifacename, ifacename)
	fmt.Fprintf(dst, "\nb.target.reference = v")
	fmt.Fprintf(dst, "\nreturn b")
	fmt.Fprintf(dst, "\n}")

	if err := writeFormattedSource(&buf, filename); err != nil {
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func generateMutatorsFromEntity(e interface{}) error {
	rv := reflect.ValueOf(e)
	filename := fmt.Sprintf("%s_mutator_gen.go", stringutil.Snake(rv.Type().Name()))
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, "openapi")
	codegen.WriteImports(dst, "log")

	ifacename := codegen.ExportedName(rv.Type().Name())
	structname := rv.Type().Name()

	fmt.Fprintf(dst, "\n\n// %sMutator is used to build an instance of %s. The user must", ifacename, ifacename)
	fmt.Fprintf(dst, "\n// call `Do()` after providing all the necessary information to")
	fmt.Fprintf(dst, "\n// the new instance of %s with new values", ifacename)
	fmt.Fprintf(dst, "\ntype %sMutator struct {", ifacename)
	fmt.Fprintf(dst, "\nproxy *%s", structname)
	fmt.Fprintf(dst, "\ntarget *%s", structname)
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Do finalizes the matuation process for %s and returns the result", ifacename)
	fmt.Fprintf(dst, "\nfunc (b *%sMutator) Do() error {", ifacename)
	// TODO: validation
	fmt.Fprintf(dst, "\n*b.target = *b.proxy")
	fmt.Fprintf(dst, "\nreturn nil")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\n// Mutate%s creates a new mutator object for %s", ifacename, ifacename)
	fmt.Fprintf(dst, "\nfunc Mutate%s(v %s) *%sMutator {", ifacename, ifacename, ifacename)
	fmt.Fprintf(dst, "\nreturn &%sMutator{", ifacename)
	fmt.Fprintf(dst, "\ntarget: v.(*%s),", structname)
	fmt.Fprintf(dst, "\nproxy: v.Clone().(*%s),", structname)
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Type().Field(i)
		exportedName := codegen.ExportedName(fv.Name)

		if fv.Tag.Get("mutator") == "-" {
			log.Printf(" * [SKIP] %s", exportedName)
			continue
		}
		log.Printf(" * %s", exportedName)

		unexportedName := codegen.UnexportedName(fv.Name)
		fieldType := fv.Type.Name()
		switch {
		case isMap(fieldType):
			fmt.Fprintf(dst, "\n\nfunc (b *%sMutator) Clear%s() *%sMutator {", ifacename, exportedName, ifacename)
			fmt.Fprintf(dst, "\nb.proxy.%s.Clear()", unexportedName)
			fmt.Fprintf(dst, "\nreturn b")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n\nfunc (b *%sMutator) %s(key %sKey, value %s) *%sMutator {", ifacename, inflection.Singular(exportedName), fieldType, typname(fv.Type.Elem()), ifacename)
			fmt.Fprintf(dst, "\nif b.proxy.%s == nil {", unexportedName)
			fmt.Fprintf(dst, "\nb.proxy.%s = %s{}", unexportedName, fieldType)
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n\nb.proxy.%s[key] = value", unexportedName)
			if isEntity(fv.Type.Elem().Name()) {
				fmt.Fprintf(dst, ".Clone()")
			}
			fmt.Fprintf(dst, "\nreturn b")
			fmt.Fprintf(dst, "\n}")
		case isList(fieldType):
			fmt.Fprintf(dst, "\n\nfunc (b *%sMutator) Clear%s() *%sMutator {", ifacename, exportedName, ifacename)
			fmt.Fprintf(dst, "\nb.proxy.%s.Clear()", unexportedName)
			fmt.Fprintf(dst, "\nreturn b")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n\nfunc (b *%sMutator) %s(value %s) *%sMutator {", ifacename, inflection.Singular(exportedName), typname(fv.Type.Elem()), ifacename)
			fmt.Fprintf(dst, "\nb.proxy.%s = append(b.proxy.%s, value)", unexportedName, unexportedName)
			fmt.Fprintf(dst, "\nreturn b")
			fmt.Fprintf(dst, "\n}")
		default:
			fmt.Fprintf(dst, "\n\n// %s sets the %s field for object %s.", exportedName, exportedName, ifacename)
			fmt.Fprintf(dst, "\nfunc (b *%sMutator) %s(v %s) *%sMutator {", ifacename, exportedName, typname(fv.Type), ifacename)
			fmt.Fprintf(dst, "\nb.proxy.%s = v", unexportedName)
			fmt.Fprintf(dst, "\nreturn b")
			fmt.Fprintf(dst, "\n}")
		}
	}
	if err := writeFormattedSource(&buf, filename); err != nil {
		return errors.Wrap(err, `failed to write result to file`)
	}
	return nil
}

func generateClonersFromEntity(entities []interface{}) error {
	filename := "cloner_gen.go"
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, "openapi")

	for _, e := range entities {
		rv := reflect.ValueOf(e)
		fmt.Fprintf(dst, "\n\nfunc (v *%s) Clone() %s {", rv.Type().Name(), codegen.ExportedName(rv.Type().Name()))
		fmt.Fprintf(dst, "\nvar dst %s", rv.Type().Name())
		fmt.Fprintf(dst, "\ndst = *v")
		fmt.Fprintf(dst, "\nreturn &dst")
		fmt.Fprintf(dst, "\n}")
	}

	return writeFormattedSource(&buf, filename)
}

func generateIteratorsFromEntity(entities []interface{}) error {
	filename := "iterator_gen.go"
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, "openapi")
	codegen.WriteImports(dst, "sync")

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

	return writeFormattedSource(&buf, filename)
}

func generateContainers(c interface{}) error {
	rv := reflect.ValueOf(c)

	filename := fmt.Sprintf("%s_gen.go", stringutil.Snake(rv.Type().Name()))
	log.Printf("Generating %s", filename)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	codegen.WritePreamble(dst, "openapi")
	codegen.WriteImports(dst, "encoding/json", "github.com/pkg/errors")

	typeName := rv.Type().Name()
	switch {
	case isList(typeName):
		fmt.Fprintf(dst, "\n\nfunc (v *%s) Clear() error {", typeName)
		fmt.Fprintf(dst, "\n*v = %s(nil)", rv.Type().Name())
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")

		elemType := codegen.UnexportedName(rv.Type().Elem().Name())

		fmt.Fprintf(dst, "\n\nfunc (v %s) Resolve(resolver *Resolver) error {", typeName)
		if _, ok := entityTypes[elemType]; ok {
			fmt.Fprintf(dst, "\nif len(v) > 0 {")
			fmt.Fprintf(dst, "\nfor i, elem := range v {")
			fmt.Fprintf(dst, "\nif err := elem.Resolve(resolver); err != nil {")
			fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to resolve %s (index = %%d)`, i)", typeName)
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n}")
		}
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")

		if rv.Type().Elem().Name() != "" && rv.Type().Elem().Kind() == reflect.Interface {
			fmt.Fprintf(dst, "\n\nfunc (v *%s) UnmarshalJSON(data []byte) error {", typeName)
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
	case isMap(rv.Type().Name()):
		fmt.Fprintf(dst, "\n\nfunc (v *%s) Clear() error {", rv.Type().Name())
		fmt.Fprintf(dst, "\n*v = make(%s)", rv.Type().Name())
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")
		fmt.Fprintf(dst, "\n\nfunc (v %s) Resolve(resolver *Resolver) error {", rv.Type().Name())
		if _, ok := entityTypes[codegen.UnexportedName(rv.Type().Elem().Name())]; ok {
			fmt.Fprintf(dst, "\nif len(v) > 0 {")
			fmt.Fprintf(dst, "\nfor name, elem := range v {")
			fmt.Fprintf(dst, "\nif err := elem.Resolve(resolver); err != nil {")
			fmt.Fprintf(dst, "\nreturn errors.Wrapf(err, `failed to resolve %s (key = %%s)`, name)", rv.Type().Name())
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n}")
		}
		fmt.Fprintf(dst, "\nreturn nil")
		fmt.Fprintf(dst, "\n}")

		fmt.Fprintf(dst, "\n\nfunc (v %s) QueryJSON(path string) (ret interface{}, ok bool) {", rv.Type().Name())
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
			fmt.Fprintf(dst, "\n\nfunc (v *%s) UnmarshalJSON(data []byte) error {", typeName)
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

	return writeFormattedSource(&buf, filename)
}

func writeFormattedSource(buf *bytes.Buffer, filename string) error {
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.String())
		return errors.Wrap(err, `failed to format source`)
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrapf(err, `failed to open file %s`, filename)
	}
	defer f.Close()

	f.Write(formatted)

	return nil
}

func isMap(s string) bool {
	return strings.HasSuffix(s, "Map")
}

func isList(s string) bool {
	return strings.HasSuffix(s, "List")
}

func isEntity(s string) bool {
	_, ok := entityTypes[s]
	return ok
}

func isContainer(s string) bool {
	_, ok := containerTypes[s]
	return ok
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
	if !isContainer(name) && !isEntity(t.Elem().Name()) {
		name = typname(t.Elem())
		if strings.HasPrefix(name, "[]") {
			name = strings.TrimPrefix(name, "[]") + "List"
		}
		name = strcase.ToCamel(name)
	}

	if name == "" {
		panic("empty name " + t.String())
	}

	return strcase.ToCamel(name + "Iterator")
}
