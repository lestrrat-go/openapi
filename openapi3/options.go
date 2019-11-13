package openapi3

import "github.com/lestrrat-go/openapi/internal/option"

const (
	optkeyValidate = "validate"
)

type Option = option.Interface

// WithValidate specifies if validation should be performed on the
// object. This option can be passed to `ParseYAML`, `ParseJSON`, 
// and `Do` methods for builders and mutators.
func WithValidate(v bool) Option {
	return option.New(optkeyValidate, v)
}
