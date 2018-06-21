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

type mediaTypeMarshalProxy struct {
	Reference string      `json:"$ref,omitempty"`
	Schema    Schema      `json:"schema,omitempty"`
	Examples  ExampleMap  `json:"examples,omitempty"`
	Encoding  EncodingMap `json:"encoding,omitempty"`
}

type mediaTypeUnmarshalProxy struct {
	Reference string          `json:"$ref,omitempty"`
	Schema    json.RawMessage `json:"schema,omitempty"`
	Examples  ExampleMap      `json:"examples,omitempty"`
	Encoding  EncodingMap     `json:"encoding,omitempty"`
}

func (v *mediaType) MarshalJSON() ([]byte, error) {
	var proxy mediaTypeMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Schema = v.schema
	proxy.Examples = v.examples
	proxy.Encoding = v.encoding
	return json.Marshal(proxy)
}

func (v *mediaType) UnmarshalJSON(data []byte) error {
	var proxy mediaTypeUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}

	if len(proxy.Schema) > 0 {
		var decoded schema
		if err := json.Unmarshal(proxy.Schema, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Schema`)
		}

		v.schema = &decoded
	}
	v.examples = proxy.Examples
	v.encoding = proxy.Encoding
	return nil
}

func (v *mediaType) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*mediaType)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type MediaType, but got %T`, resolved)
		}
		mutator := MutateMediaType(v)
		mutator.Name(asserted.Name())
		mutator.Mime(asserted.Mime())
		mutator.Schema(asserted.Schema())
		for iter := asserted.Examples(); iter.Next(); {
			key, item := iter.Item()
			mutator.Example(key, item)
		}
		for iter := asserted.Encoding(); iter.Next(); {
			key, item := iter.Item()
			mutator.Encoding(key, item)
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
	if v.encoding != nil {
		if err := v.encoding.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Encoding`)
		}
	}
	return nil
}

func (v *mediaType) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "schema":
		target = v.schema
	case "examples":
		target = v.examples
	case "encoding":
		target = v.encoding
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
