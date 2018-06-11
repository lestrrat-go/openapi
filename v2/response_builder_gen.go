package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// ResponseBuilder is used to build an instance of Response. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Response
type ResponseBuilder struct {
	target *response
}

// Do finalizes the building process for Response and returns the result
func (b *ResponseBuilder) Do(options ...Option) (Response, error) {
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

// Name sets the Name field for object Response.
func (b *ResponseBuilder) Name(v string) *ResponseBuilder {
	b.target.name = v
	return b
}

// StatusCode sets the StatusCode field for object Response.
func (b *ResponseBuilder) StatusCode(v string) *ResponseBuilder {
	b.target.statusCode = v
	return b
}

// Schema sets the Schema field for object Response.
func (b *ResponseBuilder) Schema(v Schema) *ResponseBuilder {
	b.target.schema = v
	return b
}

// Headers sets the Headers field for object Response.
func (b *ResponseBuilder) Headers(v HeaderMap) *ResponseBuilder {
	b.target.headers = v
	return b
}

// Examples sets the Examples field for object Response.
func (b *ResponseBuilder) Examples(v ExampleMap) *ResponseBuilder {
	b.target.examples = v
	return b
}

// Reference sets the $ref (reference) field for object Response.
func (b *ResponseBuilder) Reference(v string) *ResponseBuilder {
	b.target.reference = v
	return b
}
