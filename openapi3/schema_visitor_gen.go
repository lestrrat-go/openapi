package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// SchemaVisitor is an interface for objects that knows
// how to process Schema elements while traversing the OpenAPI structure
type SchemaVisitor interface {
	VisitSchema(context.Context, Schema) error
}

func visitSchema(ctx context.Context, elem Schema) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(schemaVisitorCtxKey{}).(SchemaVisitor); ok {
		if err := v.VisitSchema(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Schema element`)
		}
	}

	if child := elem.Not(); child != nil {
		if err := visitSchema(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Not element for Schema`)
		}
	}

	if child := elem.Items(); child != nil {
		if err := visitSchema(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Items element for Schema`)
		}
	}

	if child := elem.Discriminator(); child != nil {
		if err := visitDiscriminator(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Discriminator element for Schema`)
		}
	}

	if child := elem.ExternalDocs(); child != nil {
		if err := visitExternalDocumentation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit ExternalDocs element for Schema`)
		}
	}
	return nil
}
