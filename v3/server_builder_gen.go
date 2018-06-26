package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// ServerBuilder is used to build an instance of Server. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Server
type ServerBuilder struct {
	target *server
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *ServerBuilder) MustBuild(options ...Option) Server {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Server and returns the result
func (b *ServerBuilder) Build(options ...Option) (Server, error) {
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

// NewServer creates a new builder object for Server
func NewServer(url string) *ServerBuilder {
	return &ServerBuilder{
		target: &server{
			url: url,
		},
	}
}

// Description sets the Description field for object Server.
func (b *ServerBuilder) Description(v string) *ServerBuilder {
	b.target.description = v
	return b
}

// Variables sets the Variables field for object Server.
func (b *ServerBuilder) Variables(v map[string]ServerVariable) *ServerBuilder {
	b.target.variables = v
	return b
}

// Reference sets the $ref (reference) field for object Server.
func (b *ServerBuilder) Reference(v string) *ServerBuilder {
	b.target.reference = v
	return b
}
