package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// ComponentsMutator is used to build an instance of Components. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Components with new values
type ComponentsMutator struct {
	proxy  *components
	target *components
}

// Do finalizes the matuation process for Components and returns the result
func (b *ComponentsMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateComponents creates a new mutator object for Components
func MutateComponents(v Components) *ComponentsMutator {
	return &ComponentsMutator{
		target: v.(*components),
		proxy:  v.Clone().(*components),
	}
}

func (b *ComponentsMutator) ClearSchemas() *ComponentsMutator {
	b.proxy.schemas.Clear()
	return b
}

func (b *ComponentsMutator) Schema(key SchemaMapKey, value Schema) *ComponentsMutator {
	if b.proxy.schemas == nil {
		b.proxy.schemas = SchemaMap{}
	}

	b.proxy.schemas[key] = value
	return b
}

func (b *ComponentsMutator) ClearResponses() *ComponentsMutator {
	b.proxy.responses.Clear()
	return b
}

func (b *ComponentsMutator) Response(key ResponseMapKey, value Response) *ComponentsMutator {
	if b.proxy.responses == nil {
		b.proxy.responses = ResponseMap{}
	}

	b.proxy.responses[key] = value
	return b
}

func (b *ComponentsMutator) ClearParameters() *ComponentsMutator {
	b.proxy.parameters.Clear()
	return b
}

func (b *ComponentsMutator) Parameter(key ParameterMapKey, value Parameter) *ComponentsMutator {
	if b.proxy.parameters == nil {
		b.proxy.parameters = ParameterMap{}
	}

	b.proxy.parameters[key] = value
	return b
}

func (b *ComponentsMutator) ClearExamples() *ComponentsMutator {
	b.proxy.examples.Clear()
	return b
}

func (b *ComponentsMutator) Example(key ExampleMapKey, value Example) *ComponentsMutator {
	if b.proxy.examples == nil {
		b.proxy.examples = ExampleMap{}
	}

	b.proxy.examples[key] = value
	return b
}

func (b *ComponentsMutator) ClearRequestBodies() *ComponentsMutator {
	b.proxy.requestBodies.Clear()
	return b
}

func (b *ComponentsMutator) RequestBody(key RequestBodyMapKey, value RequestBody) *ComponentsMutator {
	if b.proxy.requestBodies == nil {
		b.proxy.requestBodies = RequestBodyMap{}
	}

	b.proxy.requestBodies[key] = value
	return b
}

func (b *ComponentsMutator) ClearHeaders() *ComponentsMutator {
	b.proxy.headers.Clear()
	return b
}

func (b *ComponentsMutator) Header(key HeaderMapKey, value Header) *ComponentsMutator {
	if b.proxy.headers == nil {
		b.proxy.headers = HeaderMap{}
	}

	b.proxy.headers[key] = value
	return b
}

func (b *ComponentsMutator) ClearSecuritySchemes() *ComponentsMutator {
	b.proxy.securitySchemes.Clear()
	return b
}

func (b *ComponentsMutator) SecurityScheme(key SecuritySchemeMapKey, value SecurityScheme) *ComponentsMutator {
	if b.proxy.securitySchemes == nil {
		b.proxy.securitySchemes = SecuritySchemeMap{}
	}

	b.proxy.securitySchemes[key] = value
	return b
}

func (b *ComponentsMutator) ClearLinks() *ComponentsMutator {
	b.proxy.links.Clear()
	return b
}

func (b *ComponentsMutator) Link(key LinkMapKey, value Link) *ComponentsMutator {
	if b.proxy.links == nil {
		b.proxy.links = LinkMap{}
	}

	b.proxy.links[key] = value
	return b
}

func (b *ComponentsMutator) ClearCallbacks() *ComponentsMutator {
	b.proxy.callbacks.Clear()
	return b
}

func (b *ComponentsMutator) Callback(key CallbackMapKey, value Callback) *ComponentsMutator {
	if b.proxy.callbacks == nil {
		b.proxy.callbacks = CallbackMap{}
	}

	b.proxy.callbacks[key] = value
	return b
}
