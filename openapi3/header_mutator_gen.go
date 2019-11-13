package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// HeaderMutator is used to build an instance of Header. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Header with new values
type HeaderMutator struct {
	proxy  *header
	target *header
}

// Do finalizes the matuation process for Header and returns the result
func (b *HeaderMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateHeader creates a new mutator object for Header
func MutateHeader(v Header) *HeaderMutator {
	return &HeaderMutator{
		target: v.(*header),
		proxy:  v.Clone().(*header),
	}
}

// In sets the In field for object Header.
func (b *HeaderMutator) In(v Location) *HeaderMutator {
	b.proxy.in = v
	return b
}

// Required sets the Required field for object Header.
func (b *HeaderMutator) Required(v bool) *HeaderMutator {
	b.proxy.required = v
	return b
}

// Description sets the Description field for object Header.
func (b *HeaderMutator) Description(v string) *HeaderMutator {
	b.proxy.description = v
	return b
}

// Deprecated sets the Deprecated field for object Header.
func (b *HeaderMutator) Deprecated(v bool) *HeaderMutator {
	b.proxy.deprecated = v
	return b
}

// AllowEmptyValue sets the AllowEmptyValue field for object Header.
func (b *HeaderMutator) AllowEmptyValue(v bool) *HeaderMutator {
	b.proxy.allowEmptyValue = v
	return b
}

// Explode sets the Explode field for object Header.
func (b *HeaderMutator) Explode(v bool) *HeaderMutator {
	b.proxy.explode = v
	return b
}

// AllowReserved sets the AllowReserved field for object Header.
func (b *HeaderMutator) AllowReserved(v bool) *HeaderMutator {
	b.proxy.allowReserved = v
	return b
}

// Schema sets the Schema field for object Header.
func (b *HeaderMutator) Schema(v Schema) *HeaderMutator {
	b.proxy.schema = v
	return b
}

func (b *HeaderMutator) ClearExamples() *HeaderMutator {
	b.proxy.examples.Clear()
	return b
}

func (b *HeaderMutator) Example(key ExampleMapKey, value Example) *HeaderMutator {
	if b.proxy.examples == nil {
		b.proxy.examples = ExampleMap{}
	}

	b.proxy.examples[key] = value
	return b
}

func (b *HeaderMutator) ClearContent() *HeaderMutator {
	b.proxy.content.Clear()
	return b
}

func (b *HeaderMutator) Content(key MediaTypeMapKey, value MediaType) *HeaderMutator {
	if b.proxy.content == nil {
		b.proxy.content = MediaTypeMap{}
	}

	b.proxy.content[key] = value
	return b
}
