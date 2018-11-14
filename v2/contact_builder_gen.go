package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
	"sync"
)

var _ = errors.Cause

// ContactBuilder is used to build an instance of Contact. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Contact.
// Builders may NOT be reused. It must be created for every instance
// of Contact that you want to create
type ContactBuilder struct {
	lock   sync.Locker
	target *contact
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *ContactBuilder) MustBuild(options ...Option) Contact {
	v, err := b.Build(options...)
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Contact and returns the result
// By default, Build() will validate if the given structure is valid
func (b *ContactBuilder) Build(options ...Option) (Contact, error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return nil, errors.New(`builder has already been used`)
	}
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
	defer func() { b.target = nil }()
	return b.target, nil
}

// NewContact creates a new builder object for Contact
func NewContact(options ...Option) *ContactBuilder {
	var lock sync.Locker = &sync.Mutex{}
	for _, option := range options {
		switch option.Name() {
		case optkeyLocker:
			lock = option.Value().(sync.Locker)
		}
	}
	var b ContactBuilder
	if lock == nil {
		lock = nilLock{}
	}
	b.lock = lock
	b.target = &contact{}
	return &b
}

// Name sets the name field for object Contact.
func (b *ContactBuilder) Name(v string) *ContactBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.name = v
	return b
}

// URL sets the url field for object Contact.
func (b *ContactBuilder) URL(v string) *ContactBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.url = v
	return b
}

// Email sets the email field for object Contact.
func (b *ContactBuilder) Email(v string) *ContactBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.email = v
	return b
}

// Reference sets the $ref (reference) field for object Contact.
func (b *ContactBuilder) Reference(v string) *ContactBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.reference = v
	return b
}

// Extension sets an arbitrary element (an extension) to the
// object Contact. The extension name should start with a "x-"
func (b *ContactBuilder) Extension(name string, value interface{}) *ContactBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.extensions[name] = value
	return b
}
