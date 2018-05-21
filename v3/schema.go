package openapi

func (v *schema) setName(s string) {
	v.name = s
}

func (v *schema) Properties() *SchemaListIterator {
	var items []interface{}
	for _, s := range v.properties {
		items = append(items, s)
	}

	var iter SchemaListIterator
	iter.items = items
	return &iter
}
