package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = json.Unmarshal
var _ = errors.Cause

// Clear removes all values from ParameterList
func (v *ParameterList) Clear() error {
	*v = ParameterList(nil)
	return nil
}

// Validate checks for the values for correctness. If `recurse`
// is specified, child elements are also validated
func (v *ParameterList) Validate(recurse bool) error {
	for i, elem := range *v {
		if validator, ok := elem.(Validator); ok {
			if err := validator.Validate(recurse); err != nil {
				return errors.Wrapf(err, `failed to validate element %d`, i)
			}
		}
	}
	return nil
}

// UnmarshalJSON defines how ParameterList is deserialized from JSON
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
