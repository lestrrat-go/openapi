package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// PathItemBuilder is used to build an instance of PathItem. The user must
// call `Build()` after providing all the necessary information to
// build an instance of PathItem
type PathItemBuilder struct {
	target *pathItem
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *PathItemBuilder) MustBuild(options ...Option) PathItem {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for PathItem and returns the result
func (b *PathItemBuilder) Build(options ...Option) (PathItem, error) {
	validate := true
	for _, option := range options {
		switch option.Name() {
		case optkeyValidate:
			validate = option.Value().(bool)
		}
	}
	if validate {
		if err := b.target.Validate(false); err != nil {
			return nil, errors.Wrap(err, `validation failed`)
		}
	}
	return b.target, nil
}

// NewPathItem creates a new builder object for PathItem
func NewPathItem() *PathItemBuilder {
	return &PathItemBuilder{
		target: &pathItem{},
	}
}

// Path sets the Path field for object PathItem.
func (b *PathItemBuilder) Path(v string) *PathItemBuilder {
	b.target.path = v
	return b
}

// Summary sets the Summary field for object PathItem.
func (b *PathItemBuilder) Summary(v string) *PathItemBuilder {
	b.target.summary = v
	return b
}

// Description sets the Description field for object PathItem.
func (b *PathItemBuilder) Description(v string) *PathItemBuilder {
	b.target.description = v
	return b
}

// Servers sets the Servers field for object PathItem.
func (b *PathItemBuilder) Servers(v []Server) *PathItemBuilder {
	b.target.servers = v
	return b
}

// Parameters sets the Parameters field for object PathItem.
func (b *PathItemBuilder) Parameters(v []Parameter) *PathItemBuilder {
	b.target.parameters = v
	return b
}

// Reference sets the $ref (reference) field for object PathItem.
func (b *PathItemBuilder) Reference(v string) *PathItemBuilder {
	b.target.reference = v
	return b
}
