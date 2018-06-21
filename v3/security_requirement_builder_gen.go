package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// SecurityRequirementBuilder is used to build an instance of SecurityRequirement. The user must
// call `Do()` after providing all the necessary information to
// build an instance of SecurityRequirement
type SecurityRequirementBuilder struct {
	target *securityRequirement
}

// Do finalizes the building process for SecurityRequirement and returns the result
func (b *SecurityRequirementBuilder) Do() SecurityRequirement {
	return b.target
}

// NewSecurityRequirement creates a new builder object for SecurityRequirement
func NewSecurityRequirement() *SecurityRequirementBuilder {
	return &SecurityRequirementBuilder{
		target: &securityRequirement{},
	}
}

// Schemes sets the Schemes field for object SecurityRequirement.
func (b *SecurityRequirementBuilder) Schemes(v map[string][]string) *SecurityRequirementBuilder {
	b.target.schemes = v
	return b
}

// Reference sets the $ref (reference) field for object SecurityRequirement.
func (b *SecurityRequirementBuilder) Reference(v string) *SecurityRequirementBuilder {
	b.target.reference = v
	return b
}
