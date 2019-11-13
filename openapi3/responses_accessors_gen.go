package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

func (v *responses) Default() Response {
	return v.defaultValue
}

func (v *responses) Responses() *ResponseMapIterator {
	var items []interface{}
	for key, item := range v.responses {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ResponseMapIterator
	iter.list.items = items
	return &iter
}

func (v *responses) Reference() string {
	return v.reference
}

func (v *responses) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *responses) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
