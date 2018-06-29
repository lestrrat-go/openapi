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

type callbackMarshalProxy struct {
	Reference string              `json:"$ref,omitempty"`
	URLs      map[string]PathItem `json:""`
}

type callbackUnmarshalProxy struct {
	Reference string                     `json:"$ref,omitempty"`
	URLs      map[string]json.RawMessage `json:""`
}

func (v *callback) MarshalJSON() ([]byte, error) {
	var proxy callbackMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.URLs = v.urls
	return json.Marshal(proxy)
}

func (v *callback) UnmarshalJSON(data []byte) error {
	var proxy callbackUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}

	if len(proxy.URLs) > 0 {
		m := make(map[string]PathItem)
		for key, pv := range proxy.URLs {
			var decoded pathItem
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field URLs`, key)
			}
			m[key] = &decoded
		}
		v.urls = m
	}
	return nil
}

func (v *callback) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "urls":
		target = v.urls
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

// CallbackFromJSON constructs a Callback from JSON buffer. `dst` must
// be a pointer to `Callback`
func CallbackFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Callback)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Callback, but got %T`, dst)
	}
	var tmp callback
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Callback`)
	}
	*v = &tmp
	return nil
}
