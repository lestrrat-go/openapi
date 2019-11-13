package openapi2

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// ItemsVisitor is an interface for objects that knows
// how to process Items elements while traversing the OpenAPI structure
type ItemsVisitor interface {
	VisitItems(context.Context, Items) error
}

func visitItems(ctx context.Context, elem Items) error {
	if checker, ok := elem.(interface{ IsValid() bool }); ok {
		if !checker.IsValid() {
			return nil
		}
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(itemsVisitorCtxKey{}).(ItemsVisitor); ok {
		if err := v.VisitItems(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Items element`)
		}
	}

	if child := elem.Items(); child != nil {
		if err := visitItems(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Items element for Items`)
		}
	}
	return nil
}
