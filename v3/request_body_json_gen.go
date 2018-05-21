package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
)

var _ = errors.Cause

type requestBodyMarshalProxy struct {
	Description string               `json:"description,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
	Required    bool                 `json:"required,omitempty"`
}

type requestBodyUnmarshalProxy struct {
	Description string                     `json:"description,omitempty"`
	Content     map[string]json.RawMessage `json:"content,omitempty"`
	Required    bool                       `json:"required,omitempty"`
}

func (v *requestBody) MarshalJSON() ([]byte, error) {
	var proxy requestBodyMarshalProxy
	proxy.Description = v.description
	proxy.Content = v.content
	proxy.Required = v.required
	return json.Marshal(proxy)
}

func (v *requestBody) UnmarshalJSON(data []byte) error {
	var proxy requestBodyUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.description = proxy.Description

	if len(proxy.Content) > 0 {
		m := make(map[string]MediaType)
		for key, pv := range proxy.Content {
			var decoded mediaType
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Content`, key)
			}
			m[key] = &decoded
		}
		v.content = m
	}
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
