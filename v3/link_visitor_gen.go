package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// LinkVisitor is an interface for objects that knows
// how to process Link elements while traversing the OpenAPI structure
type LinkVisitor interface {
	VisitLink(context.Context, Link) error
}

func visitLink(ctx context.Context, elem Link) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(linkVisitorCtxKey{}).(LinkVisitor); ok {
		if err := v.VisitLink(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Link element`)
		}
	}

	if child := elem.Server(); child != nil {
		if err := visitServer(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Server element for Link`)
		}
	}
	return nil
}
