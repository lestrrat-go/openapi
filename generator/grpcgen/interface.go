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
	optkeyPackageName  = "package-name"
)

type Generator struct{}

type messageContainer interface {
	AddMessage(*Message)
	LookupMessage(string) (*Message, bool)
}

type genCtx struct {
	context.Context
	annotate    bool
	dst         io.Writer
	isCompiling map[string]struct{}
	isResolving map[interface{}]struct{}
	indent      string
	types       map[string]Type
	resolver    openapi.Resolver
	root        openapi.OpenAPI
	proto       *Protobuf
	parent      messageContainer
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
	ResolveIncomplete(*genCtx) (Type, error)
}

type Builtin string

type Array struct {
	element Type
}

type Message struct {
	name      string
	reference string
	fields    []*Field
	messages  map[string]*Message
}

type Field struct {
	id   int
	body bool
	name string
	typ  Type
}

type Service struct {
	name string
	rpcs []*RPC
}

type RPC struct {
	name        string
	in          Type
	out         Type
	path        string
	verb        string
	description string
}

// Incomplete is a type that can only be fulfilled once all of the
// definitions have been compiled
type Incomplete string
