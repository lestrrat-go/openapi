package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *header) In() Location {
	return v.in
}

func (v *header) Required() bool {
	return v.required
}

func (v *header) Description() string {
	return v.description
}

func (v *header) Deprecated() bool {
	return v.deprecated
}

func (v *header) AllowEmptyValue() bool {
	return v.allowEmptyValue
}

func (v *header) Explode() bool {
	return v.explode
}

func (v *header) AllowReserved() bool {
	return v.allowReserved
}

func (v *header) Schema() Schema {
	return v.schema
}

func (v *header) Example() interface{} {
	return v.example
}

func (v *header) Examples() map[string]Example {
	return v.examples
}

func (v *header) Content() map[string]MediaType {
	return v.content
}
