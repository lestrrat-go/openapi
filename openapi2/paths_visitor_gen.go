package openapi2

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

// PathsVisitor is an interface for objects that knows
// how to process Paths elements while traversing the OpenAPI structure
type PathsVisitor interface {
	VisitPaths(context.Context, Paths) error
}
