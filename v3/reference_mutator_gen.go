package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// ReferenceMutator is used to build an instance of Reference. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Reference with new values
type ReferenceMutator struct {
	proxy  *reference
	target *reference
}

// Do finalizes the matuation process for Reference and returns the result
func (b *ReferenceMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateReference creates a new mutator object for Reference
func MutateReference(v Reference) *ReferenceMutator {
	return &ReferenceMutator{
		target: v.(*reference),
		proxy:  v.Clone().(*reference),
	}
}

// URL sets the URL field for object Reference.
func (b *ReferenceMutator) URL(v string) *ReferenceMutator {
	b.proxy.uRL = v
	return b
}
