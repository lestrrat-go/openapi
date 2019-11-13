package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// SecurityRequirementMutator is used to build an instance of SecurityRequirement. The user must
// call `Do()` after providing all the necessary information to
// the new instance of SecurityRequirement with new values
type SecurityRequirementMutator struct {
	proxy  *securityRequirement
	target *securityRequirement
}

// Do finalizes the matuation process for SecurityRequirement and returns the result
func (b *SecurityRequirementMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateSecurityRequirement creates a new mutator object for SecurityRequirement
func MutateSecurityRequirement(v SecurityRequirement) *SecurityRequirementMutator {
	return &SecurityRequirementMutator{
		target: v.(*securityRequirement),
		proxy:  v.Clone().(*securityRequirement),
	}
}

func (b *SecurityRequirementMutator) ClearSchemes() *SecurityRequirementMutator {
	b.proxy.schemes.Clear()
	return b
}

func (b *SecurityRequirementMutator) Scheme(key StringListMapKey, value []string) *SecurityRequirementMutator {
	if b.proxy.schemes == nil {
		b.proxy.schemes = StringListMap{}
	}

	b.proxy.schemes[key] = value
	return b
}
