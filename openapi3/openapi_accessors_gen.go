package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

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
	return Visit(context.Background(), newValidator(recurse), v)
}
