package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// SecuritySchemeBuilder is used to build an instance of SecurityScheme. The user must
// call `Do()` after providing all the necessary information to
// build an instance of SecurityScheme
type SecuritySchemeBuilder struct {
	target *securityScheme
}

// Do finalizes the building process for SecurityScheme and returns the result
func (b *SecuritySchemeBuilder) Do() SecurityScheme {
	return b.target
}

// NewSecurityScheme creates a new builder object for SecurityScheme
func NewSecurityScheme(typ string, name string, in string, scheme string, flows OAuthFlows, openIDConnectURL string) *SecuritySchemeBuilder {
	return &SecuritySchemeBuilder{
		target: &securityScheme{
			typ:              typ,
			name:             name,
			in:               in,
			scheme:           scheme,
			flows:            flows,
			openIDConnectURL: openIDConnectURL,
		},
	}
}

// Description sets the Description field for object SecurityScheme.
func (b *SecuritySchemeBuilder) Description(v string) *SecuritySchemeBuilder {
	b.target.description = v
	return b
}

// BearerFormat sets the BearerFormat field for object SecurityScheme.
func (b *SecuritySchemeBuilder) BearerFormat(v string) *SecuritySchemeBuilder {
	b.target.bearerFormat = v
	return b
}

// Reference sets the $ref (reference) field for object SecurityScheme.
func (b *SecuritySchemeBuilder) Reference(v string) *SecuritySchemeBuilder {
	b.target.reference = v
	return b
}
