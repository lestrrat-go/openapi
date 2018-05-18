package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Schema struct {
	target *entity.Schema
}

func (b *Schema) Build() *entity.Schema {
	return b.target
}

func (b *Builder) NewSchema() *Schema {
	return &Schema{
		target: &entity.Schema{},
	}
}

func (b *Schema) Reference(v string) *Schema {
	b.target.Reference = v
	return b
}

func (b *Schema) Title(v string) *Schema {
	b.target.Title = v
	return b
}

func (b *Schema) MultipleOf(v float64) *Schema {
	b.target.MultipleOf = v
	return b
}

func (b *Schema) Maximum(v float64) *Schema {
	b.target.Maximum = v
	return b
}

func (b *Schema) ExclusiveMaximum(v float64) *Schema {
	b.target.ExclusiveMaximum = v
	return b
}

func (b *Schema) Minimum(v float64) *Schema {
	b.target.Minimum = v
	return b
}

func (b *Schema) ExclusiveMinimum(v float64) *Schema {
	b.target.ExclusiveMinimum = v
	return b
}

func (b *Schema) MaxLength(v int) *Schema {
	b.target.MaxLength = v
	return b
}

func (b *Schema) MinLength(v int) *Schema {
	b.target.MinLength = v
	return b
}

func (b *Schema) Pattern(v string) *Schema {
	b.target.Pattern = v
	return b
}

func (b *Schema) MaxItems(v int) *Schema {
	b.target.MaxItems = v
	return b
}

func (b *Schema) MinItems(v int) *Schema {
	b.target.MinItems = v
	return b
}

func (b *Schema) UniqueItems(v bool) *Schema {
	b.target.UniqueItems = v
	return b
}

func (b *Schema) MaxProperties(v int) *Schema {
	b.target.MaxProperties = v
	return b
}

func (b *Schema) MinProperties(v int) *Schema {
	b.target.MinProperties = v
	return b
}

func (b *Schema) Required(v []string) *Schema {
	b.target.Required = v
	return b
}

func (b *Schema) Enum(v []interface{}) *Schema {
	b.target.Enum = v
	return b
}

func (b *Schema) Type(v entity.PrimitiveType) *Schema {
	b.target.Type = v
	return b
}

func (b *Schema) AllOf(v []*entity.Schema) *Schema {
	b.target.AllOf = v
	return b
}

func (b *Schema) OneOf(v []*entity.Schema) *Schema {
	b.target.OneOf = v
	return b
}

func (b *Schema) AnyOf(v []*entity.Schema) *Schema {
	b.target.AnyOf = v
	return b
}

func (b *Schema) Not(v *entity.Schema) *Schema {
	b.target.Not = v
	return b
}

func (b *Schema) Items(v *entity.Schema) *Schema {
	b.target.Items = v
	return b
}

func (b *Schema) Properties(v map[string]*entity.Schema) *Schema {
	b.target.Properties = v
	return b
}

func (b *Schema) Format(v string) *Schema {
	b.target.Format = v
	return b
}

func (b *Schema) Default(v interface{}) *Schema {
	b.target.Default = v
	return b
}

func (b *Schema) Nullable(v bool) *Schema {
	b.target.Nullable = v
	return b
}

func (b *Schema) Discriminator(v *entity.Discriminator) *Schema {
	b.target.Discriminator = v
	return b
}

func (b *Schema) ReadOnly(v bool) *Schema {
	b.target.ReadOnly = v
	return b
}

func (b *Schema) WriteOnly(v bool) *Schema {
	b.target.WriteOnly = v
	return b
}

func (b *Schema) ExternalDocs(v *entity.ExternalDocumentation) *Schema {
	b.target.ExternalDocs = v
	return b
}

func (b *Schema) Example(v interface{}) *Schema {
	b.target.Example = v
	return b
}

func (b *Schema) Deprecated(v bool) *Schema {
	b.target.Deprecated = v
	return b
}
