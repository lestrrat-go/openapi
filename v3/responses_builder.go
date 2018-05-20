// +build !gogenerate

package openapi

// Response sets the response for status code `code` to `v`
func (b *ResponsesBuilder) Response(code string, v Response) *ResponsesBuilder {
	if b.target.responses == nil {
		b.target.responses = make(map[string]Response)
	}
	b.target.responses[code] = v.Clone()
	return b
}

