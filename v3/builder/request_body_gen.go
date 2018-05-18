package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type RequestBody struct {
	target *entity.RequestBody
}

func (b *RequestBody) Build() *entity.RequestBody {
	return b.target
}

func (b *Builder) NewRequestBody() *RequestBody {
	return &RequestBody{
		target: &entity.RequestBody{},
	}
}

func (b *RequestBody) Description(v string) *RequestBody {
	b.target.Description = v
	return b
}

func (b *RequestBody) Content(v map[string]*entity.MediaType) *RequestBody {
	b.target.Content = v
	return b
}

func (b *RequestBody) Required(v bool) *RequestBody {
	b.target.Required = v
	return b
}
