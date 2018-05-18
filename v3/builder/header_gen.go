package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Header struct {
	target *entity.Header
}

func (b *Header) Build() *entity.Header {
	return b.target
}

func (b *Builder) NewHeader() *Header {
	return &Header{
		target: &entity.Header{
			In: entity.InHeader,
		},
	}
}

func (b *Header) In(v entity.Location) *Header {
	b.target.In = v
	return b
}

func (b *Header) Required(v bool) *Header {
	b.target.Required = v
	return b
}

func (b *Header) Description(v string) *Header {
	b.target.Description = v
	return b
}

func (b *Header) Deprecated(v bool) *Header {
	b.target.Deprecated = v
	return b
}

func (b *Header) AllowEmptyValue(v bool) *Header {
	b.target.AllowEmptyValue = v
	return b
}

func (b *Header) Explode(v bool) *Header {
	b.target.Explode = v
	return b
}

func (b *Header) AllowReserved(v bool) *Header {
	b.target.AllowReserved = v
	return b
}

func (b *Header) Schema(v *entity.Schema) *Header {
	b.target.Schema = v
	return b
}

func (b *Header) Example(v interface{}) *Header {
	b.target.Example = v
	return b
}

func (b *Header) Examples(v map[string]*entity.Example) *Header {
	b.target.Examples = v
	return b
}

func (b *Header) Content(v map[string]*entity.MediaType) *Header {
	b.target.Content = v
	return b
}
