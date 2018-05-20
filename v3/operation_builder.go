// +build !gogenerate

package openapi

// Tag adds tag `s` to the operation
func (b *OperationBuilder) Tag(s string) *OperationBuilder {
	b.target.tags = append(b.target.tags, s)
	return b
}

// Parameter adds parameter `p` to the operation
func (b *OperationBuilder) Parameter(p Parameter) *OperationBuilder {
	b.target.parameters = append(b.target.parameters, p.Clone())
	return b
}

