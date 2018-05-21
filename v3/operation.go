package openapi

func (v *operation) setPathItem(pi PathItem) {
	v.pathItem = pi
}

func (v *operation) setVerb(s string) {
	v.verb = s
}

func (v *operation) Detached() bool {
	return v.pathItem == nil
}

func (v *operation) Path() string {
	if v.Detached() {
		return ""
	}
	return v.pathItem.Path()
}

func (v *operation) Parameters() *ParameterListIterator{
	var iter ParameterListIterator
	for _, p := range v.parameters {
		iter.items = append(iter.items, p)
	}
	return &iter
}
