package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "github.com/pkg/errors"

var _ = errors.Cause

// OperationBuilder is used to build an instance of Operation. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Operation
type OperationBuilder struct {
	target *operation
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *OperationBuilder) MustBuild(options ...Option) Operation {
	v, err := b.Build()
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Operation and returns the result
func (b *OperationBuilder) Build(options ...Option) (Operation, error) {
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

// NewOperation creates a new builder object for Operation
func NewOperation(responses Responses) *OperationBuilder {
	return &OperationBuilder{
		target: &operation{
			responses: responses,
		},
	}
}

// Tags sets the Tags field for object Operation.
func (b *OperationBuilder) Tags(v []string) *OperationBuilder {
	b.target.tags = v
	return b
}

// Summary sets the Summary field for object Operation.
func (b *OperationBuilder) Summary(v string) *OperationBuilder {
	b.target.summary = v
	return b
}

// Description sets the Description field for object Operation.
func (b *OperationBuilder) Description(v string) *OperationBuilder {
	b.target.description = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object Operation.
func (b *OperationBuilder) ExternalDocs(v ExternalDocumentation) *OperationBuilder {
	b.target.externalDocs = v
	return b
}

// OperationID sets the OperationID field for object Operation.
func (b *OperationBuilder) OperationID(v string) *OperationBuilder {
	b.target.operationID = v
	return b
}

// Parameters sets the Parameters field for object Operation.
func (b *OperationBuilder) Parameters(v []Parameter) *OperationBuilder {
	b.target.parameters = v
	return b
}

// RequestBody sets the RequestBody field for object Operation.
func (b *OperationBuilder) RequestBody(v RequestBody) *OperationBuilder {
	b.target.requestBody = v
	return b
}

// Callbacks sets the Callbacks field for object Operation.
func (b *OperationBuilder) Callbacks(v map[string]Callback) *OperationBuilder {
	b.target.callbacks = v
	return b
}

// Deprecated sets the Deprecated field for object Operation.
func (b *OperationBuilder) Deprecated(v bool) *OperationBuilder {
	b.target.deprecated = v
	return b
}

// Security sets the Security field for object Operation.
func (b *OperationBuilder) Security(v []SecurityRequirement) *OperationBuilder {
	b.target.security = v
	return b
}

// Servers sets the Servers field for object Operation.
func (b *OperationBuilder) Servers(v []Server) *OperationBuilder {
	b.target.servers = v
	return b
}

// Reference sets the $ref (reference) field for object Operation.
func (b *OperationBuilder) Reference(v string) *OperationBuilder {
	b.target.reference = v
	return b
}
