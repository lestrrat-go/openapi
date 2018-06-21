package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"log"
	"strings"

	"github.com/pkg/errors"
)

var _ = log.Printf
var _ = errors.Cause

func (v *paths) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*paths)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Paths, but got %T`, resolved)
		}
		mutator := MutatePaths(v)
		for iter := asserted.Paths(); iter.Next(); {
			key, item := iter.Item()
			mutator.Path(key, item)
		}
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.paths != nil {
		if err := v.paths.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Paths`)
		}
	}
	return nil
}

func (v *paths) QueryJSON(path string) (ret interface{}, ok bool) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	if path == "" {
		return v, true
	}

	var frag string
	if i := strings.Index(path, "/"); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ""
	}

	var target interface{}

	switch frag {
	default:
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}
	if path == "" {
		return target, true
	}
	return nil, false
}
