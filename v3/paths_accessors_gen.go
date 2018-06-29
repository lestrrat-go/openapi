package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

func (v *paths) Paths() *PathItemMapIterator {
	var items []interface{}
	for key, item := range v.paths {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter PathItemMapIterator
	iter.list.items = items
	return &iter
}

func (v *paths) Reference() string {
	return v.reference
}

func (v *paths) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *paths) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
