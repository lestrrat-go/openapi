package openapi

import "github.com/pkg/errors"

func (v *parameter) Validate() error {
	// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#parameterObject
	if v.name == "" {
		return errors.New(`invalid parameter: "name" field is required`)
	}

	if !validLocation(v.in) {
		return errors.Errorf(`invalid parameter: invalid value in "in" field: %s`, v.in)
	}

	if v.in == InBody {
		if v.schema == nil {
			return errors.New(`invalid parameter: for parameters with {in: body}, you must specify the "schema" field`)
		}
		return nil
	}

	if v.allowEmptyValue {
		if v.in != InQuery && v.in != InForm {
			return errors.Errorf(`invalid parameter: {allowEmptyValue: true} is only applicable for "query" or "formData" parameters: got %s`, v.in)
		}
	}

	switch v.typ {
	case String, Number, Integer, Boolean:
	case Array:
		if v.items == nil {
			return errors.Errorf(`invalid paramter: for {type: array}, "items" field must be specified`)
		}
	case File:
		if v.in != InForm {
			return errors.Errorf(`invalid parameter: for {type: file}, "in" field must be "formData" (got %s)`, v.in)
		}
	default:
		return errors.Errorf(`invalid parameter: type must be one of "string", "number", "integer", "boolean", "array" or "file" (got %s)`, v.typ)
	}
	return nil
}
