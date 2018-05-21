package openapi

import "net/http"

func (v *pathItem) setPath(s string) {
	v.path = s
}

func (v *pathItem) setVerb(verb string, oper Operation) {
	if oper == nil {
		return
	}
	oper.setVerb(verb)
	oper.setPathItem(v)
}

func (v *pathItem) postUnmarshalJSON() {
	v.setVerb(http.MethodGet, v.get)
	v.setVerb(http.MethodPut, v.put)
	v.setVerb(http.MethodPost, v.post)
	v.setVerb(http.MethodDelete, v.delete)
	v.setVerb(http.MethodOptions, v.options)
	v.setVerb(http.MethodHead, v.head)
	v.setVerb(http.MethodPatch, v.patch)
	v.setVerb(http.MethodTrace, v.trace)
}

// Operations returns an iterator that you can use to iterate through
// all non-nil operations 
func (v *pathItem) Operations() *OperationListIterator {
	var items []interface{}
	for _, oper := range []Operation{v.get, v.put, v.post, v.delete, v.options, v.head, v.patch, v.trace} {
		if oper != nil {
			items = append(items, oper)
		}
	}

	var iter OperationListIterator
	iter.items = items
	return &iter
}
