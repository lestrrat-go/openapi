package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:57:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *components) Schemas() map[string]Schema {
	return v.schemas
}

func (v *components) Responses() map[string]Response {
	return v.responses
}

func (v *components) Parameters() map[string]Parameter {
	return v.parameters
}

func (v *components) Examples() map[string]Example {
	return v.examples
}

func (v *components) RequestBodies() map[string]RequestBody {
	return v.requestBodies
}

func (v *components) Headers() map[string]Header {
	return v.headers
}

func (v *components) SecuritySchemes() map[string]SecurityScheme {
	return v.securitySchemes
}

func (v *components) Links() map[string]Link {
	return v.links
}

func (v *components) Callbacks() map[string]Callback {
	return v.callbacks
}
