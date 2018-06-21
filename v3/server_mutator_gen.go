package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// ServerMutator is used to build an instance of Server. The user must
// call `Do()` after providing all the necessary information to
// the new instance of Server with new values
type ServerMutator struct {
	proxy  *server
	target *server
}

// Do finalizes the matuation process for Server and returns the result
func (b *ServerMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutateServer creates a new mutator object for Server
func MutateServer(v Server) *ServerMutator {
	return &ServerMutator{
		target: v.(*server),
		proxy:  v.Clone().(*server),
	}
}

// UrL sets the UrL field for object Server.
func (b *ServerMutator) UrL(v string) *ServerMutator {
	b.proxy.urL = v
	return b
}

// Description sets the Description field for object Server.
func (b *ServerMutator) Description(v string) *ServerMutator {
	b.proxy.description = v
	return b
}

func (b *ServerMutator) ClearVariables() *ServerMutator {
	b.proxy.variables.Clear()
	return b
}

func (b *ServerMutator) Variable(key ServerVariableMapKey, value ServerVariable) *ServerMutator {
	if b.proxy.variables == nil {
		b.proxy.variables = ServerVariableMap{}
	}

	b.proxy.variables[key] = value
	return b
}
