package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *oAuthFlow) AuthorizationURL() string {
	return v.authorizationURL
}

func (v *oAuthFlow) TokenURL() string {
	return v.tokenURL
}

func (v *oAuthFlow) RefreshURL() string {
	return v.refreshURL
}

func (v *oAuthFlow) Scopes() *ScopeMapIterator {
	var items []interface{}
	for key, item := range v.scopes {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ScopeMapIterator
	iter.list.items = items
	return &iter
}

func (v *oAuthFlow) Reference() string {
	return v.reference
}

func (v *oAuthFlow) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
