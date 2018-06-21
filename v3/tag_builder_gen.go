package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// TagBuilder is used to build an instance of Tag. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Tag
type TagBuilder struct {
	target *tag
}

// Do finalizes the building process for Tag and returns the result
func (b *TagBuilder) Do() Tag {
	return b.target
}

// NewTag creates a new builder object for Tag
func NewTag(name string) *TagBuilder {
	return &TagBuilder{
		target: &tag{
			name: name,
		},
	}
}

// Description sets the Description field for object Tag.
func (b *TagBuilder) Description(v string) *TagBuilder {
	b.target.description = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object Tag.
func (b *TagBuilder) ExternalDocs(v ExternalDocumentation) *TagBuilder {
	b.target.externalDocs = v
	return b
}

// Reference sets the $ref (reference) field for object Tag.
func (b *TagBuilder) Reference(v string) *TagBuilder {
	b.target.reference = v
	return b
}
