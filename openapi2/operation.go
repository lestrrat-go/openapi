package openapi2

func (v *operation) setPathItem(pi PathItem) {
	v.pathItem = pi
}

func (v *operation) setVerb(s string) {
	v.verb = s
}
