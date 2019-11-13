// +build !gogenerate

package openapi3

// Header sets the header `name` to `hdr`
func (b *ResponseBuilder) Header(name string, hdr Header) *ResponseBuilder {
	if b.target.headers == nil {
		b.target.headers = make(map[string]Header)
	}

	hdr = hdr.Clone()
	b.target.headers[name] = hdr
	hdr.setName(name)
	return b
}

// Content sets th content type `mime` to `desc`
func (b *ResponseBuilder) Content(mime string, desc MediaType) *ResponseBuilder {
	if b.target.content == nil {
		b.target.content = make(map[string]MediaType)
	}

	desc = desc.Clone()
	b.target.content[mime] = desc
	desc.setMime(mime)
	return b
}
