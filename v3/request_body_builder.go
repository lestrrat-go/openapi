package openapi

// Content sets the Content field for object RequestBody.
func (b *RequestBodyBuilder) Content(mime MediaTypeMapKey, mt MediaType) *RequestBodyBuilder {
	b.target.content[mime] = mt.Clone()
	b.target.content[mime].setMime(mime)
	return b
}


// Content sets the Content field for object RequestBody.
func (b *RequestBodyMutator) Content(mime MediaTypeMapKey, mt MediaType) *RequestBodyMutator {
	b.proxy.content[mime] = mt.Clone()
	b.proxy.content[mime].setMime(mime)
	return b
}


