package grpcgen

import (
	"context"
	"io"

	"github.com/lestrrat-go/openapi/internal/option"
	openapi "github.com/lestrrat-go/openapi/v2"
)

type Option = option.Interface

const (
	optkeyDestination = "destination"
	optkeyAnnotation  = "annotation"
	optkeyPackageName = "package"
)

type Generator struct{}

type genCtx struct {
	context.Context
	dst      io.Writer
	messages map[string]*Message
	resolver openapi.Resolver
	root     openapi.OpenAPI
	proto    *Protobuf
}

type Protobuf struct {
	packageName string
	imports     map[string]struct{}
	options     map[string]string
	messages    map[string]*Message
	services    map[string]*Service
}

type Message struct {
	name   string
	fields []*Field
}

type Field struct {
	id       int
	name     string
	repeated bool
	typ      string
}

type Service struct {
	name string
	rpcs []*RPC
}

type RPC struct {
	name string
	in   string
	out  string
	path string
}
