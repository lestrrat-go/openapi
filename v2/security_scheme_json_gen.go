package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

var _ = json.Unmarshal
var _ = fmt.Fprintf
var _ = log.Printf
var _ = strconv.ParseInt
var _ = errors.Cause

type securitySchemeMarshalProxy struct {
	Reference        string    `json:"$ref,omitempty"`
	Type             string    `json:"type"`
	Description      string    `json:"description,omitempty"`
	Name             string    `json:"name,omitempty"`
	In               string    `json:"in,omitempty"`
	Flow             string    `json:"flow,omitempty"`
	AuthorizationURL string    `json:"authorizationUrl,omitempty"`
	TokenURL         string    `json:"tokenUrl,omitempty"`
	Scopes           StringMap `json:"scopes,omitempty"`
}

func (v *securityScheme) MarshalJSON() ([]byte, error) {
	var proxy securitySchemeMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Type = v.typ
	proxy.Description = v.description
	proxy.Name = v.name
	proxy.In = v.in
	proxy.Flow = v.flow
	proxy.AuthorizationURL = v.authorizationURL
	proxy.TokenURL = v.tokenURL
	proxy.Scopes = v.scopes
	buf, err := json.Marshal(proxy)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal struct`)
	}
	if len(v.extensions) > 0 {
		extBuf, err := json.Marshal(v.extensions)
		if err != nil || len(extBuf) <= 2 {
			return nil, errors.Wrap(err, `failed to marshal struct (extensions)`)
		}
		buf = append(append(buf[:len(buf)-1], ','), extBuf[1:]...)
	}
	return buf, nil
}

// UnmarshalJSON defines how securityScheme is deserialized from JSON
func (v *securityScheme) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if raw, ok := proxy["$ref"]; ok {
		if err := json.Unmarshal(raw, &v.reference); err != nil {
			return errors.Wrap(err, `failed to unmarshal $ref`)
		}
		return nil
	}

	mutator := MutateSecurityScheme(v)

	const typMapKey = "type"
	if raw, ok := proxy[typMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field type`)
		}
		mutator.Type(decoded)
		delete(proxy, typMapKey)
	}

	const descriptionMapKey = "description"
	if raw, ok := proxy[descriptionMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field description`)
		}
		mutator.Description(decoded)
		delete(proxy, descriptionMapKey)
	}

	const nameMapKey = "name"
	if raw, ok := proxy[nameMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field name`)
		}
		mutator.Name(decoded)
		delete(proxy, nameMapKey)
	}

	const inMapKey = "in"
	if raw, ok := proxy[inMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field in`)
		}
		mutator.In(decoded)
		delete(proxy, inMapKey)
	}

	const flowMapKey = "flow"
	if raw, ok := proxy[flowMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field flow`)
		}
		mutator.Flow(decoded)
		delete(proxy, flowMapKey)
	}

	const authorizationURLMapKey = "authorizationUrl"
	if raw, ok := proxy[authorizationURLMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field authorizationUrl`)
		}
		mutator.AuthorizationURL(decoded)
		delete(proxy, authorizationURLMapKey)
	}

	const tokenURLMapKey = "tokenUrl"
	if raw, ok := proxy[tokenURLMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field tokenUrl`)
		}
		mutator.TokenURL(decoded)
		delete(proxy, tokenURLMapKey)
	}

	const scopesMapKey = "scopes"
	if raw, ok := proxy[scopesMapKey]; ok {
		var decoded StringMap
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Scopes`)
		}
		for key, elem := range decoded {
			mutator.Scope(key, elem)
		}
		delete(proxy, scopesMapKey)
	}

	for name, raw := range proxy {
		if strings.HasPrefix(name, `x-`) {
			var ext interface{}
			if err := json.Unmarshal(raw, &ext); err != nil {
				return errors.Wrapf(err, `failed to unmarshal field %s`, name)
			}
			mutator.Extension(name, ext)
		}
	}

	if err := mutator.Apply(); err != nil {
		return errors.Wrap(err, `failed to  unmarshal JSON`)
	}
	return nil
}

func (v *securityScheme) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "type":
		target = v.typ
	case "description":
		target = v.description
	case "name":
		target = v.name
	case "in":
		target = v.in
	case "flow":
		target = v.flow
	case "authorizationUrl":
		target = v.authorizationURL
	case "tokenUrl":
		target = v.tokenURL
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

// SecuritySchemeFromJSON constructs a SecurityScheme from JSON buffer. `dst` must
// be a pointer to `SecurityScheme`
func SecuritySchemeFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*SecurityScheme)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to SecurityScheme, but got %T`, dst)
	}
	var tmp securityScheme
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal SecurityScheme`)
	}
	*v = &tmp
	return nil
}
