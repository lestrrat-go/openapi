package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// ServerVariableVisitor is an interface for objects that knows
// how to process ServerVariable elements while traversing the OpenAPI structure
type ServerVariableVisitor interface {
	VisitServerVariable(context.Context, ServerVariable) error
}

func visitServerVariable(ctx context.Context, elem ServerVariable) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(serverVariableVisitorCtxKey{}).(ServerVariableVisitor); ok {
		if err := v.VisitServerVariable(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit ServerVariable element`)
		}
	}
	return nil
}
