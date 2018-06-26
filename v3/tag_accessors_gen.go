package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

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
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *tag) recurseValidate() error {
	if elem := v.externalDocs; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "externalDocs"`)
		}
	}
	return nil
}
