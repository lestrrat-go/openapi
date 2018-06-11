package openapi

import "github.com/pkg/errors"

func (v *operation) setPathItem(pi PathItem) {
	v.pathItem = pi
}

func (v *operation) setVerb(s string) {
	v.verb = s
}

func (v *operation) Validate(recurse bool) error {
	inMap := make(map[Location][]string) // map of parameter location to param name
	for piter := v.Parameters(); piter.Next(); {
		param := piter.Item()
		inMap[param.In()] = append(inMap[param.In()], param.Name())
	}

	if names, ok := inMap[InBody]; ok {
		if len(names) > 1 {
			return errors.Errorf(`there can only be 1 body parameter got %v`, names)
		}

		// XXX this check is transitive. this case alone will suffice
		if formNames, ok := inMap[InForm]; ok {
			return errors.Errorf(`both "body" and "formData" parameters are present (can only have either) body = %v, formData = %v`, names, formNames)
		}
	}

	if recurse {
		return v.recurseValidate()
	}
	return nil
}
