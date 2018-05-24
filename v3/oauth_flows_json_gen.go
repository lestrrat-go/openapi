package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T13:02:46+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var _ = log.Printf
var _ = errors.Cause

type oAuthFlowsMarshalProxy struct {
	Reference         string    `json:"$ref,omitempty"`
	Implicit          OAuthFlow `json:"implicit"`
	Password          OAuthFlow `json:"password"`
	ClientCredentials OAuthFlow `json:"clientCredentials"`
	AuthorizationCode OAuthFlow `json:"authorizationCode"`
}

type oAuthFlowsUnmarshalProxy struct {
	Reference         string          `json:"$ref,omitempty"`
	Implicit          json.RawMessage `json:"implicit"`
	Password          json.RawMessage `json:"password"`
	ClientCredentials json.RawMessage `json:"clientCredentials"`
	AuthorizationCode json.RawMessage `json:"authorizationCode"`
}

func (v *oAuthFlows) MarshalJSON() ([]byte, error) {
	var proxy oAuthFlowsMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Implicit = v.implicit
	proxy.Password = v.password
	proxy.ClientCredentials = v.clientCredentials
	proxy.AuthorizationCode = v.authorizationCode
	return json.Marshal(proxy)
}

func (v *oAuthFlows) UnmarshalJSON(data []byte) error {
	var proxy oAuthFlowsUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}

	if len(proxy.Implicit) > 0 {
		var decoded oAuthFlow
		if err := json.Unmarshal(proxy.Implicit, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Implicit`)
		}

		v.implicit = &decoded
	}

	if len(proxy.Password) > 0 {
		var decoded oAuthFlow
		if err := json.Unmarshal(proxy.Password, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Password`)
		}

		v.password = &decoded
	}

	if len(proxy.ClientCredentials) > 0 {
		var decoded oAuthFlow
		if err := json.Unmarshal(proxy.ClientCredentials, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ClientCredentials`)
		}

		v.clientCredentials = &decoded
	}

	if len(proxy.AuthorizationCode) > 0 {
		var decoded oAuthFlow
		if err := json.Unmarshal(proxy.AuthorizationCode, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field AuthorizationCode`)
		}

		v.authorizationCode = &decoded
	}
	return nil
}

func (v *oAuthFlows) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*oAuthFlows)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type OAuthFlows, but got %T`, resolved)
		}
		mutator := MutateOAuthFlows(v)
		mutator.Implicit(asserted.Implicit())
		mutator.Password(asserted.Password())
		mutator.ClientCredentials(asserted.ClientCredentials())
		mutator.AuthorizationCode(asserted.AuthorizationCode())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.implicit != nil {
		if err := v.implicit.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Implicit`)
		}
	}
	if v.password != nil {
		if err := v.password.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Password`)
		}
	}
	if v.clientCredentials != nil {
		if err := v.clientCredentials.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve ClientCredentials`)
		}
	}
	if v.authorizationCode != nil {
		if err := v.authorizationCode.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve AuthorizationCode`)
		}
	}
	return nil
}

func (v *oAuthFlows) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "implicit":
		target = v.implicit
	case "password":
		target = v.password
	case "clientCredentials":
		target = v.clientCredentials
	case "authorizationCode":
		target = v.authorizationCode
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