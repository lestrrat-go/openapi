package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// CallbackVisitor is an interface for objects that knows
// how to process Callback elements while traversing the OpenAPI structure
type CallbackVisitor interface {
	VisitCallback(context.Context, Callback) error
}

func visitCallback(ctx context.Context, elem Callback) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(callbackVisitorCtxKey{}).(CallbackVisitor); ok {
		if err := v.VisitCallback(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Callback element`)
		}
	}
	return nil
}
