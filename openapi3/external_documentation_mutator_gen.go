package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// ExternalDocumentationMutator is used to build an instance of ExternalDocumentation. The user must
// call `Do()` after providing all the necessary information to
// the new instance of ExternalDocumentation with new values
type ExternalDocumentationMutator struct {
	proxy  *externalDocumentation
	target *externalDocumentation
}

// Do finalizes the matuation process for ExternalDocumentation and returns the result
func (b *ExternalDocumentationMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateExternalDocumentation creates a new mutator object for ExternalDocumentation
func MutateExternalDocumentation(v ExternalDocumentation) *ExternalDocumentationMutator {
	return &ExternalDocumentationMutator{
		target: v.(*externalDocumentation),
		proxy:  v.Clone().(*externalDocumentation),
	}
}

// Description sets the Description field for object ExternalDocumentation.
func (b *ExternalDocumentationMutator) Description(v string) *ExternalDocumentationMutator {
	b.proxy.description = v
	return b
}

// URL sets the URL field for object ExternalDocumentation.
func (b *ExternalDocumentationMutator) URL(v string) *ExternalDocumentationMutator {
	b.proxy.url = v
	return b
}
