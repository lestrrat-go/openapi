package openapi

import "net/http"

func (v *pathItem) setPath(s string) {
	v.path = s
	v.setVerb(http.MethodGet, v.get)
	v.setVerb(http.MethodPut, v.put)
	v.setVerb(http.MethodPost, v.post)
	v.setVerb(http.MethodDelete, v.delete)
	v.setVerb(http.MethodOptions, v.options)
	v.setVerb(http.MethodHead, v.head)
	v.setVerb(http.MethodPatch, v.patch)
}

func (v *pathItem) setVerb(verb string, oper Operation) {
	if oper == nil {
		return
	}
	oper.setVerb(verb)
	oper.setPathItem(v)
}

func (b *PathItemBuilder) Get(oper Operation) *PathItemBuilder {
	b.target.acceptOperation(http.MethodGet, oper)
	return b
}

func (b *PathItemBuilder) Put(v Operation) *PathItemBuilder {
	b.target.acceptOperation(http.MethodPut, v)
	return b
}

func (b *PathItemBuilder) Post(v Operation) *PathItemBuilder {
	b.target.acceptOperation(http.MethodPost, v)
	return b
}

func (b *PathItemBuilder) Delete(v Operation) *PathItemBuilder {
	b.target.acceptOperation(http.MethodDelete, v)
	return b
}

func (b *PathItemBuilder) Options(v Operation) *PathItemBuilder {
	b.target.acceptOperation(http.MethodOptions, v)
	return b
}

func (b *PathItemBuilder) Head(v Operation) *PathItemBuilder {
	b.target.acceptOperation(http.MethodHead, v)
	return b
}

func (b *PathItemBuilder) Patch(v Operation) *PathItemBuilder {
	b.target.acceptOperation(http.MethodPatch, v)
	return b
}
