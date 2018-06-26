package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *responses) Default() Response {
	return v.defaultValue
}

func (v *responses) Responses() *ResponseMapIterator {
	var items []interface{}
	for key, item := range v.responses {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ResponseMapIterator
	iter.list.items = items
	return &iter
}

func (v *responses) Reference() string {
	return v.reference
}

func (v *responses) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *responses) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *responses) recurseValidate() error {
	if elem := v.defaultValue; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "defaultValue"`)
		}
	}
	return nil
}
