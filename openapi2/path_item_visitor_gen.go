package openapi2

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// PathItemVisitor is an interface for objects that knows
// how to process PathItem elements while traversing the OpenAPI structure
type PathItemVisitor interface {
	VisitPathItem(context.Context, PathItem) error
}

func visitPathItem(ctx context.Context, elem PathItem) error {
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

	if v, ok := ctx.Value(pathItemVisitorCtxKey{}).(PathItemVisitor); ok {
		if err := v.VisitPathItem(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit PathItem element`)
		}
	}

	if child := elem.Get(); child != nil {
		if err := visitOperation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Get element for PathItem`)
		}
	}

	if child := elem.Put(); child != nil {
		if err := visitOperation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Put element for PathItem`)
		}
	}

	if child := elem.Post(); child != nil {
		if err := visitOperation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Post element for PathItem`)
		}
	}

	if child := elem.Delete(); child != nil {
		if err := visitOperation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Delete element for PathItem`)
		}
	}

	if child := elem.Options(); child != nil {
		if err := visitOperation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Options element for PathItem`)
		}
	}

	if child := elem.Head(); child != nil {
		if err := visitOperation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Head element for PathItem`)
		}
	}

	if child := elem.Patch(); child != nil {
		if err := visitOperation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Patch element for PathItem`)
		}
	}

	for i, iter := 0, elem.Parameters(); iter.Next(); {
		if err := visitParameter(ctx, iter.Item()); err != nil {
			return errors.Wrapf(err, `failed to visit element %d for PathItem`, i)
		}
		i++
	}
	return nil
}
