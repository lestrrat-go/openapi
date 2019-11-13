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
		return errors.Wrapf(err, `failed to unmarshal tag`)
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

// TagFromJSON constructs a Tag from JSON buffer. `dst` must
// be a pointer to `Tag`
func TagFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Tag)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Tag, but got %T`, dst)
	}
	var tmp tag
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Tag`)
	}
	*v = &tmp
	return nil
}
