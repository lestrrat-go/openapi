package openapi

// Path sets the path item for path `path` to `item`
func (b *PathsBuilder) Path(path string, item PathItem) *PathsBuilder {
	b.target.addPathItem(path, item)
	return b
}
