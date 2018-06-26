package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *openAPI) Version() string {
	return v.version
}

func (v *openAPI) Info() Info {
	return v.info
}

func (v *openAPI) Servers() *ServerListIterator {
	var items []interface{}
	for _, item := range v.servers {
		items = append(items, item)
	}
	var iter ServerListIterator
	iter.items = items
	return &iter
}

func (v *openAPI) Paths() Paths {
	return v.paths
}

func (v *openAPI) Components() Components {
	return v.components
}

func (v *openAPI) Security() SecurityRequirement {
	return v.security
}

func (v *openAPI) Tags() *TagListIterator {
	var items []interface{}
	for _, item := range v.tags {
		items = append(items, item)
	}
	var iter TagListIterator
	iter.items = items
	return &iter
}

func (v *openAPI) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

func (v *openAPI) Reference() string {
	return v.reference
}

func (v *openAPI) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *openAPI) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *openAPI) recurseValidate() error {
	if elem := v.info; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "info"`)
		}
	}
	if elem := v.servers; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "servers"`)
		}
	}
	if elem := v.paths; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "paths"`)
		}
	}
	if elem := v.components; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "components"`)
		}
	}
	if elem := v.security; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "security"`)
		}
	}
	if elem := v.tags; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "tags"`)
		}
	}
	if elem := v.externalDocs; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "externalDocs"`)
		}
	}
	return nil
}
