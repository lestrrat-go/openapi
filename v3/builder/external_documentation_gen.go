package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type ExternalDocumentation struct {
	target *entity.ExternalDocumentation
}

func (b *ExternalDocumentation) Build() *entity.ExternalDocumentation {
	return b.target
}

func (b *Builder) NewExternalDocumentation() *ExternalDocumentation {
	return &ExternalDocumentation{
		target: &entity.ExternalDocumentation{},
	}
}

func (b *ExternalDocumentation) Description(v string) *ExternalDocumentation {
	b.target.Description = v
	return b
}

func (b *ExternalDocumentation) URL(v string) *ExternalDocumentation {
	b.target.URL = v
	return b
}
