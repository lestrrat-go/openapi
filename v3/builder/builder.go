package builder

import "github.com/lestrrat-go/openapi/v3/entity"

// Package builder provides objects to safely build a set of objects to
// represent an OpenAPI object.

func New() *Builder {
	return &Builder{}
}

/*
func (b *Builder) NewOpenAPI() *OpenAPI {
	return &OpenAPI{
		target: &entity.OpenAPI{
			OpenAPI: "3.0.1",
		},
	}
}
*/

// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.1.md#parameterObject
// required (boolean)
//    Determines whether this parameter is mandatory. If the parameter
//    location is "path", this property is REQUIRED and its value MUST
//    be true. Otherwise, the property MAY be included and its default
//    value is false.
func defaultParameterRequiredFromLocation(in entity.Location) bool {
	return in == entity.InPath
}
