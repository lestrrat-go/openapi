package openapi

// Content sets the Content field for object RequestBody.
func (b *RequestBodyBuilder) Content(mime string, mt MediaType) *RequestBodyBuilder {
	b.target.content[mime] = mt.Clone()
	b.target.content[mime].setMime(mime)
	return b
}

