package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// OpenAPIBuilder is used to build an instance of OpenAPI. The user must
// call `Build()` after providing all the necessary information to
// build an instance of OpenAPI
type OpenAPIBuilder struct {
	target *openAPI
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *OpenAPIBuilder) MustBuild(options ...Option) OpenAPI {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for OpenAPI and returns the result
func (b *OpenAPIBuilder) Build(options ...Option) (OpenAPI, error) {
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

// NewOpenAPI creates a new builder object for OpenAPI
func NewOpenAPI(info Info, paths Paths) *OpenAPIBuilder {
	return &OpenAPIBuilder{
		target: &openAPI{
			version: DefaultVersion,
			info:    info,
			paths:   paths,
		},
	}
}

// Version sets the Version field for object OpenAPI. If this is not called,
// a default value (DefaultVersion) is assigned to this field
func (b *OpenAPIBuilder) Version(v string) *OpenAPIBuilder {
	b.target.version = v
	return b
}

// Servers sets the Servers field for object OpenAPI.
func (b *OpenAPIBuilder) Servers(v []Server) *OpenAPIBuilder {
	b.target.servers = v
	return b
}

// Components sets the Components field for object OpenAPI.
func (b *OpenAPIBuilder) Components(v Components) *OpenAPIBuilder {
	b.target.components = v
	return b
}

// Security sets the Security field for object OpenAPI.
func (b *OpenAPIBuilder) Security(v SecurityRequirement) *OpenAPIBuilder {
	b.target.security = v
	return b
}

// Tags sets the Tags field for object OpenAPI.
func (b *OpenAPIBuilder) Tags(v []Tag) *OpenAPIBuilder {
	b.target.tags = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object OpenAPI.
func (b *OpenAPIBuilder) ExternalDocs(v ExternalDocumentation) *OpenAPIBuilder {
	b.target.externalDocs = v
	return b
}

// Reference sets the $ref (reference) field for object OpenAPI.
func (b *OpenAPIBuilder) Reference(v string) *OpenAPIBuilder {
	b.target.reference = v
	return b
}
