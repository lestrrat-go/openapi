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

func (v *schema) Name() string {
	return v.name
}

func (v *schema) Type() PrimitiveType {
	return v.typ
}

func (v *schema) Format() string {
	return v.format
}

func (v *schema) Title() string {
	return v.title
}

// HasMultipleOf returns true if the value for multipleOf has been
// explicitly specified
func (v *schema) HasMultipleOf() bool {
	return v.multipleOf != nil
}

// MultipleOf returns the value of multipleOf. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) MultipleOf() float64 {
	if !v.HasMultipleOf() {
		return 0
	}
	return *v.multipleOf
}

// HasMaximum returns true if the value for maximum has been
// explicitly specified
func (v *schema) HasMaximum() bool {
	return v.maximum != nil
}

// Maximum returns the value of maximum. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) Maximum() float64 {
	if !v.HasMaximum() {
		return 0
	}
	return *v.maximum
}

// HasExclusiveMaximum returns true if the value for exclusiveMaximum has been
// explicitly specified
func (v *schema) HasExclusiveMaximum() bool {
	return v.exclusiveMaximum != nil
}

// ExclusiveMaximum returns the value of exclusiveMaximum. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) ExclusiveMaximum() float64 {
	if !v.HasExclusiveMaximum() {
		return 0
	}
	return *v.exclusiveMaximum
}

// HasMinimum returns true if the value for minimum has been
// explicitly specified
func (v *schema) HasMinimum() bool {
	return v.minimum != nil
}

// Minimum returns the value of minimum. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) Minimum() float64 {
	if !v.HasMinimum() {
		return 0
	}
	return *v.minimum
}

// HasExclusiveMinimum returns true if the value for exclusiveMinimum has been
// explicitly specified
func (v *schema) HasExclusiveMinimum() bool {
	return v.exclusiveMinimum != nil
}

// ExclusiveMinimum returns the value of exclusiveMinimum. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) ExclusiveMinimum() float64 {
	if !v.HasExclusiveMinimum() {
		return 0
	}
	return *v.exclusiveMinimum
}

// HasMaxLength returns true if the value for maxLength has been
// explicitly specified
func (v *schema) HasMaxLength() bool {
	return v.maxLength != nil
}

// MaxLength returns the value of maxLength. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) MaxLength() int {
	if !v.HasMaxLength() {
		return 0
	}
	return *v.maxLength
}

// HasMinLength returns true if the value for minLength has been
// explicitly specified
func (v *schema) HasMinLength() bool {
	return v.minLength != nil
}

// MinLength returns the value of minLength. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) MinLength() int {
	if !v.HasMinLength() {
		return 0
	}
	return *v.minLength
}

func (v *schema) Pattern() string {
	return v.pattern
}

// HasMaxItems returns true if the value for maxItems has been
// explicitly specified
func (v *schema) HasMaxItems() bool {
	return v.maxItems != nil
}

// MaxItems returns the value of maxItems. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) MaxItems() int {
	if !v.HasMaxItems() {
		return 0
	}
	return *v.maxItems
}

// HasMinItems returns true if the value for minItems has been
// explicitly specified
func (v *schema) HasMinItems() bool {
	return v.minItems != nil
}

// MinItems returns the value of minItems. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) MinItems() int {
	if !v.HasMinItems() {
		return 0
	}
	return *v.minItems
}

func (v *schema) UniqueItems() bool {
	return v.uniqueItems
}

// HasMaxProperties returns true if the value for maxProperties has been
// explicitly specified
func (v *schema) HasMaxProperties() bool {
	return v.maxProperties != nil
}

// MaxProperties returns the value of maxProperties. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) MaxProperties() int {
	if !v.HasMaxProperties() {
		return 0
	}
	return *v.maxProperties
}

// HasMinProperties returns true if the value for minProperties has been
// explicitly specified
func (v *schema) HasMinProperties() bool {
	return v.minProperties != nil
}

// MinProperties returns the value of minProperties. If the value has not
// been explicitly, set, the zero value will be returned
func (v *schema) MinProperties() int {
	if !v.HasMinProperties() {
		return 0
	}
	return *v.minProperties
}

func (v *schema) Required() *StringListIterator {
	var items []interface{}
	for _, item := range v.required {
		items = append(items, item)
	}
	var iter StringListIterator
	iter.items = items
	return &iter
}

func (v *schema) Enum() *InterfaceListIterator {
	var items []interface{}
	for _, item := range v.enum {
		items = append(items, item)
	}
	var iter InterfaceListIterator
	iter.items = items
	return &iter
}

func (v *schema) AllOf() *SchemaListIterator {
	var items []interface{}
	for _, item := range v.allOf {
		items = append(items, item)
	}
	var iter SchemaListIterator
	iter.items = items
	return &iter
}

func (v *schema) Items() Schema {
	return v.items
}

func (v *schema) Properties() *SchemaMapIterator {
	var keys []string
	for key := range v.properties {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.properties[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SchemaMapIterator
	iter.list.items = items
	return &iter
}

func (v *schema) AdditionaProperties() *SchemaMapIterator {
	var keys []string
	for key := range v.additionaProperties {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.additionaProperties[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SchemaMapIterator
	iter.list.items = items
	return &iter
}

func (v *schema) Default() interface{} {
	return v.defaultValue
}

func (v *schema) Discriminator() string {
	return v.discriminator
}

func (v *schema) ReadOnly() bool {
	return v.readOnly
}

func (v *schema) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

func (v *schema) Example() interface{} {
	return v.example
}

func (v *schema) Deprecated() bool {
	return v.deprecated
}

func (v *schema) XML() XML {
	return v.xml
}

// Reference returns the value of `$ref` field
func (v *schema) Reference() string {
	return v.reference
}

func (v *schema) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *schema) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *schema) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *schema) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
