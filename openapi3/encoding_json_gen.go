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

type encodingMarshalProxy struct {
	Reference     string    `json:"$ref,omitempty"`
	ContentType   string    `json:"contentType"`
	Headers       HeaderMap `json:"headers"`
	Explode       bool      `json:"explode"`
	AllowReserved bool      `json:"allowReserved"`
}

type encodingUnmarshalProxy struct {
	Reference     string    `json:"$ref,omitempty"`
	ContentType   string    `json:"contentType"`
	Headers       HeaderMap `json:"headers"`
	Explode       bool      `json:"explode"`
	AllowReserved bool      `json:"allowReserved"`
}

func (v *encoding) MarshalJSON() ([]byte, error) {
	var proxy encodingMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.ContentType = v.contentType
	proxy.Headers = v.headers
	proxy.Explode = v.explode
	proxy.AllowReserved = v.allowReserved
	return json.Marshal(proxy)
}

func (v *encoding) UnmarshalJSON(data []byte) error {
	var proxy encodingUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrapf(err, `failed to unmarshal encoding`)
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.contentType = proxy.ContentType
	v.headers = proxy.Headers
	v.explode = proxy.Explode
	v.allowReserved = proxy.AllowReserved
	return nil
}

func (v *encoding) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "contentType":
		target = v.contentType
	case "headers":
		target = v.headers
	case "explode":
		target = v.explode
	case "allowReserved":
		target = v.allowReserved
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

// EncodingFromJSON constructs a Encoding from JSON buffer. `dst` must
// be a pointer to `Encoding`
func EncodingFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Encoding)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Encoding, but got %T`, dst)
	}
	var tmp encoding
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Encoding`)
	}
	*v = &tmp
	return nil
}
