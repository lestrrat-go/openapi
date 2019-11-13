package openapi3

func (v *schema) setName(s string) {
	v.name = s
}

func (v *schema) Type() PrimitiveType {
	if v.typ != "" {
		return v.typ
	}

	// Guess
	if len(v.properties) > 0 {
		return Object
	}

	if v.items != nil {
		return Array
	}

	return Invalid
}
