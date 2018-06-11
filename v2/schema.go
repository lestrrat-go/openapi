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