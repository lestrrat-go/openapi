package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// ContactVisitor is an interface for objects that knows
// how to process Contact elements while traversing the OpenAPI structure
type ContactVisitor interface {
	VisitContact(context.Context, Contact) error
}

func visitContact(ctx context.Context, elem Contact) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(contactVisitorCtxKey{}).(ContactVisitor); ok {
		if err := v.VisitContact(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Contact element`)
		}
	}
	return nil
}
