package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type License struct {
	target *entity.License
}

func (b *License) Build() *entity.License {
	return b.target
}

func (b *Builder) NewLicense(name string) *License {
	return &License{
		target: &entity.License{
			Name: name,
		},
	}
}

func (b *License) URL(v string) *License {
	b.target.URL = v
	return b
}
