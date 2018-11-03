package openapi

type MapQueryJSON map[string]interface{}
type SliceQueryJSON []interface{}

type QueryJSONer interface {
	QueryJSON(string) (interface{}, bool)
}

const (
	Int32  = "int32"
	Int64  = "int64"
	Float  = "float"
	Double = "double"
)
