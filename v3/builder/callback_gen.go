package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Callback struct {
	target *entity.Callback
}

func (b *Callback) Build() *entity.Callback {
	return b.target
}

func (b *Builder) NewCallback() *Callback {
	return &Callback{
		target: &entity.Callback{},
	}
}

func (b *Callback) URLs(v map[string]*entity.PathItem) *Callback {
	b.target.URLs = v
	return b
}
