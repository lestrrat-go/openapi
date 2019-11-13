package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// ExampleMutator is used to build an instance of Example. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Example with new values
type ExampleMutator struct {
	proxy  *example
	target *example
}

// Do finalizes the matuation process for Example and returns the result
func (b *ExampleMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateExample creates a new mutator object for Example
func MutateExample(v Example) *ExampleMutator {
	return &ExampleMutator{
		target: v.(*example),
		proxy:  v.Clone().(*example),
	}
}

// Name sets the Name field for object Example.
func (b *ExampleMutator) Name(v string) *ExampleMutator {
	b.proxy.name = v
	return b
}

// Description sets the Description field for object Example.
func (b *ExampleMutator) Description(v string) *ExampleMutator {
	b.proxy.description = v
	return b
}

// Value sets the Value field for object Example.
func (b *ExampleMutator) Value(v interface{}) *ExampleMutator {
	b.proxy.value = v
	return b
}

// ExternalValue sets the ExternalValue field for object Example.
func (b *ExampleMutator) ExternalValue(v string) *ExampleMutator {
	b.proxy.externalValue = v
	return b
}
