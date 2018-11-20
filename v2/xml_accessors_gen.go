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

func (v *xml) IsValid() bool {
	return v != nil
}

func (v *xml) Name() string {
	return v.name
}

func (v *xml) Namespace() string {
	return v.namespace
}

func (v *xml) Prefix() string {
	return v.prefix
}

func (v *xml) Attribute() bool {
	return v.attribute
}

func (v *xml) Wrapped() bool {
	return v.wrapped
}

// Reference returns the value of `$ref` field
func (v *xml) Reference() string {
	return v.reference
}

func (v *xml) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *xml) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *xml) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.size = len(items)
	iter.list.items = items
	return &iter
}

func (v *xml) Validate(recurse bool) error {
	return newValidator(recurse).Validate(context.Background(), v)
}
