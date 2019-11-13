package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// HeaderBuilder is used to build an instance of Header. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Header
type HeaderBuilder struct {
	target *header
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *HeaderBuilder) MustBuild(options ...Option) Header {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Header and returns the result
func (b *HeaderBuilder) Build(options ...Option) (Header, error) {
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

// NewHeader creates a new builder object for Header
func NewHeader() *HeaderBuilder {
	return &HeaderBuilder{
		target: &header{
			in: InHeader,
		},
	}
}

// In sets the In field for object Header. If this is not called,
// a default value (InHeader) is assigned to this field
func (b *HeaderBuilder) In(v Location) *HeaderBuilder {
	b.target.in = v
	return b
}

// Required sets the Required field for object Header.
func (b *HeaderBuilder) Required(v bool) *HeaderBuilder {
	b.target.required = v
	return b
}

// Description sets the Description field for object Header.
func (b *HeaderBuilder) Description(v string) *HeaderBuilder {
	b.target.description = v
	return b
}

// Deprecated sets the Deprecated field for object Header.
func (b *HeaderBuilder) Deprecated(v bool) *HeaderBuilder {
	b.target.deprecated = v
	return b
}

// AllowEmptyValue sets the AllowEmptyValue field for object Header.
func (b *HeaderBuilder) AllowEmptyValue(v bool) *HeaderBuilder {
	b.target.allowEmptyValue = v
	return b
}

// Explode sets the Explode field for object Header.
func (b *HeaderBuilder) Explode(v bool) *HeaderBuilder {
	b.target.explode = v
	return b
}

// AllowReserved sets the AllowReserved field for object Header.
func (b *HeaderBuilder) AllowReserved(v bool) *HeaderBuilder {
	b.target.allowReserved = v
	return b
}

// Schema sets the Schema field for object Header.
func (b *HeaderBuilder) Schema(v Schema) *HeaderBuilder {
	b.target.schema = v
	return b
}

// Examples sets the Examples field for object Header.
func (b *HeaderBuilder) Examples(v map[string]Example) *HeaderBuilder {
	b.target.examples = v
	return b
}

// Content sets the Content field for object Header.
func (b *HeaderBuilder) Content(v map[string]MediaType) *HeaderBuilder {
	b.target.content = v
	return b
}

// Reference sets the $ref (reference) field for object Header.
func (b *HeaderBuilder) Reference(v string) *HeaderBuilder {
	b.target.reference = v
	return b
}
