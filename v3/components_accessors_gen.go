package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *components) Schemas() *SchemaMapIterator {
	var items []interface{}
	for key, item := range v.schemas {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SchemaMapIterator
	iter.list.items = items
	return &iter
}

func (v *components) Responses() *ResponseMapIterator {
	var items []interface{}
	for key, item := range v.responses {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ResponseMapIterator
	iter.list.items = items
	return &iter
}

func (v *components) Parameters() *ParameterMapIterator {
	var items []interface{}
	for key, item := range v.parameters {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ParameterMapIterator
	iter.list.items = items
	return &iter
}

func (v *components) Examples() *ExampleMapIterator {
	var items []interface{}
	for key, item := range v.examples {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExampleMapIterator
	iter.list.items = items
	return &iter
}

func (v *components) RequestBodies() *RequestBodyMapIterator {
	var items []interface{}
	for key, item := range v.requestBodies {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter RequestBodyMapIterator
	iter.list.items = items
	return &iter
}

func (v *components) Headers() *HeaderMapIterator {
	var items []interface{}
	for key, item := range v.headers {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter HeaderMapIterator
	iter.list.items = items
	return &iter
}

func (v *components) SecuritySchemes() *SecuritySchemeMapIterator {
	var items []interface{}
	for key, item := range v.securitySchemes {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter SecuritySchemeMapIterator
	iter.list.items = items
	return &iter
}

func (v *components) Links() *LinkMapIterator {
	var items []interface{}
	for key, item := range v.links {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter LinkMapIterator
	iter.list.items = items
	return &iter
}

func (v *components) Callbacks() *CallbackMapIterator {
	var items []interface{}
	for key, item := range v.callbacks {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter CallbackMapIterator
	iter.list.items = items
	return &iter
}

func (v *components) Reference() string {
	return v.reference
}

func (v *components) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *components) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *components) recurseValidate() error {
	if elem := v.schemas; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "schemas"`)
		}
	}
	if elem := v.responses; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "responses"`)
		}
	}
	if elem := v.parameters; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "parameters"`)
		}
	}
	if elem := v.examples; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "examples"`)
		}
	}
	if elem := v.requestBodies; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "requestBodies"`)
		}
	}
	if elem := v.headers; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "headers"`)
		}
	}
	if elem := v.securitySchemes; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "securitySchemes"`)
		}
	}
	if elem := v.links; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "links"`)
		}
	}
	if elem := v.callbacks; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "callbacks"`)
		}
	}
	return nil
}
