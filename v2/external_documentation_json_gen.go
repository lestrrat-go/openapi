package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T19:20:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var _ = log.Printf
var _ = json.Unmarshal
var _ = errors.Cause

type externalDocumentationMarshalProxy struct {
	Reference   string `json:"$ref,omitempty"`
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

type externalDocumentationUnmarshalProxy struct {
	Reference   string `json:"$ref,omitempty"`
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

func (v *externalDocumentation) MarshalJSON() ([]byte, error) {
	var proxy externalDocumentationMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.URL = v.uRL
	proxy.Description = v.description
	return json.Marshal(proxy)
}

func (v *externalDocumentation) UnmarshalJSON(data []byte) error {
	var proxy externalDocumentationUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.uRL = proxy.URL
	v.description = proxy.Description
	return nil
}

func (v *externalDocumentation) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*externalDocumentation)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type ExternalDocumentation, but got %T`, resolved)
		}
		mutator := MutateExternalDocumentation(v)
		mutator.URL(asserted.URL())
		mutator.Description(asserted.Description())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	return nil
}

func (v *externalDocumentation) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "url":
		target = v.uRL
	case "description":
		target = v.description
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
