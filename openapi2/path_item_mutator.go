package openapi2

import "net/http"

func (m *PathItemMutator) Get(v Operation) *PathItemMutator {
	m.proxy.acceptOperation(http.MethodGet, v)
	return m
}

func (m *PathItemMutator) Put(v Operation) *PathItemMutator {
	m.proxy.acceptOperation(http.MethodPut, v)
	return m
}

func (m *PathItemMutator) Post(v Operation) *PathItemMutator {
	m.proxy.acceptOperation(http.MethodPost, v)
	return m
}

func (m *PathItemMutator) Delete(v Operation) *PathItemMutator {
	m.proxy.acceptOperation(http.MethodDelete, v)
	return m
}

func (m *PathItemMutator) Options(v Operation) *PathItemMutator {
	m.proxy.acceptOperation(http.MethodOptions, v)
	return m
}

func (m *PathItemMutator) Head(v Operation) *PathItemMutator {
	m.proxy.acceptOperation(http.MethodHead, v)
	return m
}

func (m *PathItemMutator) Patch(v Operation) *PathItemMutator {
	m.proxy.acceptOperation(http.MethodPatch, v)
	return m
}

func (m *PathItemMutator) Trace(v Operation) *PathItemMutator {
	m.proxy.acceptOperation(http.MethodTrace, v)
	return m
}
