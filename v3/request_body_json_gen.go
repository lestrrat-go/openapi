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

type requestBodyMarshalProxy struct {
	Reference   string       `json:"$ref,omitempty"`
	Description string       `json:"description,omitempty"`
	Content     MediaTypeMap `json:"content,omitempty"`
	Required    bool         `json:"required,omitempty"`
}

type requestBodyUnmarshalProxy struct {
	Reference   string       `json:"$ref,omitempty"`
	Description string       `json:"description,omitempty"`
	Content     MediaTypeMap `json:"content,omitempty"`
	Required    bool         `json:"required,omitempty"`
}

func (v *requestBody) MarshalJSON() ([]byte, error) {
	var proxy requestBodyMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Description = v.description
	proxy.Content = v.content
	proxy.Required = v.required
	return json.Marshal(proxy)
}

func (v *requestBody) UnmarshalJSON(data []byte) error {
	var proxy requestBodyUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrapf(err, `failed to unmarshal requestBody`)
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.description = proxy.Description
	v.content = proxy.Content
	v.required = proxy.Required

	v.postUnmarshalJSON()
	return nil
}

func (v *requestBody) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "content":
		target = v.content
	case "required":
		target = v.required
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

// RequestBodyFromJSON constructs a RequestBody from JSON buffer. `dst` must
// be a pointer to `RequestBody`
func RequestBodyFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*RequestBody)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to RequestBody, but got %T`, dst)
	}
	var tmp requestBody
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal RequestBody`)
	}
	*v = &tmp
	return nil
}
