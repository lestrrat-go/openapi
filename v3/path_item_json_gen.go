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

type pathItemMarshalProxy struct {
	Reference   string        `json:"$ref,omitempty"`
	Summary     string        `json:"summary,omitempty"`
	Description string        `json:"description,omitempty"`
	Get         Operation     `json:"get,omitempty"`
	Put         Operation     `json:"put,omitempty"`
	Post        Operation     `json:"post,omitempty"`
	Delete      Operation     `json:"delete,omitempty"`
	Options     Operation     `json:"options,omitempty"`
	Head        Operation     `json:"head,omitempty"`
	Patch       Operation     `json:"patch,omitempty"`
	Trace       Operation     `json:"trace,omitempty"`
	Servers     ServerList    `json:"servers,omitempty"`
	Parameters  ParameterList `json:"parameters,omitempty"`
}

type pathItemUnmarshalProxy struct {
	Reference   string          `json:"$ref,omitempty"`
	Summary     string          `json:"summary,omitempty"`
	Description string          `json:"description,omitempty"`
	Get         json.RawMessage `json:"get,omitempty"`
	Put         json.RawMessage `json:"put,omitempty"`
	Post        json.RawMessage `json:"post,omitempty"`
	Delete      json.RawMessage `json:"delete,omitempty"`
	Options     json.RawMessage `json:"options,omitempty"`
	Head        json.RawMessage `json:"head,omitempty"`
	Patch       json.RawMessage `json:"patch,omitempty"`
	Trace       json.RawMessage `json:"trace,omitempty"`
	Servers     ServerList      `json:"servers,omitempty"`
	Parameters  ParameterList   `json:"parameters,omitempty"`
}

func (v *pathItem) MarshalJSON() ([]byte, error) {
	var proxy pathItemMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Summary = v.summary
	proxy.Description = v.description
	proxy.Get = v.get
	proxy.Put = v.put
	proxy.Post = v.post
	proxy.Delete = v.delete
	proxy.Options = v.options
	proxy.Head = v.head
	proxy.Patch = v.patch
	proxy.Trace = v.trace
	proxy.Servers = v.servers
	proxy.Parameters = v.parameters
	return json.Marshal(proxy)
}

func (v *pathItem) UnmarshalJSON(data []byte) error {
	var proxy pathItemUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.summary = proxy.Summary
	v.description = proxy.Description

	if len(proxy.Get) > 0 {
		var decoded operation
		if err := json.Unmarshal(proxy.Get, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Get`)
		}

		v.get = &decoded
	}

	if len(proxy.Put) > 0 {
		var decoded operation
		if err := json.Unmarshal(proxy.Put, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Put`)
		}

		v.put = &decoded
	}

	if len(proxy.Post) > 0 {
		var decoded operation
		if err := json.Unmarshal(proxy.Post, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Post`)
		}

		v.post = &decoded
	}

	if len(proxy.Delete) > 0 {
		var decoded operation
		if err := json.Unmarshal(proxy.Delete, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Delete`)
		}

		v.delete = &decoded
	}

	if len(proxy.Options) > 0 {
		var decoded operation
		if err := json.Unmarshal(proxy.Options, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Options`)
		}

		v.options = &decoded
	}

	if len(proxy.Head) > 0 {
		var decoded operation
		if err := json.Unmarshal(proxy.Head, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Head`)
		}

		v.head = &decoded
	}

	if len(proxy.Patch) > 0 {
		var decoded operation
		if err := json.Unmarshal(proxy.Patch, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Patch`)
		}

		v.patch = &decoded
	}

	if len(proxy.Trace) > 0 {
		var decoded operation
		if err := json.Unmarshal(proxy.Trace, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Trace`)
		}

		v.trace = &decoded
	}
	v.servers = proxy.Servers
	v.parameters = proxy.Parameters

	v.postUnmarshalJSON()
	return nil
}

func (v *pathItem) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "summary":
		target = v.summary
	case "description":
		target = v.description
	case "get":
		target = v.get
	case "put":
		target = v.put
	case "post":
		target = v.post
	case "delete":
		target = v.delete
	case "options":
		target = v.options
	case "head":
		target = v.head
	case "patch":
		target = v.patch
	case "trace":
		target = v.trace
	case "servers":
		target = v.servers
	case "parameters":
		target = v.parameters
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

// PathItemFromJSON constructs a PathItem from JSON buffer. `dst` must
// be a pointer to `PathItem`
func PathItemFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*PathItem)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to PathItem, but got %T`, dst)
	}
	var tmp pathItem
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal PathItem`)
	}
	*v = &tmp
	return nil
}
