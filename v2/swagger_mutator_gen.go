package openapi

// This file was automatically generated by gentyeps.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"log"
)

var _ = log.Printf

// SwaggerMutator is used to build an instance of Swagger. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Swagger with new values
type SwaggerMutator struct {
	proxy  *swagger
	target *swagger
}

// Do finalizes the matuation process for Swagger and returns the result
func (b *SwaggerMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateSwagger creates a new mutator object for Swagger
func MutateSwagger(v Swagger) *SwaggerMutator {
	return &SwaggerMutator{
		target: v.(*swagger),
		proxy:  v.Clone().(*swagger),
	}
}

// Version sets the Version field for object Swagger.
func (b *SwaggerMutator) Version(v string) *SwaggerMutator {
	b.proxy.version = v
	return b
}

// Info sets the Info field for object Swagger.
func (b *SwaggerMutator) Info(v Info) *SwaggerMutator {
	b.proxy.info = v
	return b
}

// Host sets the Host field for object Swagger.
func (b *SwaggerMutator) Host(v string) *SwaggerMutator {
	b.proxy.host = v
	return b
}

// BasePath sets the BasePath field for object Swagger.
func (b *SwaggerMutator) BasePath(v string) *SwaggerMutator {
	b.proxy.basePath = v
	return b
}

func (b *SwaggerMutator) ClearSchemes() *SwaggerMutator {
	b.proxy.schemes.Clear()
	return b
}

func (b *SwaggerMutator) Scheme(value string) *SwaggerMutator {
	b.proxy.schemes = append(b.proxy.schemes, value)
	return b
}

func (b *SwaggerMutator) ClearConsumes() *SwaggerMutator {
	b.proxy.consumes.Clear()
	return b
}

func (b *SwaggerMutator) Consume(value string) *SwaggerMutator {
	b.proxy.consumes = append(b.proxy.consumes, value)
	return b
}

func (b *SwaggerMutator) ClearProduces() *SwaggerMutator {
	b.proxy.produces.Clear()
	return b
}

func (b *SwaggerMutator) Produce(value string) *SwaggerMutator {
	b.proxy.produces = append(b.proxy.produces, value)
	return b
}

// Paths sets the Paths field for object Swagger.
func (b *SwaggerMutator) Paths(v Paths) *SwaggerMutator {
	b.proxy.paths = v
	return b
}

func (b *SwaggerMutator) ClearDefinitions() *SwaggerMutator {
	b.proxy.definitions.Clear()
	return b
}

func (b *SwaggerMutator) Definition(key SchemaMapKey, value Schema) *SwaggerMutator {
	if b.proxy.definitions == nil {
		b.proxy.definitions = SchemaMap{}
	}

	b.proxy.definitions[key] = value
	return b
}

func (b *SwaggerMutator) ClearParameters() *SwaggerMutator {
	b.proxy.parameters.Clear()
	return b
}

func (b *SwaggerMutator) Parameter(key ParameterMapKey, value Parameter) *SwaggerMutator {
	if b.proxy.parameters == nil {
		b.proxy.parameters = ParameterMap{}
	}

	b.proxy.parameters[key] = value
	return b
}

func (b *SwaggerMutator) ClearResponses() *SwaggerMutator {
	b.proxy.responses.Clear()
	return b
}

func (b *SwaggerMutator) Response(key ResponseMapKey, value Response) *SwaggerMutator {
	if b.proxy.responses == nil {
		b.proxy.responses = ResponseMap{}
	}

	b.proxy.responses[key] = value
	return b
}

func (b *SwaggerMutator) ClearSecurityDefinitions() *SwaggerMutator {
	b.proxy.securityDefinitions.Clear()
	return b
}

func (b *SwaggerMutator) SecurityDefinition(key SecuritySchemeMapKey, value SecurityScheme) *SwaggerMutator {
	if b.proxy.securityDefinitions == nil {
		b.proxy.securityDefinitions = SecuritySchemeMap{}
	}

	b.proxy.securityDefinitions[key] = value
	return b
}

func (b *SwaggerMutator) ClearSecurity() *SwaggerMutator {
	b.proxy.security.Clear()
	return b
}

func (b *SwaggerMutator) Security(value SecurityRequirement) *SwaggerMutator {
	b.proxy.security = append(b.proxy.security, value)
	return b
}

func (b *SwaggerMutator) ClearTags() *SwaggerMutator {
	b.proxy.tags.Clear()
	return b
}

func (b *SwaggerMutator) Tag(value Tag) *SwaggerMutator {
	b.proxy.tags = append(b.proxy.tags, value)
	return b
}

// ExternalDocs sets the ExternalDocs field for object Swagger.
func (b *SwaggerMutator) ExternalDocs(v ExternalDocumentation) *SwaggerMutator {
	b.proxy.externalDocs = v
	return b
}
func (b *SwaggerMutator) Extension(name string, value interface{}) *SwaggerMutator {
	if b.proxy.extensions == nil {
		b.proxy.extensions = Extensions{}
	}
	b.proxy.extensions[name] = value
	return b
}
