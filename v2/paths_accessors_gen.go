package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sort"
)

var _ = sort.Strings

func (v *paths) Paths() *PathItemMapIterator {
	var keys []string
	for key := range v.paths {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var items []interface{}
	for _, key := range keys {
		item := v.paths[key]
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter PathItemMapIterator
	iter.list.items = items
	return &iter
}

func (v *paths) Reference() string {
	return v.reference
}

func (v *paths) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *paths) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *paths) Validate() error {
	return nil
}
