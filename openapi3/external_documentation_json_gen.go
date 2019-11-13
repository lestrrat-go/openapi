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

type externalDocumentationMarshalProxy struct {
	Reference   string `json:"$ref,omitempty"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type externalDocumentationUnmarshalProxy struct {
	Reference   string `json:"$ref,omitempty"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func (v *externalDocumentation) MarshalJSON() ([]byte, error) {
	var proxy externalDocumentationMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Description = v.description
	proxy.URL = v.url
	return json.Marshal(proxy)
}

func (v *externalDocumentation) UnmarshalJSON(data []byte) error {
	var proxy externalDocumentationUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrapf(err, `failed to unmarshal externalDocumentation`)
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.description = proxy.Description
	v.url = proxy.URL
	return nil
}

func (v *externalDocumentation) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "url":
		target = v.url
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

// ExternalDocumentationFromJSON constructs a ExternalDocumentation from JSON buffer. `dst` must
// be a pointer to `ExternalDocumentation`
func ExternalDocumentationFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*ExternalDocumentation)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to ExternalDocumentation, but got %T`, dst)
	}
	var tmp externalDocumentation
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal ExternalDocumentation`)
	}
	*v = &tmp
	return nil
}
