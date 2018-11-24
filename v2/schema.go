package openapi

import (
	"github.com/pkg/errors"
)

func (v *schema) setName(s string) {
	v.name = s
}

func (v *schema) IsRequiredProperty(prop string) bool {
	for _, name := range v.required {
		if name == prop {
			return true
		}
	}
	return false
}

// ConvertToSchema fulfills the SchemaCoverter interface.
// schema is already a Schema object, but it's useful to
// align the interface with other Schema-like objects.
// This method just returns itself
func (v *schema) ConvertToSchema() (Schema, error) {
	return v, nil
}

func isSupportedTypeForMergeSchemas(t PrimitiveType) error {
	switch t {
	case Object, Array:
		return nil
	default:
		return errors.Errorf(`primitive type %s is not supported for merging`, t)
	}
}

func chooseStringForMergeSchemas(left, right string) (string, error) {
	if left == right {
		return left, nil
	}

	if left == "" && right != "" {
		return right, nil
	}

	if left != "" && right == "" {
		return left, nil
	}

	return "", errors.Errorf(`both %s and %s are specified`, left, right)
}

func chooseIntForMergeSchemas(hasLeft bool, left int, hasRight bool, right int) (int, bool, error) {
	if !hasLeft && !hasRight {
		return 0, false, nil
	}

	if left == right {
		return left, true, nil
	}

	return 0, false, errors.Errorf(`both values are specified but do not match: %d <-> %d`, left, right)
}

func chooseBoolForMergeSchemas(hasLeft, left, hasRight, right bool) (bool, bool, error) {
	if !hasLeft && !hasRight {
		return false, false, nil
	}

	if left == right {
		return left, true, nil
	}

	return false, false, errors.Errorf(`both values are specified but do not match: %t <-> %t`, left, right)
}

// Merges two schemas, and creates a new one.
//
// Schemas that have JSON references are not supported.
//
// Only objects and arrays are supported. Both of the specified
// schemas must be of the same primitive type.
//
// If there are ambiquities such as duplicate fields defined
// in both schemas, it returns an error.
//
// The resulting schema is always unnamed.
func MergeSchemas(left, right Schema) (Schema, error) {
	if left == nil || !left.IsValid() {
		return right, nil
	}

	if right == nil || !right.IsValid() {
		return left, nil
	}

	if left.Reference() != "" || right.Reference() != "" {
		return nil, errors.Errorf(`references are not supported`)
	}

	leftType := GuessSchemaType(left)
	rightType := GuessSchemaType(right)
	if err := isSupportedTypeForMergeSchemas(leftType); err != nil {
		return nil, errors.Wrap(err, `type of left schema to be merged is not supported`)
	}

	if err := isSupportedTypeForMergeSchemas(rightType); err != nil {
		return nil, errors.Wrap(err, `type of right schema to be merged is not supported`)
	}

	if leftType != rightType {
		return nil, errors.Errorf(`primitive types of left and right schema do not match`)
	}

	builder := NewSchema()
	builder.Type(left.Type())

	switch GuessSchemaType(left) {
	case Object:
		if err := mergeSchemasObjectAttributes(builder, left, right); err != nil {
			return nil, errors.Wrap(err, `failed to merge object attributes`)
		}
	case Array:
		if err := mergeSchemasArrayAttributes(builder, left, right); err != nil {
			return nil, errors.Wrap(err, `failed to merge array attributes`)
		}
	}

	return builder.Build()
}

func mergeSchemasObjectAttributes(builder *SchemaBuilder, left, right Schema) error {
	maxProperties, hasMaxProperties, err := chooseIntForMergeSchemas(left.HasMaxProperties(), left.MaxProperties(), right.HasMaxProperties(), right.MaxProperties())
	if err != nil {
		return errors.Wrap(err, `failed to merge maxProperties`)
	}

	if hasMaxProperties {
		builder.MaxProperties(maxProperties)
	}

	minProperties, hasMinProperties, err := chooseIntForMergeSchemas(left.HasMinProperties(), left.MinProperties(), right.HasMinProperties(), right.MinProperties())
	if err != nil {
		return errors.Wrap(err, `failed to merge minProperties`)
	}

	if hasMinProperties {
		builder.MinProperties(minProperties)
	}

	var required []string
	for iter := left.Required(); iter.Next(); {
		required = append(required, iter.Item())
	}
	for iter := right.Required(); iter.Next(); {
		required = append(required, iter.Item())
	}
	if len(required) > 0 {
		builder.Required(required...)
	}

	props := SchemaMap{}
	for iter := left.Properties(); iter.Next(); {
		name, prop := iter.Item()
		props[name] = prop
	}
	for iter := right.Properties(); iter.Next(); {
		name, prop := iter.Item()
		if _, ok := props[name]; ok {
			return errors.Errorf(`property %s already exists in destination`, name)
		}
		props[name] = prop
	}
	for name, prop := range props {
		builder.Property(name, prop)
	}
	return nil
}

func mergeSchemasArrayAttributes(builder *SchemaBuilder, left, right Schema) error {
	maxItems, hasMaxItems, err := chooseIntForMergeSchemas(left.HasMaxItems(), left.MaxItems(), right.HasMaxItems(), right.MaxItems())
	if err != nil {
		return errors.Wrap(err, `failed to merge maxItems`)
	}
	if hasMaxItems {
		builder.MaxItems(maxItems)
	}

	minItems, hasMinItems, err := chooseIntForMergeSchemas(left.HasMinItems(), left.MinItems(), right.HasMinItems(), right.MinItems())
	if err != nil {
		return errors.Wrap(err, `failed to merge minItems`)
	}
	if hasMinItems {
		builder.MinItems(minItems)
	}

	uniqueItems, hasUniqueItems, err := chooseBoolForMergeSchemas(left.HasUniqueItems(), left.UniqueItems(), right.HasUniqueItems(), right.UniqueItems())
	if hasUniqueItems {
		builder.UniqueItems(uniqueItems)
	}

	/*
		if
		Items() Schema

		AllOf() *SchemaListIterator
		AdditionaProperties() *SchemaMapIterator
		Default() interface{}
		Discriminator() string
		ReadOnly() bool
		ExternalDocs() ExternalDocumentation
		Example() interface{}
		Deprecated() bool
		XML() XML
		Extension(string) (interface{}, bool)
		Extensions() *ExtensionsIterator
		Clone() Schema
		IsUnresolved() bool
		IsRequiredProperty(string) bool
	*/
	return nil
}

// GuessSchemaType returns the type declared in the Schema, or if not
// present, attempts to guess by the presence of certain fields.
// In case it can't determine the type, return Invalid
func GuessSchemaType(s Schema) PrimitiveType {
	if t := s.Type(); t != "" {
		return t
	}

	if iter := s.Properties(); iter.Size() > 0 {
		return Object
	}
	if s.HasMaxProperties() || s.HasMinProperties() {
		return Object
	}

	if t := s.Items(); t != nil {
		return Array
	}
	if s.HasMaxItems() || s.HasMinItems() || s.HasUniqueItems() {
		return Array
	}

	if t := s.Default(); t == "true" || t == "false" {
		return Boolean
	}

	if t := s.Pattern(); t != "" {
		return String
	}
	if s.HasMaxLength() || s.HasMinLength() {
		return String
	}

	if s.HasMultipleOf() || s.HasMaximum() || s.HasExclusiveMaximum() || s.HasMinimum() || s.HasExclusiveMinimum() {
		return Number
	}

	return Invalid
}
