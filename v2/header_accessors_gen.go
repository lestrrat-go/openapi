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

func (v *header) Name() string {
	return v.name
}

func (v *header) Description() string {
	return v.description
}

func (v *header) Type() string {
	return v.typ
}

func (v *header) Format() string {
	return v.format
}

func (v *header) Items() Items {
	return v.items
}

func (v *header) CollectionFormat() CollectionFormat {
	return v.collectionFormat
}

func (v *header) Default() interface{} {
	return v.defaultValue
}

func (v *header) Maximum() float64 {
	return v.maximum
}

func (v *header) ExclusiveMaximum() float64 {
	return v.exclusiveMaximum
}

func (v *header) Minimum() float64 {
	return v.minimum
}

func (v *header) ExclusiveMinimum() float64 {
	return v.exclusiveMinimum
}

func (v *header) MaxLength() int {
	return v.maxLength
}

func (v *header) MinLength() int {
	return v.minLength
}

func (v *header) Pattern() string {
	return v.pattern
}

func (v *header) MaxItems() int {
	return v.maxItems
}

func (v *header) MinItems() int {
	return v.minItems
}

func (v *header) UniqueItems() bool {
	return v.uniqueItems
}

func (v *header) Enum() *InterfaceListIterator {
	var items []interface{}
	for _, item := range v.enum {
		items = append(items, item)
	}
	var iter InterfaceListIterator
	iter.items = items
	return &iter
}

func (v *header) MultipleOf() float64 {
	return v.multipleOf
}

// Reference returns the value of `$ref` field
func (v *header) Reference() string {
	return v.reference
}

func (v *header) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *header) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *header) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *header) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
