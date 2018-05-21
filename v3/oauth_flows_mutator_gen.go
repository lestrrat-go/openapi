package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// OAuthFlowsMutator is used to build an instance of OAuthFlows. The user must
// call `Do()` after providing all the necessary information to
// the new instance of OAuthFlows with new values
type OAuthFlowsMutator struct {
	proxy  *oAuthFlows
	target *oAuthFlows
}

// Do finalizes the matuation process for OAuthFlows and returns the result
func (b *OAuthFlowsMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateOAuthFlows creates a new mutator object for OAuthFlows
func MutateOAuthFlows(v OAuthFlows) *OAuthFlowsMutator {
	return &OAuthFlowsMutator{
		target: v.(*oAuthFlows),
		proxy:  v.Clone().(*oAuthFlows),
	}
}

// Implicit sets the Implicit field for object OAuthFlows.
func (b *OAuthFlowsMutator) Implicit(v OAuthFlow) *OAuthFlowsMutator {
	b.proxy.implicit = v
	return b
}

// Password sets the Password field for object OAuthFlows.
func (b *OAuthFlowsMutator) Password(v OAuthFlow) *OAuthFlowsMutator {
	b.proxy.password = v
	return b
}

// ClientCredentials sets the ClientCredentials field for object OAuthFlows.
func (b *OAuthFlowsMutator) ClientCredentials(v OAuthFlow) *OAuthFlowsMutator {
	b.proxy.clientCredentials = v
	return b
}

// AuthorizationCode sets the AuthorizationCode field for object OAuthFlows.
func (b *OAuthFlowsMutator) AuthorizationCode(v OAuthFlow) *OAuthFlowsMutator {
	b.proxy.authorizationCode = v
	return b
}
