package openapi2


// Path sets the path item for path `path` to `item`
func (m *PathsMutator) Path(path string, item PathItem) *PathsMutator {
	m.proxy.addPathItem(path, item)
	return m
}
