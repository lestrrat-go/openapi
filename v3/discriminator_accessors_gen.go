package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

func (v *discriminator) PropertyName() string {
	return v.propertyName
}

func (v *discriminator) Mapping() *StringMapIterator {
	var items []interface{}
	for key, item := range v.mapping {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter StringMapIterator
	iter.list.items = items
	return &iter
}

func (v *discriminator) Reference() string {
	return v.reference
}

func (v *discriminator) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *discriminator) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
