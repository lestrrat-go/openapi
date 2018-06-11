package openapi

func (b *SchemaBuilder) Property(name string, prop Schema) *SchemaBuilder {
	if b.target.properties == nil {
		b.target.properties = SchemaMap{}
	}
	b.target.properties[name] = prop
	return b
}