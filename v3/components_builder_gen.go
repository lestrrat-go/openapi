package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// ComponentsBuilder is used to build an instance of Components. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Components
type ComponentsBuilder struct {
	target *components
}

// Do finalizes the building process for Components and returns the result
func (b *ComponentsBuilder) Do() Components {
	return b.target
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
