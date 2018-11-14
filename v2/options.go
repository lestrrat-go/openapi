package openapi

import (
	"sync"

	"github.com/lestrrat-go/openapi/internal/option"
)

const (
	optkeyLocker   = "locker"
	optkeyValidate = "validate"
)

type Option = option.Interface

// WithValidate specifies if validation should be performed on the
// object. This option can be passed to `ParseYAML`, `ParseJSON`,
// and `Do` methods for builders and mutators.
func WithValidate(v bool) Option {
	return option.New(optkeyValidate, v)
}

// WithLocker specifies the lock object to be used for applicable
// constructs, such as a Builder object. If a nil value is passed,
// usually it effectively becomes a non-locking operation. Please
// Consult the specific construct you're working with for exact
// semantics
func WithLocker(v sync.Locker) Option {
	return option.New(optkeyLocker, v)
}
