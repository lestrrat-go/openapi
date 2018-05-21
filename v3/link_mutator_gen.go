package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// LinkMutator is used to build an instance of Link. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Link with new values
type LinkMutator struct {
	proxy  *link
	target *link
}

// Do finalizes the matuation process for Link and returns the result
func (b *LinkMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateLink creates a new mutator object for Link
func MutateLink(v Link) *LinkMutator {
	return &LinkMutator{
		target: v.(*link),
		proxy:  v.Clone().(*link),
	}
}

// OperationRef sets the OperationRef field for object Link.
func (b *LinkMutator) OperationRef(v string) *LinkMutator {
	b.proxy.operationRef = v
	return b
}

// OperationID sets the OperationID field for object Link.
func (b *LinkMutator) OperationID(v string) *LinkMutator {
	b.proxy.operationID = v
	return b
}

// Parameters sets the Parameters field for object Link.
func (b *LinkMutator) Parameters(v map[string]interface{}) *LinkMutator {
	b.proxy.parameters = v
	return b
}

// RequestBody sets the RequestBody field for object Link.
func (b *LinkMutator) RequestBody(v interface{}) *LinkMutator {
	b.proxy.requestBody = v
	return b
}

// Description sets the Description field for object Link.
func (b *LinkMutator) Description(v string) *LinkMutator {
	b.proxy.description = v
	return b
}

// Server sets the Server field for object Link.
func (b *LinkMutator) Server(v Server) *LinkMutator {
	b.proxy.server = v
	return b
}
