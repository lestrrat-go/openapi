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

type openAPIMarshalProxy struct {
	Reference    string                `json:"$ref,omitempty"`
	Version      string                `json:"openapi"`
	Info         Info                  `json:"info"`
	Servers      []Server              `json:"servers,omitempty"`
	Paths        Paths                 `json:"paths"`
	Components   Components            `json:"components,omitempty"`
	Security     SecurityRequirement   `json:"security,omitempty"`
	Tags         []Tag                 `json:"tags,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type openAPIUnmarshalProxy struct {
	Reference    string            `json:"$ref,omitempty"`
	Version      string            `json:"openapi"`
	Info         json.RawMessage   `json:"info"`
	Servers      []json.RawMessage `json:"servers,omitempty"`
	Paths        json.RawMessage   `json:"paths"`
	Components   json.RawMessage   `json:"components,omitempty"`
	Security     json.RawMessage   `json:"security,omitempty"`
	Tags         []json.RawMessage `json:"tags,omitempty"`
	ExternalDocs json.RawMessage   `json:"externalDocs,omitempty"`
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
		return err
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

	if len(proxy.Servers) > 0 {
		var list []Server
		for i, pv := range proxy.Servers {
			var decoded server
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %d of field Servers`, i)
			}
			list = append(list, &decoded)
		}
		v.servers = list
	}

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

	if len(proxy.Tags) > 0 {
		var list []Tag
		for i, pv := range proxy.Tags {
			var decoded tag
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %d of field Tags`, i)
			}
			list = append(list, &decoded)
		}
		v.tags = list
	}

	if len(proxy.ExternalDocs) > 0 {
		var decoded externalDocumentation
		if err := json.Unmarshal(proxy.ExternalDocs, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		v.externalDocs = &decoded
	}
	return nil
}

func (v *openAPI) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*openAPI)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type OpenAPI, but got %T`, resolved)
		}
		mutator := MutateOpenAPI(v)
		mutator.Version(asserted.Version())
		mutator.Info(asserted.Info())
		for iter := asserted.Servers(); iter.Next(); {
			item := iter.Item()
			mutator.Server(item)
		}
		mutator.Paths(asserted.Paths())
		mutator.Components(asserted.Components())
		mutator.Security(asserted.Security())
		for iter := asserted.Tags(); iter.Next(); {
			item := iter.Item()
			mutator.Tag(item)
		}
		mutator.ExternalDocs(asserted.ExternalDocs())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.info != nil {
		if err := v.info.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Info`)
		}
	}
	if v.servers != nil {
		if err := v.servers.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Servers`)
		}
	}
	if v.paths != nil {
		if err := v.paths.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Paths`)
		}
	}
	if v.components != nil {
		if err := v.components.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Components`)
		}
	}
	if v.security != nil {
		if err := v.security.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Security`)
		}
	}
	if v.tags != nil {
		if err := v.tags.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Tags`)
		}
	}
	if v.externalDocs != nil {
		if err := v.externalDocs.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve ExternalDocs`)
		}
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