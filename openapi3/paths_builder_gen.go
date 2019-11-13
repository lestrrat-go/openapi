package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// PathsBuilder is used to build an instance of Paths. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Paths
type PathsBuilder struct {
	target *paths
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *PathsBuilder) MustBuild(options ...Option) Paths {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Paths and returns the result
func (b *PathsBuilder) Build(options ...Option) (Paths, error) {
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

// NewPaths creates a new builder object for Paths
func NewPaths() *PathsBuilder {
	return &PathsBuilder{
		target: &paths{},
	}
}

// Paths sets the Paths field for object Paths.
func (b *PathsBuilder) Paths(v map[string]PathItem) *PathsBuilder {
	b.target.paths = v
	return b
}

// Reference sets the $ref (reference) field for object Paths.
func (b *PathsBuilder) Reference(v string) *PathsBuilder {
	b.target.reference = v
	return b
}
