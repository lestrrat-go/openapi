package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// ResponseBuilder is used to build an instance of Response. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Response
type ResponseBuilder struct {
	target *response
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *ResponseBuilder) MustBuild(options ...Option) Response {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Response and returns the result
func (b *ResponseBuilder) Build(options ...Option) (Response, error) {
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

// NewResponse creates a new builder object for Response
func NewResponse(description string) *ResponseBuilder {
	return &ResponseBuilder{
		target: &response{
			description: description,
		},
	}
}

// Headers sets the Headers field for object Response.
func (b *ResponseBuilder) Headers(v map[string]Header) *ResponseBuilder {
	b.target.headers = v
	return b
}

// Reference sets the $ref (reference) field for object Response.
func (b *ResponseBuilder) Reference(v string) *ResponseBuilder {
	b.target.reference = v
	return b
}
