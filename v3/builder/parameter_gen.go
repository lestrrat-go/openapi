package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Parameter struct {
	target *entity.Parameter
}

func (b *Parameter) Build() *entity.Parameter {
	return b.target
}

func (b *Builder) NewParameter(name string, in entity.Location) *Parameter {
	return &Parameter{
		target: &entity.Parameter{
			Required: defaultParameterRequiredFromLocation(in),
			Name:     name,
			In:       in,
		},
	}
}

func (b *Parameter) Required(v bool) *Parameter {
	b.target.Required = v
	return b
}

func (b *Parameter) Description(v string) *Parameter {
	b.target.Description = v
	return b
}

func (b *Parameter) Deprecated(v bool) *Parameter {
	b.target.Deprecated = v
	return b
}

func (b *Parameter) AllowEmptyValue(v bool) *Parameter {
	b.target.AllowEmptyValue = v
	return b
}

func (b *Parameter) Explode(v bool) *Parameter {
	b.target.Explode = v
	return b
}

func (b *Parameter) AllowReserved(v bool) *Parameter {
	b.target.AllowReserved = v
	return b
}

func (b *Parameter) Schema(v *entity.Schema) *Parameter {
	b.target.Schema = v
	return b
}

func (b *Parameter) Example(v interface{}) *Parameter {
	b.target.Example = v
	return b
}

func (b *Parameter) Examples(v map[string]*entity.Example) *Parameter {
	b.target.Examples = v
	return b
}

func (b *Parameter) Content(v map[string]*entity.MediaType) *Parameter {
	b.target.Content = v
	return b
}
