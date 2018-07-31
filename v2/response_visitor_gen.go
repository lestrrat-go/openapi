package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// ResponseVisitor is an interface for objects that knows
// how to process Response elements while traversing the OpenAPI structure
type ResponseVisitor interface {
	VisitResponse(context.Context, Response) error
}

func visitResponse(ctx context.Context, elem Response) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(responseVisitorCtxKey{}).(ResponseVisitor); ok {
		if err := v.VisitResponse(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Response element`)
		}
	}

	if child := elem.Schema(); child != nil {
		if err := visitSchema(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Schema element for Response`)
		}
	}

	for iter := elem.Headers(); iter.Next(); {
		key, value := iter.Item()
		if err := visitHeader(context.WithValue(ctx, headerMapKeyVisitorCtxKey{}, key), value); err != nil {
			return errors.Wrap(err, `failed to visit Headers element for Response`)
		}
	}
	return nil
}
