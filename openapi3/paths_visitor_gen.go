package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// PathsVisitor is an interface for objects that knows
// how to process Paths elements while traversing the OpenAPI structure
type PathsVisitor interface {
	VisitPaths(context.Context, Paths) error
}

func visitPaths(ctx context.Context, elem Paths) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(pathsVisitorCtxKey{}).(PathsVisitor); ok {
		if err := v.VisitPaths(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Paths element`)
		}
	}
	return nil
}
