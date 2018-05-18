package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type MediaType struct {
	target *entity.MediaType
}

func (b *MediaType) Build() *entity.MediaType {
	return b.target
}

func (b *Builder) NewMediaType() *MediaType {
	return &MediaType{
		target: &entity.MediaType{},
	}
}

func (b *MediaType) Schema(v *entity.Schema) *MediaType {
	b.target.Schema = v
	return b
}

func (b *MediaType) Example(v interface{}) *MediaType {
	b.target.Example = v
	return b
}

func (b *MediaType) Examples(v map[string]*entity.Example) *MediaType {
	b.target.Examples = v
	return b
}

func (b *MediaType) Encoding(v map[string]*entity.Encoding) *MediaType {
	b.target.Encoding = v
	return b
}
