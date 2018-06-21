package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *oauthFlow) AuthorizationURL() string {
	return v.authorizationURL
}

func (v *oauthFlow) TokenURL() string {
	return v.tokenURL
}

func (v *oauthFlow) RefreshURL() string {
	return v.refreshURL
}

func (v *oauthFlow) Scopes() *ScopeMapIterator {
	var items []interface{}
	for key, item := range v.scopes {
		items = append(items, &mapIteratorItem{key: key, item: item})
	}
	var iter ScopeMapIterator
	iter.list.items = items
	return &iter
}

func (v *oauthFlow) Reference() string {
	return v.reference
}

func (v *oauthFlow) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
