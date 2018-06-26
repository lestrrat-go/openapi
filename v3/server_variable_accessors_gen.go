package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *serverVariable) Name() string {
	return v.name
}

func (v *serverVariable) Enum() *StringListIterator {
	var items []interface{}
	for _, item := range v.enum {
		items = append(items, item)
	}
	var iter StringListIterator
	iter.items = items
	return &iter
}

func (v *serverVariable) Default() string {
	return v.defaultValue
}

func (v *serverVariable) Description() string {
	return v.description
}

func (v *serverVariable) Reference() string {
	return v.reference
}

func (v *serverVariable) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *serverVariable) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *serverVariable) recurseValidate() error {
	if elem := v.enum; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "enum"`)
		}
	}
	return nil
}
