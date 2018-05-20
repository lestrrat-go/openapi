// +build !gogenerate

package openapi

// Header sets the header `name` to `hdr`
func (b *ResponseBuilder) Header(name string, hdr Header) *ResponseBuilder {
	if b.target.headers == nil {
		b.target.headers = make(map[string]Header)
	}

	b.target.headers[name] = hdr.Clone()
	return b
}

// Content sets th content type `name` to `desc`
func (b *ResponseBuilder) Content(name string, desc MediaType) *ResponseBuilder {
	if b.target.content == nil {
		b.target.content = make(map[string]MediaType)
	}

	b.target.content[name] = desc.Clone()
	return b
}

