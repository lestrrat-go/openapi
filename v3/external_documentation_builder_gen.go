package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// ExternalDocumentationBuilder is used to build an instance of ExternalDocumentation. The user must
// call `Do()` after providing all the necessary information to
// build an instance of ExternalDocumentation
type ExternalDocumentationBuilder struct {
	target *externalDocumentation
}

// Do finalizes the building process for ExternalDocumentation and returns the result
func (b *ExternalDocumentationBuilder) Do() ExternalDocumentation {
	return b.target
}

// NewExternalDocumentation creates a new builder object for ExternalDocumentation
func NewExternalDocumentation(url string) *ExternalDocumentationBuilder {
	return &ExternalDocumentationBuilder{
		target: &externalDocumentation{
			url: url,
		},
	}
}

// Description sets the Description field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Description(v string) *ExternalDocumentationBuilder {
	b.target.description = v
	return b
}

// Reference sets the $ref (reference) field for object ExternalDocumentation.
func (b *ExternalDocumentationBuilder) Reference(v string) *ExternalDocumentationBuilder {
	b.target.reference = v
	return b
}
