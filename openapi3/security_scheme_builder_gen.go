package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// SecuritySchemeBuilder is used to build an instance of SecurityScheme. The user must
// call `Build()` after providing all the necessary information to
// build an instance of SecurityScheme
type SecuritySchemeBuilder struct {
	target *securityScheme
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *SecuritySchemeBuilder) MustBuild(options ...Option) SecurityScheme {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for SecurityScheme and returns the result
func (b *SecuritySchemeBuilder) Build(options ...Option) (SecurityScheme, error) {
	validate := true
	for _, option := range options {
		switch option.Name() {
		case optkeyValidate:
			validate = option.Value().(bool)
		}
	}
	if validate {
		if err := b.target.Validate(false); err != nil {
			return nil, errors.Wrap(err, `validation failed`)
		}
	}
	return b.target, nil
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
