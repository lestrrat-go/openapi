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

type serverMarshalProxy struct {
	Reference   string            `json:"$ref,omitempty"`
	URL         string            `json:"url"`
	Description string            `json:"description,omitempty"`
	Variables   ServerVariableMap `json:"variables,omitempty"`
}

type serverUnmarshalProxy struct {
	Reference   string            `json:"$ref,omitempty"`
	URL         string            `json:"url"`
	Description string            `json:"description,omitempty"`
	Variables   ServerVariableMap `json:"variables,omitempty"`
}

func (v *server) MarshalJSON() ([]byte, error) {
	var proxy serverMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.URL = v.url
	proxy.Description = v.description
	proxy.Variables = v.variables
	return json.Marshal(proxy)
}

func (v *server) UnmarshalJSON(data []byte) error {
	var proxy serverUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.url = proxy.URL
	v.description = proxy.Description
	v.variables = proxy.Variables
	return nil
}

func (v *server) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "url":
		target = v.url
	case "description":
		target = v.description
	case "variables":
		target = v.variables
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
