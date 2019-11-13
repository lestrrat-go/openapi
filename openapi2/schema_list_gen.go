package openapi2

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

// Clear removes all values from SchemaList
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

// UnmarshalJSON defines how SchemaList is deserialized from JSON
func (v *SchemaList) UnmarshalJSON(data []byte) error {
	var proxy []*schema
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
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
