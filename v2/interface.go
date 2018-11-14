package openapi

type MapQueryJSON map[string]interface{}
type SliceQueryJSON []interface{}

type QueryJSONer interface {
	QueryJSON(string) (interface{}, bool)
}

type nilLock struct{}
func (l nilLock) Lock() {}
func (l nilLock) Unlock() {}

const (
	Int32  = "int32"
	Int64  = "int64"
	Float  = "float"
	Double = "double"
)
