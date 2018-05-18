package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Operation struct {
	target *entity.Operation
}

func (b *Operation) Build() *entity.Operation {
	return b.target
}

func (b *Builder) NewOperation(responses *entity.Responses) *Operation {
	return &Operation{
		target: &entity.Operation{
			Responses: responses,
		},
	}
}

func (b *Operation) Tags(v []string) *Operation {
	b.target.Tags = v
	return b
}

func (b *Operation) Summary(v string) *Operation {
	b.target.Summary = v
	return b
}

func (b *Operation) Description(v string) *Operation {
	b.target.Description = v
	return b
}

func (b *Operation) ExternalDocs(v *entity.ExternalDocumentation) *Operation {
	b.target.ExternalDocs = v
	return b
}

func (b *Operation) OperationID(v string) *Operation {
	b.target.OperationID = v
	return b
}

func (b *Operation) Parameters(v []*entity.Parameter) *Operation {
	b.target.Parameters = v
	return b
}

func (b *Operation) RequestBody(v *entity.RequestBody) *Operation {
	b.target.RequestBody = v
	return b
}

func (b *Operation) Callbacks(v map[string]*entity.Callback) *Operation {
	b.target.Callbacks = v
	return b
}

func (b *Operation) Deprecated(v bool) *Operation {
	b.target.Deprecated = v
	return b
}

func (b *Operation) Security(v []*entity.SecurityRequirement) *Operation {
	b.target.Security = v
	return b
}

func (b *Operation) Servers(v []*entity.Server) *Operation {
	b.target.Servers = v
	return b
}
