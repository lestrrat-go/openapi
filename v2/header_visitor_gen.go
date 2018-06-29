package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// HeaderVisitor is an interface for objects that knows
// how to process Header elements while traversing the OpenAPI structure
type HeaderVisitor interface {
	VisitHeader(context.Context, Header) error
}

func visitHeader(ctx context.Context, elem Header) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(headerVisitorCtxKey{}).(HeaderVisitor); ok {
		if err := v.VisitHeader(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Header element`)
		}
	}

	if child := elem.Items(); child != nil {
		if err := visitItems(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Items element for Header`)
		}
	}
	return nil
}
