package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *securityRequirement) Schemes() *StringListMapIterator {
	var items []interface{}
	for key, item := range v.schemes {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter StringListMapIterator
	iter.list.items = items
	return &iter
}

func (v *securityRequirement) Reference() string {
	return v.reference
}

func (v *securityRequirement) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *securityRequirement) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *securityRequirement) recurseValidate() error {
	if elem := v.schemes; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "schemes"`)
		}
	}
	return nil
}
