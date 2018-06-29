package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

// ParameterVisitor is an interface for objects that knows
// how to process Parameter elements while traversing the OpenAPI structure
type ParameterVisitor interface {
	VisitParameter(context.Context, Parameter) error
}

func visitParameter(ctx context.Context, elem Parameter) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(parameterVisitorCtxKey{}).(ParameterVisitor); ok {
		if err := v.VisitParameter(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Parameter element`)
		}
	}

	if child := elem.Schema(); child != nil {
		if err := visitSchema(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Schema element for Parameter`)
		}
	}
	return nil
}
