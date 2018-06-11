package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
	"sort"
)

var _ = sort.Strings
var _ = errors.Cause

func (v *tag) Name() string {
	return v.name
}

func (v *tag) Description() string {
	return v.description
}

func (v *tag) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

func (v *tag) Reference() string {
	return v.reference
}

func (v *tag) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *tag) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

func (v *tag) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *tag) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *tag) recurseValidate() error {
	return nil
	return nil
}
