package openapi

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
