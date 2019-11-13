package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// ResponseMutator is used to build an instance of Response. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Response with new values
type ResponseMutator struct {
	proxy  *response
	target *response
}

// Do finalizes the matuation process for Response and returns the result
func (b *ResponseMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateResponse creates a new mutator object for Response
func MutateResponse(v Response) *ResponseMutator {
	return &ResponseMutator{
		target: v.(*response),
		proxy:  v.Clone().(*response),
	}
}

// Name sets the Name field for object Response.
func (b *ResponseMutator) Name(v string) *ResponseMutator {
	b.proxy.name = v
	return b
}

// Description sets the Description field for object Response.
func (b *ResponseMutator) Description(v string) *ResponseMutator {
	b.proxy.description = v
	return b
}

func (b *ResponseMutator) ClearHeaders() *ResponseMutator {
	b.proxy.headers.Clear()
	return b
}

func (b *ResponseMutator) Header(key HeaderMapKey, value Header) *ResponseMutator {
	if b.proxy.headers == nil {
		b.proxy.headers = HeaderMap{}
	}

	b.proxy.headers[key] = value
	return b
}

func (b *ResponseMutator) ClearContent() *ResponseMutator {
	b.proxy.content.Clear()
	return b
}

func (b *ResponseMutator) Content(key MediaTypeMapKey, value MediaType) *ResponseMutator {
	if b.proxy.content == nil {
		b.proxy.content = MediaTypeMap{}
	}

	b.proxy.content[key] = value
	return b
}

func (b *ResponseMutator) ClearLinks() *ResponseMutator {
	b.proxy.links.Clear()
	return b
}

func (b *ResponseMutator) Link(key LinkMapKey, value Link) *ResponseMutator {
	if b.proxy.links == nil {
		b.proxy.links = LinkMap{}
	}

	b.proxy.links[key] = value
	return b
}
