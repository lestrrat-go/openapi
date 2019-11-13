package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// ParameterBuilder is used to build an instance of Parameter. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Parameter
type ParameterBuilder struct {
	target *parameter
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *ParameterBuilder) MustBuild(options ...Option) Parameter {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Parameter and returns the result
func (b *ParameterBuilder) Build(options ...Option) (Parameter, error) {
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

// NewParameter creates a new builder object for Parameter
func NewParameter(name string, in Location) *ParameterBuilder {
	return &ParameterBuilder{
		target: &parameter{
			required: defaultParameterRequiredFromLocation(in),
			name:     name,
			in:       in,
		},
	}
}

// Required sets the Required field for object Parameter. If this is not called,
// a default value (defaultParameterRequiredFromLocation(in)) is assigned to this field
func (b *ParameterBuilder) Required(v bool) *ParameterBuilder {
	b.target.required = v
	return b
}

// Description sets the Description field for object Parameter.
func (b *ParameterBuilder) Description(v string) *ParameterBuilder {
	b.target.description = v
	return b
}

// Deprecated sets the Deprecated field for object Parameter.
func (b *ParameterBuilder) Deprecated(v bool) *ParameterBuilder {
	b.target.deprecated = v
	return b
}

// AllowEmptyValue sets the AllowEmptyValue field for object Parameter.
func (b *ParameterBuilder) AllowEmptyValue(v bool) *ParameterBuilder {
	b.target.allowEmptyValue = v
	return b
}

// Explode sets the Explode field for object Parameter.
func (b *ParameterBuilder) Explode(v bool) *ParameterBuilder {
	b.target.explode = v
	return b
}

// AllowReserved sets the AllowReserved field for object Parameter.
func (b *ParameterBuilder) AllowReserved(v bool) *ParameterBuilder {
	b.target.allowReserved = v
	return b
}

// Schema sets the Schema field for object Parameter.
func (b *ParameterBuilder) Schema(v Schema) *ParameterBuilder {
	b.target.schema = v
	return b
}

// Examples sets the Examples field for object Parameter.
func (b *ParameterBuilder) Examples(v map[string]Example) *ParameterBuilder {
	b.target.examples = v
	return b
}

// Content sets the Content field for object Parameter.
func (b *ParameterBuilder) Content(v map[string]MediaType) *ParameterBuilder {
	b.target.content = v
	return b
}

// Reference sets the $ref (reference) field for object Parameter.
func (b *ParameterBuilder) Reference(v string) *ParameterBuilder {
	b.target.reference = v
	return b
}
