package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

func (v *server) URL() string {
	return v.url
}

func (v *server) Description() string {
	return v.description
}

func (v *server) Variables() *ServerVariableMapIterator {
	var items []interface{}
	for key, item := range v.variables {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ServerVariableMapIterator
	iter.list.items = items
	return &iter
}

func (v *server) Reference() string {
	return v.reference
}

func (v *server) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *server) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
