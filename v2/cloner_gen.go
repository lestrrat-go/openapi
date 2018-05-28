package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T19:20:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *swagger) Clone() Swagger {
	var dst swagger
	dst = *v
	return &dst
}

func (v *info) Clone() Info {
	var dst info
	dst = *v
	return &dst
}

func (v *contact) Clone() Contact {
	var dst contact
	dst = *v
	return &dst
}

func (v *license) Clone() License {
	var dst license
	dst = *v
	return &dst
}

func (v *paths) Clone() Paths {
	var dst paths
	dst = *v
	return &dst
}

func (v *pathItem) Clone() PathItem {
	var dst pathItem
	dst = *v
	return &dst
}

func (v *operation) Clone() Operation {
	var dst operation
	dst = *v
	return &dst
}

func (v *externalDocumentation) Clone() ExternalDocumentation {
	var dst externalDocumentation
	dst = *v
	return &dst
}

func (v *parameter) Clone() Parameter {
	var dst parameter
	dst = *v
	return &dst
}

func (v *items) Clone() Items {
	var dst items
	dst = *v
	return &dst
}

func (v *responses) Clone() Responses {
	var dst responses
	dst = *v
	return &dst
}

func (v *response) Clone() Response {
	var dst response
	dst = *v
	return &dst
}

func (v *header) Clone() Header {
	var dst header
	dst = *v
	return &dst
}

func (v *schema) Clone() Schema {
	var dst schema
	dst = *v
	return &dst
}

func (v *xml) Clone() XML {
	var dst xml
	dst = *v
	return &dst
}

func (v *securityScheme) Clone() SecurityScheme {
	var dst securityScheme
	dst = *v
	return &dst
}

func (v *securityRequirement) Clone() SecurityRequirement {
	var dst securityRequirement
	dst = *v
	return &dst
}

func (v *tag) Clone() Tag {
	var dst tag
	dst = *v
	return &dst
}
