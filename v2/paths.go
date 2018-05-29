package openapi

func (v *paths) addPathItem(path string, item PathItem) {
  if v.paths == nil {
    v.paths = make(map[string]PathItem)
  }

  v.paths[path] = item.Clone()
  v.paths[path].setPath(path)
}
