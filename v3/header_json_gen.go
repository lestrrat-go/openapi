package openapi

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

type headerMarshalProxy struct {
	Reference       string       `json:"$ref,omitempty"`
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

type headerUnmarshalProxy struct {
	Reference       string          `json:"$ref,omitempty"`
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

func (v *header) MarshalJSON() ([]byte, error) {
	var proxy headerMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
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

func (v *header) UnmarshalJSON(data []byte) error {
	var proxy headerUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
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

func (v *header) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*header)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Header, but got %T`, resolved)
		}
		mutator := MutateHeader(v)
		mutator.In(asserted.In())
		mutator.Required(asserted.Required())
		mutator.Description(asserted.Description())
		mutator.Deprecated(asserted.Deprecated())
		mutator.AllowEmptyValue(asserted.AllowEmptyValue())
		mutator.Explode(asserted.Explode())
		mutator.AllowReserved(asserted.AllowReserved())
		mutator.Schema(asserted.Schema())
		for iter := asserted.Examples(); iter.Next(); {
			key, item := iter.Item()
			mutator.Example(key, item)
		}
		for iter := asserted.Content(); iter.Next(); {
			key, item := iter.Item()
			mutator.Content(key, item)
		}
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.schema != nil {
		if err := v.schema.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Schema`)
		}
	}
	if v.examples != nil {
		if err := v.examples.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Examples`)
		}
	}
	if v.content != nil {
		if err := v.content.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Content`)
		}
	}
	return nil
}

func (v *header) QueryJSON(path string) (ret interface{}, ok bool) {
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
