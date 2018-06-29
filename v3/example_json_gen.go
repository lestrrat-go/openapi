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

type exampleMarshalProxy struct {
	Reference     string      `json:"$ref,omitempty"`
	Description   string      `json:"description"`
	Value         interface{} `json:"value"`
	ExternalValue string      `json:"externalValue"`
}

type exampleUnmarshalProxy struct {
	Reference     string      `json:"$ref,omitempty"`
	Description   string      `json:"description"`
	Value         interface{} `json:"value"`
	ExternalValue string      `json:"externalValue"`
}

func (v *example) MarshalJSON() ([]byte, error) {
	var proxy exampleMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Description = v.description
	proxy.Value = v.value
	proxy.ExternalValue = v.externalValue
	return json.Marshal(proxy)
}

func (v *example) UnmarshalJSON(data []byte) error {
	var proxy exampleUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.description = proxy.Description
	v.value = proxy.Value
	v.externalValue = proxy.ExternalValue
	return nil
}

func (v *example) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "description":
		target = v.description
	case "value":
		target = v.value
	case "externalValue":
		target = v.externalValue
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

// ExampleFromJSON constructs a Example from JSON buffer. `dst` must
// be a pointer to `Example`
func ExampleFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Example)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Example, but got %T`, dst)
	}
	var tmp example
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Example`)
	}
	*v = &tmp
	return nil
}
