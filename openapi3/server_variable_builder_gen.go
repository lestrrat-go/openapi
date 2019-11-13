package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// ServerVariableBuilder is used to build an instance of ServerVariable. The user must
// call `Build()` after providing all the necessary information to
// build an instance of ServerVariable
type ServerVariableBuilder struct {
	target *serverVariable
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *ServerVariableBuilder) MustBuild(options ...Option) ServerVariable {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for ServerVariable and returns the result
func (b *ServerVariableBuilder) Build(options ...Option) (ServerVariable, error) {
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

// NewServerVariable creates a new builder object for ServerVariable
func NewServerVariable(defaultValue string) *ServerVariableBuilder {
	return &ServerVariableBuilder{
		target: &serverVariable{
			defaultValue: defaultValue,
		},
	}
}

// Enum sets the Enum field for object ServerVariable.
func (b *ServerVariableBuilder) Enum(v []string) *ServerVariableBuilder {
	b.target.enum = v
	return b
}

// Description sets the Description field for object ServerVariable.
func (b *ServerVariableBuilder) Description(v string) *ServerVariableBuilder {
	b.target.description = v
	return b
}

// Reference sets the $ref (reference) field for object ServerVariable.
func (b *ServerVariableBuilder) Reference(v string) *ServerVariableBuilder {
	b.target.reference = v
	return b
}
