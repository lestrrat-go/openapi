package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *callback) Name() string {
	return v.name
}

func (v *callback) URLs() map[string]PathItem {
	return v.urls
}

func (v *callback) Reference() string {
	return v.reference
}

func (v *callback) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *callback) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *callback) recurseValidate() error {
	return nil
}
