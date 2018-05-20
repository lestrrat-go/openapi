package openapi

// Property sets property `name` to `s`
func (b *SchemaBuilder) Property(name string, s Schema) *SchemaBuilder {
	if b.target.properties == nil {
		b.target.properties = make(map[string]Schema)
	}
	b.target.properties[name] = s.Clone()
	return b
}
