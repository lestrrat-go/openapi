package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *SecuritySchemeMap) Clear() error {
	*v = make(SecuritySchemeMap)
	return nil
}

// Validate checks the correctness of values in SecuritySchemeMap
func (v *SecuritySchemeMap) Validate(recurse bool) error {
	for name, elem := range *v {
		if validator, ok := elem.(Validator); ok {
			if err := validator.Validate(recurse); err != nil {
				return errors.Wrapf(err, `failed to validate element %v`, name)
			}
		}
	}
	return nil
}

func (v SecuritySchemeMap) QueryJSON(path string) (ret interface{}, ok bool) {
	if path == `` {
		return v, true
	}

	var frag string
	frag, path = extractFragFromPath(path)
	target, ok := v[frag]
	if !ok {
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}

	if path == `` {
		return target, true
	}
	return nil, false
}

func (v *SecuritySchemeMap) UnmarshalJSON(data []byte) error {
	var proxy map[string]*securityScheme
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal SecuritySchemeMap`)
	}
	tmp := make(map[string]SecurityScheme)
	for name, value := range proxy {
		tmp[name] = value
	}
	*v = tmp
	return nil
}
