package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// EncodingBuilder is used to build an instance of Encoding. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Encoding
type EncodingBuilder struct {
	target *encoding
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *EncodingBuilder) MustBuild(options ...Option) Encoding {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Encoding and returns the result
func (b *EncodingBuilder) Build(options ...Option) (Encoding, error) {
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

// NewEncoding creates a new builder object for Encoding
func NewEncoding() *EncodingBuilder {
	return &EncodingBuilder{
		target: &encoding{},
	}
}

// ContentType sets the ContentType field for object Encoding.
func (b *EncodingBuilder) ContentType(v string) *EncodingBuilder {
	b.target.contentType = v
	return b
}

// Headers sets the Headers field for object Encoding.
func (b *EncodingBuilder) Headers(v map[string]Header) *EncodingBuilder {
	b.target.headers = v
	return b
}

// Explode sets the Explode field for object Encoding.
func (b *EncodingBuilder) Explode(v bool) *EncodingBuilder {
	b.target.explode = v
	return b
}

// AllowReserved sets the AllowReserved field for object Encoding.
func (b *EncodingBuilder) AllowReserved(v bool) *EncodingBuilder {
	b.target.allowReserved = v
	return b
}

// Reference sets the $ref (reference) field for object Encoding.
func (b *EncodingBuilder) Reference(v string) *EncodingBuilder {
	b.target.reference = v
	return b
}
