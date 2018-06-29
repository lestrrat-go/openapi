package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// SwaggerVisitor is an interface for objects that knows
// how to process Swagger elements while traversing the OpenAPI structure
type SwaggerVisitor interface {
	VisitSwagger(context.Context, Swagger) error
}

func visitSwagger(ctx context.Context, elem Swagger) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(swaggerVisitorCtxKey{}).(SwaggerVisitor); ok {
		if err := v.VisitSwagger(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Swagger element`)
		}
	}

	if child := elem.Info(); child != nil {
		if err := visitInfo(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Info element for Swagger`)
		}
	}

	if child := elem.Paths(); child != nil {
		if err := visitPaths(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Paths element for Swagger`)
		}
	}

	if child := elem.ExternalDocs(); child != nil {
		if err := visitExternalDocumentation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit ExternalDocs element for Swagger`)
		}
	}
	return nil
}
