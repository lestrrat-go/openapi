package openapi2

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// ExternalDocumentationVisitor is an interface for objects that knows
// how to process ExternalDocumentation elements while traversing the OpenAPI structure
type ExternalDocumentationVisitor interface {
	VisitExternalDocumentation(context.Context, ExternalDocumentation) error
}

func visitExternalDocumentation(ctx context.Context, elem ExternalDocumentation) error {
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

	if v, ok := ctx.Value(externalDocumentationVisitorCtxKey{}).(ExternalDocumentationVisitor); ok {
		if err := v.VisitExternalDocumentation(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit ExternalDocumentation element`)
		}
	}
	return nil
}
