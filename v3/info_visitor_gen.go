package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

// InfoVisitor is an interface for objects that knows
// how to process Info elements while traversing the OpenAPI structure
type InfoVisitor interface {
	VisitInfo(context.Context, Info) error
}

func visitInfo(ctx context.Context, elem Info) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(infoVisitorCtxKey{}).(InfoVisitor); ok {
		if err := v.VisitInfo(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Info element`)
		}
	}

	if child := elem.Contact(); child != nil {
		if err := visitContact(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit Contact element for Info`)
		}
	}

	if child := elem.License(); child != nil {
		if err := visitLicense(ctx, child); err != nil {
			return errors.Wrap(err, `failed to visit License element for Info`)
		}
	}
	return nil
}
