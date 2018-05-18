package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type PathItem struct {
	target *entity.PathItem
}

func (b *PathItem) Build() *entity.PathItem {
	return b.target
}

func (b *Builder) NewPathItem() *PathItem {
	return &PathItem{
		target: &entity.PathItem{},
	}
}

func (b *PathItem) Reference(v string) *PathItem {
	b.target.Reference = v
	return b
}

func (b *PathItem) Summary(v string) *PathItem {
	b.target.Summary = v
	return b
}

func (b *PathItem) Description(v string) *PathItem {
	b.target.Description = v
	return b
}

func (b *PathItem) Get(v *entity.Operation) *PathItem {
	b.target.Get = v
	return b
}

func (b *PathItem) Put(v *entity.Operation) *PathItem {
	b.target.Put = v
	return b
}

func (b *PathItem) Post(v *entity.Operation) *PathItem {
	b.target.Post = v
	return b
}

func (b *PathItem) Delete(v *entity.Operation) *PathItem {
	b.target.Delete = v
	return b
}

func (b *PathItem) Options(v *entity.Operation) *PathItem {
	b.target.Options = v
	return b
}

func (b *PathItem) Head(v *entity.Operation) *PathItem {
	b.target.Head = v
	return b
}

func (b *PathItem) Patch(v *entity.Operation) *PathItem {
	b.target.Patch = v
	return b
}

func (b *PathItem) Trace(v *entity.Operation) *PathItem {
	b.target.Trace = v
	return b
}

func (b *PathItem) Servers(v []*entity.Server) *PathItem {
	b.target.Servers = v
	return b
}

func (b *PathItem) Parameters(v []*entity.Parameter) *PathItem {
	b.target.Parameters = v
	return b
}
