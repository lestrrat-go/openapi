package openapi

import "net/http"

func (m *PathItemMutator) acceptOperation(method string, v Operation) Operation {
	cloned := v.Clone()
	cloned.setVerb(method)
	cloned.setPathItem(m.proxy)
	return cloned
}

func (m *PathItemMutator) Get(v Operation) *PathItemMutator {
	m.proxy.get = m.acceptOperation(http.MethodGet, v)
	return m
}

func (m *PathItemMutator) Put(v Operation) *PathItemMutator {
	m.proxy.put = m.acceptOperation(http.MethodPut, v)
	return m
}

func (m *PathItemMutator) Post(v Operation) *PathItemMutator {
	m.proxy.post = m.acceptOperation(http.MethodPost, v)
	return m
}

func (m *PathItemMutator) Delete(v Operation) *PathItemMutator {
	m.proxy.delete = m.acceptOperation(http.MethodDelete, v)
	return m
}

func (m *PathItemMutator) Options(v Operation) *PathItemMutator {
	m.proxy.options = m.acceptOperation(http.MethodOptions, v)
	return m
}

func (m *PathItemMutator) Head(v Operation) *PathItemMutator {
	m.proxy.head = m.acceptOperation(http.MethodHead, v)
	return m
}

func (m *PathItemMutator) Patch(v Operation) *PathItemMutator {
	m.proxy.patch = m.acceptOperation(http.MethodPatch, v)
	return m
}

func (m *PathItemMutator) Trace(v Operation) *PathItemMutator {
	m.proxy.trace = m.acceptOperation(http.MethodTrace, v)
	return m
}
