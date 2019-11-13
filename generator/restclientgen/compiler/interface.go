package compiler

import (
	"github.com/lestrrat-go/openapi/openapi2"
)

type compileCtx struct {
	client             *ClientDefinition
	compiling          map[string]struct{}
	currentCall        *Call
	isCompiling        map[interface{}]struct{}
	isResolving        map[interface{}]struct{}
	defaultServiceName string
	resolver           openapi2.Resolver
	root               openapi2.Swagger
	consumes           []string
	security           map[string]openapi2.SecurityScheme
}

// TypeDefinition gives you the context where this type was generated
type TypeDefinition struct {
	Path    string
	Context string
	Type    Type
}

type ClientDefinition struct {
	services    map[string]*Service
	definitions map[string]TypeDefinition
	types       map[string]Type
}

type Type interface {
	Name() string
	SetName(string)
	ResolveIncomplete(ctx *compileCtx) (Type, error)
}

type Incomplete string

type Builtin string

type Array struct {
	name string
	elem Type
}

type Struct struct {
	name   string
	fields []*Field
}

type Service struct {
	name  string
	calls []*Call
}

type Call struct {
	name             string
	method           string
	requestPath      string
	verb             string
	consumes         []string
	allFields        []*Field // only populated after a success compile
	requireds        []*Field
	optionals        []*Field
	query            Type
	header           Type
	body             Type
	path             Type
	responses        []*Response
	securitySettings []*SecuritySettings
}

type Response struct {
	code string
	typ  string
}

type Hints struct {
	GoName string // camelCase name
	GoTag  string
	JsName string
}

type Field struct {
	name     string // raw name
	hints    Hints
	typ      Type
	required bool
	in       openapi2.Location
}

type SecuritySettings struct {
	definition openapi2.SecurityScheme
	scopes     []string
}
