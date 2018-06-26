package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// CallbackBuilder is used to build an instance of Callback. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Callback
type CallbackBuilder struct {
	target *callback
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *CallbackBuilder) MustBuild(options ...Option) Callback {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Callback and returns the result
func (b *CallbackBuilder) Build(options ...Option) (Callback, error) {
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

// NewCallback creates a new builder object for Callback
func NewCallback() *CallbackBuilder {
	return &CallbackBuilder{
		target: &callback{},
	}
}

// URLs sets the URLs field for object Callback.
func (b *CallbackBuilder) URLs(v map[string]PathItem) *CallbackBuilder {
	b.target.urls = v
	return b
}

// Reference sets the $ref (reference) field for object Callback.
func (b *CallbackBuilder) Reference(v string) *CallbackBuilder {
	b.target.reference = v
	return b
}
