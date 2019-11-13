package openapi3

type MapQueryJSON map[string]interface{}
type SliceQueryJSON []interface{}

type QueryJSONer interface {
	QueryJSON(string) (interface{}, bool)
}