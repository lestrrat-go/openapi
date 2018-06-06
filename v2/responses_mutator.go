package openapi

func (b *ResponsesMutator) Response(key ResponseMapKey, value Response) *ResponsesMutator {
	b.proxy.setResponse(key, value)
	return b
}

