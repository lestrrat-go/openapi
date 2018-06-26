package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

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
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *mediaType) recurseValidate() error {
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
	if elem := v.encoding; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "encoding"`)
		}
	}
	return nil
}
