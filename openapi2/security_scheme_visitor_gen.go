package openapi2

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// SecuritySchemeVisitor is an interface for objects that knows
// how to process SecurityScheme elements while traversing the OpenAPI structure
type SecuritySchemeVisitor interface {
	VisitSecurityScheme(context.Context, SecurityScheme) error
}

func visitSecurityScheme(ctx context.Context, elem SecurityScheme) error {
	if checker, ok := elem.(interface{ IsValid() bool }); ok {
		if !checker.IsValid() {
			return nil
		}
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(securitySchemeVisitorCtxKey{}).(SecuritySchemeVisitor); ok {
		if err := v.VisitSecurityScheme(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit SecurityScheme element`)
		}
	}
	return nil
}
