package openapi

import "net/http"

func (v *pathItem) setName(s string) {
	v.name = s
}

func (v *pathItem) acceptOperation(method string, oper Operation) {
	cloned := oper.Clone()
	cloned.setVerb(method)
	cloned.setPathItem(v)
	switch method {
	case http.MethodGet:
		v.get = cloned
	case http.MethodPut:
		v.put = cloned
	case http.MethodPost:
		v.post = cloned
	case http.MethodDelete:
		v.delete = cloned
	case http.MethodOptions:
		v.options = cloned
	case http.MethodHead:
		v.head = cloned
	case http.MethodPatch:
		v.patch = cloned
	}
}

// Operations returns an iterator that you can use to iterate through
// all non-nil operations
func (v *pathItem) Operations() *OperationListIterator {
	var items []interface{}
	for _, oper := range []Operation{v.get, v.put, v.post, v.delete, v.options, v.head, v.patch} {
		if oper != nil {
			items = append(items, oper)
		}
	}

	var iter OperationListIterator
	iter.items = items
	return &iter
}
