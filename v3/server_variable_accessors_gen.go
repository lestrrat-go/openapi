package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:57:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *serverVariable) Enum() []string {
	return v.enum
}

func (v *serverVariable) Default() string {
	return v.defaultValue
}

func (v *serverVariable) Description() string {
	return v.description
}
