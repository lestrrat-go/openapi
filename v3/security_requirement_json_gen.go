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

type securityRequirementMarshalProxy struct {
	Reference string        `json:"$ref,omitempty"`
	Schemes   StringListMap `json:""`
}

type securityRequirementUnmarshalProxy struct {
	Reference string        `json:"$ref,omitempty"`
	Schemes   StringListMap `json:""`
}

func (v *securityRequirement) MarshalJSON() ([]byte, error) {
	var proxy securityRequirementMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Schemes = v.schemes
	return json.Marshal(proxy)
}

func (v *securityRequirement) UnmarshalJSON(data []byte) error {
	var proxy securityRequirementUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrapf(err, `failed to unmarshal securityRequirement`)
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.schemes = proxy.Schemes
	return nil
}

func (v *securityRequirement) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "schemes":
		target = v.schemes
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

// SecurityRequirementFromJSON constructs a SecurityRequirement from JSON buffer. `dst` must
// be a pointer to `SecurityRequirement`
func SecurityRequirementFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*SecurityRequirement)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to SecurityRequirement, but got %T`, dst)
	}
	var tmp securityRequirement
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal SecurityRequirement`)
	}
	*v = &tmp
	return nil
}
