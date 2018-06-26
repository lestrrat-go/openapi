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

type operationMarshalProxy struct {
	Reference    string                  `json:"$ref,omitempty"`
	Tags         StringList              `json:"tags,omitempty"`
	Summary      string                  `json:"summary,omitempty"`
	Description  string                  `json:"description,omitempty"`
	ExternalDocs ExternalDocumentation   `json:"externalDocs,omitempty"`
	OperationID  string                  `json:"operationId,omitempty"`
	Parameters   ParameterList           `json:"parameters,omitempty"`
	RequestBody  RequestBody             `json:"requestBody,omitempty"`
	Responses    Responses               `json:"responses"`
	Callbacks    CallbackMap             `json:"callbacks,omitempty"`
	Deprecated   bool                    `json:"deprecated,omitempty"`
	Security     SecurityRequirementList `json:"security,omitempty"`
	Servers      ServerList              `json:"servers,omitempty"`
}

type operationUnmarshalProxy struct {
	Reference    string                  `json:"$ref,omitempty"`
	Tags         StringList              `json:"tags,omitempty"`
	Summary      string                  `json:"summary,omitempty"`
	Description  string                  `json:"description,omitempty"`
	ExternalDocs json.RawMessage         `json:"externalDocs,omitempty"`
	OperationID  string                  `json:"operationId,omitempty"`
	Parameters   ParameterList           `json:"parameters,omitempty"`
	RequestBody  json.RawMessage         `json:"requestBody,omitempty"`
	Responses    json.RawMessage         `json:"responses"`
	Callbacks    CallbackMap             `json:"callbacks,omitempty"`
	Deprecated   bool                    `json:"deprecated,omitempty"`
	Security     SecurityRequirementList `json:"security,omitempty"`
	Servers      ServerList              `json:"servers,omitempty"`
}

func (v *operation) MarshalJSON() ([]byte, error) {
	var proxy operationMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Tags = v.tags
	proxy.Summary = v.summary
	proxy.Description = v.description
	proxy.ExternalDocs = v.externalDocs
	proxy.OperationID = v.operationID
	proxy.Parameters = v.parameters
	proxy.RequestBody = v.requestBody
	proxy.Responses = v.responses
	proxy.Callbacks = v.callbacks
	proxy.Deprecated = v.deprecated
	proxy.Security = v.security
	proxy.Servers = v.servers
	return json.Marshal(proxy)
}

func (v *operation) UnmarshalJSON(data []byte) error {
	var proxy operationUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.tags = proxy.Tags
	v.summary = proxy.Summary
	v.description = proxy.Description

	if len(proxy.ExternalDocs) > 0 {
		var decoded externalDocumentation
		if err := json.Unmarshal(proxy.ExternalDocs, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		v.externalDocs = &decoded
	}
	v.operationID = proxy.OperationID
	v.parameters = proxy.Parameters

	if len(proxy.RequestBody) > 0 {
		var decoded requestBody
		if err := json.Unmarshal(proxy.RequestBody, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field RequestBody`)
		}

		v.requestBody = &decoded
	}

	if len(proxy.Responses) > 0 {
		var decoded responses
		if err := json.Unmarshal(proxy.Responses, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Responses`)
		}

		v.responses = &decoded
	}
	v.callbacks = proxy.Callbacks
	v.deprecated = proxy.Deprecated
	v.security = proxy.Security
	v.servers = proxy.Servers
	return nil
}

func (v *operation) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "tags":
		target = v.tags
	case "summary":
		target = v.summary
	case "description":
		target = v.description
	case "externalDocs":
		target = v.externalDocs
	case "operationId":
		target = v.operationID
	case "parameters":
		target = v.parameters
	case "requestBody":
		target = v.requestBody
	case "responses":
		target = v.responses
	case "callbacks":
		target = v.callbacks
	case "deprecated":
		target = v.deprecated
	case "security":
		target = v.security
	case "servers":
		target = v.servers
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
