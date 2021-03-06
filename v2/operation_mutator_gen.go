package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"sync"
)

// OperationMutator is used to build an instance of Operation. The user must
// call `Apply()` after providing all the necessary information to
// the new instance of Operation with new values
type OperationMutator struct {
	lock   sync.Locker
	proxy  *operation
	target *operation
}

// Apply finalizes the matuation process for Operation and returns the result
func (m *OperationMutator) Apply() error {
	m.lock.Lock()
	defer m.lock.Unlock()
	*m.target = *m.proxy
	return nil
}

// MutateOperation creates a new mutator object for Operation
// Operations on the mutator are safe to be used concurrently, except for
// when calling `Apply()`, where the user is responsible for restricting access
// to the target object to be mutated
func MutateOperation(v Operation, options ...Option) *OperationMutator {
	var lock sync.Locker = &sync.Mutex{}
	for _, option := range options {
		switch option.Name() {
		case optkeyLocker:
			lock = option.Value().(sync.Locker)
		}
	}
	if lock == nil {
		lock = nilLock{}
	}
	return &OperationMutator{
		lock:   lock,
		target: v.(*operation),
		proxy:  v.Clone().(*operation),
	}
}

// ClearTags clears all elements in tags
func (m *OperationMutator) ClearTags() *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	_ = m.proxy.tags.Clear()
	return m
}

// Tag appends a value to tags
func (m *OperationMutator) Tag(value string) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.tags = append(m.proxy.tags, value)
	return m
}

// Summary sets the Summary field for object Operation.
func (m *OperationMutator) Summary(v string) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.summary = v
	return m
}

// Description sets the Description field for object Operation.
func (m *OperationMutator) Description(v string) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.description = v
	return m
}

// ExternalDocs sets the ExternalDocs field for object Operation.
func (m *OperationMutator) ExternalDocs(v ExternalDocumentation) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.externalDocs = v
	return m
}

// OperationID sets the OperationID field for object Operation.
func (m *OperationMutator) OperationID(v string) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.operationID = v
	return m
}

// ClearConsumes clears all elements in consumes
func (m *OperationMutator) ClearConsumes() *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	_ = m.proxy.consumes.Clear()
	return m
}

// Consume appends a value to consumes
func (m *OperationMutator) Consume(value string) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.consumes = append(m.proxy.consumes, value)
	return m
}

// ClearProduces clears all elements in produces
func (m *OperationMutator) ClearProduces() *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	_ = m.proxy.produces.Clear()
	return m
}

// Produce appends a value to produces
func (m *OperationMutator) Produce(value string) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.produces = append(m.proxy.produces, value)
	return m
}

// ClearParameters clears all elements in parameters
func (m *OperationMutator) ClearParameters() *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	_ = m.proxy.parameters.Clear()
	return m
}

// Parameter appends a value to parameters
func (m *OperationMutator) Parameter(value Parameter) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.parameters = append(m.proxy.parameters, value)
	return m
}

// Responses sets the Responses field for object Operation.
func (m *OperationMutator) Responses(v Responses) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.responses = v
	return m
}

// ClearSchemes clears all elements in schemes
func (m *OperationMutator) ClearSchemes() *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	_ = m.proxy.schemes.Clear()
	return m
}

// Scheme appends a value to schemes
func (m *OperationMutator) Scheme(value string) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.schemes = append(m.proxy.schemes, value)
	return m
}

// Deprecated sets the Deprecated field for object Operation.
func (m *OperationMutator) Deprecated(v bool) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.deprecated = v
	return m
}

// ClearSecurity clears all elements in security
func (m *OperationMutator) ClearSecurity() *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	_ = m.proxy.security.Clear()
	return m
}

// Security appends a value to security
func (m *OperationMutator) Security(value SecurityRequirement) *OperationMutator {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.proxy.security = append(m.proxy.security, value)
	return m
}

// Extension sets an arbitrary extension field in Operation
func (m *OperationMutator) Extension(name string, value interface{}) *OperationMutator {
	if m.proxy.extensions == nil {
		m.proxy.extensions = Extensions{}
	}
	m.proxy.extensions[name] = value
	return m
}
