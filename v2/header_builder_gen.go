package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"github.com/pkg/errors"
	"sync"
)

var _ = errors.Cause

// HeaderBuilder is used to build an instance of Header. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Header.
// Builders may NOT be reused. It must be created for every instance
// of Header that you want to create
type HeaderBuilder struct {
	lock   sync.Locker
	target *header
}

// MustBuild is a convenience function for those time when you know that
// the result of the builder must be successful
func (b *HeaderBuilder) MustBuild(options ...Option) Header {
	v, err := b.Build(options...)
	if err != nil {
		panic(err)
	}
	return v
}

// Build finalizes the building process for Header and returns the result
// By default, Build() will validate if the given structure is valid
func (b *HeaderBuilder) Build(options ...Option) (Header, error) {
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

// NewHeader creates a new builder object for Header
func NewHeader(typ string, options ...Option) *HeaderBuilder {
	var lock sync.Locker = &sync.Mutex{}
	for _, option := range options {
		switch option.Name() {
		case optkeyLocker:
			lock = option.Value().(sync.Locker)
		}
	}
	var b HeaderBuilder
	if lock == nil {
		lock = nilLock{}
	}
	b.lock = lock
	b.target = &header{
		typ: typ,
	}
	return &b
}

// Name sets the name field for object Header.
func (b *HeaderBuilder) Name(v string) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.name = v
	return b
}

// Description sets the description field for object Header.
func (b *HeaderBuilder) Description(v string) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.description = v
	return b
}

// Format sets the format field for object Header.
func (b *HeaderBuilder) Format(v string) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.format = v
	return b
}

// Items sets the items field for object Header.
func (b *HeaderBuilder) Items(v Items) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.items = v
	return b
}

// CollectionFormat sets the collectionFormat field for object Header.
func (b *HeaderBuilder) CollectionFormat(v CollectionFormat) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.collectionFormat = v
	return b
}

// Default sets the defaultValue field for object Header.
func (b *HeaderBuilder) Default(v interface{}) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.defaultValue = v
	return b
}

// Maximum sets the maximum field for object Header.
func (b *HeaderBuilder) Maximum(v float64) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.maximum = v
	return b
}

// ExclusiveMaximum sets the exclusiveMaximum field for object Header.
func (b *HeaderBuilder) ExclusiveMaximum(v float64) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.exclusiveMaximum = v
	return b
}

// Minimum sets the minimum field for object Header.
func (b *HeaderBuilder) Minimum(v float64) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.minimum = v
	return b
}

// ExclusiveMinimum sets the exclusiveMinimum field for object Header.
func (b *HeaderBuilder) ExclusiveMinimum(v float64) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.exclusiveMinimum = v
	return b
}

// MaxLength sets the maxLength field for object Header.
func (b *HeaderBuilder) MaxLength(v int) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.maxLength = v
	return b
}

// MinLength sets the minLength field for object Header.
func (b *HeaderBuilder) MinLength(v int) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.minLength = v
	return b
}

// Pattern sets the pattern field for object Header.
func (b *HeaderBuilder) Pattern(v string) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.pattern = v
	return b
}

// MaxItems sets the maxItems field for object Header.
func (b *HeaderBuilder) MaxItems(v int) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.maxItems = v
	return b
}

// MinItems sets the minItems field for object Header.
func (b *HeaderBuilder) MinItems(v int) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.minItems = v
	return b
}

// UniqueItems sets the uniqueItems field for object Header.
func (b *HeaderBuilder) UniqueItems(v bool) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.uniqueItems = v
	return b
}

// Enum sets the enum field for object Header.
func (b *HeaderBuilder) Enum(v ...interface{}) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.enum = v
	return b
}

// MultipleOf sets the multipleOf field for object Header.
func (b *HeaderBuilder) MultipleOf(v float64) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.multipleOf = v
	return b
}

// Reference sets the $ref (reference) field for object Header.
func (b *HeaderBuilder) Reference(v string) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.reference = v
	return b
}

// Extension sets an arbitrary element (an extension) to the
// object Header. The extension name should start with a "x-"
func (b *HeaderBuilder) Extension(name string, value interface{}) *HeaderBuilder {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.target == nil {
		return b
	}
	b.target.extensions[name] = value
	return b
}
