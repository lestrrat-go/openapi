package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T13:02:46+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// ResponsesBuilder is used to build an instance of Responses. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Responses
type ResponsesBuilder struct {
	target *responses
}

// Do finalizes the building process for Responses and returns the result
func (b *ResponsesBuilder) Do() Responses {
	return b.target
}

// NewResponses creates a new builder object for Responses
func NewResponses() *ResponsesBuilder {
	return &ResponsesBuilder{
		target: &responses{},
	}
}

// Default sets the Default field for object Responses.
func (b *ResponsesBuilder) Default(v Response) *ResponsesBuilder {
	b.target.defaultValue = v
	return b
}

// Reference sets the $ref (reference) field for object Responses.
func (b *ResponsesBuilder) Reference(v string) *ResponsesBuilder {
	b.target.reference = v
	return b
}