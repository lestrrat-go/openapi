package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import "log"

var _ = log.Printf

// PathItemMutator is used to build an instance of PathItem. The user must
// call `Do()` after providing all the necessary information to
// the new instance of PathItem with new values
type PathItemMutator struct {
	proxy  *pathItem
	target *pathItem
}

// Do finalizes the matuation process for PathItem and returns the result
func (b *PathItemMutator) Do() error {
	*b.target = *b.proxy
	return nil
}

// MutatePathItem creates a new mutator object for PathItem
func MutatePathItem(v PathItem) *PathItemMutator {
	return &PathItemMutator{
		target: v.(*pathItem),
		proxy:  v.Clone().(*pathItem),
	}
}

// Name sets the Name field for object PathItem.
func (b *PathItemMutator) Name(v string) *PathItemMutator {
	b.proxy.name = v
	return b
}

// Path sets the Path field for object PathItem.
func (b *PathItemMutator) Path(v string) *PathItemMutator {
	b.proxy.path = v
	return b
}

// Summary sets the Summary field for object PathItem.
func (b *PathItemMutator) Summary(v string) *PathItemMutator {
	b.proxy.summary = v
	return b
}

// Description sets the Description field for object PathItem.
func (b *PathItemMutator) Description(v string) *PathItemMutator {
	b.proxy.description = v
	return b
}

func (b *PathItemMutator) ClearServers() *PathItemMutator {
	b.proxy.servers.Clear()
	return b
}

func (b *PathItemMutator) Server(value Server) *PathItemMutator {
	b.proxy.servers = append(b.proxy.servers, value)
	return b
}

func (b *PathItemMutator) ClearParameters() *PathItemMutator {
	b.proxy.parameters.Clear()
	return b
}

func (b *PathItemMutator) Parameter(value Parameter) *PathItemMutator {
	b.proxy.parameters = append(b.proxy.parameters, value)
	return b
}
