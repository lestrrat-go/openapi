package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *response) Name() string {
	return v.name
}

func (v *response) Description() string {
	return v.description
}

func (v *response) Headers() *HeaderMapIterator {
	var items []interface{}
	for key, item := range v.headers {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter HeaderMapIterator
	iter.list.items = items
	return &iter
}

func (v *response) Content() *MediaTypeMapIterator {
	var items []interface{}
	for key, item := range v.content {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter MediaTypeMapIterator
	iter.list.items = items
	return &iter
}

func (v *response) Links() *LinkMapIterator {
	var items []interface{}
	for key, item := range v.links {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter LinkMapIterator
	iter.list.items = items
	return &iter
}

func (v *response) Reference() string {
	return v.reference
}

func (v *response) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *response) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *response) recurseValidate() error {
	if elem := v.headers; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "headers"`)
		}
	}
	if elem := v.content; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "content"`)
		}
	}
	if elem := v.links; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "links"`)
		}
	}
	return nil
}
