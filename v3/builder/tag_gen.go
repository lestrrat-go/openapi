package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Tag struct {
	target *entity.Tag
}

func (b *Tag) Build() *entity.Tag {
	return b.target
}

func (b *Builder) NewTag(name string) *Tag {
	return &Tag{
		target: &entity.Tag{
			Name: name,
		},
	}
}

func (b *Tag) Description(v string) *Tag {
	b.target.Description = v
	return b
}

func (b *Tag) ExternalDocs(v *entity.ExternalDocumentation) *Tag {
	b.target.ExternalDocs = v
	return b
}
