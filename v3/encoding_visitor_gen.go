package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

// EncodingVisitor is an interface for objects that knows
// how to process Encoding elements while traversing the OpenAPI structure
type EncodingVisitor interface {
	VisitEncoding(context.Context, Encoding) error
}

func visitEncoding(ctx context.Context, elem Encoding) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(encodingVisitorCtxKey{}).(EncodingVisitor); ok {
		if err := v.VisitEncoding(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Encoding element`)
		}
	}
	return nil
}
