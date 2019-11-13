package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// MediaTypeVisitor is an interface for objects that knows
// how to process MediaType elements while traversing the OpenAPI structure
type MediaTypeVisitor interface {
	VisitMediaType(context.Context, MediaType) error
}

func visitMediaType(ctx context.Context, elem MediaType) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(mediaTypeVisitorCtxKey{}).(MediaTypeVisitor); ok {
		if err := v.VisitMediaType(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit MediaType element`)
		}
	}

	if child := elem.Schema(); child != nil {
		if err := visitSchema(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Schema element for MediaType`)
		}
	}
	return nil
}
