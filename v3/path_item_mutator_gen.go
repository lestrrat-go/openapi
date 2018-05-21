package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

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

// Path sets the Path field for object PathItem.
func (b *PathItemMutator) Path(v string) *PathItemMutator {
	b.proxy.path = v
	return b
}

// Reference sets the Reference field for object PathItem.
func (b *PathItemMutator) Reference(v string) *PathItemMutator {
	b.proxy.reference = v
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

// Get sets the Get field for object PathItem.
func (b *PathItemMutator) Get(v Operation) *PathItemMutator {
	b.proxy.get = v
	return b
}

// Put sets the Put field for object PathItem.
func (b *PathItemMutator) Put(v Operation) *PathItemMutator {
	b.proxy.put = v
	return b
}

// Post sets the Post field for object PathItem.
func (b *PathItemMutator) Post(v Operation) *PathItemMutator {
	b.proxy.post = v
	return b
}

// Delete sets the Delete field for object PathItem.
func (b *PathItemMutator) Delete(v Operation) *PathItemMutator {
	b.proxy.delete = v
	return b
}

// Options sets the Options field for object PathItem.
func (b *PathItemMutator) Options(v Operation) *PathItemMutator {
	b.proxy.options = v
	return b
}

// Head sets the Head field for object PathItem.
func (b *PathItemMutator) Head(v Operation) *PathItemMutator {
	b.proxy.head = v
	return b
}

// Patch sets the Patch field for object PathItem.
func (b *PathItemMutator) Patch(v Operation) *PathItemMutator {
	b.proxy.patch = v
	return b
}

// Trace sets the Trace field for object PathItem.
func (b *PathItemMutator) Trace(v Operation) *PathItemMutator {
	b.proxy.trace = v
	return b
}

// Servers sets the Servers field for object PathItem.
func (b *PathItemMutator) Servers(v []Server) *PathItemMutator {
	b.proxy.servers = v
	return b
}

// Parameters sets the Parameters field for object PathItem.
func (b *PathItemMutator) Parameters(v []Parameter) *PathItemMutator {
	b.proxy.parameters = v
	return b
}
