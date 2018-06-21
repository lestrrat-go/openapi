package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// ResponseBuilder is used to build an instance of Response. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Response
type ResponseBuilder struct {
	target *response
}

// Do finalizes the building process for Response and returns the result
func (b *ResponseBuilder) Do() Response {
	return b.target
}

// NewResponse creates a new builder object for Response
func NewResponse(description string) *ResponseBuilder {
	return &ResponseBuilder{
		target: &response{
			description: description,
		},
	}
}

// Headers sets the Headers field for object Response.
func (b *ResponseBuilder) Headers(v map[string]Header) *ResponseBuilder {
	b.target.headers = v
	return b
}

// Reference sets the $ref (reference) field for object Response.
func (b *ResponseBuilder) Reference(v string) *ResponseBuilder {
	b.target.reference = v
	return b
}
