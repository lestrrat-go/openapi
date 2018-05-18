package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Example struct {
	target *entity.Example
}

func (b *Example) Build() *entity.Example {
	return b.target
}

func (b *Builder) NewExample() *Example {
	return &Example{
		target: &entity.Example{},
	}
}

func (b *Example) Description(v string) *Example {
	b.target.Description = v
	return b
}

func (b *Example) Value(v interface{}) *Example {
	b.target.Value = v
	return b
}

func (b *Example) ExternalValue(v string) *Example {
	b.target.ExternalValue = v
	return b
}
