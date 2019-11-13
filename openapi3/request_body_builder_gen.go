package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// RequestBodyBuilder is used to build an instance of RequestBody. The user must
// call `Build()` after providing all the necessary information to
// build an instance of RequestBody
type RequestBodyBuilder struct {
	target *requestBody
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *RequestBodyBuilder) MustBuild(options ...Option) RequestBody {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for RequestBody and returns the result
func (b *RequestBodyBuilder) Build(options ...Option) (RequestBody, error) {
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

// NewRequestBody creates a new builder object for RequestBody
func NewRequestBody() *RequestBodyBuilder {
	return &RequestBodyBuilder{
		target: &requestBody{},
	}
}

// Description sets the Description field for object RequestBody.
func (b *RequestBodyBuilder) Description(v string) *RequestBodyBuilder {
	b.target.description = v
	return b
}

// Required sets the Required field for object RequestBody.
func (b *RequestBodyBuilder) Required(v bool) *RequestBodyBuilder {
	b.target.required = v
	return b
}

// Reference sets the $ref (reference) field for object RequestBody.
func (b *RequestBodyBuilder) Reference(v string) *RequestBodyBuilder {
	b.target.reference = v
	return b
}
