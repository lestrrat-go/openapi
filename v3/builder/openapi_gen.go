package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type OpenAPI struct {
	target *entity.OpenAPI
}

func (b *OpenAPI) Build() *entity.OpenAPI {
	return b.target
}

func (b *Builder) NewOpenAPI(info *entity.Info, paths *entity.Paths) *OpenAPI {
	return &OpenAPI{
		target: &entity.OpenAPI{
			OpenAPI: DefaultVersion,
			Info:    info,
			Paths:   paths,
		},
	}
}

func (b *OpenAPI) OpenAPI(v string) *OpenAPI {
	b.target.OpenAPI = v
	return b
}

func (b *OpenAPI) Servers(v []entity.Server) *OpenAPI {
	b.target.Servers = v
	return b
}

func (b *OpenAPI) Components(v *entity.Components) *OpenAPI {
	b.target.Components = v
	return b
}

func (b *OpenAPI) Security(v *entity.SecurityRequirement) *OpenAPI {
	b.target.Security = v
	return b
}

func (b *OpenAPI) Tags(v []*entity.Tag) *OpenAPI {
	b.target.Tags = v
	return b
}

func (b *OpenAPI) ExternalDocs(v *entity.ExternalDocumentation) *OpenAPI {
	b.target.ExternalDocs = v
	return b
}
