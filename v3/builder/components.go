package builder

import "github.com/lestrrat-go/openapi/v3/entity"

func (b *Components) Schema(name string, s *entity.Schema) *Components {
	if b.target.Schemas == nil {
		b.target.Schemas = make(map[string]*entity.Schema)
	}
	b.target.Schemas[name] = s
	return b
}
