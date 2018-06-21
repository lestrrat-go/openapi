package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *info) Title() string {
	return v.title
}

func (v *info) Description() string {
	return v.description
}

func (v *info) TermsOfService() string {
	return v.termsOfService
}

func (v *info) Contact() Contact {
	return v.contact
}

func (v *info) License() License {
	return v.license
}

func (v *info) Version() string {
	return v.version
}

func (v *info) Reference() string {
	return v.reference
}

func (v *info) IsUnresolved() bool {
	return v.reference != "" && !v.resolved
}
