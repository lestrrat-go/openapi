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

// Clear removes all elements from SchemaMap
func (v *SchemaMap) Clear() error {
	*v = make(SchemaMap)
	return nil
}

// Validate checks the correctness of values in SchemaMap
func (v *SchemaMap) Validate(recurse bool) error {
	return Visit(context.Background(), newValidator(recurse), v)
}

// QueryJSON is used to query an element within the document
// Using jsonref
func (v SchemaMap) QueryJSON(path string) (ret interface{}, ok bool) {
	if path == `` {
		return v, true
	}

	var frag string
	frag, path = extractFragFromPath(path)
	target, ok := v[frag]
	if !ok {
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}

	if path == `` {
		return target, true
	}
	return nil, false
}

// UnmarshalJSON takes a JSON buffer and properly populates `v`
func (v *SchemaMap) UnmarshalJSON(data []byte) error {
	var proxy map[string]*schema
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal`)
	}
	tmp := make(map[string]Schema)
	for name, value := range proxy {
		value.setName(name)
		tmp[name] = value
	}
	*v = tmp
	return nil
}
