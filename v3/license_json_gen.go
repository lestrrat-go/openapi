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

type licenseMarshalProxy struct {
	Reference string `json:"$ref,omitempty"`
	Name      string `json:"name"`
	URL       string `json:"url,omitempty"`
}

type licenseUnmarshalProxy struct {
	Reference string `json:"$ref,omitempty"`
	Name      string `json:"name"`
	URL       string `json:"url,omitempty"`
}

func (v *license) MarshalJSON() ([]byte, error) {
	var proxy licenseMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Name = v.name
	proxy.URL = v.url
	return json.Marshal(proxy)
}

func (v *license) UnmarshalJSON(data []byte) error {
	var proxy licenseUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.name = proxy.Name
	v.url = proxy.URL
	return nil
}

func (v *license) QueryJSON(path string) (ret interface{}, ok bool) {
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

// LicenseFromJSON constructs a License from JSON buffer. `dst` must
// be a pointer to `License`
func LicenseFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*License)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to License, but got %T`, dst)
	}
	var tmp license
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal License`)
	}
	*v = &tmp
	return nil
}
