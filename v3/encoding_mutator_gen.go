package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// EncodingMutator is used to build an instance of Encoding. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Encoding with new values
type EncodingMutator struct {
	proxy  *encoding
	target *encoding
}

// Do finalizes the matuation process for Encoding and returns the result
func (b *EncodingMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateEncoding creates a new mutator object for Encoding
func MutateEncoding(v Encoding) *EncodingMutator {
	return &EncodingMutator{
		target: v.(*encoding),
		proxy:  v.Clone().(*encoding),
	}
}

// Name sets the Name field for object Encoding.
func (b *EncodingMutator) Name(v string) *EncodingMutator {
	b.proxy.name = v
	return b
}

// ContentType sets the ContentType field for object Encoding.
func (b *EncodingMutator) ContentType(v string) *EncodingMutator {
	b.proxy.contentType = v
	return b
}

func (b *EncodingMutator) ClearHeaders() *EncodingMutator {
	b.proxy.headers.Clear()
	return b
}

func (b *EncodingMutator) Header(key HeaderMapKey, value Header) *EncodingMutator {
	if b.proxy.headers == nil {
		b.proxy.headers = HeaderMap{}
	}

	b.proxy.headers[key] = value
	return b
}

// Explode sets the Explode field for object Encoding.
func (b *EncodingMutator) Explode(v bool) *EncodingMutator {
	b.proxy.explode = v
	return b
}

// AllowReserved sets the AllowReserved field for object Encoding.
func (b *EncodingMutator) AllowReserved(v bool) *EncodingMutator {
	b.proxy.allowReserved = v
	return b
}
