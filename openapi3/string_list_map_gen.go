package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *StringListMap) Clear() error {
	*v = make(StringListMap)
	return nil
}

// Validate checks the correctness of values in StringListMap
func (v *StringListMap) Validate(recurse bool) error {
	return nil
}

func (v StringListMap) QueryJSON(path string) (ret interface{}, ok bool) {
	if path == `` {
		return v, true
	}

	var frag string
	frag, path = extractFragFromPath(path)
	target, ok := v[frag]
	if !ok {
		return nil, false
	}

	if path == `` {
		return target, true
	}
	return nil, false
}
