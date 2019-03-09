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

type responseMarshalProxy struct {
	Reference   string       `json:"$ref,omitempty"`
	Description string       `json:"description"`
	Headers     HeaderMap    `json:"headers,omitempty"`
	Content     MediaTypeMap `json:"content,omitempty"`
	Links       LinkMap      `json:"links,omitempty"`
}

type responseUnmarshalProxy struct {
	Reference   string       `json:"$ref,omitempty"`
	Description string       `json:"description"`
	Headers     HeaderMap    `json:"headers,omitempty"`
	Content     MediaTypeMap `json:"content,omitempty"`
	Links       LinkMap      `json:"links,omitempty"`
}

func (v *response) MarshalJSON() ([]byte, error) {
	var proxy responseMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Description = v.description
	proxy.Headers = v.headers
	proxy.Content = v.content
	proxy.Links = v.links
	return json.Marshal(proxy)
}

func (v *response) UnmarshalJSON(data []byte) error {
	var proxy responseUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrapf(err, `failed to unmarshal response`)
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.description = proxy.Description
	v.headers = proxy.Headers
	v.content = proxy.Content
	v.links = proxy.Links
	return nil
}

func (v *response) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "headers":
		target = v.headers
	case "content":
		target = v.content
	case "links":
		target = v.links
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

// ResponseFromJSON constructs a Response from JSON buffer. `dst` must
// be a pointer to `Response`
func ResponseFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Response)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Response, but got %T`, dst)
	}
	var tmp response
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Response`)
	}
	*v = &tmp
	return nil
}
