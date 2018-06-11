package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// PathsBuilder is used to build an instance of Paths. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Paths
type PathsBuilder struct {
	target *paths
}

// Do finalizes the building process for Paths and returns the result
func (b *PathsBuilder) Do(options ...Option) (Paths, error) {
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
func (b *PathsBuilder) Paths(v PathItemMap) *PathsBuilder {
	b.target.paths = v
	return b
}

// Reference sets the $ref (reference) field for object Paths.
func (b *PathsBuilder) Reference(v string) *PathsBuilder {
	b.target.reference = v
	return b
}
