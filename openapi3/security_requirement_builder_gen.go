package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// SecurityRequirementBuilder is used to build an instance of SecurityRequirement. The user must
// call `Build()` after providing all the necessary information to
// build an instance of SecurityRequirement
type SecurityRequirementBuilder struct {
	target *securityRequirement
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *SecurityRequirementBuilder) MustBuild(options ...Option) SecurityRequirement {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for SecurityRequirement and returns the result
func (b *SecurityRequirementBuilder) Build(options ...Option) (SecurityRequirement, error) {
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

// NewSecurityRequirement creates a new builder object for SecurityRequirement
func NewSecurityRequirement() *SecurityRequirementBuilder {
	return &SecurityRequirementBuilder{
		target: &securityRequirement{},
	}
}

// Schemes sets the Schemes field for object SecurityRequirement.
func (b *SecurityRequirementBuilder) Schemes(v map[string][]string) *SecurityRequirementBuilder {
	b.target.schemes = v
	return b
}

// Reference sets the $ref (reference) field for object SecurityRequirement.
func (b *SecurityRequirementBuilder) Reference(v string) *SecurityRequirementBuilder {
	b.target.reference = v
	return b
}
