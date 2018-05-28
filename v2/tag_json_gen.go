package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T19:20:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var _ = log.Printf
var _ = json.Unmarshal
var _ = errors.Cause

type tagMarshalProxy struct {
	Reference    string                `json:"$ref,omitempty"`
	Name         string                `json:"name"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type tagUnmarshalProxy struct {
	Reference    string          `json:"$ref,omitempty"`
	Name         string          `json:"name"`
	Description  string          `json:"description,omitempty"`
	ExternalDocs json.RawMessage `json:"externalDocs,omitempty"`
}

func (v *tag) MarshalJSON() ([]byte, error) {
	var proxy tagMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Name = v.name
	proxy.Description = v.description
	proxy.ExternalDocs = v.externalDocs
	return json.Marshal(proxy)
}

func (v *tag) UnmarshalJSON(data []byte) error {
	var proxy tagUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.name = proxy.Name
	v.description = proxy.Description

	if len(proxy.ExternalDocs) > 0 {
		var decoded externalDocumentation
		if err := json.Unmarshal(proxy.ExternalDocs, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		v.externalDocs = &decoded
	}
	return nil
}

func (v *tag) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*tag)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Tag, but got %T`, resolved)
		}
		mutator := MutateTag(v)
		mutator.Name(asserted.Name())
		mutator.Description(asserted.Description())
		mutator.ExternalDocs(asserted.ExternalDocs())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.externalDocs != nil {
		if err := v.externalDocs.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve ExternalDocs`)
		}
	}
	return nil
}

func (v *tag) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "description":
		target = v.description
	case "externalDocs":
		target = v.externalDocs
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
