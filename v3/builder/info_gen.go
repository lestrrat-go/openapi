package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Info struct {
	target *entity.Info
}

func (b *Info) Build() *entity.Info {
	return b.target
}

func (b *Builder) NewInfo(title string) *Info {
	return &Info{
		target: &entity.Info{
			Version: DefaultSpecVersion,
			Title:   title,
		},
	}
}

func (b *Info) Description(v string) *Info {
	b.target.Description = v
	return b
}

func (b *Info) TermsOfService(v string) *Info {
	b.target.TermsOfService = v
	return b
}

func (b *Info) Contact(v *entity.Contact) *Info {
	b.target.Contact = v
	return b
}

func (b *Info) License(v *entity.License) *Info {
	b.target.License = v
	return b
}

func (b *Info) Version(v string) *Info {
	b.target.Version = v
	return b
}
