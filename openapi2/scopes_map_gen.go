package openapi2

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = json.Unmarshal
var _ = errors.Cause

// Clear removes all elements from ScopesMap
func (v *ScopesMap) Clear() error {
	*v = make(ScopesMap)
	return nil
}

// Validate checks the correctness of values in ScopesMap
func (v *ScopesMap) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}

// QueryJSON is used to query an element within the document
// Using jsonref
func (v ScopesMap) QueryJSON(path string) (ret interface{}, ok bool) {
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
