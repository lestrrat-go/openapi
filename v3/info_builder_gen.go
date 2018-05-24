package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T13:02:46+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// InfoBuilder is used to build an instance of Info. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Info
type InfoBuilder struct {
	target *info
}

// Do finalizes the building process for Info and returns the result
func (b *InfoBuilder) Do() Info {
	return b.target
}

// NewInfo creates a new builder object for Info
func NewInfo(title string) *InfoBuilder {
	return &InfoBuilder{
		target: &info{
			version: DefaultSpecVersion,
			title:   title,
		},
	}
}

// Description sets the Description field for object Info.
func (b *InfoBuilder) Description(v string) *InfoBuilder {
	b.target.description = v
	return b
}

// TermsOfService sets the TermsOfService field for object Info.
func (b *InfoBuilder) TermsOfService(v string) *InfoBuilder {
	b.target.termsOfService = v
	return b
}

// Contact sets the Contact field for object Info.
func (b *InfoBuilder) Contact(v Contact) *InfoBuilder {
	b.target.contact = v
	return b
}

// License sets the License field for object Info.
func (b *InfoBuilder) License(v License) *InfoBuilder {
	b.target.license = v
	return b
}

// Version sets the Version field for object Info. If this is not called,
// a default value (DefaultSpecVersion) is assigned to this field
func (b *InfoBuilder) Version(v string) *InfoBuilder {
	b.target.version = v
	return b
}

// Reference sets the $ref (reference) field for object Info.
func (b *InfoBuilder) Reference(v string) *InfoBuilder {
	b.target.reference = v
	return b
}