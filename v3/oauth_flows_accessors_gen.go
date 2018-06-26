package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

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

func (v *oauthFlows) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *oauthFlows) recurseValidate() error {
	if elem := v.implicit; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "implicit"`)
		}
	}
	if elem := v.password; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "password"`)
		}
	}
	if elem := v.clientCredentials; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "clientCredentials"`)
		}
	}
	if elem := v.authorizationCode; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "authorizationCode"`)
		}
	}
	return nil
}
