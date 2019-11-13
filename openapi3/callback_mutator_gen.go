package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

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

// Name sets the Name field for object Callback.
func (b *CallbackMutator) Name(v string) *CallbackMutator {
	b.proxy.name = v
	return b
}

// URLs sets the URLs field for object Callback.
func (b *CallbackMutator) URLs(v map[string]PathItem) *CallbackMutator {
	b.proxy.urls = v
	return b
}
