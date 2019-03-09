package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *SchemaList) Clear() error {
	*v = SchemaList(nil)
	return nil
}

// Validate checks for the values for correctness. If `recurse`
// is specified, child elements are also validated
func (v *SchemaList) Validate(recurse bool) error {
	for i, elem := range *v {
		if validator, ok := elem.(Validator); ok {
			if err := validator.Validate(recurse); err != nil {
				return errors.Wrapf(err, `failed to validate element %d`, i)
			}
		}
	}
	return nil
}

func (v *SchemaList) UnmarshalJSON(data []byte) error {
	var proxy []*schema
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal SchemaList`)
	}

	if len(proxy) == 0 {
		*v = SchemaList(nil)
		return nil
	}

	tmp := make(SchemaList, len(proxy))
	for i, value := range proxy {
		tmp[i] = value
	}
	*v = tmp
	return nil
}
