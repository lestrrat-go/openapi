package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:57:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type componentsMarshalProxy struct {
	Schemas         map[string]Schema         `json:"schemas,omitempty"`
	Responses       map[string]Response       `json:"responses,omitempty"`
	Parameters      map[string]Parameter      `json:"parameters,omitempty"`
	Examples        map[string]Example        `json:"examples,omitempty"`
	RequestBodies   map[string]RequestBody    `json:"requestBodies,omitempty"`
	Headers         map[string]Header         `json:"headers,omitempty"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty"`
	Links           map[string]Link           `json:"links,omitempty"`
	Callbacks       map[string]Callback       `json:"callbacks,omitempty"`
}

type componentsUnmarshalProxy struct {
	Schemas         map[string]json.RawMessage `json:"schemas,omitempty"`
	Responses       map[string]json.RawMessage `json:"responses,omitempty"`
	Parameters      map[string]json.RawMessage `json:"parameters,omitempty"`
	Examples        map[string]json.RawMessage `json:"examples,omitempty"`
	RequestBodies   map[string]json.RawMessage `json:"requestBodies,omitempty"`
	Headers         map[string]json.RawMessage `json:"headers,omitempty"`
	SecuritySchemes map[string]json.RawMessage `json:"securitySchemes,omitempty"`
	Links           map[string]json.RawMessage `json:"links,omitempty"`
	Callbacks       map[string]json.RawMessage `json:"callbacks,omitempty"`
}

func (v *components) MarshalJSON() ([]byte, error) {
	var proxy componentsMarshalProxy
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

	if len(proxy.Schemas) > 0 {
		m := make(map[string]Schema)
		for key, pv := range proxy.Schemas {
			var decoded schema
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Schemas`, key)
			}
			m[key] = &decoded
		}
		v.schemas = m
	}

	if len(proxy.Responses) > 0 {
		m := make(map[string]Response)
		for key, pv := range proxy.Responses {
			var decoded response
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Responses`, key)
			}
			m[key] = &decoded
		}
		v.responses = m
	}

	if len(proxy.Parameters) > 0 {
		m := make(map[string]Parameter)
		for key, pv := range proxy.Parameters {
			var decoded parameter
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Parameters`, key)
			}
			m[key] = &decoded
		}
		v.parameters = m
	}

	if len(proxy.Examples) > 0 {
		m := make(map[string]Example)
		for key, pv := range proxy.Examples {
			var decoded example
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Examples`, key)
			}
			m[key] = &decoded
		}
		v.examples = m
	}

	if len(proxy.RequestBodies) > 0 {
		m := make(map[string]RequestBody)
		for key, pv := range proxy.RequestBodies {
			var decoded requestBody
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field RequestBodies`, key)
			}
			m[key] = &decoded
		}
		v.requestBodies = m
	}

	if len(proxy.Headers) > 0 {
		m := make(map[string]Header)
		for key, pv := range proxy.Headers {
			var decoded header
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Headers`, key)
			}
			m[key] = &decoded
		}
		v.headers = m
	}

	if len(proxy.SecuritySchemes) > 0 {
		m := make(map[string]SecurityScheme)
		for key, pv := range proxy.SecuritySchemes {
			var decoded securityScheme
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field SecuritySchemes`, key)
			}
			m[key] = &decoded
		}
		v.securitySchemes = m
	}

	if len(proxy.Links) > 0 {
		m := make(map[string]Link)
		for key, pv := range proxy.Links {
			var decoded link
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Links`, key)
			}
			m[key] = &decoded
		}
		v.links = m
	}

	if len(proxy.Callbacks) > 0 {
		m := make(map[string]Callback)
		for key, pv := range proxy.Callbacks {
			var decoded callback
			if err := json.Unmarshal(pv, &decoded); err != nil {
				return errors.Wrapf(err, `failed to unmasrhal element %s of field Callbacks`, key)
			}
			m[key] = &decoded
		}
		v.callbacks = m
	}
	return nil
}
