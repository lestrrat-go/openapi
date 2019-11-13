package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

func (v *serverVariable) Name() string {
	return v.name
}

func (v *serverVariable) Enum() *StringListIterator {
	var items []interface{}
	for _, item := range v.enum {
		items = append(items, item)
	}
	var iter StringListIterator
	iter.items = items
	return &iter
}

func (v *serverVariable) Default() string {
	return v.defaultValue
}

func (v *serverVariable) Description() string {
	return v.description
}

func (v *serverVariable) Reference() string {
	return v.reference
}

func (v *serverVariable) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *serverVariable) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
