package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T12:53:29+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *serverVariable) Enum() *StringListIterator {
	var items []interface{}
	for _, item := range v.enum {
		items = append(items, item)
	}
	var iter StringListIterator
	iter.items = items
	return &iter
}

func (v *serverVariable) Default() string {
	return v.defaultValue
}

func (v *serverVariable) Description() string {
	return v.description
}

func (v *serverVariable) Reference() string {
	return v.reference
}

func (v *serverVariable) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
