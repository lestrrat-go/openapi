package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// DiscriminatorBuilder is used to build an instance of Discriminator. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Discriminator
type DiscriminatorBuilder struct {
	target *discriminator
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *DiscriminatorBuilder) MustBuild(options ...Option) Discriminator {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Discriminator and returns the result
func (b *DiscriminatorBuilder) Build(options ...Option) (Discriminator, error) {
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

// NewDiscriminator creates a new builder object for Discriminator
func NewDiscriminator(propertyName string) *DiscriminatorBuilder {
	return &DiscriminatorBuilder{
		target: &discriminator{
			propertyName: propertyName,
		},
	}
}

// Mapping sets the Mapping field for object Discriminator.
func (b *DiscriminatorBuilder) Mapping(v map[string]string) *DiscriminatorBuilder {
	b.target.mapping = v
	return b
}

// Reference sets the $ref (reference) field for object Discriminator.
func (b *DiscriminatorBuilder) Reference(v string) *DiscriminatorBuilder {
	b.target.reference = v
	return b
}
