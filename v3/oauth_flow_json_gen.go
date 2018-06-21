package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/pkg/errors"
)

var _ = log.Printf
var _ = json.Unmarshal
var _ = errors.Cause

type oauthFlowMarshalProxy struct {
	Reference        string   `json:"$ref,omitempty"`
	AuthorizationURL string   `json:"authorizationUrl"`
	TokenURL         string   `json:"tokenUrl"`
	RefreshURL       string   `json:"refreshUrl"`
	Scopes           ScopeMap `json:"scopes"`
}

type oauthFlowUnmarshalProxy struct {
	Reference        string   `json:"$ref,omitempty"`
	AuthorizationURL string   `json:"authorizationUrl"`
	TokenURL         string   `json:"tokenUrl"`
	RefreshURL       string   `json:"refreshUrl"`
	Scopes           ScopeMap `json:"scopes"`
}

func (v *oauthFlow) MarshalJSON() ([]byte, error) {
	var proxy oauthFlowMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.AuthorizationURL = v.authorizationURL
	proxy.TokenURL = v.tokenURL
	proxy.RefreshURL = v.refreshURL
	proxy.Scopes = v.scopes
	return json.Marshal(proxy)
}

func (v *oauthFlow) UnmarshalJSON(data []byte) error {
	var proxy oauthFlowUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.authorizationURL = proxy.AuthorizationURL
	v.tokenURL = proxy.TokenURL
	v.refreshURL = proxy.RefreshURL
	v.scopes = proxy.Scopes
	return nil
}

func (v *oauthFlow) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*oauthFlow)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type OAuthFlow, but got %T`, resolved)
		}
		mutator := MutateOAuthFlow(v)
		mutator.AuthorizationURL(asserted.AuthorizationURL())
		mutator.TokenURL(asserted.TokenURL())
		mutator.RefreshURL(asserted.RefreshURL())
		for iter := asserted.Scopes(); iter.Next(); {
			key, item := iter.Item()
			mutator.Scope(key, item)
		}
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.scopes != nil {
		if err := v.scopes.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Scopes`)
		}
	}
	return nil
}

func (v *oauthFlow) QueryJSON(path string) (ret interface{}, ok bool) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	if path == "" {
		return v, true
	}

	var frag string
	if i := strings.Index(path, "/"); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ""
	}

	var target interface{}

	switch frag {
	case "authorizationUrl":
		target = v.authorizationURL
	case "tokenUrl":
		target = v.tokenURL
	case "refreshUrl":
		target = v.refreshURL
	case "scopes":
		target = v.scopes
	default:
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}
	if path == "" {
		return target, true
	}
	return nil, false
}
