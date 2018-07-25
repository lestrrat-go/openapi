package compiler

func (call *Call) SecuritySettings() []*SecuritySettings {
	return call.securitySettings
}

func (call *Call) Name() string {
	return call.name
}

func (call *Call) Responses() []*Response {
	return call.responses
}

func (call *Call) Verb() string {
	return call.verb
}

func (call *Call) Method() string {
	return call.method
}

func (call *Call) Body() Type {
	return call.body
}

func (call *Call) Query() Type {
	return call.query
}

func (call *Call) Header() Type {
	return call.header
}

func (call *Call) Path() Type {
	return call.path
}

func (call *Call) RequestPath() string {
	return call.requestPath
}

func (call *Call) Optionals() []*Field {
	return call.optionals
}
func (call *Call) Requireds() []*Field {
	return call.requireds
}

func (call *Call) AllFields() []*Field {
	return call.allFields
}

func (call *Call) DefaultConsumes() string {
	// default to "application/x-www-form-urlencoded"
	if len(call.consumes) == 0 {
		return "application/x-www-form-urlencoded"
	}
	return call.consumes[0]
}

func (call *Call) Consumes() []string {
	return call.consumes
}

