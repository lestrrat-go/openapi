package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// ContactMutator is used to build an instance of Contact. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Contact with new values
type ContactMutator struct {
	proxy  *contact
	target *contact
}

// Do finalizes the matuation process for Contact and returns the result
func (b *ContactMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateContact creates a new mutator object for Contact
func MutateContact(v Contact) *ContactMutator {
	return &ContactMutator{
		target: v.(*contact),
		proxy:  v.Clone().(*contact),
	}
}

// Name sets the Name field for object Contact.
func (b *ContactMutator) Name(v string) *ContactMutator {
	b.proxy.name = v
	return b
}

// URL sets the URL field for object Contact.
func (b *ContactMutator) URL(v string) *ContactMutator {
	b.proxy.url = v
	return b
}

// Email sets the Email field for object Contact.
func (b *ContactMutator) Email(v string) *ContactMutator {
	b.proxy.email = v
	return b
}
