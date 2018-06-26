package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

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

func (v *operation) Parameters() *ParameterListIterator {
	var items []interface{}
	for _, item := range v.parameters {
		items = append(items, item)
	}
	var iter ParameterListIterator
	iter.items = items
	return &iter
}

func (v *operation) RequestBody() RequestBody {
	return v.requestBody
}

func (v *operation) Responses() Responses {
	return v.responses
}

func (v *operation) Callbacks() *CallbackMapIterator {
	var items []interface{}
	for key, item := range v.callbacks {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter CallbackMapIterator
	iter.list.items = items
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
	iter.items = items
	return &iter
}

func (v *operation) Servers() *ServerListIterator {
	var items []interface{}
	for _, item := range v.servers {
		items = append(items, item)
	}
	var iter ServerListIterator
	iter.items = items
	return &iter
}

func (v *operation) Reference() string {
	return v.reference
}

func (v *operation) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *operation) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *operation) recurseValidate() error {
	if elem := v.tags; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "tags"`)
		}
	}
	if elem := v.externalDocs; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "externalDocs"`)
		}
	}
	if elem := v.parameters; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "parameters"`)
		}
	}
	if elem := v.requestBody; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "requestBody"`)
		}
	}
	if elem := v.responses; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "responses"`)
		}
	}
	if elem := v.callbacks; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "callbacks"`)
		}
	}
	if elem := v.security; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "security"`)
		}
	}
	if elem := v.servers; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "servers"`)
		}
	}
	return nil
}
