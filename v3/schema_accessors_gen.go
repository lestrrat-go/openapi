package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *schema) Name() string {
	return v.name
}

func (v *schema) Title() string {
	return v.title
}

func (v *schema) MultipleOf() float64 {
	return v.multipleOf
}

func (v *schema) Maximum() float64 {
	return v.maximum
}

func (v *schema) ExclusiveMaximum() float64 {
	return v.exclusiveMaximum
}

func (v *schema) Minimum() float64 {
	return v.minimum
}

func (v *schema) ExclusiveMinimum() float64 {
	return v.exclusiveMinimum
}

func (v *schema) MaxLength() int {
	return v.maxLength
}

func (v *schema) MinLength() int {
	return v.minLength
}

func (v *schema) Pattern() string {
	return v.pattern
}

func (v *schema) MaxItems() int {
	return v.maxItems
}

func (v *schema) MinItems() int {
	return v.minItems
}

func (v *schema) UniqueItems() bool {
	return v.uniqueItems
}

func (v *schema) MaxProperties() int {
	return v.maxProperties
}

func (v *schema) MinProperties() int {
	return v.minProperties
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

func (v *schema) OneOf() *SchemaListIterator {
	var items []interface{}
	for _, item := range v.oneOf {
		items = append(items, item)
	}
	var iter SchemaListIterator
	iter.items = items
	return &iter
}

func (v *schema) AnyOf() *SchemaListIterator {
	var items []interface{}
	for _, item := range v.anyOf {
		items = append(items, item)
	}
	var iter SchemaListIterator
	iter.items = items
	return &iter
}

func (v *schema) Not() Schema {
	return v.not
}

func (v *schema) Items() Schema {
	return v.items
}

func (v *schema) Properties() *SchemaMapIterator {
	var items []interface{}
	for key, item := range v.properties {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SchemaMapIterator
	iter.list.items = items
	return &iter
}

func (v *schema) Format() string {
	return v.format
}

func (v *schema) Default() interface{} {
	return v.defaultValue
}

func (v *schema) Nullable() bool {
	return v.nullable
}

func (v *schema) Discriminator() Discriminator {
	return v.discriminator
}

func (v *schema) ReadOnly() bool {
	return v.readOnly
}

func (v *schema) WriteOnly() bool {
	return v.writeOnly
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

func (v *schema) Reference() string {
	return v.reference
}

func (v *schema) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *schema) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *schema) recurseValidate() error {
	if elem := v.required; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "required"`)
		}
	}
	if elem := v.enum; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "enum"`)
		}
	}
	if elem := v.allOf; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "allOf"`)
		}
	}
	if elem := v.oneOf; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "oneOf"`)
		}
	}
	if elem := v.anyOf; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "anyOf"`)
		}
	}
	if elem := v.not; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "not"`)
		}
	}
	if elem := v.items; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "items"`)
		}
	}
	if elem := v.properties; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "properties"`)
		}
	}
	if elem := v.discriminator; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "discriminator"`)
		}
	}
	if elem := v.externalDocs; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "externalDocs"`)
		}
	}
	return nil
}
