package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/pkg/errors"
)

var _ = log.Printf
var _ = json.Unmarshal
var _ = errors.Cause

type parameterMarshalProxy struct {
	Reference       string       `json:"$ref,omitempty"`
	Name            string       `json:"name,omitempty"`
	In              Location     `json:"in"`
	Required        bool         `json:"required,omitempty"`
	Description     string       `json:"description,omitempty"`
	Deprecated      bool         `json:"deprecated,omitempty"`
	AllowEmptyValue bool         `json:"allowEmptyValue,omitempty"`
	Explode         bool         `json:"explode,omitempty"`
	AllowReserved   bool         `json:"allowReserved,omitempty"`
	Schema          Schema       `json:"schema,omitempty"`
	Examples        ExampleMap   `json:"examples,omitempty"`
	Content         MediaTypeMap `json:"content,omitempty"`
}

type parameterUnmarshalProxy struct {
	Reference       string          `json:"$ref,omitempty"`
	Name            string          `json:"name,omitempty"`
	In              Location        `json:"in"`
	Required        bool            `json:"required,omitempty"`
	Description     string          `json:"description,omitempty"`
	Deprecated      bool            `json:"deprecated,omitempty"`
	AllowEmptyValue bool            `json:"allowEmptyValue,omitempty"`
	Explode         bool            `json:"explode,omitempty"`
	AllowReserved   bool            `json:"allowReserved,omitempty"`
	Schema          json.RawMessage `json:"schema,omitempty"`
	Examples        ExampleMap      `json:"examples,omitempty"`
	Content         MediaTypeMap    `json:"content,omitempty"`
}

func (v *parameter) MarshalJSON() ([]byte, error) {
	var proxy parameterMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Name = v.name
	proxy.In = v.in
	proxy.Required = v.required
	proxy.Description = v.description
	proxy.Deprecated = v.deprecated
	proxy.AllowEmptyValue = v.allowEmptyValue
	proxy.Explode = v.explode
	proxy.AllowReserved = v.allowReserved
	proxy.Schema = v.schema
	proxy.Examples = v.examples
	proxy.Content = v.content
	return json.Marshal(proxy)
}

func (v *parameter) UnmarshalJSON(data []byte) error {
	var proxy parameterUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrapf(err, `failed to unmarshal parameter`)
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.name = proxy.Name
	v.in = proxy.In
	v.required = proxy.Required
	v.description = proxy.Description
	v.deprecated = proxy.Deprecated
	v.allowEmptyValue = proxy.AllowEmptyValue
	v.explode = proxy.Explode
	v.allowReserved = proxy.AllowReserved

	if len(proxy.Schema) > 0 {
		var decoded schema
		if err := json.Unmarshal(proxy.Schema, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Schema`)
		}

		v.schema = &decoded
	}
	v.examples = proxy.Examples
	v.content = proxy.Content
	return nil
}

func (v *parameter) QueryJSON(path string) (ret interface{}, ok bool) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	if path == "" {
		return v, true
	}

	var frag string
	if i := strings.Index(path, "/"); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ""
	}

	var target interface{}

	switch frag {
	case "name":
		target = v.name
	case "in":
		target = v.in
	case "required":
		target = v.required
	case "description":
		target = v.description
	case "deprecated":
		target = v.deprecated
	case "allowEmptyValue":
		target = v.allowEmptyValue
	case "explode":
		target = v.explode
	case "allowReserved":
		target = v.allowReserved
	case "schema":
		target = v.schema
	case "examples":
		target = v.examples
	case "content":
		target = v.content
	default:
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}
	if path == "" {
		return target, true
	}
	return nil, false
}

// ParameterFromJSON constructs a Parameter from JSON buffer. `dst` must
// be a pointer to `Parameter`
func ParameterFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Parameter)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Parameter, but got %T`, dst)
	}
	var tmp parameter
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Parameter`)
	}
	*v = &tmp
	return nil
}
