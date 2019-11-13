// +build !gogenerate

package openapi3

func (v *paths) addPathItem(path string, item PathItem) {
	if v.paths == nil {
		v.paths = make(map[string]PathItem)
	}

	v.paths[path] = item.Clone()
	v.paths[path].setPath(path)
}

// Path sets the path item for path `path` to `item`
func (b *PathsBuilder) Path(path string, item PathItem) *PathsBuilder {
	b.target.addPathItem(path, item)
	return b
}

// Path sets the path item for path `path` to `item`
func (m *PathsMutator) Path(path string, item PathItem) *PathsMutator {
	m.proxy.addPathItem(path, item)
	return m
}


