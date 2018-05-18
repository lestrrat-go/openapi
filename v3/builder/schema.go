package builder

import "github.com/lestrrat-go/openapi/v3/entity"

func (b *Schema) Property(name string, s *entity.Schema) *Schema {
	if b.target.Properties == nil {
		b.target.Properties = make(map[string]*entity.Schema)
	}
	b.target.Properties[name] = s
	return b
}
