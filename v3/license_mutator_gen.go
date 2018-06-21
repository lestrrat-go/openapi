package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// LicenseMutator is used to build an instance of License. The user must
// call `Do()` after providing all the necessary information to
// the new instance of License with new values
type LicenseMutator struct {
	proxy  *license
	target *license
}

// Do finalizes the matuation process for License and returns the result
func (b *LicenseMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateLicense creates a new mutator object for License
func MutateLicense(v License) *LicenseMutator {
	return &LicenseMutator{
		target: v.(*license),
		proxy:  v.Clone().(*license),
	}
}

// Name sets the Name field for object License.
func (b *LicenseMutator) Name(v string) *LicenseMutator {
	b.proxy.name = v
	return b
}

// UrL sets the UrL field for object License.
func (b *LicenseMutator) UrL(v string) *LicenseMutator {
	b.proxy.urL = v
	return b
}
