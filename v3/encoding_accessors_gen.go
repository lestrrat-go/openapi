package openapi

// This file was automatically generated by genbuilders.go on 2018-05-20T20:57:22+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

func (v *encoding) ContentType() string {
	return v.contentType
}

func (v *encoding) Headers() map[string]Header {
	return v.headers
}

func (v *encoding) Explode() bool {
	return v.explode
}

func (v *encoding) AllowReserved() bool {
	return v.allowReserved
}
