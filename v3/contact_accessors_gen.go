package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

func (v *contact) Name() string {
	return v.name
}

func (v *contact) URL() string {
	return v.url
}

func (v *contact) Email() string {
	return v.email
}

func (v *contact) Reference() string {
	return v.reference
}

func (v *contact) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *contact) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
