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

func (v *pathItem) Name() string {
	return v.name
}

func (v *pathItem) Path() string {
	return v.path
}

func (v *pathItem) Get() Operation {
	return v.get
}

func (v *pathItem) Put() Operation {
	return v.put
}

func (v *pathItem) Post() Operation {
	return v.post
}

func (v *pathItem) Delete() Operation {
	return v.delete
}

func (v *pathItem) Options() Operation {
	return v.options
}

func (v *pathItem) Head() Operation {
	return v.head
}

func (v *pathItem) Patch() Operation {
	return v.patch
}

func (v *pathItem) Parameters() *ParameterListIterator {
	var items []interface{}
	for _, item := range v.parameters {
		items = append(items, item)
	}
	var iter ParameterListIterator
	iter.items = items
	return &iter
}

// Reference returns the value of `$ref` field
func (v *pathItem) Reference() string {
	return v.reference
}

func (v *pathItem) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *pathItem) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *pathItem) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *pathItem) Validate(recurse bool) error {
	return newValidator(recurse).Validate(context.Background(), v)
}
