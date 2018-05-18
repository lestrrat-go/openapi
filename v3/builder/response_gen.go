package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Response struct {
	target *entity.Response
}

func (b *Response) Build() *entity.Response {
	return b.target
}

func (b *Builder) NewResponse(description string) *Response {
	return &Response{
		target: &entity.Response{
			Description: description,
		},
	}
}
