package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// SchemaBuilder is used to build an instance of Schema. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Schema
type SchemaBuilder struct {
	target *schema
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *SchemaBuilder) MustBuild(options ...Option) Schema {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Schema and returns the result
func (b *SchemaBuilder) Build(options ...Option) (Schema, error) {
	validate := true
	for _, option := range options {
		switch option.Name() {
		case optkeyValidate:
			validate = option.Value().(bool)
		}
	}
	if validate {
		if err := b.target.Validate(false); err != nil {
			return nil, errors.Wrap(err, `validation failed`)
		}
	}
	return b.target, nil
}

// NewSchema creates a new builder object for Schema
func NewSchema() *SchemaBuilder {
	return &SchemaBuilder{
		target: &schema{},
	}
}

// Title sets the Title field for object Schema.
func (b *SchemaBuilder) Title(v string) *SchemaBuilder {
	b.target.title = v
	return b
}

// MultipleOf sets the MultipleOf field for object Schema.
func (b *SchemaBuilder) MultipleOf(v float64) *SchemaBuilder {
	b.target.multipleOf = v
	return b
}

// Maximum sets the Maximum field for object Schema.
func (b *SchemaBuilder) Maximum(v float64) *SchemaBuilder {
	b.target.maximum = v
	return b
}

// ExclusiveMaximum sets the ExclusiveMaximum field for object Schema.
func (b *SchemaBuilder) ExclusiveMaximum(v float64) *SchemaBuilder {
	b.target.exclusiveMaximum = v
	return b
}

// Minimum sets the Minimum field for object Schema.
func (b *SchemaBuilder) Minimum(v float64) *SchemaBuilder {
	b.target.minimum = v
	return b
}

// ExclusiveMinimum sets the ExclusiveMinimum field for object Schema.
func (b *SchemaBuilder) ExclusiveMinimum(v float64) *SchemaBuilder {
	b.target.exclusiveMinimum = v
	return b
}

// MaxLength sets the MaxLength field for object Schema.
func (b *SchemaBuilder) MaxLength(v int) *SchemaBuilder {
	b.target.maxLength = v
	return b
}

// MinLength sets the MinLength field for object Schema.
func (b *SchemaBuilder) MinLength(v int) *SchemaBuilder {
	b.target.minLength = v
	return b
}

// Pattern sets the Pattern field for object Schema.
func (b *SchemaBuilder) Pattern(v string) *SchemaBuilder {
	b.target.pattern = v
	return b
}

// MaxItems sets the MaxItems field for object Schema.
func (b *SchemaBuilder) MaxItems(v int) *SchemaBuilder {
	b.target.maxItems = v
	return b
}

// MinItems sets the MinItems field for object Schema.
func (b *SchemaBuilder) MinItems(v int) *SchemaBuilder {
	b.target.minItems = v
	return b
}

// UniqueItems sets the UniqueItems field for object Schema.
func (b *SchemaBuilder) UniqueItems(v bool) *SchemaBuilder {
	b.target.uniqueItems = v
	return b
}

// MaxProperties sets the MaxProperties field for object Schema.
func (b *SchemaBuilder) MaxProperties(v int) *SchemaBuilder {
	b.target.maxProperties = v
	return b
}

// MinProperties sets the MinProperties field for object Schema.
func (b *SchemaBuilder) MinProperties(v int) *SchemaBuilder {
	b.target.minProperties = v
	return b
}

// Required sets the Required field for object Schema.
func (b *SchemaBuilder) Required(v []string) *SchemaBuilder {
	b.target.required = v
	return b
}

// Enum sets the Enum field for object Schema.
func (b *SchemaBuilder) Enum(v []interface{}) *SchemaBuilder {
	b.target.enum = v
	return b
}

// Type sets the Type field for object Schema.
func (b *SchemaBuilder) Type(v PrimitiveType) *SchemaBuilder {
	b.target.typ = v
	return b
}

// AllOf sets the AllOf field for object Schema.
func (b *SchemaBuilder) AllOf(v []Schema) *SchemaBuilder {
	b.target.allOf = v
	return b
}

// OneOf sets the OneOf field for object Schema.
func (b *SchemaBuilder) OneOf(v []Schema) *SchemaBuilder {
	b.target.oneOf = v
	return b
}

// AnyOf sets the AnyOf field for object Schema.
func (b *SchemaBuilder) AnyOf(v []Schema) *SchemaBuilder {
	b.target.anyOf = v
	return b
}

// Not sets the Not field for object Schema.
func (b *SchemaBuilder) Not(v Schema) *SchemaBuilder {
	b.target.not = v
	return b
}

// Items sets the Items field for object Schema.
func (b *SchemaBuilder) Items(v Schema) *SchemaBuilder {
	b.target.items = v
	return b
}

// Properties sets the Properties field for object Schema.
func (b *SchemaBuilder) Properties(v map[string]Schema) *SchemaBuilder {
	b.target.properties = v
	return b
}

// Format sets the Format field for object Schema.
func (b *SchemaBuilder) Format(v string) *SchemaBuilder {
	b.target.format = v
	return b
}

// Default sets the Default field for object Schema.
func (b *SchemaBuilder) Default(v interface{}) *SchemaBuilder {
	b.target.defaultValue = v
	return b
}

// Nullable sets the Nullable field for object Schema.
func (b *SchemaBuilder) Nullable(v bool) *SchemaBuilder {
	b.target.nullable = v
	return b
}

// Discriminator sets the Discriminator field for object Schema.
func (b *SchemaBuilder) Discriminator(v Discriminator) *SchemaBuilder {
	b.target.discriminator = v
	return b
}

// ReadOnly sets the ReadOnly field for object Schema.
func (b *SchemaBuilder) ReadOnly(v bool) *SchemaBuilder {
	b.target.readOnly = v
	return b
}

// WriteOnly sets the WriteOnly field for object Schema.
func (b *SchemaBuilder) WriteOnly(v bool) *SchemaBuilder {
	b.target.writeOnly = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object Schema.
func (b *SchemaBuilder) ExternalDocs(v ExternalDocumentation) *SchemaBuilder {
	b.target.externalDocs = v
	return b
}

// Example sets the Example field for object Schema.
func (b *SchemaBuilder) Example(v interface{}) *SchemaBuilder {
	b.target.example = v
	return b
}

// Deprecated sets the Deprecated field for object Schema.
func (b *SchemaBuilder) Deprecated(v bool) *SchemaBuilder {
	b.target.deprecated = v
	return b
}

// Reference sets the $ref (reference) field for object Schema.
func (b *SchemaBuilder) Reference(v string) *SchemaBuilder {
	b.target.reference = v
	return b
}
