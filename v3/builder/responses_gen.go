package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Responses struct {
	target *entity.Responses
}

func (b *Responses) Build() *entity.Responses {
	return b.target
}

func (b *Builder) NewResponses() *Responses {
	return &Responses{
		target: &entity.Responses{},
	}
}

func (b *Responses) Default(v *entity.Response) *Responses {
	b.target.Default = v
	return b
}
