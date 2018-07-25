package es6flow

import (
	"mime"

	"github.com/lestrrat-go/openapi/generator/restclientgen/compiler"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

type Context struct {
	client             *compiler.ClientDefinition
	clientName         string
	compiling          map[string]struct{}
	dir                string
	packageName        string
	defaultServiceName string
	resolver           openapi.Resolver
	root               openapi.Swagger
	types              map[string]typeDefinition
	consumes           []string
	produces           []string
}

type typeDefinition struct {
	Path    string
	Context string
	Type    Type
}

type Client struct {
	services map[string]*Service
	types    map[string]Type
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

type Struct struct {
	name   string
	fields []*Field
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
		return "Array<" + a.elem + ">"
	}
	return a.name
}

func (o *Struct) SetName(s string) {
	o.name = s
}

func (o *Struct) Name() string {
	return o.name
}

func (c *Client) getServiceFor(name string) *Service {
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

func (s *Service) addCall(call *Call) {
	s.calls = append(s.calls, call)
}

type Call struct {
	name        string
	method      string
	path        string
	verb        string
	consumes    []string
	produces    []string
	requireds   []*Field
	optionals   []*Field
	pathparams  []*Field
	queryparams []*Field
	bodyType    Type
	responses   []*Response
}

func (call *Call) DefaultConsumes() string {
	// default to "application/x-www-form-urlencoded"
	if len(call.consumes) == 0 {
		return "application/x-www-form-urlencoded"
	}
	return call.consumes[0]
}

func (call *Call) Consumes() []string {
	return call.consumes
}

type Response struct {
	code string
	typ  string
}

type Field struct {
	name     string // raw name
	jsName   string // lowerCameCase name
	typ      string
	tag      string
	required bool
	inBody   bool
}

func canonicalConsumesList(iter *openapi.MIMETypeListIterator) ([]string, error) {
	consumesSeen := map[string]struct{}{}

	var consumesList []string
	for iter.Next() {
		mt := iter.Item()
		if _, ok := consumesSeen[mt]; ok {
			continue
		}
		consumesList = append(consumesList, mt)
		consumesSeen[mt] = struct{}{}
	}

	// Make sure the consumes list is valid
	for i, v := range consumesList {
		mt, _, err := mime.ParseMediaType(v)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to parse "consumes" value %s`, v)
		}
		// Use the canonical mime type (the parsed one)
		consumesList[i] = mt
	}
	return consumesList, nil
}
