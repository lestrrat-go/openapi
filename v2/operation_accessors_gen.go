package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"
	"sort"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = sort.Strings
var _ = errors.Cause

func (v *operation) IsValid() bool {
	return v != nil
}

func (v *operation) Verb() string {
	return v.verb
}

func (v *operation) PathItem() PathItem {
	return v.pathItem
}

func (v *operation) Tags() *StringListIterator {
	var items []interface{}
	for _, item := range v.tags {
		items = append(items, item)
	}
	var iter StringListIterator
	iter.size = len(items)
	iter.items = items
	return &iter
}

func (v *operation) Summary() string {
	return v.summary
}

func (v *operation) Description() string {
	return v.description
}

func (v *operation) ExternalDocs() ExternalDocumentation {
	return v.externalDocs
}

func (v *operation) OperationID() string {
	return v.operationID
}

func (v *operation) Consumes() *MIMETypeListIterator {
	var items []interface{}
	for _, item := range v.consumes {
		items = append(items, item)
	}
	var iter MIMETypeListIterator
	iter.size = len(items)
	iter.items = items
	return &iter
}

func (v *operation) Produces() *MIMETypeListIterator {
	var items []interface{}
	for _, item := range v.produces {
		items = append(items, item)
	}
	var iter MIMETypeListIterator
	iter.size = len(items)
	iter.items = items
	return &iter
}

func (v *operation) Parameters() *ParameterListIterator {
	var items []interface{}
	for _, item := range v.parameters {
		items = append(items, item)
	}
	var iter ParameterListIterator
	iter.size = len(items)
	iter.items = items
	return &iter
}

func (v *operation) Responses() Responses {
	return v.responses
}

func (v *operation) Schemes() *SchemeListIterator {
	var items []interface{}
	for _, item := range v.schemes {
		items = append(items, item)
	}
	var iter SchemeListIterator
	iter.size = len(items)
	iter.items = items
	return &iter
}

func (v *operation) Deprecated() bool {
	return v.deprecated
}

func (v *operation) Security() *SecurityRequirementListIterator {
	var items []interface{}
	for _, item := range v.security {
		items = append(items, item)
	}
	var iter SecurityRequirementListIterator
	iter.size = len(items)
	iter.items = items
	return &iter
}

// Reference returns the value of `$ref` field
func (v *operation) Reference() string {
	return v.reference
}

func (v *operation) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

// Extension returns the value of an arbitrary extension
func (v *operation) Extension(key string) (interface{}, bool) {
	e, ok := v.extensions[key]
	return e, ok
}

// Extensions return an iterator to iterate over all extensions
func (v *operation) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.size = len(items)
	iter.list.items = items
	return &iter
}

func (v *operation) Validate(recurse bool) error {
	return newValidator(recurse).Validate(context.Background(), v)
}
