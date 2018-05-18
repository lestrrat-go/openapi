package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Encoding struct {
	target *entity.Encoding
}

func (b *Encoding) Build() *entity.Encoding {
	return b.target
}

func (b *Builder) NewEncoding() *Encoding {
	return &Encoding{
		target: &entity.Encoding{},
	}
}

func (b *Encoding) ContentType(v string) *Encoding {
	b.target.ContentType = v
	return b
}

func (b *Encoding) Headers(v map[string]*entity.Header) *Encoding {
	b.target.Headers = v
	return b
}

func (b *Encoding) Explode(v bool) *Encoding {
	b.target.Explode = v
	return b
}

func (b *Encoding) AllowReserved(v bool) *Encoding {
	b.target.AllowReserved = v
	return b
}
