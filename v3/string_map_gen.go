package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *StringMap) Clear() error {
	*v = make(StringMap)
	return nil
}

// Validate checks the correctness of values in StringMap
func (v *StringMap) Validate(recurse bool) error {
	return nil
}

func (v StringMap) QueryJSON(path string) (ret interface{}, ok bool) {
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
