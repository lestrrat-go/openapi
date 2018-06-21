package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"log"
)

var _ = log.Printf

// OpenApiMutator is used to build an instance of OpenApi. The user must
// call `Do()` after providing all the necessary information to
// the new instance of OpenApi with new values
type OpenApiMutator struct {
	proxy  *openAPI
	target *openAPI
}

// Do finalizes the matuation process for OpenApi and returns the result
func (b *OpenApiMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateOpenApi creates a new mutator object for OpenApi
func MutateOpenApi(v OpenApi) *OpenApiMutator {
	return &OpenApiMutator{
		target: v.(*openAPI),
		proxy:  v.Clone().(*openAPI),
	}
}

// Version sets the Version field for object OpenApi.
func (b *OpenApiMutator) Version(v string) *OpenApiMutator {
	b.proxy.version = v
	return b
}

// Info sets the Info field for object OpenApi.
func (b *OpenApiMutator) Info(v Info) *OpenApiMutator {
	b.proxy.info = v
	return b
}

func (b *OpenApiMutator) ClearServers() *OpenApiMutator {
	b.proxy.servers.Clear()
	return b
}

func (b *OpenApiMutator) Server(value Server) *OpenApiMutator {
	b.proxy.servers = append(b.proxy.servers, value)
	return b
}

// Paths sets the Paths field for object OpenApi.
func (b *OpenApiMutator) Paths(v Paths) *OpenApiMutator {
	b.proxy.paths = v
	return b
}

// Components sets the Components field for object OpenApi.
func (b *OpenApiMutator) Components(v Components) *OpenApiMutator {
	b.proxy.components = v
	return b
}

// Security sets the Security field for object OpenApi.
func (b *OpenApiMutator) Security(v SecurityRequirement) *OpenApiMutator {
	b.proxy.security = v
	return b
}

func (b *OpenApiMutator) ClearTags() *OpenApiMutator {
	b.proxy.tags.Clear()
	return b
}

func (b *OpenApiMutator) Tag(value Tag) *OpenApiMutator {
	b.proxy.tags = append(b.proxy.tags, value)
	return b
}

// ExternalDocs sets the ExternalDocs field for object OpenApi.
func (b *OpenApiMutator) ExternalDocs(v ExternalDocumentation) *OpenApiMutator {
	b.proxy.externalDocs = v
	return b
}
