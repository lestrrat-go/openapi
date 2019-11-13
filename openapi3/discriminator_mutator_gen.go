package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// DiscriminatorMutator is used to build an instance of Discriminator. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Discriminator with new values
type DiscriminatorMutator struct {
	proxy  *discriminator
	target *discriminator
}

// Do finalizes the matuation process for Discriminator and returns the result
func (b *DiscriminatorMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateDiscriminator creates a new mutator object for Discriminator
func MutateDiscriminator(v Discriminator) *DiscriminatorMutator {
	return &DiscriminatorMutator{
		target: v.(*discriminator),
		proxy:  v.Clone().(*discriminator),
	}
}

// PropertyName sets the PropertyName field for object Discriminator.
func (b *DiscriminatorMutator) PropertyName(v string) *DiscriminatorMutator {
	b.proxy.propertyName = v
	return b
}

func (b *DiscriminatorMutator) ClearMapping() *DiscriminatorMutator {
	b.proxy.mapping.Clear()
	return b
}

func (b *DiscriminatorMutator) Mapping(key StringMapKey, value string) *DiscriminatorMutator {
	if b.proxy.mapping == nil {
		b.proxy.mapping = StringMap{}
	}

	b.proxy.mapping[key] = value
	return b
}
