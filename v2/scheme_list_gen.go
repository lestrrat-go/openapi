package openapi

// This file was automatically generated by gentyeps.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *SchemeList) Clear() error {
	*v = SchemeList(nil)
	return nil
}
