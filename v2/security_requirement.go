package openapi

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
)

func (v *securityRequirement) setName(s string) {
	v.name = s
}

func (v *securityRequirement) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return errors.Wrap(err, `failed to unmarshal JSON data`)
	}

	if ref, ok := m["$ref"]; ok {
		if sref, ok := ref.(string); ok {
			v.reference = sref
		} else {
			return errors.Errorf(`invalid value for $ref: %T`, ref)
		}
		return nil
	}

	scopes := make(map[string][]string)
	for name, value := range m {
		list, ok := value.([]interface{})
		if !ok {
			return errors.Errorf(`expected hash of arrays in security requirements, got %T`, value)
		}

		slist := make([]string, len(list))
		for i, scope := range list {
			sscope, ok := scope.(string)
			if !ok {
				return errors.Errorf(`expected list of scopes (strings) in security requirements, got %T`, scope)
			}
			slist[i] = sscope
		}

		scopes[name] = slist
	}

	v.scopes = scopes
	return nil
}

func (v *securityRequirement) MarshalJSON() ([]byte, error) {
	if ref := v.reference; ref != "" {
		return []byte(`{"$ref": ` + strconv.Quote(ref) + `}`), nil
	}

	return json.Marshal(v.scopes)
}
