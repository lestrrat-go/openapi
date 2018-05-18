package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type SecurityRequirement struct {
	target *entity.SecurityRequirement
}

func (b *SecurityRequirement) Build() *entity.SecurityRequirement {
	return b.target
}

func (b *Builder) NewSecurityRequirement() *SecurityRequirement {
	return &SecurityRequirement{
		target: &entity.SecurityRequirement{},
	}
}

func (b *SecurityRequirement) Schemes(v map[string][]string) *SecurityRequirement {
	b.target.Schemes = v
	return b
}
