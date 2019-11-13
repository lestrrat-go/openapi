//go:generate go run internal/cmd/gentypes/gentypes.go

// Package openapi implements OpenAPI 3.x.
//
// This package contains objects that comprise an OpenAPI spec, but only
// their interfaces are exported. For example, openapi3.OpenAPI (the root
// node that encompasses an entire spec) is an interface, paths are
// stored under the openapi3.Paths interface, and so forth.
//
// The real objects are hidden within this package so that users cannot
// accidentally create an incomplete object (for example, the spec states
// that the root element must contain an "info" key, but if the openapi3.OpenAPI
// was a struct you would be able to create an empty openapi3.OpenAPI{},
// which would be invalid)
//
// The API limits your ability to accidentally create empty values (and also
// the ability to assign arbitrary values inot fields) because
// the whole OpenAPI spec is rather complex -- that is, when we translate
// the document structure into Go, and we try to work with it, it's not
// just a matter of processing each objects, but we need context. For example
// when working with Operation objects, we really need to know the path
// that we're dealing with, as well as the HTTP verb that go with it.
// However, these values are not defined within the Operation objects: they
// belong to the Paths object and PathItem objects. When processing through
// these objects we would either have to pass all of the objects, or we
// would need to somehow remember their relationships, so we can query them.
// This is where the limitation comes into play. By limiting unregulated
// access to fields, we provide automatic hooks to keep track of these
// relationships.
//
// In doing so, we are able to provide API like this
//
//   oper := ... // Operation object
//   if oper.Detached() { // has this been assigned to a PathItem?
//     path := oper.Path() // path that this Operation belongs to
//     verb := oper.Verb() // HTTP verb that this Operation is assigned to
//   }
//
// Building objects must happen through a builder object: The constructor for
// the object builder asks the user to pass the required parameters that have
// no default values, and everything else can be optionally passed via methods.
// When all values are handed, you call `Do()` to obtain the object.
//
//   openapi3.NewParameter(name). // name parameter is required
//     Required(true). // required parameter is optional
//     Do()
//
// The objects are *generally* immutable and therefore provide no mutator
// methods on themselves. However we realize that there are times when you
// just want to change them: To workaround this issue, we provide mutators
// which work almost identically as the builders.
//
//   openapi3.MutatePrameter(p).
//     Required(false).
//     Do()
//
// When we assign entity objects, the objects are automatically cloned to
// avoid mutation from outside of the OpenAPI tree
package openapi3

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.1.md#parameterObject
// required (boolean)
//    Determines whether this parameter is mandatory. If the parameter
//    location is "path", this property is REQUIRED and its value MUST
//    be true. Otherwise, the property MAY be included and its default
//    value is false.
func defaultParameterRequiredFromLocation(in Location) bool {
	return in == InPath
}

func extractFragFromPath(path string) (string, string) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	var frag string
	if i := strings.Index(path, `/`); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ``
	}
	return frag, path
}

func parse(unmarshaler func([]byte, interface{}) error, src io.Reader, options ...Option) (OpenAPI, error) {
	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, errors.Wrap(err, `failed to read from source`)
	}

	var spec openAPI
	if err := unmarshaler(buf, &spec); err != nil {
		return nil, errors.Wrap(err, `failed to unmarshal spec`)
	}

	validate := true
	for _, option := range options {
		switch option.Name() {
		case optkeyValidate:
			validate = option.Value().(bool)
		}
	}

	if validate {
		if err := spec.Validate(true); err != nil {
			return nil, errors.Wrap(err, `failed to validate spec`)
		}
	}
	return &spec, nil
}

func yamlUnmarshal(buf []byte, v interface{}) error {
  return yaml.Unmarshal(buf, v)
}

func ParseYAML(src io.Reader, options ...Option) (OpenAPI, error) {
	return parse(yamlUnmarshal, src, options...)
}

func ParseJSON(src io.Reader, options ...Option) (OpenAPI, error) {
	return parse(json.Unmarshal, src, options...)
}

