package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

func (v *example) Name() string {
	return v.name
}

func (v *example) Description() string {
	return v.description
}

func (v *example) Value() interface{} {
	return v.value
}

func (v *example) ExternalValue() string {
	return v.externalValue
}

func (v *example) Reference() string {
	return v.reference
}

func (v *example) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *example) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
