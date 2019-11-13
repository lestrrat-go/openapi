package openapi3

// Property sets property `name` to `s`
func (b *SchemaBuilder) Property(name string, s Schema) *SchemaBuilder {
	if b.target.properties == nil {
		b.target.properties = make(map[string]Schema)
	}
	s = s.Clone()
	b.target.properties[name] = s
	s.setName(name)
	return b
}
