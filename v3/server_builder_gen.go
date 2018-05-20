package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:57:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

// ServerBuilder is used to build an instance of Server. The user must
// call `Build()` after providing all the necessary information to
// build an instance of Server
type ServerBuilder struct {
	target *server
}

// Build finalizes the building process for Server and returns the result
func (b *ServerBuilder) Build() Server {
	return b.target
}

// NewServer creates a new builder object for Server
func NewServer(uRL string) *ServerBuilder {
	return &ServerBuilder{
		target: &server{
			uRL: uRL,
		},
	}
}

// Description sets the Description field for object Server.
func (b *ServerBuilder) Description(v string) *ServerBuilder {
	b.target.description = v
	return b
}

// Variables sets the Variables field for object Server.
func (b *ServerBuilder) Variables(v map[string]ServerVariable) *ServerBuilder {
	b.target.variables = v
	return b
}
