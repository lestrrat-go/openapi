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

func (m *PathItemBuilder) Put(v Operation) *PathItemBuilder {
	m.target.acceptOperation(http.MethodPut, v)
	return m
}

func (m *PathItemBuilder) Post(v Operation) *PathItemBuilder {
	m.target.acceptOperation(http.MethodPost, v)
	return m
}

func (m *PathItemBuilder) Delete(v Operation) *PathItemBuilder {
	m.target.acceptOperation(http.MethodDelete, v)
	return m
}

func (m *PathItemBuilder) Options(v Operation) *PathItemBuilder {
	m.target.acceptOperation(http.MethodOptions, v)
	return m
}

func (m *PathItemBuilder) Head(v Operation) *PathItemBuilder {
	m.target.acceptOperation(http.MethodHead, v)
	return m
}

func (m *PathItemBuilder) Patch(v Operation) *PathItemBuilder {
	m.target.acceptOperation(http.MethodPatch, v)
	return m
}
