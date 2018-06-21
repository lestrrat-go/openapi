package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// CallbackBuilder is used to build an instance of Callback. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Callback
type CallbackBuilder struct {
	target *callback
}

// Do finalizes the building process for Callback and returns the result
func (b *CallbackBuilder) Do() Callback {
	return b.target
}

// NewCallback creates a new builder object for Callback
func NewCallback() *CallbackBuilder {
	return &CallbackBuilder{
		target: &callback{},
	}
}

// URLs sets the URLs field for object Callback.
func (b *CallbackBuilder) URLs(v map[string]PathItem) *CallbackBuilder {
	b.target.urls = v
	return b
}

// Reference sets the $ref (reference) field for object Callback.
func (b *CallbackBuilder) Reference(v string) *CallbackBuilder {
	b.target.reference = v
	return b
}
