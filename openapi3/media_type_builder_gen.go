package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// MediaTypeBuilder is used to build an instance of MediaType. The user must
// call `Build()` after providing all the necessary information to
// build an instance of MediaType
type MediaTypeBuilder struct {
	target *mediaType
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *MediaTypeBuilder) MustBuild(options ...Option) MediaType {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for MediaType and returns the result
func (b *MediaTypeBuilder) Build(options ...Option) (MediaType, error) {
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

// NewMediaType creates a new builder object for MediaType
func NewMediaType() *MediaTypeBuilder {
	return &MediaTypeBuilder{
		target: &mediaType{},
	}
}

// Schema sets the Schema field for object MediaType.
func (b *MediaTypeBuilder) Schema(v Schema) *MediaTypeBuilder {
	b.target.schema = v
	return b
}

// Examples sets the Examples field for object MediaType.
func (b *MediaTypeBuilder) Examples(v map[string]Example) *MediaTypeBuilder {
	b.target.examples = v
	return b
}

// Encoding sets the Encoding field for object MediaType.
func (b *MediaTypeBuilder) Encoding(v map[string]Encoding) *MediaTypeBuilder {
	b.target.encoding = v
	return b
}

// Reference sets the $ref (reference) field for object MediaType.
func (b *MediaTypeBuilder) Reference(v string) *MediaTypeBuilder {
	b.target.reference = v
	return b
}
