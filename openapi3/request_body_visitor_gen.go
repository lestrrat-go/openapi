package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// RequestBodyVisitor is an interface for objects that knows
// how to process RequestBody elements while traversing the OpenAPI structure
type RequestBodyVisitor interface {
	VisitRequestBody(context.Context, RequestBody) error
}

func visitRequestBody(ctx context.Context, elem RequestBody) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(requestBodyVisitorCtxKey{}).(RequestBodyVisitor); ok {
		if err := v.VisitRequestBody(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit RequestBody element`)
		}
	}
	return nil
}
