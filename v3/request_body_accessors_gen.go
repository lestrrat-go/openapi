package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

func (v *requestBody) Name() string {
	return v.name
}

func (v *requestBody) Description() string {
	return v.description
}

func (v *requestBody) Content() *MediaTypeMapIterator {
	var items []interface{}
	for key, item := range v.content {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter MediaTypeMapIterator
	iter.list.items = items
	return &iter
}

func (v *requestBody) Required() bool {
	return v.required
}

func (v *requestBody) Reference() string {
	return v.reference
}

func (v *requestBody) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *requestBody) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
