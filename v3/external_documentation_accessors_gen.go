package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

func (v *externalDocumentation) Description() string {
	return v.description
}

func (v *externalDocumentation) URL() string {
	return v.url
}

func (v *externalDocumentation) Reference() string {
	return v.reference
}

func (v *externalDocumentation) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *externalDocumentation) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
