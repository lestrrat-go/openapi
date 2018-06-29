package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// XMLVisitor is an interface for objects that knows
// how to process XML elements while traversing the OpenAPI structure
type XMLVisitor interface {
	VisitXML(context.Context, XML) error
}

func visitXML(ctx context.Context, elem XML) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(xmlVisitorCtxKey{}).(XMLVisitor); ok {
		if err := v.VisitXML(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit XML element`)
		}
	}
	return nil
}
