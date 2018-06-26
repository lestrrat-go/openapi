package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

func (v *pathItem) Name() string {
	return v.name
}

func (v *pathItem) Path() string {
	return v.path
}

func (v *pathItem) Summary() string {
	return v.summary
}

func (v *pathItem) Description() string {
	return v.description
}

func (v *pathItem) Get() Operation {
	return v.get
}

func (v *pathItem) Put() Operation {
	return v.put
}

func (v *pathItem) Post() Operation {
	return v.post
}

func (v *pathItem) Delete() Operation {
	return v.delete
}

func (v *pathItem) Options() Operation {
	return v.options
}

func (v *pathItem) Head() Operation {
	return v.head
}

func (v *pathItem) Patch() Operation {
	return v.patch
}

func (v *pathItem) Trace() Operation {
	return v.trace
}

func (v *pathItem) Servers() *ServerListIterator {
	var items []interface{}
	for _, item := range v.servers {
		items = append(items, item)
	}
	var iter ServerListIterator
	iter.items = items
	return &iter
}

func (v *pathItem) Parameters() *ParameterListIterator {
	var items []interface{}
	for _, item := range v.parameters {
		items = append(items, item)
	}
	var iter ParameterListIterator
	iter.items = items
	return &iter
}

func (v *pathItem) Reference() string {
	return v.reference
}

func (v *pathItem) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *pathItem) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *pathItem) recurseValidate() error {
	if elem := v.get; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "get"`)
		}
	}
	if elem := v.put; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "put"`)
		}
	}
	if elem := v.post; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "post"`)
		}
	}
	if elem := v.delete; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "delete"`)
		}
	}
	if elem := v.options; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "options"`)
		}
	}
	if elem := v.head; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "head"`)
		}
	}
	if elem := v.patch; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "patch"`)
		}
	}
	if elem := v.trace; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "trace"`)
		}
	}
	if elem := v.servers; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "servers"`)
		}
	}
	if elem := v.parameters; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "parameters"`)
		}
	}
	return nil
}
