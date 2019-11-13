package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// ComponentsVisitor is an interface for objects that knows
// how to process Components elements while traversing the OpenAPI structure
type ComponentsVisitor interface {
	VisitComponents(context.Context, Components) error
}

func visitComponents(ctx context.Context, elem Components) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(componentsVisitorCtxKey{}).(ComponentsVisitor); ok {
		if err := v.VisitComponents(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Components element`)
		}
	}
	return nil
}
