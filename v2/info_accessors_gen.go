package openapi

// This file was automatically generated by gentyeps.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sort"
)

var _ = sort.Strings

func (v *info) Title() string {
	return v.title
}

func (v *info) Version() string {
	return v.version
}

func (v *info) Description() string {
	return v.description
}

func (v *info) TermsOfService() string {
	return v.termsOfService
}

func (v *info) Contact() Contact {
	return v.contact
}

func (v *info) License() License {
	return v.license
}

func (v *info) Reference() string {
	return v.reference
}

func (v *info) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}

func (v *info) Extensions() *ExtensionsIterator {
	var items []interface{}
	for key, item := range v.extensions {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ExtensionsIterator
	iter.list.items = items
	return &iter
}

func (v *info) Validate() error {
	return nil
}
