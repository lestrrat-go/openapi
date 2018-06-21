package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *ParameterList) Clear() error {
	*v = ParameterList(nil)
	return nil
}

func (v ParameterList) Resolve(resolver *Resolver) error {
	if len(v) > 0 {
		for i, elem := range v {
			if err := elem.Resolve(resolver); err != nil {
				return errors.Wrapf(err, `failed to resolve ParameterList (index = %d)`, i)
			}
		}
	}
	return nil
}

func (v *ParameterList) UnmarshalJSON(data []byte) error {
	var proxy []*parameter
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}

	if len(proxy) == 0 {
		*v = ParameterList(nil)
		return nil
	}

	tmp := make(ParameterList, len(proxy))
	for i, value := range proxy {
		tmp[i] = value
	}
	*v = tmp
	return nil
}
