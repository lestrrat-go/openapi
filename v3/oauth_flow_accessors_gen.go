package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

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

func (v *oauthFlow) Validate(recurse bool) error {
	if recurse {
		return v.recurseValidate()
	}
	return nil
}

func (v *oauthFlow) recurseValidate() error {
	if elem := v.scopes; elem != nil {
		if err := elem.Validate(true); err != nil {
			return errors.Wrap(err, `failed to validate field "scopes"`)
		}
	}
	return nil
}
