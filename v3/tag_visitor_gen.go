package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// TagVisitor is an interface for objects that knows
// how to process Tag elements while traversing the OpenAPI structure
type TagVisitor interface {
	VisitTag(context.Context, Tag) error
}

func visitTag(ctx context.Context, elem Tag) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(tagVisitorCtxKey{}).(TagVisitor); ok {
		if err := v.VisitTag(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Tag element`)
		}
	}

	if child := elem.ExternalDocs(); child != nil {
		if err := visitExternalDocumentation(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit ExternalDocs element for Tag`)
		}
	}
	return nil
}
