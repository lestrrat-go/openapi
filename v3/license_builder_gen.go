package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// LicenseBuilder is used to build an instance of License. The user must
// call `Do()` after providing all the necessary information to
// build an instance of License
type LicenseBuilder struct {
	target *license
}

// Do finalizes the building process for License and returns the result
func (b *LicenseBuilder) Do() License {
	return b.target
}

// NewLicense creates a new builder object for License
func NewLicense(name string) *LicenseBuilder {
	return &LicenseBuilder{
		target: &license{
			name: name,
		},
	}
}

// UrL sets the UrL field for object License.
func (b *LicenseBuilder) UrL(v string) *LicenseBuilder {
	b.target.uRL = v
	return b
}

// Reference sets the $ref (reference) field for object License.
func (b *LicenseBuilder) Reference(v string) *LicenseBuilder {
	b.target.reference = v
	return b
}
