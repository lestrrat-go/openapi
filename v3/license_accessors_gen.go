package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

func (v *license) Name() string {
	return v.name
}

func (v *license) URL() string {
	return v.url
}

func (v *license) Reference() string {
	return v.reference
}

func (v *license) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *license) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
