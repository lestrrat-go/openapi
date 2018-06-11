package openapi

import "github.com/pkg/errors"

func (v *paths) addPathItem(path string, item PathItem) {
	if v.paths == nil {
		v.paths = make(map[string]PathItem)
	}

	v.paths[path] = item.Clone()
	v.paths[path].setPath(path)
}

func (v *paths) Validate(recurse bool) error {
	for path, item := range v.paths {
		if err := item.Validate(recurse); err != nil {
			return errors.Wrapf(err, `failed to validate path %v`, path)
		}
	}
	return nil
}
