package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

// SecurityRequirementVisitor is an interface for objects that knows
// how to process SecurityRequirement elements while traversing the OpenAPI structure
type SecurityRequirementVisitor interface {
	VisitSecurityRequirement(context.Context, SecurityRequirement) error
}

func visitSecurityRequirement(ctx context.Context, elem SecurityRequirement) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(securityRequirementVisitorCtxKey{}).(SecurityRequirementVisitor); ok {
		if err := v.VisitSecurityRequirement(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit SecurityRequirement element`)
		}
	}
	return nil
}
