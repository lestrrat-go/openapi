package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *externalDocumentation) Description() string {
	return v.description
}

func (v *externalDocumentation) UrL() string {
	return v.urL
}

func (v *externalDocumentation) Reference() string {
	return v.reference
}

func (v *externalDocumentation) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
