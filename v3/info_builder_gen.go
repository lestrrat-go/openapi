package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// InfoBuilder is used to build an instance of Info. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Info
type InfoBuilder struct {
	target *info
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *InfoBuilder) MustBuild(options ...Option) Info {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Info and returns the result
func (b *InfoBuilder) Build(options ...Option) (Info, error) {
	validate := true
	for _, option := range options {
		switch option.Name() {
		case optkeyValidate:
			validate = option.Value().(bool)
		}
	}
	if validate {
		if err := b.target.Validate(false); err != nil {
			return nil, errors.Wrap(err, `validation failed`)
		}
	}
	return b.target, nil
}

// NewInfo creates a new builder object for Info
func NewInfo(title string) *InfoBuilder {
	return &InfoBuilder{
		target: &info{
			version: DefaultSpecVersion,
			title:   title,
		},
	}
}

// Description sets the Description field for object Info.
func (b *InfoBuilder) Description(v string) *InfoBuilder {
	b.target.description = v
	return b
}

// TermsOfService sets the TermsOfService field for object Info.
func (b *InfoBuilder) TermsOfService(v string) *InfoBuilder {
	b.target.termsOfService = v
	return b
}

// Contact sets the Contact field for object Info.
func (b *InfoBuilder) Contact(v Contact) *InfoBuilder {
	b.target.contact = v
	return b
}

// License sets the License field for object Info.
func (b *InfoBuilder) License(v License) *InfoBuilder {
	b.target.license = v
	return b
}

// Version sets the Version field for object Info. If this is not called,
// a default value (DefaultSpecVersion) is assigned to this field
func (b *InfoBuilder) Version(v string) *InfoBuilder {
	b.target.version = v
	return b
}

// Reference sets the $ref (reference) field for object Info.
func (b *InfoBuilder) Reference(v string) *InfoBuilder {
	b.target.reference = v
	return b
}
