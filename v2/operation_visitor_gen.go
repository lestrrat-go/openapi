package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// OperationVisitor is an interface for objects that knows
// how to process Operation elements while traversing the OpenAPI structure
type OperationVisitor interface {
	VisitOperation(context.Context, Operation) error
}

func visitOperation(ctx context.Context, elem Operation) error {
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

	if v, ok := ctx.Value(operationVisitorCtxKey{}).(OperationVisitor); ok {
		if err := v.VisitOperation(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Operation element`)
		}
	}

	if child := elem.ExternalDocs(); child != nil {
		if err := visitExternalDocumentation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit ExternalDocs element for Operation`)
		}
	}

	for i, iter := 0, elem.Parameters(); iter.Next(); {
		if err := visitParameter(ctx, iter.Item()); err != nil {
			return errors.Wrapf(err, `failed to visit element %d for Operation`, i)
		}
		i++
	}

	if child := elem.Responses(); child != nil {
		if err := visitResponses(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Responses element for Operation`)
		}
	}

	for i, iter := 0, elem.Security(); iter.Next(); {
		if err := visitSecurityRequirement(ctx, iter.Item()); err != nil {
			return errors.Wrapf(err, `failed to visit element %d for Operation`, i)
		}
		i++
	}
	return nil
}
