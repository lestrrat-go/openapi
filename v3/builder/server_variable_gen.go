package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type ServerVariable struct {
	target *entity.ServerVariable
}

func (b *ServerVariable) Build() *entity.ServerVariable {
	return b.target
}

func (b *Builder) NewServerVariable(defaultValue string) *ServerVariable {
	return &ServerVariable{
		target: &entity.ServerVariable{
			Default: defaultValue,
		},
	}
}

func (b *ServerVariable) Enum(v []string) *ServerVariable {
	b.target.Enum = v
	return b
}

func (b *ServerVariable) Description(v string) *ServerVariable {
	b.target.Description = v
	return b
}
