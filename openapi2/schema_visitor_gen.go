package openapi2

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

	if v, ok := ctx.Value(schemaVisitorCtxKey{}).(SchemaVisitor); ok {
		if err := v.VisitSchema(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Schema element`)
		}
	}

	for i, iter := 0, elem.AllOf(); iter.Next(); {
		if err := visitSchema(ctx, iter.Item()); err != nil {
			return errors.Wrapf(err, `failed to visit element %d for Schema`, i)
		}
		i++
	}

	if child := elem.Items(); child != nil {
		if err := visitSchema(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Items element for Schema`)
		}
	}

	for iter := elem.Properties(); iter.Next(); {
		key, value := iter.Item()
		if err := visitSchema(context.WithValue(ctx, schemaMapKeyVisitorCtxKey{}, key), value); err != nil {
			return errors.Wrap(err, `failed to visit Properties element for Schema`)
		}
	}

	for iter := elem.AdditionaProperties(); iter.Next(); {
		key, value := iter.Item()
		if err := visitSchema(context.WithValue(ctx, schemaMapKeyVisitorCtxKey{}, key), value); err != nil {
			return errors.Wrap(err, `failed to visit AdditionaProperties element for Schema`)
		}
	}

	if child := elem.ExternalDocs(); child != nil {
		if err := visitExternalDocumentation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit ExternalDocs element for Schema`)
		}
	}

	if child := elem.XML(); child != nil {
		if err := visitXML(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit XML element for Schema`)
		}
	}
	return nil
}
