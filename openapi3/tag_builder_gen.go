package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// TagBuilder is used to build an instance of Tag. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Tag
type TagBuilder struct {
	target *tag
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *TagBuilder) MustBuild(options ...Option) Tag {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Tag and returns the result
func (b *TagBuilder) Build(options ...Option) (Tag, error) {
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

// NewTag creates a new builder object for Tag
func NewTag(name string) *TagBuilder {
	return &TagBuilder{
		target: &tag{
			name: name,
		},
	}
}

// Description sets the Description field for object Tag.
func (b *TagBuilder) Description(v string) *TagBuilder {
	b.target.description = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object Tag.
func (b *TagBuilder) ExternalDocs(v ExternalDocumentation) *TagBuilder {
	b.target.externalDocs = v
	return b
}

// Reference sets the $ref (reference) field for object Tag.
func (b *TagBuilder) Reference(v string) *TagBuilder {
	b.target.reference = v
	return b
}
