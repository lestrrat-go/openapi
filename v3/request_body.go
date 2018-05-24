package openapi

func (v *requestBody) setName(s string) {
	v.name = s
}

func (v *requestBody) postUnmarshalJSON() {
	for mime, content := range v.content {
		content.setMime(mime)
	}
}

func (v *requestBody) Contents() *MediaTypeListIterator {
	var items []interface{}
	for _, mt := range v.content {
		items = append(items, mt)
	}

	var iter MediaTypeListIterator
	iter.items = items
	return &iter
}
