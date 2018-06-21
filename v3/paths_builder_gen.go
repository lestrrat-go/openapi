package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// PathsBuilder is used to build an instance of Paths. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Paths
type PathsBuilder struct {
	target *paths
}

// Do finalizes the building process for Paths and returns the result
func (b *PathsBuilder) Do() Paths {
	return b.target
}

// NewPaths creates a new builder object for Paths
func NewPaths() *PathsBuilder {
	return &PathsBuilder{
		target: &paths{},
	}
}

// Paths sets the Paths field for object Paths.
func (b *PathsBuilder) Paths(v map[string]PathItem) *PathsBuilder {
	b.target.paths = v
	return b
}

// Reference sets the $ref (reference) field for object Paths.
func (b *PathsBuilder) Reference(v string) *PathsBuilder {
	b.target.reference = v
	return b
}
