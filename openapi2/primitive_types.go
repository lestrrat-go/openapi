package openapi2

func (p PrimitiveType) IsValid() bool {
	switch p {
	case Integer, Number, String, Boolean, Object, Array, File, Null:
		return true
	}
	return false
}
