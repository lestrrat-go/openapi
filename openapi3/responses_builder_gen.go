package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// ResponsesBuilder is used to build an instance of Responses. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Responses
type ResponsesBuilder struct {
	target *responses
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *ResponsesBuilder) MustBuild(options ...Option) Responses {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Responses and returns the result
func (b *ResponsesBuilder) Build(options ...Option) (Responses, error) {
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

// NewResponses creates a new builder object for Responses
func NewResponses() *ResponsesBuilder {
	return &ResponsesBuilder{
		target: &responses{},
	}
}

// Default sets the Default field for object Responses.
func (b *ResponsesBuilder) Default(v Response) *ResponsesBuilder {
	b.target.defaultValue = v
	return b
}

// Reference sets the $ref (reference) field for object Responses.
func (b *ResponsesBuilder) Reference(v string) *ResponsesBuilder {
	b.target.reference = v
	return b
}
