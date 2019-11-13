package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// OpenAPIMutator is used to build an instance of OpenAPI. The user must
// call `Do()` after providing all the necessary information to
// the new instance of OpenAPI with new values
type OpenAPIMutator struct {
	proxy  *openAPI
	target *openAPI
}

// Do finalizes the matuation process for OpenAPI and returns the result
func (b *OpenAPIMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateOpenAPI creates a new mutator object for OpenAPI
func MutateOpenAPI(v OpenAPI) *OpenAPIMutator {
	return &OpenAPIMutator{
		target: v.(*openAPI),
		proxy:  v.Clone().(*openAPI),
	}
}

// Version sets the Version field for object OpenAPI.
func (b *OpenAPIMutator) Version(v string) *OpenAPIMutator {
	b.proxy.version = v
	return b
}

// Info sets the Info field for object OpenAPI.
func (b *OpenAPIMutator) Info(v Info) *OpenAPIMutator {
	b.proxy.info = v
	return b
}

func (b *OpenAPIMutator) ClearServers() *OpenAPIMutator {
	b.proxy.servers.Clear()
	return b
}

func (b *OpenAPIMutator) Server(value Server) *OpenAPIMutator {
	b.proxy.servers = append(b.proxy.servers, value)
	return b
}

// Paths sets the Paths field for object OpenAPI.
func (b *OpenAPIMutator) Paths(v Paths) *OpenAPIMutator {
	b.proxy.paths = v
	return b
}

// Components sets the Components field for object OpenAPI.
func (b *OpenAPIMutator) Components(v Components) *OpenAPIMutator {
	b.proxy.components = v
	return b
}

// Security sets the Security field for object OpenAPI.
func (b *OpenAPIMutator) Security(v SecurityRequirement) *OpenAPIMutator {
	b.proxy.security = v
	return b
}

func (b *OpenAPIMutator) ClearTags() *OpenAPIMutator {
	b.proxy.tags.Clear()
	return b
}

func (b *OpenAPIMutator) Tag(value Tag) *OpenAPIMutator {
	b.proxy.tags = append(b.proxy.tags, value)
	return b
}

// ExternalDocs sets the ExternalDocs field for object OpenAPI.
func (b *OpenAPIMutator) ExternalDocs(v ExternalDocumentation) *OpenAPIMutator {
	b.proxy.externalDocs = v
	return b
}
