package golang

import (
	"github.com/lestrrat-go/openapi/generator/restclientgen/compiler"
)

type Context struct {
	client      *compiler.ClientDefinition
	dir         string
	packageName string
	exportNew   bool
}

type Type interface {
	Name() string
	SetName(string)
}

type Builtin string

func (b Builtin) Name() string {
	return string(b)
}

func (b Builtin) SetName(s string) {
	panic("oops?")
}

type Array struct {
	name string
	elem string
}

func (a *Array) SetName(s string) {
	a.name = s
}

func isBuiltinType(s string) bool {
	switch s {
	case "string",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "byte", "rune", "bool":
		return true
	default:
		return false
	}
}

func (a *Array) Name() string {
	if a.name == "" {
		if isBuiltinType(a.elem) {
			return "[]" + a.elem
		} else {
			return "[]*" + a.elem
		}
	}
	return a.name
}
