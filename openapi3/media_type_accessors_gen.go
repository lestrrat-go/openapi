package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

func (v *mediaType) Name() string {
	return v.name
}

func (v *mediaType) Mime() string {
	return v.mime
}

func (v *mediaType) Schema() Schema {
	return v.schema
}

func (v *mediaType) Examples() *ExampleMapIterator {
	var items []interface{}
	for key, item := range v.examples {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExampleMapIterator
	iter.list.items = items
	return &iter
}

func (v *mediaType) Encoding() *EncodingMapIterator {
	var items []interface{}
	for key, item := range v.encoding {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter EncodingMapIterator
	iter.list.items = items
	return &iter
}

func (v *mediaType) Reference() string {
	return v.reference
}

func (v *mediaType) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *mediaType) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
