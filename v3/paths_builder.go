// +build !gogenerate

package openapi

// Path sets the path item for path `path` to `item`
func (b *PathsBuilder) Path(path string, item PathItem) *PathsBuilder {
	if b.target.paths == nil {
		b.target.paths = make(map[string]PathItem)
	}

	b.target.paths[path] = item.Clone()
	b.target.paths[path].setPath(path)

	return b
}

