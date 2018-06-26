package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// OAuthFlowsBuilder is used to build an instance of OAuthFlows. The user must
// call `Build()` after providing all the necessary information to
// build an instance of OAuthFlows
type OAuthFlowsBuilder struct {
	target *oauthFlows
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *OAuthFlowsBuilder) MustBuild(options ...Option) OAuthFlows {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for OAuthFlows and returns the result
func (b *OAuthFlowsBuilder) Build(options ...Option) (OAuthFlows, error) {
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

// NewOAuthFlows creates a new builder object for OAuthFlows
func NewOAuthFlows() *OAuthFlowsBuilder {
	return &OAuthFlowsBuilder{
		target: &oauthFlows{},
	}
}

// Implicit sets the Implicit field for object OAuthFlows.
func (b *OAuthFlowsBuilder) Implicit(v OAuthFlow) *OAuthFlowsBuilder {
	b.target.implicit = v
	return b
}

// Password sets the Password field for object OAuthFlows.
func (b *OAuthFlowsBuilder) Password(v OAuthFlow) *OAuthFlowsBuilder {
	b.target.password = v
	return b
}

// ClientCredentials sets the ClientCredentials field for object OAuthFlows.
func (b *OAuthFlowsBuilder) ClientCredentials(v OAuthFlow) *OAuthFlowsBuilder {
	b.target.clientCredentials = v
	return b
}

// AuthorizationCode sets the AuthorizationCode field for object OAuthFlows.
func (b *OAuthFlowsBuilder) AuthorizationCode(v OAuthFlow) *OAuthFlowsBuilder {
	b.target.authorizationCode = v
	return b
}

// Reference sets the $ref (reference) field for object OAuthFlows.
func (b *OAuthFlowsBuilder) Reference(v string) *OAuthFlowsBuilder {
	b.target.reference = v
	return b
}
