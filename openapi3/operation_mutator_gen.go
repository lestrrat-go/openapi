package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// OperationMutator is used to build an instance of Operation. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Operation with new values
type OperationMutator struct {
	proxy  *operation
	target *operation
}

// Do finalizes the matuation process for Operation and returns the result
func (b *OperationMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateOperation creates a new mutator object for Operation
func MutateOperation(v Operation) *OperationMutator {
	return &OperationMutator{
		target: v.(*operation),
		proxy:  v.Clone().(*operation),
	}
}

func (b *OperationMutator) ClearTags() *OperationMutator {
	b.proxy.tags.Clear()
	return b
}

func (b *OperationMutator) Tag(value string) *OperationMutator {
	b.proxy.tags = append(b.proxy.tags, value)
	return b
}

// Summary sets the Summary field for object Operation.
func (b *OperationMutator) Summary(v string) *OperationMutator {
	b.proxy.summary = v
	return b
}

// Description sets the Description field for object Operation.
func (b *OperationMutator) Description(v string) *OperationMutator {
	b.proxy.description = v
	return b
}

// ExternalDocs sets the ExternalDocs field for object Operation.
func (b *OperationMutator) ExternalDocs(v ExternalDocumentation) *OperationMutator {
	b.proxy.externalDocs = v
	return b
}

// OperationID sets the OperationID field for object Operation.
func (b *OperationMutator) OperationID(v string) *OperationMutator {
	b.proxy.operationID = v
	return b
}

func (b *OperationMutator) ClearParameters() *OperationMutator {
	b.proxy.parameters.Clear()
	return b
}

func (b *OperationMutator) Parameter(value Parameter) *OperationMutator {
	b.proxy.parameters = append(b.proxy.parameters, value)
	return b
}

// RequestBody sets the RequestBody field for object Operation.
func (b *OperationMutator) RequestBody(v RequestBody) *OperationMutator {
	b.proxy.requestBody = v
	return b
}

// Responses sets the Responses field for object Operation.
func (b *OperationMutator) Responses(v Responses) *OperationMutator {
	b.proxy.responses = v
	return b
}

func (b *OperationMutator) ClearCallbacks() *OperationMutator {
	b.proxy.callbacks.Clear()
	return b
}

func (b *OperationMutator) Callback(key CallbackMapKey, value Callback) *OperationMutator {
	if b.proxy.callbacks == nil {
		b.proxy.callbacks = CallbackMap{}
	}

	b.proxy.callbacks[key] = value
	return b
}

// Deprecated sets the Deprecated field for object Operation.
func (b *OperationMutator) Deprecated(v bool) *OperationMutator {
	b.proxy.deprecated = v
	return b
}

func (b *OperationMutator) ClearSecurity() *OperationMutator {
	b.proxy.security.Clear()
	return b
}

func (b *OperationMutator) Security(value SecurityRequirement) *OperationMutator {
	b.proxy.security = append(b.proxy.security, value)
	return b
}

func (b *OperationMutator) ClearServers() *OperationMutator {
	b.proxy.servers.Clear()
	return b
}

func (b *OperationMutator) Server(value Server) *OperationMutator {
	b.proxy.servers = append(b.proxy.servers, value)
	return b
}
