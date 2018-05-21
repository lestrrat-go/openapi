package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
)

var _ = errors.Cause

type oAuthFlowMarshalProxy struct {
	AuthorizationURL string            `json:"authorizationUrl"`
	TokenURL         string            `json:"tokenUrl"`
	RefreshURL       string            `json:"refreshUrl"`
	Scopes           map[string]string `json:"scopes"`
}

type oAuthFlowUnmarshalProxy struct {
	AuthorizationURL string `json:"authorizationUrl"`
	TokenURL         string `json:"tokenUrl"`
	RefreshURL       string `json:"refreshUrl"`
}

func (v *oAuthFlow) MarshalJSON() ([]byte, error) {
	var proxy oAuthFlowMarshalProxy
	proxy.AuthorizationURL = v.authorizationURL
	proxy.TokenURL = v.tokenURL
	proxy.RefreshURL = v.refreshURL
	proxy.Scopes = v.scopes
	return json.Marshal(proxy)
}

func (v *oAuthFlow) UnmarshalJSON(data []byte) error {
	var proxy oAuthFlowUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.authorizationURL = proxy.AuthorizationURL
	v.tokenURL = proxy.TokenURL
	v.refreshURL = proxy.RefreshURL
	return nil
}

func (v *oAuthFlow) QueryJSON(path string) (ret interface{}, ok bool) {
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
