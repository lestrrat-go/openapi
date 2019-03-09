package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// ServerVisitor is an interface for objects that knows
// how to process Server elements while traversing the OpenAPI structure
type ServerVisitor interface {
	VisitServer(context.Context, Server) error
}

func visitServer(ctx context.Context, elem Server) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(serverVisitorCtxKey{}).(ServerVisitor); ok {
		if err := v.VisitServer(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Server element`)
		}
	}
	return nil
}
