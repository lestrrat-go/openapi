package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *SecurityRequirementList) Clear() error {
	*v = SecurityRequirementList(nil)
	return nil
}

// Validate checks for the values for correctness. If `recurse`
// is specified, child elements are also validated
func (v *SecurityRequirementList) Validate(recurse bool) error {
	for i, elem := range *v {
		if validator, ok := elem.(Validator); ok {
			if err := validator.Validate(recurse); err != nil {
				return errors.Wrapf(err, `failed to validate element %d`, i)
			}
		}
	}
	return nil
}

func (v *SecurityRequirementList) UnmarshalJSON(data []byte) error {
	var proxy []*securityRequirement
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}

	if len(proxy) == 0 {
		*v = SecurityRequirementList(nil)
		return nil
	}

	tmp := make(SecurityRequirementList, len(proxy))
	for i, value := range proxy {
		tmp[i] = value
	}
	*v = tmp
	return nil
}
