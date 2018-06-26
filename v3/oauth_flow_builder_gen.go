package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// OAuthFlowBuilder is used to build an instance of OAuthFlow. The user must
// call `Build()` after providing all the necessary information to
// build an instance of OAuthFlow
type OAuthFlowBuilder struct {
	target *oauthFlow
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *OAuthFlowBuilder) MustBuild(options ...Option) OAuthFlow {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for OAuthFlow and returns the result
func (b *OAuthFlowBuilder) Build(options ...Option) (OAuthFlow, error) {
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

// NewOAuthFlow creates a new builder object for OAuthFlow
func NewOAuthFlow() *OAuthFlowBuilder {
	return &OAuthFlowBuilder{
		target: &oauthFlow{},
	}
}

// AuthorizationURL sets the AuthorizationURL field for object OAuthFlow.
func (b *OAuthFlowBuilder) AuthorizationURL(v string) *OAuthFlowBuilder {
	b.target.authorizationURL = v
	return b
}

// TokenURL sets the TokenURL field for object OAuthFlow.
func (b *OAuthFlowBuilder) TokenURL(v string) *OAuthFlowBuilder {
	b.target.tokenURL = v
	return b
}

// RefreshURL sets the RefreshURL field for object OAuthFlow.
func (b *OAuthFlowBuilder) RefreshURL(v string) *OAuthFlowBuilder {
	b.target.refreshURL = v
	return b
}

// Scopes sets the Scopes field for object OAuthFlow.
func (b *OAuthFlowBuilder) Scopes(v map[string]string) *OAuthFlowBuilder {
	b.target.scopes = v
	return b
}

// Reference sets the $ref (reference) field for object OAuthFlow.
func (b *OAuthFlowBuilder) Reference(v string) *OAuthFlowBuilder {
	b.target.reference = v
	return b
}
