package grpcgen

import (
	"context"
	"io"

	"github.com/lestrrat-go/openapi/internal/option"
	openapi "github.com/lestrrat-go/openapi/v2"
)

type Option = option.Interface

const (
	optkeyAnnotation   = "annotation"
	optkeyDestination  = "destination"
	optkeyGlobalOption = "global-option"
)

type Generator struct{}

type genCtx struct {
	context.Context
	dst      io.Writer
	indent   string
	types    map[string]Type
	resolver openapi.Resolver
	root     openapi.OpenAPI
	proto    *Protobuf
}

type Protobuf struct {
	globalOptions []*globalOption
	imports       map[string]struct{}
	messages      map[string]*Message
	options       map[string]string
	packageName   string
	services      map[string]*Service
}

type Type interface {
	Name() string
}

type Builtin string

type Array struct {
	element Type
}

type Message struct {
	name   string
	fields []*Field
}

type Field struct {
	id   int
	name string
	typ  Type
}

type Service struct {
	name string
	rpcs []*RPC
}

type RPC struct {
	name        string
	in          string
	out         string
	path        string
	description string
}
