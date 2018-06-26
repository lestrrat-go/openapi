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

type infoMarshalProxy struct {
	Reference      string  `json:"$ref,omitempty"`
	Title          string  `json:"title"`
	Description    string  `json:"description,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
	Version        string  `json:"version"`
}

type infoUnmarshalProxy struct {
	Reference      string          `json:"$ref,omitempty"`
	Title          string          `json:"title"`
	Description    string          `json:"description,omitempty"`
	TermsOfService string          `json:"termsOfService,omitempty"`
	Contact        json.RawMessage `json:"contact,omitempty"`
	License        json.RawMessage `json:"license,omitempty"`
	Version        string          `json:"version"`
}

func (v *info) MarshalJSON() ([]byte, error) {
	var proxy infoMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Title = v.title
	proxy.Description = v.description
	proxy.TermsOfService = v.termsOfService
	proxy.Contact = v.contact
	proxy.License = v.license
	proxy.Version = v.version
	return json.Marshal(proxy)
}

func (v *info) UnmarshalJSON(data []byte) error {
	var proxy infoUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.title = proxy.Title
	v.description = proxy.Description
	v.termsOfService = proxy.TermsOfService

	if len(proxy.Contact) > 0 {
		var decoded contact
		if err := json.Unmarshal(proxy.Contact, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Contact`)
		}

		v.contact = &decoded
	}

	if len(proxy.License) > 0 {
		var decoded license
		if err := json.Unmarshal(proxy.License, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field License`)
		}

		v.license = &decoded
	}
	v.version = proxy.Version
	return nil
}

func (v *info) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "title":
		target = v.title
	case "description":
		target = v.description
	case "termsOfService":
		target = v.termsOfService
	case "contact":
		target = v.contact
	case "license":
		target = v.license
	case "version":
		target = v.version
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
