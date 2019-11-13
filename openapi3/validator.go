package openapi3

import (
	"context"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type validator struct {
	recurse bool
}

func newValidator(recurse bool) *validator {
	return &validator{recurse: recurse}
}

func looksLikeSemanticVersion(s string) error {
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return r == '.'
	})
	if len(fields) != 3 {
		return errors.Errorf(`invalid number of fields in semantic version (%d)`, len(fields))
	}
	if fields[0] != "3" {
		return errors.Errorf(`invalid major version in semantic version (%s)`, fields[0])
	}

	for i := 1; i < 3; i++ {
		if _, err := strconv.Atoi(fields[i]); err != nil {
			return errors.Wrapf(err, `invalid semantic version component %s`, fields[i])
		}
	}
	return nil
}

func (val *validator) VisitOpenAPI(ctx context.Context, v OpenAPI) error {
	if err := looksLikeSemanticVersion(v.Version()); err != nil {
		return errors.Wrapf(err, `openapi field must be a semantic version 3.x.x (got %s)`, strconv.Quote(v.Version()))
	}

	if v.Info() == nil {
		return errors.New(`info is required`)
	}

	if v.Paths() == nil {
		return errors.New(`paths is required`)
	}

	return nil
}

func (val *validator) VisitOperation(ctx context.Context, v Operation) error {
	if v.Responses() == nil {
		return errors.New(`missing required field "responses"`)
	}

	return nil
}
