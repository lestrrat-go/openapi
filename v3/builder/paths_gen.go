package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Paths struct {
	target *entity.Paths
}

func (b *Paths) Build() *entity.Paths {
	return b.target
}

func (b *Builder) NewPaths() *Paths {
	return &Paths{
		target: &entity.Paths{},
	}
}

func (b *Paths) Paths(v map[string]*entity.PathItem) *Paths {
	b.target.Paths = v
	return b
}
