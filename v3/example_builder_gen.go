package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// ExampleBuilder is used to build an instance of Example. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Example
type ExampleBuilder struct {
	target *example
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *ExampleBuilder) MustBuild(options ...Option) Example {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Example and returns the result
func (b *ExampleBuilder) Build(options ...Option) (Example, error) {
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

// NewExample creates a new builder object for Example
func NewExample() *ExampleBuilder {
	return &ExampleBuilder{
		target: &example{},
	}
}

// Description sets the Description field for object Example.
func (b *ExampleBuilder) Description(v string) *ExampleBuilder {
	b.target.description = v
	return b
}

// Value sets the Value field for object Example.
func (b *ExampleBuilder) Value(v interface{}) *ExampleBuilder {
	b.target.value = v
	return b
}

// ExternalValue sets the ExternalValue field for object Example.
func (b *ExampleBuilder) ExternalValue(v string) *ExampleBuilder {
	b.target.externalValue = v
	return b
}

// Reference sets the $ref (reference) field for object Example.
func (b *ExampleBuilder) Reference(v string) *ExampleBuilder {
	b.target.reference = v
	return b
}
