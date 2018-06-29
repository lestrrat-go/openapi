package openapi

import (
	"context"

	"github.com/pkg/errors"
)

func (v *paths) addPathItem(path string, item PathItem) {
	if v.paths == nil {
		v.paths = make(map[string]PathItem)
	}

	v.paths[path] = item.Clone()
	v.paths[path].setPath(path)
}

type pathItemKeyVisitorKey struct{}

func visitPaths(ctx context.Context, v Paths) error {
	for iter := v.Paths(); iter.Next(); {
		path, item := iter.Item()
		if err := visitPathItem(context.WithValue(ctx, pathItemKeyVisitorKey{}, path), item); err != nil {
			return errors.Wrapf(err, `failed to visit path %v`, path)
		}
	}
	return nil
}
