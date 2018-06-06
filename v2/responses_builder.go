package openapi

// Response sets the response for status code `code` to `v`
func (b *ResponsesBuilder) Response(code string, v Response) *ResponsesBuilder {
	b.target.setResponse(code, v)
	return b
}
