package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *oAuthFlows) Implicit() OAuthFlow {
	return v.implicit
}

func (v *oAuthFlows) Password() OAuthFlow {
	return v.password
}

func (v *oAuthFlows) ClientCredentials() OAuthFlow {
	return v.clientCredentials
}

func (v *oAuthFlows) AuthorizationCode() OAuthFlow {
	return v.authorizationCode
}

func (v *oAuthFlows) Reference() string {
	return v.reference
}

func (v *oAuthFlows) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
