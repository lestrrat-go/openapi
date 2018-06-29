package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

// DiscriminatorVisitor is an interface for objects that knows
// how to process Discriminator elements while traversing the OpenAPI structure
type DiscriminatorVisitor interface {
	VisitDiscriminator(context.Context, Discriminator) error
}

func visitDiscriminator(ctx context.Context, elem Discriminator) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if v, ok := ctx.Value(discriminatorVisitorCtxKey{}).(DiscriminatorVisitor); ok {
		if err := v.VisitDiscriminator(ctx, elem); err != nil {
			if err == ErrVisitAbort {
				return nil
			}
			return errors.Wrap(err, `failed to visit Discriminator element`)
		}
	}
	return nil
}
