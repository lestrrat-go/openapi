package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// LinkBuilder is used to build an instance of Link. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Link
type LinkBuilder struct {
	target *link
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *LinkBuilder) MustBuild(options ...Option) Link {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Link and returns the result
func (b *LinkBuilder) Build(options ...Option) (Link, error) {
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

// NewLink creates a new builder object for Link
func NewLink() *LinkBuilder {
	return &LinkBuilder{
		target: &link{},
	}
}

// OperationRef sets the OperationRef field for object Link.
func (b *LinkBuilder) OperationRef(v string) *LinkBuilder {
	b.target.operationRef = v
	return b
}

// OperationID sets the OperationID field for object Link.
func (b *LinkBuilder) OperationID(v string) *LinkBuilder {
	b.target.operationID = v
	return b
}

// Parameters sets the Parameters field for object Link.
func (b *LinkBuilder) Parameters(v map[string]interface{}) *LinkBuilder {
	b.target.parameters = v
	return b
}

// RequestBody sets the RequestBody field for object Link.
func (b *LinkBuilder) RequestBody(v interface{}) *LinkBuilder {
	b.target.requestBody = v
	return b
}

// Description sets the Description field for object Link.
func (b *LinkBuilder) Description(v string) *LinkBuilder {
	b.target.description = v
	return b
}

// Server sets the Server field for object Link.
func (b *LinkBuilder) Server(v Server) *LinkBuilder {
	b.target.server = v
	return b
}

// Reference sets the $ref (reference) field for object Link.
func (b *LinkBuilder) Reference(v string) *LinkBuilder {
	b.target.reference = v
	return b
}
