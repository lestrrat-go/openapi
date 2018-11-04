package compiler

import (
	"sort"

	openapi "github.com/lestrrat-go/openapi/v2"
)

type compileCtx struct {
	client             *ClientDefinition
	compiling          map[string]struct{}
	currentCall        *Call
	isCompiling        map[interface{}]struct{}
	isResolving        map[interface{}]struct{}
	defaultServiceName string
	resolver           openapi.Resolver
	root               openapi.Swagger
	consumes           []string
	security           map[string]openapi.SecurityScheme
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

func (c *ClientDefinition) Definitions() map[string]TypeDefinition {
	return c.definitions
}

func (c *ClientDefinition) ServiceNames() []string {
	var serviceNames []string
	for name := range c.services {
		serviceNames = append(serviceNames, name)
	}
	sort.Strings(serviceNames)
	return serviceNames
}

func (c *ClientDefinition) Services() map[string]*Service {
	return c.services
}

func (c *ClientDefinition) getServiceFor(name string) *Service {
	name = name + "Service"
	svc, ok := c.services[name]
	if !ok {
		svc = &Service{name: name}
		c.services[name] = svc
	}
	return svc
}

type Service struct {
	name  string
	calls []*Call
}

func (s *Service) Name() string {
	return s.name
}

func (s *Service) Calls() []*Call {
	return s.calls
}

func (s *Service) addCall(call *Call) {
	s.calls = append(s.calls, call)
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

func (r *Response) Code() string {
	return r.code
}

func (r *Response) Type() string {
	return r.typ
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
	in       openapi.Location
}

func (f *Field) Required() bool {
	return f.required
}

func (f *Field) Hints() Hints {
	return f.hints
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Type() Type {
	return f.typ
}

func (f *Field) In() openapi.Location {
	return f.in
}

func (f *Field) ContainerName() string {
	switch f.in {
	case openapi.InBody, openapi.InForm:
		return "body"
	case openapi.InQuery:
		return "query"
	case openapi.InHeader:
		return "header"
	case openapi.InPath:
		return "path"
	default:
		// No error case, as it should've been handled in Validate()
		return "(no container)"
	}
}

type SecuritySettings struct {
	definition openapi.SecurityScheme
	scopes     []string
}

func (ss *SecuritySettings) Definition() openapi.SecurityScheme {
	return ss.definition
}

func (ss *SecuritySettings) Scopes() []string {
	return ss.scopes
}
