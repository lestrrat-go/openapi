package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:57:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ = errors.Cause

type infoMarshalProxy struct {
	Title          string  `json:"title" builder:"required"`
	Description    string  `json:"description,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
	Version        string  `json:"version" builder:"required" default:"DefaultSpecVersion"`
}

type infoUnmarshalProxy struct {
	Title          string          `json:"title" builder:"required"`
	Description    string          `json:"description,omitempty"`
	TermsOfService string          `json:"termsOfService,omitempty"`
	Contact        json.RawMessage `json:"contact,omitempty"`
	License        json.RawMessage `json:"license,omitempty"`
	Version        string          `json:"version" builder:"required" default:"DefaultSpecVersion"`
}

func (v *info) MarshalJSON() ([]byte, error) {
	var proxy infoMarshalProxy
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
