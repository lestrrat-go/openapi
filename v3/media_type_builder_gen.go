package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// MediaTypeBuilder is used to build an instance of MediaType. The user must
// call `Do()` after providing all the necessary information to
// build an instance of MediaType
type MediaTypeBuilder struct {
	target *mediaType
}

// Do finalizes the building process for MediaType and returns the result
func (b *MediaTypeBuilder) Do() MediaType {
	return b.target
}

// NewMediaType creates a new builder object for MediaType
func NewMediaType() *MediaTypeBuilder {
	return &MediaTypeBuilder{
		target: &mediaType{},
	}
}

// Schema sets the Schema field for object MediaType.
func (b *MediaTypeBuilder) Schema(v Schema) *MediaTypeBuilder {
	b.target.schema = v
	return b
}

// Examples sets the Examples field for object MediaType.
func (b *MediaTypeBuilder) Examples(v map[string]Example) *MediaTypeBuilder {
	b.target.examples = v
	return b
}

// Encoding sets the Encoding field for object MediaType.
func (b *MediaTypeBuilder) Encoding(v map[string]Encoding) *MediaTypeBuilder {
	b.target.encoding = v
	return b
}

// Reference sets the $ref (reference) field for object MediaType.
func (b *MediaTypeBuilder) Reference(v string) *MediaTypeBuilder {
	b.target.reference = v
	return b
}
