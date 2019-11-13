package openapi3

// Schema sets the schema identified by `name` to `s`
func (b *ComponentsBuilder) Schema(name string, s Schema) *ComponentsBuilder {
	if b.target.schemas == nil {
		b.target.schemas = make(map[string]Schema)
	}
	s = s.Clone()
	b.target.schemas[name] = s
	s.setName(name)
	return b
}

