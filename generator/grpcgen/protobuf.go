package grpcgen

import (
	"github.com/pkg/errors"
)

func (v *Protobuf) AddImport(lib string) {
	v.imports[lib] = struct{}{}
}

func (v *Protobuf) AddMessage(msg *Message) {
	if v.messages == nil {
		v.messages = make(map[string]*Message)
	}
	v.messages[msg.name] = msg
}

func (v *Protobuf) LookupMessage(name string) (*Message, bool) {
	if len(v.messages) == 0 {
		return nil, false
	}
	m, ok := v.messages[name]
	return m, ok
}

func (v *Protobuf) GetService(name string) *Service {
	if v.services == nil {
		v.services = make(map[string]*Service)
	}

	s, ok := v.services[name]
	if !ok {
		s = &Service{name: name}
		v.services[name] = s
	}
	return s
}

func (v *Service) AddRPC(r *RPC) {
	v.rpcs = append(v.rpcs, r)
}

func (v Builtin) Name() string {
	return string(v)
}

func (v Builtin) ResolveIncomplete(_ *genCtx) (Type, error) {
	return v, nil
}

func (v *Message) Name() string {
	return v.name
}

func (v *Message) ResolveIncomplete(ctx *genCtx) (Type, error) {
	if ctx.IsResolving(v) {
		return v, nil
	}

	cancel := ctx.MarkAsResolving(v)
	defer cancel()

	for name, m := range v.messages {
		resolved, err := m.ResolveIncomplete(ctx)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to resolve nested message %s`, name)
		}
		msg, ok := resolved.(*Message)
		if !ok {
			return nil, errors.Errorf(`expected resolved type to be a Message, but got %T`, resolved)
		}
		v.messages[name] = msg
	}

	for _, f := range v.fields {
		resolved, err := f.typ.ResolveIncomplete(ctx)
		if err != nil {
			return nil, errors.Wrapf(err, `failed to resolve field %s`, f.name)
		}
		f.typ = resolved
	}

	return v, nil
}

func (v *Message) AddMessage(m *Message) {
	// no locking, cause we know we will only be used from a single goroutine
	if len(v.messages) == 0 {
		v.messages = make(map[string]*Message)
	}
	v.messages[m.Name()] = m
}

func (v *Message) LookupMessage(name string) (*Message, bool) {
	// no locking, cause we know we will only be used from a single goroutine
	if len(v.messages) == 0 {
		return nil, false
	}
	m, ok := v.messages[name]
	return m, ok
}

func (v *Array) Name() string {
	return "repeated " + v.element.Name()
}

func (v *Array) ResolveIncomplete(ctx *genCtx) (Type, error) {
	if ctx.IsResolving(v) {
		return v, nil
	}

	cancel := ctx.MarkAsResolving(v)
	defer cancel()

	resolved, err := v.element.ResolveIncomplete(ctx)
	if err != nil {
		return nil, errors.Wrap(err, `failed to resolve array element`)
	}
	v.element = resolved
	return v, nil
}

func (v Incomplete) Name() string {
	return "#incomplete"
}

func (v Incomplete) ResolveIncomplete(ctx *genCtx) (Type, error) {
	if typ, ok := ctx.LookupType(string(v)); ok {
		return typ, nil
	}
	return nil, errors.Errorf(`invalid ref: %s`, v)
}
