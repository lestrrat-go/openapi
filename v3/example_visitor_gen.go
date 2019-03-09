package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// ExampleVisitor is an interface for objects that knows
// how to process Example elements while traversing the OpenAPI structure
type ExampleVisitor interface {
	VisitExample(context.Context, Example) error
}

func visitExample(ctx context.Context, elem Example) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(exampleVisitorCtxKey{}).(ExampleVisitor); ok {
		if err := v.VisitExample(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Example element`)
		}
	}
	return nil
}
