package compiler

import (
	"fmt"
	"io"
	"sort"

	openapi "github.com/lestrrat-go/openapi/v2"
)

type compileCtx struct {
	client             *ClientDefinition
	compiling          map[string]struct{}
	currentCall        *Call
	dir                string
	packageName        string
	defaultServiceName string
	resolver           openapi.Resolver
	root               openapi.Swagger
	exportNew          bool
	consumes           []string
	produces           []string
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

func (a *Array) Elem() string {
	return a.elem
}

type Struct struct {
	name   string
	fields []*Field
}

func (s *Struct) Name() string {
	return s.name
}

func (s *Struct) Fields() []*Field {
	return s.fields
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

func (v *Struct) SetName(s string) {
	v.name = s
}

func (v *Struct) WriteCode(dst io.Writer) error {
	fmt.Fprintf(dst, "\ntype %s struct {", v.name)
	for _, field := range v.fields {
		fmt.Fprintf(dst, "\n%s %s `%s`", field.hints.GoName, field.typ, field.hints.GoTag)
	}
	fmt.Fprintf(dst, "\n}")
	return nil
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
	produces         []string
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
}

type Field struct {
	name     string // raw name
	hints    Hints
	typ      string
	required bool
	in       openapi.Location
}

func (f *Field) Hints() Hints {
	return f.hints
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Type() string {
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
