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

// Clear removes all values from TagList
func (v *TagList) Clear() error {
	*v = TagList(nil)
	return nil
}

// Validate checks for the values for correctness. If `recurse`
// is specified, child elements are also validated
func (v *TagList) Validate(recurse bool) error {
	for i, elem := range *v {
		if validator, ok := elem.(Validator); ok {
			if err := validator.Validate(recurse); err != nil {
				return errors.Wrapf(err, `failed to validate element %d`, i)
			}
		}
	}
	return nil
}

// UnmarshalJSON defines how TagList is deserialized from JSON
func (v *TagList) UnmarshalJSON(data []byte) error {
	var proxy []*tag
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}

	if len(proxy) == 0 {
		*v = TagList(nil)
		return nil
	}

	tmp := make(TagList, len(proxy))
	for i, value := range proxy {
		tmp[i] = value
	}
	*v = tmp
	return nil
}
