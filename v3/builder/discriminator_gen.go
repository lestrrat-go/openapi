package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Discriminator struct {
	target *entity.Discriminator
}

func (b *Discriminator) Build() *entity.Discriminator {
	return b.target
}

func (b *Builder) NewDiscriminator() *Discriminator {
	return &Discriminator{
		target: &entity.Discriminator{},
	}
}

func (b *Discriminator) PropertyName(v string) *Discriminator {
	b.target.PropertyName = v
	return b
}

func (b *Discriminator) Mapping(v map[string]string) *Discriminator {
	b.target.Mapping = v
	return b
}
