package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"
	"sort"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = sort.Strings
var _ = errors.Cause

func (v *paths) Paths() *PathItemMapIterator {
	var keys []string
	for key := range v.paths {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.paths[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter PathItemMapIterator
	iter.list.items = items
	return &iter
}

// Reference returns the value of `$ref` field
func (v *paths) Reference() string {
	return v.reference
}

func (v *paths) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *paths) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *paths) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *paths) Validate(recurse bool) error {
	return newValidator(recurse).Validate(context.Background(), v)
}
