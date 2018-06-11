package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// ParameterBuilder is used to build an instance of Parameter. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Parameter
type ParameterBuilder struct {
	target *parameter
}

// Do finalizes the building process for Parameter and returns the result
func (b *ParameterBuilder) Do(options ...Option) (Parameter, error) {
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

// NewParameter creates a new builder object for Parameter
func NewParameter(name string, in Location) *ParameterBuilder {
	return &ParameterBuilder{
		target: &parameter{
			name: name,
			in:   in,
		},
	}
}

// Description sets the Description field for object Parameter.
func (b *ParameterBuilder) Description(v string) *ParameterBuilder {
	b.target.description = v
	return b
}

// Required sets the Required field for object Parameter.
func (b *ParameterBuilder) Required(v bool) *ParameterBuilder {
	b.target.required = v
	return b
}

// Schema sets the Schema field for object Parameter.
func (b *ParameterBuilder) Schema(v Schema) *ParameterBuilder {
	b.target.schema = v
	return b
}

// Type sets the Type field for object Parameter.
func (b *ParameterBuilder) Type(v PrimitiveType) *ParameterBuilder {
	b.target.typ = v
	return b
}

// Format sets the Format field for object Parameter.
func (b *ParameterBuilder) Format(v string) *ParameterBuilder {
	b.target.format = v
	return b
}

// Title sets the Title field for object Parameter.
func (b *ParameterBuilder) Title(v string) *ParameterBuilder {
	b.target.title = v
	return b
}

// AllowEmptyValue sets the AllowEmptyValue field for object Parameter.
func (b *ParameterBuilder) AllowEmptyValue(v bool) *ParameterBuilder {
	b.target.allowEmptyValue = v
	return b
}

// Items sets the Items field for object Parameter.
func (b *ParameterBuilder) Items(v Items) *ParameterBuilder {
	b.target.items = v
	return b
}

// CollectionFormat sets the CollectionFormat field for object Parameter.
func (b *ParameterBuilder) CollectionFormat(v CollectionFormat) *ParameterBuilder {
	b.target.collectionFormat = v
	return b
}

// DefaultValue sets the DefaultValue field for object Parameter.
func (b *ParameterBuilder) DefaultValue(v interface{}) *ParameterBuilder {
	b.target.defaultValue = v
	return b
}

// Maximum sets the Maximum field for object Parameter.
func (b *ParameterBuilder) Maximum(v float64) *ParameterBuilder {
	b.target.maximum = v
	return b
}

// ExclusiveMaximum sets the ExclusiveMaximum field for object Parameter.
func (b *ParameterBuilder) ExclusiveMaximum(v float64) *ParameterBuilder {
	b.target.exclusiveMaximum = v
	return b
}

// Minimum sets the Minimum field for object Parameter.
func (b *ParameterBuilder) Minimum(v float64) *ParameterBuilder {
	b.target.minimum = v
	return b
}

// ExclusiveMinimum sets the ExclusiveMinimum field for object Parameter.
func (b *ParameterBuilder) ExclusiveMinimum(v float64) *ParameterBuilder {
	b.target.exclusiveMinimum = v
	return b
}

// MaxLength sets the MaxLength field for object Parameter.
func (b *ParameterBuilder) MaxLength(v int) *ParameterBuilder {
	b.target.maxLength = v
	return b
}

// MinLength sets the MinLength field for object Parameter.
func (b *ParameterBuilder) MinLength(v int) *ParameterBuilder {
	b.target.minLength = v
	return b
}

// Pattern sets the Pattern field for object Parameter.
func (b *ParameterBuilder) Pattern(v string) *ParameterBuilder {
	b.target.pattern = v
	return b
}

// MaxItems sets the MaxItems field for object Parameter.
func (b *ParameterBuilder) MaxItems(v int) *ParameterBuilder {
	b.target.maxItems = v
	return b
}

// MinItems sets the MinItems field for object Parameter.
func (b *ParameterBuilder) MinItems(v int) *ParameterBuilder {
	b.target.minItems = v
	return b
}

// UniqueItems sets the UniqueItems field for object Parameter.
func (b *ParameterBuilder) UniqueItems(v bool) *ParameterBuilder {
	b.target.uniqueItems = v
	return b
}

// Enum sets the Enum field for object Parameter.
func (b *ParameterBuilder) Enum(v ...interface{}) *ParameterBuilder {
	b.target.enum = v
	return b
}

// MultipleOf sets the MultipleOf field for object Parameter.
func (b *ParameterBuilder) MultipleOf(v float64) *ParameterBuilder {
	b.target.multipleOf = v
	return b
}

// Reference sets the $ref (reference) field for object Parameter.
func (b *ParameterBuilder) Reference(v string) *ParameterBuilder {
	b.target.reference = v
	return b
}
