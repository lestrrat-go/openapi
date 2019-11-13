package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// ResponsesMutator is used to build an instance of Responses. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Responses with new values
type ResponsesMutator struct {
	proxy  *responses
	target *responses
}

// Do finalizes the matuation process for Responses and returns the result
func (b *ResponsesMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateResponses creates a new mutator object for Responses
func MutateResponses(v Responses) *ResponsesMutator {
	return &ResponsesMutator{
		target: v.(*responses),
		proxy:  v.Clone().(*responses),
	}
}

// Default sets the Default field for object Responses.
func (b *ResponsesMutator) Default(v Response) *ResponsesMutator {
	b.proxy.defaultValue = v
	return b
}

func (b *ResponsesMutator) ClearResponses() *ResponsesMutator {
	b.proxy.responses.Clear()
	return b
}

func (b *ResponsesMutator) Response(key ResponseMapKey, value Response) *ResponsesMutator {
	if b.proxy.responses == nil {
		b.proxy.responses = ResponseMap{}
	}

	b.proxy.responses[key] = value
	return b
}
