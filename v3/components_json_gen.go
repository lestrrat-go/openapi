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

type componentsMarshalProxy struct {
	Reference       string            `json:"$ref,omitempty"`
	Schemas         SchemaMap         `json:"schemas,omitempty"`
	Responses       ResponseMap       `json:"responses,omitempty"`
	Parameters      ParameterMap      `json:"parameters,omitempty"`
	Examples        ExampleMap        `json:"examples,omitempty"`
	RequestBodies   RequestBodyMap    `json:"requestBodies,omitempty"`
	Headers         HeaderMap         `json:"headers,omitempty"`
	SecuritySchemes SecuritySchemeMap `json:"securitySchemes,omitempty"`
	Links           LinkMap           `json:"links,omitempty"`
	Callbacks       CallbackMap       `json:"callbacks,omitempty"`
}

type componentsUnmarshalProxy struct {
	Reference       string            `json:"$ref,omitempty"`
	Schemas         SchemaMap         `json:"schemas,omitempty"`
	Responses       ResponseMap       `json:"responses,omitempty"`
	Parameters      ParameterMap      `json:"parameters,omitempty"`
	Examples        ExampleMap        `json:"examples,omitempty"`
	RequestBodies   RequestBodyMap    `json:"requestBodies,omitempty"`
	Headers         HeaderMap         `json:"headers,omitempty"`
	SecuritySchemes SecuritySchemeMap `json:"securitySchemes,omitempty"`
	Links           LinkMap           `json:"links,omitempty"`
	Callbacks       CallbackMap       `json:"callbacks,omitempty"`
}

func (v *components) MarshalJSON() ([]byte, error) {
	var proxy componentsMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Schemas = v.schemas
	proxy.Responses = v.responses
	proxy.Parameters = v.parameters
	proxy.Examples = v.examples
	proxy.RequestBodies = v.requestBodies
	proxy.Headers = v.headers
	proxy.SecuritySchemes = v.securitySchemes
	proxy.Links = v.links
	proxy.Callbacks = v.callbacks
	return json.Marshal(proxy)
}

func (v *components) UnmarshalJSON(data []byte) error {
	var proxy componentsUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.schemas = proxy.Schemas
	v.responses = proxy.Responses
	v.parameters = proxy.Parameters
	v.examples = proxy.Examples
	v.requestBodies = proxy.RequestBodies
	v.headers = proxy.Headers
	v.securitySchemes = proxy.SecuritySchemes
	v.links = proxy.Links
	v.callbacks = proxy.Callbacks
	return nil
}

func (v *components) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*components)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Components, but got %T`, resolved)
		}
		mutator := MutateComponents(v)
		for iter := asserted.Schemas(); iter.Next(); {
			key, item := iter.Item()
			mutator.Schema(key, item)
		}
		for iter := asserted.Responses(); iter.Next(); {
			key, item := iter.Item()
			mutator.Response(key, item)
		}
		for iter := asserted.Parameters(); iter.Next(); {
			key, item := iter.Item()
			mutator.Parameter(key, item)
		}
		for iter := asserted.Examples(); iter.Next(); {
			key, item := iter.Item()
			mutator.Example(key, item)
		}
		for iter := asserted.RequestBodies(); iter.Next(); {
			key, item := iter.Item()
			mutator.RequestBody(key, item)
		}
		for iter := asserted.Headers(); iter.Next(); {
			key, item := iter.Item()
			mutator.Header(key, item)
		}
		for iter := asserted.SecuritySchemes(); iter.Next(); {
			key, item := iter.Item()
			mutator.SecurityScheme(key, item)
		}
		for iter := asserted.Links(); iter.Next(); {
			key, item := iter.Item()
			mutator.Link(key, item)
		}
		for iter := asserted.Callbacks(); iter.Next(); {
			key, item := iter.Item()
			mutator.Callback(key, item)
		}
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.schemas != nil {
		if err := v.schemas.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Schemas`)
		}
	}
	if v.responses != nil {
		if err := v.responses.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Responses`)
		}
	}
	if v.parameters != nil {
		if err := v.parameters.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Parameters`)
		}
	}
	if v.examples != nil {
		if err := v.examples.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Examples`)
		}
	}
	if v.requestBodies != nil {
		if err := v.requestBodies.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve RequestBodies`)
		}
	}
	if v.headers != nil {
		if err := v.headers.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Headers`)
		}
	}
	if v.securitySchemes != nil {
		if err := v.securitySchemes.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve SecuritySchemes`)
		}
	}
	if v.links != nil {
		if err := v.links.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Links`)
		}
	}
	if v.callbacks != nil {
		if err := v.callbacks.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Callbacks`)
		}
	}
	return nil
}

func (v *components) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "schemas":
		target = v.schemas
	case "responses":
		target = v.responses
	case "parameters":
		target = v.parameters
	case "examples":
		target = v.examples
	case "requestBodies":
		target = v.requestBodies
	case "headers":
		target = v.headers
	case "securitySchemes":
		target = v.securitySchemes
	case "links":
		target = v.links
	case "callbacks":
		target = v.callbacks
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
