package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// OAuthFlowsVisitor is an interface for objects that knows
// how to process OAuthFlows elements while traversing the OpenAPI structure
type OAuthFlowsVisitor interface {
	VisitOAuthFlows(context.Context, OAuthFlows) error
}

func visitOAuthFlows(ctx context.Context, elem OAuthFlows) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(oauthFlowsVisitorCtxKey{}).(OAuthFlowsVisitor); ok {
		if err := v.VisitOAuthFlows(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit OAuthFlows element`)
		}
	}

	if child := elem.Implicit(); child != nil {
		if err := visitOAuthFlow(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Implicit element for OAuthFlows`)
		}
	}

	if child := elem.Password(); child != nil {
		if err := visitOAuthFlow(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Password element for OAuthFlows`)
		}
	}

	if child := elem.ClientCredentials(); child != nil {
		if err := visitOAuthFlow(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit ClientCredentials element for OAuthFlows`)
		}
	}

	if child := elem.AuthorizationCode(); child != nil {
		if err := visitOAuthFlow(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit AuthorizationCode element for OAuthFlows`)
		}
	}
	return nil
}
