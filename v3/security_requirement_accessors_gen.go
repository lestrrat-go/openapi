package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

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
	return Visit(context.Background(), newValidator(recurse), v)
}
