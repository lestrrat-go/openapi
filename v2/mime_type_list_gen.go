package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T19:20:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = json.Unmarshal
var _ = errors.Cause

func (v *MIMETypeList) Clear() error {
	*v = MIMETypeList(nil)
	return nil
}

func (v MIMETypeList) Resolve(resolver *Resolver) error {
	return nil
}
