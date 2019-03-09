package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause
var _ = context.Background

func (v *link) Name() string {
	return v.name
}

func (v *link) OperationRef() string {
	return v.operationRef
}

func (v *link) OperationID() string {
	return v.operationID
}

func (v *link) Parameters() *InterfaceMapIterator {
	var items []interface{}
	for key, item := range v.parameters {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter InterfaceMapIterator
	iter.list.items = items
	return &iter
}

func (v *link) RequestBody() interface{} {
	return v.requestBody
}

func (v *link) Description() string {
	return v.description
}

func (v *link) Server() Server {
	return v.server
}

func (v *link) Reference() string {
	return v.reference
}

func (v *link) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *link) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}
