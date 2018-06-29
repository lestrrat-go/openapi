package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

func (v *securityScheme) Type() string {
	return v.typ
}

func (v *securityScheme) Description() string {
	return v.description
}

func (v *securityScheme) Name() string {
	return v.name
}

func (v *securityScheme) In() string {
	return v.in
}

func (v *securityScheme) Scheme() string {
	return v.scheme
}

func (v *securityScheme) BearerFormat() string {
	return v.bearerFormat
}

func (v *securityScheme) Flows() OAuthFlows {
	return v.flows
}

func (v *securityScheme) OpenIDConnectURL() string {
	return v.openIDConnectURL
}

func (v *securityScheme) Reference() string {
	return v.reference
}

func (v *securityScheme) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *securityScheme) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
