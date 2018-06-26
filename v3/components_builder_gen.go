package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// ComponentsBuilder is used to build an instance of Components. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Components
type ComponentsBuilder struct {
	target *components
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *ComponentsBuilder) MustBuild(options ...Option) Components {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Components and returns the result
func (b *ComponentsBuilder) Build(options ...Option) (Components, error) {
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

// NewComponents creates a new builder object for Components
func NewComponents() *ComponentsBuilder {
	return &ComponentsBuilder{
		target: &components{},
	}
}

// Schemas sets the Schemas field for object Components.
func (b *ComponentsBuilder) Schemas(v map[string]Schema) *ComponentsBuilder {
	b.target.schemas = v
	return b
}

// Responses sets the Responses field for object Components.
func (b *ComponentsBuilder) Responses(v map[string]Response) *ComponentsBuilder {
	b.target.responses = v
	return b
}

// Parameters sets the Parameters field for object Components.
func (b *ComponentsBuilder) Parameters(v map[string]Parameter) *ComponentsBuilder {
	b.target.parameters = v
	return b
}

// Examples sets the Examples field for object Components.
func (b *ComponentsBuilder) Examples(v map[string]Example) *ComponentsBuilder {
	b.target.examples = v
	return b
}

// RequestBodies sets the RequestBodies field for object Components.
func (b *ComponentsBuilder) RequestBodies(v map[string]RequestBody) *ComponentsBuilder {
	b.target.requestBodies = v
	return b
}

// Headers sets the Headers field for object Components.
func (b *ComponentsBuilder) Headers(v map[string]Header) *ComponentsBuilder {
	b.target.headers = v
	return b
}

// SecuritySchemes sets the SecuritySchemes field for object Components.
func (b *ComponentsBuilder) SecuritySchemes(v map[string]SecurityScheme) *ComponentsBuilder {
	b.target.securitySchemes = v
	return b
}

// Links sets the Links field for object Components.
func (b *ComponentsBuilder) Links(v map[string]Link) *ComponentsBuilder {
	b.target.links = v
	return b
}

// Callbacks sets the Callbacks field for object Components.
func (b *ComponentsBuilder) Callbacks(v map[string]Callback) *ComponentsBuilder {
	b.target.callbacks = v
	return b
}

// Reference sets the $ref (reference) field for object Components.
func (b *ComponentsBuilder) Reference(v string) *ComponentsBuilder {
	b.target.reference = v
	return b
}
