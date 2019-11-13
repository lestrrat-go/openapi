package openapi3

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

type openAPIMarshalProxy struct {
	Reference    string                `json:"$ref,omitempty"`
	Version      string                `json:"openapi"`
	Info         Info                  `json:"info"`
	Servers      ServerList            `json:"servers,omitempty"`
	Paths        Paths                 `json:"paths"`
	Components   Components            `json:"components,omitempty"`
	Security     SecurityRequirement   `json:"security,omitempty"`
	Tags         TagList               `json:"tags,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type openAPIUnmarshalProxy struct {
	Reference    string          `json:"$ref,omitempty"`
	Version      string          `json:"openapi"`
	Info         json.RawMessage `json:"info"`
	Servers      ServerList      `json:"servers,omitempty"`
	Paths        json.RawMessage `json:"paths"`
	Components   json.RawMessage `json:"components,omitempty"`
	Security     json.RawMessage `json:"security,omitempty"`
	Tags         TagList         `json:"tags,omitempty"`
	ExternalDocs json.RawMessage `json:"externalDocs,omitempty"`
}

func (v *openAPI) MarshalJSON() ([]byte, error) {
	var proxy openAPIMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Version = v.version
	proxy.Info = v.info
	proxy.Servers = v.servers
	proxy.Paths = v.paths
	proxy.Components = v.components
	proxy.Security = v.security
	proxy.Tags = v.tags
	proxy.ExternalDocs = v.externalDocs
	return json.Marshal(proxy)
}

func (v *openAPI) UnmarshalJSON(data []byte) error {
	var proxy openAPIUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrapf(err, `failed to unmarshal openAPI`)
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.version = proxy.Version

	if len(proxy.Info) > 0 {
		var decoded info
		if err := json.Unmarshal(proxy.Info, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Info`)
		}

		v.info = &decoded
	}
	v.servers = proxy.Servers

	if len(proxy.Paths) > 0 {
		var decoded paths
		if err := json.Unmarshal(proxy.Paths, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Paths`)
		}

		v.paths = &decoded
	}

	if len(proxy.Components) > 0 {
		var decoded components
		if err := json.Unmarshal(proxy.Components, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Components`)
		}

		v.components = &decoded
	}

	if len(proxy.Security) > 0 {
		var decoded securityRequirement
		if err := json.Unmarshal(proxy.Security, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Security`)
		}

		v.security = &decoded
	}
	v.tags = proxy.Tags

	if len(proxy.ExternalDocs) > 0 {
		var decoded externalDocumentation
		if err := json.Unmarshal(proxy.ExternalDocs, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		v.externalDocs = &decoded
	}
	return nil
}

func (v *openAPI) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "openapi":
		target = v.version
	case "info":
		target = v.info
	case "servers":
		target = v.servers
	case "paths":
		target = v.paths
	case "components":
		target = v.components
	case "security":
		target = v.security
	case "tags":
		target = v.tags
	case "externalDocs":
		target = v.externalDocs
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

// OpenAPIFromJSON constructs a OpenAPI from JSON buffer. `dst` must
// be a pointer to `OpenAPI`
func OpenAPIFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*OpenAPI)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to OpenAPI, but got %T`, dst)
	}
	var tmp openAPI
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal OpenAPI`)
	}
	*v = &tmp
	return nil
}
