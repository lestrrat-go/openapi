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

type discriminatorMarshalProxy struct {
	Reference    string    `json:"$ref,omitempty"`
	PropertyName string    `json:"propertyName"`
	Mapping      StringMap `json:"mapping"`
}

type discriminatorUnmarshalProxy struct {
	Reference    string    `json:"$ref,omitempty"`
	PropertyName string    `json:"propertyName"`
	Mapping      StringMap `json:"mapping"`
}

func (v *discriminator) MarshalJSON() ([]byte, error) {
	var proxy discriminatorMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.PropertyName = v.propertyName
	proxy.Mapping = v.mapping
	return json.Marshal(proxy)
}

func (v *discriminator) UnmarshalJSON(data []byte) error {
	var proxy discriminatorUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrapf(err, `failed to unmarshal discriminator`)
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.propertyName = proxy.PropertyName
	v.mapping = proxy.Mapping
	return nil
}

func (v *discriminator) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "propertyName":
		target = v.propertyName
	case "mapping":
		target = v.mapping
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

// DiscriminatorFromJSON constructs a Discriminator from JSON buffer. `dst` must
// be a pointer to `Discriminator`
func DiscriminatorFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Discriminator)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Discriminator, but got %T`, dst)
	}
	var tmp discriminator
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Discriminator`)
	}
	*v = &tmp
	return nil
}
