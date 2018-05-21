package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
)

var _ = errors.Cause

type tagMarshalProxy struct {
	Name         string                `json:"name"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type tagUnmarshalProxy struct {
	Name         string          `json:"name"`
	Description  string          `json:"description,omitempty"`
	ExternalDocs json.RawMessage `json:"externalDocs,omitempty"`
}

func (v *tag) MarshalJSON() ([]byte, error) {
	var proxy tagMarshalProxy
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
