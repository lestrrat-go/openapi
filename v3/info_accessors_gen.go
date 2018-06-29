package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

func (v *info) Title() string {
	return v.title
}

func (v *info) Description() string {
	return v.description
}

func (v *info) TermsOfService() string {
	return v.termsOfService
}

func (v *info) Contact() Contact {
	return v.contact
}

func (v *info) License() License {
	return v.license
}

func (v *info) Version() string {
	return v.version
}

func (v *info) Reference() string {
	return v.reference
}

func (v *info) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *info) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
