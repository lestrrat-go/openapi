package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// OpenAPIVisitor is an interface for objects that knows
// how to process OpenAPI elements while traversing the OpenAPI structure
type OpenAPIVisitor interface {
	VisitOpenAPI(context.Context, OpenAPI) error
}

func visitOpenAPI(ctx context.Context, elem OpenAPI) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(openapiVisitorCtxKey{}).(OpenAPIVisitor); ok {
		if err := v.VisitOpenAPI(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit OpenAPI element`)
		}
	}

	if child := elem.Info(); child != nil {
		if err := visitInfo(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Info element for OpenAPI`)
		}
	}

	if child := elem.Paths(); child != nil {
		if err := visitPaths(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Paths element for OpenAPI`)
		}
	}

	if child := elem.Components(); child != nil {
		if err := visitComponents(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Components element for OpenAPI`)
		}
	}

	if child := elem.Security(); child != nil {
		if err := visitSecurityRequirement(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Security element for OpenAPI`)
		}
	}

	if child := elem.ExternalDocs(); child != nil {
		if err := visitExternalDocumentation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit ExternalDocs element for OpenAPI`)
		}
	}
	return nil
}
