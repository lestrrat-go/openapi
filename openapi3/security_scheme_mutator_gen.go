package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// SecuritySchemeMutator is used to build an instance of SecurityScheme. The user must
// call `Do()` after providing all the necessary information to
// the new instance of SecurityScheme with new values
type SecuritySchemeMutator struct {
	proxy  *securityScheme
	target *securityScheme
}

// Do finalizes the matuation process for SecurityScheme and returns the result
func (b *SecuritySchemeMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateSecurityScheme creates a new mutator object for SecurityScheme
func MutateSecurityScheme(v SecurityScheme) *SecuritySchemeMutator {
	return &SecuritySchemeMutator{
		target: v.(*securityScheme),
		proxy:  v.Clone().(*securityScheme),
	}
}

// Type sets the Type field for object SecurityScheme.
func (b *SecuritySchemeMutator) Type(v string) *SecuritySchemeMutator {
	b.proxy.typ = v
	return b
}

// Description sets the Description field for object SecurityScheme.
func (b *SecuritySchemeMutator) Description(v string) *SecuritySchemeMutator {
	b.proxy.description = v
	return b
}

// Name sets the Name field for object SecurityScheme.
func (b *SecuritySchemeMutator) Name(v string) *SecuritySchemeMutator {
	b.proxy.name = v
	return b
}

// In sets the In field for object SecurityScheme.
func (b *SecuritySchemeMutator) In(v string) *SecuritySchemeMutator {
	b.proxy.in = v
	return b
}

// Scheme sets the Scheme field for object SecurityScheme.
func (b *SecuritySchemeMutator) Scheme(v string) *SecuritySchemeMutator {
	b.proxy.scheme = v
	return b
}

// BearerFormat sets the BearerFormat field for object SecurityScheme.
func (b *SecuritySchemeMutator) BearerFormat(v string) *SecuritySchemeMutator {
	b.proxy.bearerFormat = v
	return b
}

// Flows sets the Flows field for object SecurityScheme.
func (b *SecuritySchemeMutator) Flows(v OAuthFlows) *SecuritySchemeMutator {
	b.proxy.flows = v
	return b
}

// OpenIDConnectURL sets the OpenIDConnectURL field for object SecurityScheme.
func (b *SecuritySchemeMutator) OpenIDConnectURL(v string) *SecuritySchemeMutator {
	b.proxy.openIDConnectURL = v
	return b
}
