package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *header) Name() string {
	return v.name
}

func (v *header) In() Location {
	return v.in
}

func (v *header) Required() bool {
	return v.required
}

func (v *header) Description() string {
	return v.description
}

func (v *header) Deprecated() bool {
	return v.deprecated
}

func (v *header) AllowEmptyValue() bool {
	return v.allowEmptyValue
}

func (v *header) Explode() bool {
	return v.explode
}

func (v *header) AllowReserved() bool {
	return v.allowReserved
}

func (v *header) Schema() Schema {
	return v.schema
}

func (v *header) Examples() *ExampleMapIterator {
	var items []interface{}
	for key, item := range v.examples {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExampleMapIterator
	iter.list.items = items
	return &iter
}

func (v *header) Content() *MediaTypeMapIterator {
	var items []interface{}
	for key, item := range v.content {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter MediaTypeMapIterator
	iter.list.items = items
	return &iter
}

func (v *header) Reference() string {
	return v.reference
}

func (v *header) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
