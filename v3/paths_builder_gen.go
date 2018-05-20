package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:57:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// PathsBuilder is used to build an instance of Paths. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Paths
type PathsBuilder struct {
	target *paths
}

// Build finalizes the building process for Paths and returns the result
func (b *PathsBuilder) Build() Paths {
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
