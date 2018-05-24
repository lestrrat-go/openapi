package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T12:53:29+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
)

var _ = errors.Cause

func (v *ServerList) Clear() error {
	*v = ServerList(nil)
	return nil
}

func (v ServerList) Resolve(resolver *Resolver) error {
	if len(v) > 0 {
		for i, elem := range v {
			if err := elem.Resolve(resolver); err != nil {
				return errors.Wrapf(err, `failed to resolve ServerList (index = %d)`, i)
			}
		}
	}
	return nil
}
