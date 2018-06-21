package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

// LinkBuilder is used to build an instance of Link. The user must
// call `Do()` after providing all the necessary information to
// build an instance of Link
type LinkBuilder struct {
	target *link
}

// Do finalizes the building process for Link and returns the result
func (b *LinkBuilder) Do() Link {
	return b.target
}

// NewLink creates a new builder object for Link
func NewLink() *LinkBuilder {
	return &LinkBuilder{
		target: &link{},
	}
}

// OperationRef sets the OperationRef field for object Link.
func (b *LinkBuilder) OperationRef(v string) *LinkBuilder {
	b.target.operationRef = v
	return b
}

// OperationID sets the OperationID field for object Link.
func (b *LinkBuilder) OperationID(v string) *LinkBuilder {
	b.target.operationID = v
	return b
}

// Parameters sets the Parameters field for object Link.
func (b *LinkBuilder) Parameters(v map[string]interface{}) *LinkBuilder {
	b.target.parameters = v
	return b
}

// RequestBody sets the RequestBody field for object Link.
func (b *LinkBuilder) RequestBody(v interface{}) *LinkBuilder {
	b.target.requestBody = v
	return b
}

// Description sets the Description field for object Link.
func (b *LinkBuilder) Description(v string) *LinkBuilder {
	b.target.description = v
	return b
}

// Server sets the Server field for object Link.
func (b *LinkBuilder) Server(v Server) *LinkBuilder {
	b.target.server = v
	return b
}

// Reference sets the $ref (reference) field for object Link.
func (b *LinkBuilder) Reference(v string) *LinkBuilder {
	b.target.reference = v
	return b
}
