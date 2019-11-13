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

type contactMarshalProxy struct {
	Reference string `json:"$ref,omitempty"`
	Name      string `json:"name,omitempty"`
	URL       string `json:"url,omitempty"`
	Email     string `json:"email,omitempty"`
}

type contactUnmarshalProxy struct {
	Reference string `json:"$ref,omitempty"`
	Name      string `json:"name,omitempty"`
	URL       string `json:"url,omitempty"`
	Email     string `json:"email,omitempty"`
}

func (v *contact) MarshalJSON() ([]byte, error) {
	var proxy contactMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Name = v.name
	proxy.URL = v.url
	proxy.Email = v.email
	return json.Marshal(proxy)
}

func (v *contact) UnmarshalJSON(data []byte) error {
	var proxy contactUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrapf(err, `failed to unmarshal contact`)
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.name = proxy.Name
	v.url = proxy.URL
	v.email = proxy.Email
	return nil
}

func (v *contact) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "url":
		target = v.url
	case "email":
		target = v.email
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

// ContactFromJSON constructs a Contact from JSON buffer. `dst` must
// be a pointer to `Contact`
func ContactFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Contact)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Contact, but got %T`, dst)
	}
	var tmp contact
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Contact`)
	}
	*v = &tmp
	return nil
}
