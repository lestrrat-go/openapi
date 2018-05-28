package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T19:20:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

// SwaggerBuilder is used to build an instance of Swagger. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Swagger
type SwaggerBuilder struct {
	target *swagger
}

// Do finalizes the building process for Swagger and returns the result
func (b *SwaggerBuilder) Do() (Swagger, error) {
	if err := b.target.Validate(); err != nil {
		return nil, errors.Wrap(err, `validation failed`)
	}
	return b.target, nil
}

// NewSwagger creates a new builder object for Swagger
func NewSwagger(info Info, paths Paths) *SwaggerBuilder {
	return &SwaggerBuilder{
		target: &swagger{
			version: defaultSwaggerVersion,
			info:    info,
			paths:   paths,
		},
	}
}

// Host sets the Host field for object Swagger.
func (b *SwaggerBuilder) Host(v string) *SwaggerBuilder {
	b.target.host = v
	return b
}

// BasePath sets the BasePath field for object Swagger.
func (b *SwaggerBuilder) BasePath(v string) *SwaggerBuilder {
	b.target.basePath = v
	return b
}

// Schemes sets the Schemes field for object Swagger.
func (b *SwaggerBuilder) Schemes(v ...string) *SwaggerBuilder {
	b.target.schemes = v
	return b
}

// Consumes sets the Consumes field for object Swagger.
func (b *SwaggerBuilder) Consumes(v ...string) *SwaggerBuilder {
	b.target.consumes = v
	return b
}

// Produces sets the Produces field for object Swagger.
func (b *SwaggerBuilder) Produces(v ...string) *SwaggerBuilder {
	b.target.produces = v
	return b
}

// Definitions sets the Definitions field for object Swagger.
func (b *SwaggerBuilder) Definitions(v map[string]Schema) *SwaggerBuilder {
	b.target.definitions = v
	return b
}

// Parameters sets the Parameters field for object Swagger.
func (b *SwaggerBuilder) Parameters(v map[string]Parameter) *SwaggerBuilder {
	b.target.parameters = v
	return b
}

// Responses sets the Responses field for object Swagger.
func (b *SwaggerBuilder) Responses(v map[string]Response) *SwaggerBuilder {
	b.target.responses = v
	return b
}

// SecurityDefinitions sets the SecurityDefinitions field for object Swagger.
func (b *SwaggerBuilder) SecurityDefinitions(v map[string]SecurityScheme) *SwaggerBuilder {
	b.target.securityDefinitions = v
	return b
}

// Security sets the Security field for object Swagger.
func (b *SwaggerBuilder) Security(v ...SecurityRequirement) *SwaggerBuilder {
	b.target.security = v
	return b
}

// Tags sets the Tags field for object Swagger.
func (b *SwaggerBuilder) Tags(v ...Tag) *SwaggerBuilder {
	b.target.tags = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object Swagger.
func (b *SwaggerBuilder) ExternalDocs(v ExternalDocumentation) *SwaggerBuilder {
	b.target.externalDocs = v
	return b
}

// Reference sets the $ref (reference) field for object Swagger.
func (b *SwaggerBuilder) Reference(v string) *SwaggerBuilder {
	b.target.reference = v
	return b
}
