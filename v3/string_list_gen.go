package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *StringList) Clear() error {
	*v = StringList(nil)
	return nil
}

func (v StringList) Resolve(resolver *Resolver) error {
	return nil
}
