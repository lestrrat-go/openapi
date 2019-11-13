package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// OAuthFlowMutator is used to build an instance of OAuthFlow. The user must
// call `Do()` after providing all the necessary information to
// the new instance of OAuthFlow with new values
type OAuthFlowMutator struct {
	proxy  *oauthFlow
	target *oauthFlow
}

// Do finalizes the matuation process for OAuthFlow and returns the result
func (b *OAuthFlowMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateOAuthFlow creates a new mutator object for OAuthFlow
func MutateOAuthFlow(v OAuthFlow) *OAuthFlowMutator {
	return &OAuthFlowMutator{
		target: v.(*oauthFlow),
		proxy:  v.Clone().(*oauthFlow),
	}
}

// AuthorizationURL sets the AuthorizationURL field for object OAuthFlow.
func (b *OAuthFlowMutator) AuthorizationURL(v string) *OAuthFlowMutator {
	b.proxy.authorizationURL = v
	return b
}

// TokenURL sets the TokenURL field for object OAuthFlow.
func (b *OAuthFlowMutator) TokenURL(v string) *OAuthFlowMutator {
	b.proxy.tokenURL = v
	return b
}

// RefreshURL sets the RefreshURL field for object OAuthFlow.
func (b *OAuthFlowMutator) RefreshURL(v string) *OAuthFlowMutator {
	b.proxy.refreshURL = v
	return b
}

func (b *OAuthFlowMutator) ClearScopes() *OAuthFlowMutator {
	b.proxy.scopes.Clear()
	return b
}

func (b *OAuthFlowMutator) Scope(key ScopeMapKey, value string) *OAuthFlowMutator {
	if b.proxy.scopes == nil {
		b.proxy.scopes = ScopeMap{}
	}

	b.proxy.scopes[key] = value
	return b
}
