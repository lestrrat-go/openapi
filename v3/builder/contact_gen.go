package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Contact struct {
	target *entity.Contact
}

func (b *Contact) Build() *entity.Contact {
	return b.target
}

func (b *Builder) NewContact() *Contact {
	return &Contact{
		target: &entity.Contact{},
	}
}

func (b *Contact) Name(v string) *Contact {
	b.target.Name = v
	return b
}

func (b *Contact) URL(v string) *Contact {
	b.target.URL = v
	return b
}

func (b *Contact) Email(v string) *Contact {
	b.target.Email = v
	return b
}
