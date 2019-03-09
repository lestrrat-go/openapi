package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

func (v *encoding) Name() string {
	return v.name
}

func (v *encoding) ContentType() string {
	return v.contentType
}

func (v *encoding) Headers() *HeaderMapIterator {
	var items []interface{}
	for key, item := range v.headers {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter HeaderMapIterator
	iter.list.items = items
	return &iter
}

func (v *encoding) Explode() bool {
	return v.explode
}

func (v *encoding) AllowReserved() bool {
	return v.allowReserved
}

func (v *encoding) Reference() string {
	return v.reference
}

func (v *encoding) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *encoding) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
