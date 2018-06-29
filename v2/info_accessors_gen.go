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

func (v *info) Title() string {
	return v.title
}

func (v *info) Version() string {
	return v.version
}

func (v *info) Description() string {
	return v.description
}

func (v *info) TermsOfService() string {
	return v.termsOfService
}

func (v *info) Contact() Contact {
	return v.contact
}

func (v *info) License() License {
	return v.license
}

// Reference returns the value of `$ref` field
func (v *info) Reference() string {
	return v.reference
}

func (v *info) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *info) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *info) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *info) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
