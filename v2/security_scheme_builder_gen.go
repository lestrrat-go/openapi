package openapi

// This file was automatically generated by gentyeps.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// SecuritySchemeBuilder is used to build an instance of SecurityScheme. The user must
// call `Do()` after providing all the necessary information to
// build an instance of SecurityScheme
type SecuritySchemeBuilder struct {
	target *securityScheme
}

// Do finalizes the building process for SecurityScheme and returns the result
func (b *SecuritySchemeBuilder) Do() (SecurityScheme, error) {
	if err := b.target.Validate(); err != nil {
		return nil, errors.Wrap(err, `validation failed`)
	}
	return b.target, nil
}

// NewSecurityScheme creates a new builder object for SecurityScheme
func NewSecurityScheme(typ string) *SecuritySchemeBuilder {
	return &SecuritySchemeBuilder{
		target: &securityScheme{
			typ: typ,
		},
	}
}

// Description sets the Description field for object SecurityScheme.
func (b *SecuritySchemeBuilder) Description(v string) *SecuritySchemeBuilder {
	b.target.description = v
	return b
}

// Name sets the Name field for object SecurityScheme.
func (b *SecuritySchemeBuilder) Name(v string) *SecuritySchemeBuilder {
	b.target.name = v
	return b
}

// In sets the In field for object SecurityScheme.
func (b *SecuritySchemeBuilder) In(v string) *SecuritySchemeBuilder {
	b.target.in = v
	return b
}

// Flow sets the Flow field for object SecurityScheme.
func (b *SecuritySchemeBuilder) Flow(v string) *SecuritySchemeBuilder {
	b.target.flow = v
	return b
}

// AuthorizationURL sets the AuthorizationURL field for object SecurityScheme.
func (b *SecuritySchemeBuilder) AuthorizationURL(v string) *SecuritySchemeBuilder {
	b.target.authorizationURL = v
	return b
}

// TokenURL sets the TokenURL field for object SecurityScheme.
func (b *SecuritySchemeBuilder) TokenURL(v string) *SecuritySchemeBuilder {
	b.target.tokenURL = v
	return b
}

// Scopes sets the Scopes field for object SecurityScheme.
func (b *SecuritySchemeBuilder) Scopes(v StringMap) *SecuritySchemeBuilder {
	b.target.scopes = v
	return b
}

// Reference sets the $ref (reference) field for object SecurityScheme.
func (b *SecuritySchemeBuilder) Reference(v string) *SecuritySchemeBuilder {
	b.target.reference = v
	return b
}
