package compiler

import (
	"sort"

	"github.com/lestrrat-go/openapi/openapi2"
	"github.com/pkg/errors"
)

func (b Builtin) Name() string {
	return string(b)
}

func (b Builtin) SetName(s string) {
	panic("oops?")
}

func (b Builtin) ResolveIncomplete(ctx *compileCtx) (Type, error) {
	return b, nil
}

func (a *Array) Name() string {
	/*
		if a.name == "" {
			if isBuiltinType(a.elem.Name()) {
				return "[]" + a.elem.Name()
			} else {
				return "[]*" + a.elem.Name()
			}
		}
	*/
	return a.name
}

func (a *Array) SetName(s string) {
	a.name = s
}

func (a *Array) Elem() string {
	return a.elem.Name()
}

func (a *Array) ResolveIncomplete(ctx *compileCtx) (Type, error) {
	if ctx.IsResolving(a.elem) {
		return a.elem, nil
	}
	cancel := ctx.MarkAsResolving(a.elem)
	defer cancel()

	typ, err := a.elem.ResolveIncomplete(ctx)
	if err != nil {
		return nil, errors.Wrap(err, `failed to resolve array element type`)
	}
	a.elem = typ
	return a, nil
}

func (s *Struct) Name() string {
	return s.name
}

func (v *Struct) SetName(s string) {
	v.name = s
}

func (s *Struct) Fields() []*Field {
	return s.fields
}

func (v *Struct) ResolveIncomplete(ctx *compileCtx) (Type, error) {
	for _, field := range v.fields {
		if ctx.IsResolving(field.typ) {
			return v, nil
		}

		cancel := ctx.MarkAsResolving(field.typ)
		typ, err := field.typ.ResolveIncomplete(ctx)
		if err != nil {
			cancel()
			return nil, errors.Wrapf(err, `failed to resolve incomplete type for field %s`, field.Name())
		}
		field.typ = typ
		cancel()
	}
	return v, nil
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

func (v Incomplete) Name() string {
	return string(v)
}

func (v Incomplete) SetName(s string) {
}

func (v Incomplete) ResolveIncomplete(ctx *compileCtx) (Type, error) {
	typ, ok := lookupReferencedType(ctx, string(v))
	if !ok {
		return nil, errors.Errorf(`invalid reference %s`, string(v))
	}
	return typ, nil
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
	svc, ok := c.services[name]
	if !ok {
		svc = &Service{name: name}
		c.services[name] = svc
	}
	return svc
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

func (r *Response) Code() string {
	return r.code
}

func (r *Response) Type() string {
	return r.typ
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

func (f *Field) In() openapi2.Location {
	return f.in
}

func (f *Field) ContainerName() string {
	switch f.in {
	case openapi2.InBody, openapi2.InForm:
		return "body"
	case openapi2.InQuery:
		return "query"
	case openapi2.InHeader:
		return "header"
	case openapi2.InPath:
		return "path"
	default:
		// No error case, as it should've been handled in Validate()
		return "(no container)"
	}
}

func (ss *SecuritySettings) Definition() openapi2.SecurityScheme {
	return ss.definition
}

func (ss *SecuritySettings) Scopes() []string {
	return ss.scopes
}
