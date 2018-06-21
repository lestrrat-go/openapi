package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *RequestBodyMap) Clear() error {
	*v = make(RequestBodyMap)
	return nil
}

func (v RequestBodyMap) Resolve(resolver *Resolver) error {
	if len(v) > 0 {
		for name, elem := range v {
			if err := elem.Resolve(resolver); err != nil {
				return errors.Wrapf(err, `failed to resolve RequestBodyMap (key = %s)`, name)
			}
		}
	}
	return nil
}

func (v RequestBodyMap) QueryJSON(path string) (ret interface{}, ok bool) {
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

func (v *RequestBodyMap) UnmarshalJSON(data []byte) error {
	var proxy map[string]*requestBody
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}
	tmp := make(map[string]RequestBody)
	for name, value := range proxy {
		value.setName(name)
		tmp[name] = value
	}
	*v = tmp
	return nil
}
