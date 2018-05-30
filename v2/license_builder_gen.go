package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// LicenseBuilder is used to build an instance of License. The user must
// call `Do()` after providing all the necessary information to
// build an instance of License
type LicenseBuilder struct {
	target *license
}

// Do finalizes the building process for License and returns the result
func (b *LicenseBuilder) Do() (License, error) {
	if err := b.target.Validate(); err != nil {
		return nil, errors.Wrap(err, `validation failed`)
	}
	return b.target, nil
}

// NewLicense creates a new builder object for License
func NewLicense(name string) *LicenseBuilder {
	return &LicenseBuilder{
		target: &license{
			name: name,
		},
	}
}

// URL sets the URL field for object License.
func (b *LicenseBuilder) URL(v string) *LicenseBuilder {
	b.target.uRL = v
	return b
}

// Reference sets the $ref (reference) field for object License.
func (b *LicenseBuilder) Reference(v string) *LicenseBuilder {
	b.target.reference = v
	return b
}
