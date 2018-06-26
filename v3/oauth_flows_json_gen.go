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

type oauthFlowsMarshalProxy struct {
	Reference         string    `json:"$ref,omitempty"`
	Implicit          OAuthFlow `json:"implicit"`
	Password          OAuthFlow `json:"password"`
	ClientCredentials OAuthFlow `json:"clientCredentials"`
	AuthorizationCode OAuthFlow `json:"authorizationCode"`
}

type oauthFlowsUnmarshalProxy struct {
	Reference         string          `json:"$ref,omitempty"`
	Implicit          json.RawMessage `json:"implicit"`
	Password          json.RawMessage `json:"password"`
	ClientCredentials json.RawMessage `json:"clientCredentials"`
	AuthorizationCode json.RawMessage `json:"authorizationCode"`
}

func (v *oauthFlows) MarshalJSON() ([]byte, error) {
	var proxy oauthFlowsMarshalProxy
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

func (v *oauthFlows) UnmarshalJSON(data []byte) error {
	var proxy oauthFlowsUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}

	if len(proxy.Implicit) > 0 {
		var decoded oauthFlow
		if err := json.Unmarshal(proxy.Implicit, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Implicit`)
		}

		v.implicit = &decoded
	}

	if len(proxy.Password) > 0 {
		var decoded oauthFlow
		if err := json.Unmarshal(proxy.Password, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Password`)
		}

		v.password = &decoded
	}

	if len(proxy.ClientCredentials) > 0 {
		var decoded oauthFlow
		if err := json.Unmarshal(proxy.ClientCredentials, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ClientCredentials`)
		}

		v.clientCredentials = &decoded
	}

	if len(proxy.AuthorizationCode) > 0 {
		var decoded oauthFlow
		if err := json.Unmarshal(proxy.AuthorizationCode, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field AuthorizationCode`)
		}

		v.authorizationCode = &decoded
	}
	return nil
}

func (v *oauthFlows) QueryJSON(path string) (ret interface{}, ok bool) {
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
