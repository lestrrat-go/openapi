package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T13:02:46+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *discriminator) PropertyName() string {
	return v.propertyName
}

func (v *discriminator) Mapping() *StringMapIterator {
	var items []interface{}
	for key, item := range v.mapping {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter StringMapIterator
	iter.list.items = items
	return &iter
}

func (v *discriminator) Reference() string {
	return v.reference
}

func (v *discriminator) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}