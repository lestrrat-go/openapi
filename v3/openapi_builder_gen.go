package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// OpenApiBuilder is used to build an instance of OpenApi. The user must
// call `Do()` after providing all the necessary information to
// build an instance of OpenApi
type OpenApiBuilder struct {
	target *openAPI
}

// Do finalizes the building process for OpenApi and returns the result
func (b *OpenApiBuilder) Do() OpenApi {
	return b.target
}

// NewOpenApi creates a new builder object for OpenApi
func NewOpenApi(info Info, paths Paths) *OpenApiBuilder {
	return &OpenApiBuilder{
		target: &openAPI{
			version: DefaultVersion,
			info:    info,
			paths:   paths,
		},
	}
}

// Version sets the Version field for object OpenApi. If this is not called,
// a default value (DefaultVersion) is assigned to this field
func (b *OpenApiBuilder) Version(v string) *OpenApiBuilder {
	b.target.version = v
	return b
}

// Servers sets the Servers field for object OpenApi.
func (b *OpenApiBuilder) Servers(v []Server) *OpenApiBuilder {
	b.target.servers = v
	return b
}

// Components sets the Components field for object OpenApi.
func (b *OpenApiBuilder) Components(v Components) *OpenApiBuilder {
	b.target.components = v
	return b
}

// Security sets the Security field for object OpenApi.
func (b *OpenApiBuilder) Security(v SecurityRequirement) *OpenApiBuilder {
	b.target.security = v
	return b
}

// Tags sets the Tags field for object OpenApi.
func (b *OpenApiBuilder) Tags(v []Tag) *OpenApiBuilder {
	b.target.tags = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object OpenApi.
func (b *OpenApiBuilder) ExternalDocs(v ExternalDocumentation) *OpenApiBuilder {
	b.target.externalDocs = v
	return b
}

// Reference sets the $ref (reference) field for object OpenApi.
func (b *OpenApiBuilder) Reference(v string) *OpenApiBuilder {
	b.target.reference = v
	return b
}
