package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T19:20:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *responses) DefaultValue() Response {
	return v.defaultValue
}

func (v *responses) Responses() *ResponseMapIterator {
	var items []interface{}
	for key, item := range v.responses {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ResponseMapIterator
	iter.list.items = items
	return &iter
}

func (v *responses) Reference() string {
	return v.reference
}

func (v *responses) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *responses) Validate() error {
	return nil
}
