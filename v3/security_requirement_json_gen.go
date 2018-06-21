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
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.schemes = proxy.Schemes
	return nil
}

func (v *securityRequirement) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*securityRequirement)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type SecurityRequirement, but got %T`, resolved)
		}
		mutator := MutateSecurityRequirement(v)
		for iter := asserted.Schemes(); iter.Next(); {
			key, item := iter.Item()
			mutator.Scheme(key, item)
		}
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.schemes != nil {
		if err := v.schemes.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Schemes`)
		}
	}
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
