package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *parameter) Name() string {
	return v.name
}

func (v *parameter) In() Location {
	return v.in
}

func (v *parameter) Required() bool {
	return v.required
}

func (v *parameter) Description() string {
	return v.description
}

func (v *parameter) Deprecated() bool {
	return v.deprecated
}

func (v *parameter) AllowEmptyValue() bool {
	return v.allowEmptyValue
}

func (v *parameter) Explode() bool {
	return v.explode
}

func (v *parameter) AllowReserved() bool {
	return v.allowReserved
}

func (v *parameter) Schema() Schema {
	return v.schema
}

func (v *parameter) Examples() *ExampleMapIterator {
	var items []interface{}
	for key, item := range v.examples {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExampleMapIterator
	iter.list.items = items
	return &iter
}

func (v *parameter) Content() *MediaTypeMapIterator {
	var items []interface{}
	for key, item := range v.content {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter MediaTypeMapIterator
	iter.list.items = items
	return &iter
}

func (v *parameter) Reference() string {
	return v.reference
}

func (v *parameter) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *parameter) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *parameter) recurseValidate() error {
	if elem := v.schema; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "schema"`)
		}
	}
	if elem := v.examples; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "examples"`)
		}
	}
	if elem := v.content; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "content"`)
		}
	}
	return nil
}
