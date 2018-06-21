package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// ServerVariableBuilder is used to build an instance of ServerVariable. The user must
// call `Do()` after providing all the necessary information to
// build an instance of ServerVariable
type ServerVariableBuilder struct {
	target *serverVariable
}

// Do finalizes the building process for ServerVariable and returns the result
func (b *ServerVariableBuilder) Do() ServerVariable {
	return b.target
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
