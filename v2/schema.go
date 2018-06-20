package openapi

import (
	"github.com/pkg/errors"
)

func (v *schema) setName(s string) {
	v.name = s
}

func (v *schema) IsRequiredProperty(prop string) bool {
	for _, name := range v.required {
		if name == prop {
			return true
		}
	}
	return false
}

// ConvertToSchema fulfills the SchemaCoverter interface.
// schema is already a Schema object, but it's useful to
// align the interface with other Schema-like objects.
// This method just returns itself
func (v *schema) ConvertToSchema() (Schema, error) {
	return v, nil
}

func (v *schema) Validate(recurse bool) error {
	if v.reference != "" {
		return nil
	}

	if v.typ != "" && !v.typ.IsValid() {
		return errors.Errorf(`invalid type in schema: %s`, v.typ)
	}

	if recurse {
		return v.recurseValidate()
	}
	return nil
}
