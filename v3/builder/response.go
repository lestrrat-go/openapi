package builder

import "github.com/lestrrat-go/openapi/v3/entity"

func (b *Response) Header(name string, hdr *entity.Header) *Response {
	if b.target.Headers == nil {
		b.target.Headers = make(map[string]*entity.Header)
	}

	b.target.Headers[name] = hdr
	return b
}

func (b *Response) Content(name string, desc *entity.MediaType) *Response {
	if b.target.Content == nil {
		b.target.Content = make(map[string]*entity.MediaType)
	}

	b.target.Content[name] = desc
	return b
}
