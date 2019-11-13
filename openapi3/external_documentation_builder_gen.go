package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// ExternalDocumentationBuilder is used to build an instance of ExternalDocumentation. The user must
// call `Build()` after providing all the necessary information to
// build an instance of ExternalDocumentation
type ExternalDocumentationBuilder struct {
	target *externalDocumentation
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *ExternalDocumentationBuilder) MustBuild(options ...Option) ExternalDocumentation {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for ExternalDocumentation and returns the result
func (b *ExternalDocumentationBuilder) Build(options ...Option) (ExternalDocumentation, error) {
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

// NewExternalDocumentation creates a new builder object for ExternalDocumentation
func NewExternalDocumentation(url string) *ExternalDocumentationBuilder {
	return &ExternalDocumentationBuilder{
		target: &externalDocumentation{
			url: url,
		},
	}
}

// Description sets the Description field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Description(v string) *ExternalDocumentationBuilder {
	b.target.description = v
	return b
}

// Reference sets the $ref (reference) field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Reference(v string) *ExternalDocumentationBuilder {
	b.target.reference = v
	return b
}
