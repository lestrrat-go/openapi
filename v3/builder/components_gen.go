package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Components struct {
	target *entity.Components
}

func (b *Components) Build() *entity.Components {
	return b.target
}

func (b *Builder) NewComponents() *Components {
	return &Components{
		target: &entity.Components{},
	}
}

func (b *Components) Schemas(v map[string]*entity.Schema) *Components {
	b.target.Schemas = v
	return b
}

func (b *Components) Responses(v map[string]*entity.Response) *Components {
	b.target.Responses = v
	return b
}

func (b *Components) Parameters(v map[string]*entity.Parameter) *Components {
	b.target.Parameters = v
	return b
}

func (b *Components) Examples(v map[string]*entity.Example) *Components {
	b.target.Examples = v
	return b
}

func (b *Components) RequestBodies(v map[string]*entity.RequestBody) *Components {
	b.target.RequestBodies = v
	return b
}

func (b *Components) Headers(v map[string]*entity.Header) *Components {
	b.target.Headers = v
	return b
}

func (b *Components) SecuritySchemes(v map[string]*entity.SecurityScheme) *Components {
	b.target.SecuritySchemes = v
	return b
}

func (b *Components) Links(v map[string]*entity.Link) *Components {
	b.target.Links = v
	return b
}

func (b *Components) Callbacks(v map[string]*entity.Callback) *Components {
	b.target.Callbacks = v
	return b
}
