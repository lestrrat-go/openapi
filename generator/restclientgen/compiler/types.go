package compiler

import (
	"fmt"
	"io"

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
	if a.name == "" {
		if isBuiltinType(a.elem.Name()) {
			return "[]" + a.elem.Name()
		} else {
			return "[]*" + a.elem.Name()
		}
	}
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

func (v *Struct) WriteCode(dst io.Writer) error {
	fmt.Fprintf(dst, "\ntype %s struct {", v.name)
	for _, field := range v.fields {
		typ := field.typ.Name()
		fmt.Fprintf(dst, "\n%s %s `%s`", field.hints.GoName, typ, field.hints.GoTag)
	}
	fmt.Fprintf(dst, "\n}")
	return nil
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
