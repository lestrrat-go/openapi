package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T13:02:46+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"log"
)

var _ = log.Printf

// CallbackMutator is used to build an instance of Callback. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Callback with new values
type CallbackMutator struct {
	proxy  *callback
	target *callback
}

// Do finalizes the matuation process for Callback and returns the result
func (b *CallbackMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateCallback creates a new mutator object for Callback
func MutateCallback(v Callback) *CallbackMutator {
	return &CallbackMutator{
		target: v.(*callback),
		proxy:  v.Clone().(*callback),
	}
}

// URLs sets the URLs field for object Callback.
func (b *CallbackMutator) URLs(v map[string]PathItem) *CallbackMutator {
	b.proxy.uRLs = v
	return b
}