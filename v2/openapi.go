//go:generate go run internal/cmd/gentypes/gentypes.go

// Package openapi implement OpenAPI Spec v2
package openapi

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

func validLocation(l Location) bool {
	switch l {
	case InPath, InQuery, InHeader, InBody, InForm:
		return true
	}
	return false
}

func extractFragFromPath(path string) (string, string) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	var frag string
	if i := strings.Index(path, `/`); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ``
	}
	return frag, path
}

func ParseYAML(src io.Reader, options ...Option) (Swagger, error) {
	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, errors.Wrap(err, `failed to read from source`)
	}

	var spec swagger
	if err := yaml.Unmarshal(buf, &spec); err != nil {
		return nil, errors.Wrap(err, `failed to unmarshal spec`)
	}

	validate := true
	for _, option := range options {
		switch option.Name() {
		case optkeyValidate:
			validate = option.Value().(bool)
		}
	}

	if validate {
		if err := spec.Validate(true); err != nil {
			return nil, errors.Wrap(err, `failed to validate spec`)
		}
	}

	return &spec, nil
}

func ParseJSON(src io.Reader, options ...Option) (Swagger, error) {
	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, errors.Wrap(err, `failed to read from source`)
	}

	var spec swagger
	if err := json.Unmarshal(buf, &spec); err != nil {
		return nil, errors.Wrap(err, `failed to unmarshal spec`)
	}

	validate := true
	for _, option := range options {
		switch option.Name() {
		case optkeyValidate:
			validate = option.Value().(bool)
		}
	}

	if validate {
		if err := spec.Validate(true); err != nil {
			return nil, errors.Wrap(err, `failed to validate spec`)
		}
	}

	return &spec, nil
}
