package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// DiscriminatorBuilder is used to build an instance of Discriminator. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Discriminator
type DiscriminatorBuilder struct {
	target *discriminator
}

// Do finalizes the building process for Discriminator and returns the result
func (b *DiscriminatorBuilder) Do() Discriminator {
	return b.target
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
