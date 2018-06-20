package golang

import (
	"fmt"
	"io"
	"mime"

	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

type Context struct {
	client             *Client
	compiling          map[string]struct{}
	currentCall        *Call
	dir                string
	packageName        string
	defaultServiceName string
	resolver           openapi.Resolver
	root               openapi.Swagger
	types              map[string]typeDefinition
	exportNew          bool
	consumes           []string
	produces           []string
	security           map[string]openapi.SecurityScheme
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

func (v *Struct) WriteCode(dst io.Writer) error {
	fmt.Fprintf(dst, "\ntype %s struct {", v.name)
	for _, field := range v.fields {
		fmt.Fprintf(dst, "\n%s %s `%s`", field.goName, field.typ, field.tag)
	}
	fmt.Fprintf(dst, "\n}")
	return nil
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
	requestPath string
	verb        string
	consumes    []string
	produces    []string
	requireds   []*Field
	optionals   []*Field
	query       Type
	path        Type
	header      Type
	body        Type
	responses   []*Response
	securitySettings []*SecuritySettings
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
	goName   string // camelCase name
	typ      string
	tag      string
	required bool
	in       openapi.Location
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

type SecuritySettings struct {
	definition openapi.SecurityScheme
	scopes     []string
}
