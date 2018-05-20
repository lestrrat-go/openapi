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
}

func (v *pathItem) postUnmarshalJSON() {
	v.setVerb(http.MethodGet, v.get)
	v.setVerb(http.MethodPost, v.post)
}

// Operations returns an iterator that you can use to iterate through
// all non-nil operations 
func (v *pathItem) Operations() *OperationIterator {
	var items []Operation
	for _, oper := range []Operation{v.get, v.put, v.post, v.delete, v.options, v.head, v.patch, v.trace} {
		if oper != nil {
			items = append(items, oper)
		}
	}

	return &OperationIterator{
		items: items,
	}
}

// Next returns true if there are more elements in this iterator
func (iter *OperationIterator) Next() bool {
	iter.mu.RLock()
	defer iter.mu.RUnlock()
	return iter.nextNoLock()
}

func (iter *OperationIterator) nextNoLock() bool {
	return len(iter.items) > 0
}

// Operation returns the next operation in this iterator
func (iter *OperationIterator) Operation() Operation {
	iter.mu.Lock()
	defer iter.mu.Unlock()

	if !iter.nextNoLock() {
		return nil
	}

	item := iter.items[0]
	iter.items = iter.items[1:]
	return item
}
