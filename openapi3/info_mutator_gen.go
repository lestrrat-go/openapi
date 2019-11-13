package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// InfoMutator is used to build an instance of Info. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Info with new values
type InfoMutator struct {
	proxy  *info
	target *info
}

// Do finalizes the matuation process for Info and returns the result
func (b *InfoMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateInfo creates a new mutator object for Info
func MutateInfo(v Info) *InfoMutator {
	return &InfoMutator{
		target: v.(*info),
		proxy:  v.Clone().(*info),
	}
}

// Title sets the Title field for object Info.
func (b *InfoMutator) Title(v string) *InfoMutator {
	b.proxy.title = v
	return b
}

// Description sets the Description field for object Info.
func (b *InfoMutator) Description(v string) *InfoMutator {
	b.proxy.description = v
	return b
}

// TermsOfService sets the TermsOfService field for object Info.
func (b *InfoMutator) TermsOfService(v string) *InfoMutator {
	b.proxy.termsOfService = v
	return b
}

// Contact sets the Contact field for object Info.
func (b *InfoMutator) Contact(v Contact) *InfoMutator {
	b.proxy.contact = v
	return b
}

// License sets the License field for object Info.
func (b *InfoMutator) License(v License) *InfoMutator {
	b.proxy.license = v
	return b
}

// Version sets the Version field for object Info.
func (b *InfoMutator) Version(v string) *InfoMutator {
	b.proxy.version = v
	return b
}
