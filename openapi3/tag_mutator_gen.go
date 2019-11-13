package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// TagMutator is used to build an instance of Tag. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Tag with new values
type TagMutator struct {
	proxy  *tag
	target *tag
}

// Do finalizes the matuation process for Tag and returns the result
func (b *TagMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateTag creates a new mutator object for Tag
func MutateTag(v Tag) *TagMutator {
	return &TagMutator{
		target: v.(*tag),
		proxy:  v.Clone().(*tag),
	}
}

// Name sets the Name field for object Tag.
func (b *TagMutator) Name(v string) *TagMutator {
	b.proxy.name = v
	return b
}

// Description sets the Description field for object Tag.
func (b *TagMutator) Description(v string) *TagMutator {
	b.proxy.description = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object Tag.
func (b *TagMutator) ExternalDocs(v ExternalDocumentation) *TagMutator {
	b.proxy.externalDocs = v
	return b
}
