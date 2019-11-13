package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// OAuthFlowVisitor is an interface for objects that knows
// how to process OAuthFlow elements while traversing the OpenAPI structure
type OAuthFlowVisitor interface {
	VisitOAuthFlow(context.Context, OAuthFlow) error
}

func visitOAuthFlow(ctx context.Context, elem OAuthFlow) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(oauthFlowVisitorCtxKey{}).(OAuthFlowVisitor); ok {
		if err := v.VisitOAuthFlow(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit OAuthFlow element`)
		}
	}
	return nil
}
