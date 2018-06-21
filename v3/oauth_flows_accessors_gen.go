package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *oauthFlows) Implicit() OAuthFlow {
	return v.implicit
}

func (v *oauthFlows) Password() OAuthFlow {
	return v.password
}

func (v *oauthFlows) ClientCredentials() OAuthFlow {
	return v.clientCredentials
}

func (v *oauthFlows) AuthorizationCode() OAuthFlow {
	return v.authorizationCode
}

func (v *oauthFlows) Reference() string {
	return v.reference
}

func (v *oauthFlows) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
