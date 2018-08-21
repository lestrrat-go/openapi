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

func (v *response) Name() string {
	return v.name
}

func (v *response) StatusCode() string {
	return v.statusCode
}

func (v *response) Description() string {
	return v.description
}

func (v *response) Schema() Schema {
	return v.schema
}

func (v *response) Headers() *HeaderMapIterator {
	var keys []string
	for key := range v.headers {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.headers[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter HeaderMapIterator
	iter.list.items = items
	return &iter
}

func (v *response) Examples() *ExampleMapIterator {
	var keys []string
	for key := range v.examples {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.examples[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExampleMapIterator
	iter.list.items = items
	return &iter
}

// Reference returns the value of `$ref` field
func (v *response) Reference() string {
	return v.reference
}

func (v *response) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *response) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *response) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *response) Validate(recurse bool) error {
	return newValidator(recurse).Validate(context.Background(), v)
}
