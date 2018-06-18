package openapi

import "github.com/pkg/errors"

func (v *response) setName(s string) {
	v.name = s
}

func (v *response) setStatusCode(s string) {
	v.statusCode = s
}

func (v *response) Validate(recurse bool) error {
	if len(v.description) == 0 {
		return errors.New(`response description is required`)
	}

	if recurse {
		return v.recurseValidate()
	}
	return nil
}
