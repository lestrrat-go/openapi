package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

func (v *tag) Name() string {
	return v.name
}

func (v *tag) Description() string {
	return v.description
}

func (v *tag) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

func (v *tag) Reference() string {
	return v.reference
}

func (v *tag) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *tag) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
