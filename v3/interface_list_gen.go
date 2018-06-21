package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *InterfaceList) Clear() error {
	*v = InterfaceList(nil)
	return nil
}

func (v InterfaceList) Resolve(resolver *Resolver) error {
	return nil
}
