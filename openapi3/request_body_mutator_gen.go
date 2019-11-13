package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// RequestBodyMutator is used to build an instance of RequestBody. The user must
// call `Do()` after providing all the necessary information to
// the new instance of RequestBody with new values
type RequestBodyMutator struct {
	proxy  *requestBody
	target *requestBody
}

// Do finalizes the matuation process for RequestBody and returns the result
func (b *RequestBodyMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateRequestBody creates a new mutator object for RequestBody
func MutateRequestBody(v RequestBody) *RequestBodyMutator {
	return &RequestBodyMutator{
		target: v.(*requestBody),
		proxy:  v.Clone().(*requestBody),
	}
}

// Name sets the Name field for object RequestBody.
func (b *RequestBodyMutator) Name(v string) *RequestBodyMutator {
	b.proxy.name = v
	return b
}

// Description sets the Description field for object RequestBody.
func (b *RequestBodyMutator) Description(v string) *RequestBodyMutator {
	b.proxy.description = v
	return b
}

// Required sets the Required field for object RequestBody.
func (b *RequestBodyMutator) Required(v bool) *RequestBodyMutator {
	b.proxy.required = v
	return b
}
