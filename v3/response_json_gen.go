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

type responseMarshalProxy struct {
	Reference   string       `json:"$ref,omitempty"`
	Description string       `json:"description"`
	Headers     HeaderMap    `json:"headers,omitempty"`
	Content     MediaTypeMap `json:"content,omitempty"`
	Links       LinkMap      `json:"links,omitempty"`
}

type responseUnmarshalProxy struct {
	Reference   string       `json:"$ref,omitempty"`
	Description string       `json:"description"`
	Headers     HeaderMap    `json:"headers,omitempty"`
	Content     MediaTypeMap `json:"content,omitempty"`
	Links       LinkMap      `json:"links,omitempty"`
}

func (v *response) MarshalJSON() ([]byte, error) {
	var proxy responseMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Description = v.description
	proxy.Headers = v.headers
	proxy.Content = v.content
	proxy.Links = v.links
	return json.Marshal(proxy)
}

func (v *response) UnmarshalJSON(data []byte) error {
	var proxy responseUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.description = proxy.Description
	v.headers = proxy.Headers
	v.content = proxy.Content
	v.links = proxy.Links
	return nil
}

func (v *response) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*response)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Response, but got %T`, resolved)
		}
		mutator := MutateResponse(v)
		mutator.Name(asserted.Name())
		mutator.Description(asserted.Description())
		for iter := asserted.Headers(); iter.Next(); {
			key, item := iter.Item()
			mutator.Header(key, item)
		}
		for iter := asserted.Content(); iter.Next(); {
			key, item := iter.Item()
			mutator.Content(key, item)
		}
		for iter := asserted.Links(); iter.Next(); {
			key, item := iter.Item()
			mutator.Link(key, item)
		}
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.headers != nil {
		if err := v.headers.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Headers`)
		}
	}
	if v.content != nil {
		if err := v.content.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Content`)
		}
	}
	if v.links != nil {
		if err := v.links.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Links`)
		}
	}
	return nil
}

func (v *response) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "description":
		target = v.description
	case "headers":
		target = v.headers
	case "content":
		target = v.content
	case "links":
		target = v.links
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
