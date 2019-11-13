package openapi2

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

func (v *externalDocumentation) IsValid() bool {
	return v != nil
}

func (v *externalDocumentation) URL() string {
	return v.url
}

func (v *externalDocumentation) Description() string {
	return v.description
}

// Reference returns the value of `$ref` field
func (v *externalDocumentation) Reference() string {
	return v.reference
}

func (v *externalDocumentation) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *externalDocumentation) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *externalDocumentation) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.size = len(items)
	iter.list.items = items
	return &iter
}

func (v *externalDocumentation) Validate(recurse bool) error {
	return newValidator(recurse).Validate(context.Background(), v)
}
