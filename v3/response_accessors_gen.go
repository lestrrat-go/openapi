package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:57:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *response) Description() string {
	return v.description
}

func (v *response) Headers() map[string]Header {
	return v.headers
}

func (v *response) Content() map[string]MediaType {
	return v.content
}

func (v *response) Links() map[string]Link {
	return v.links
}
