package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type SecurityScheme struct {
	target *entity.SecurityScheme
}

func (b *SecurityScheme) Build() *entity.SecurityScheme {
	return b.target
}

func (b *Builder) NewSecurityScheme() *SecurityScheme {
	return &SecurityScheme{
		target: &entity.SecurityScheme{},
	}
}

func (b *SecurityScheme) Type(v string) *SecurityScheme {
	b.target.Type = v
	return b
}

func (b *SecurityScheme) Description(v string) *SecurityScheme {
	b.target.Description = v
	return b
}

func (b *SecurityScheme) Name(v string) *SecurityScheme {
	b.target.Name = v
	return b
}

func (b *SecurityScheme) In(v string) *SecurityScheme {
	b.target.In = v
	return b
}

func (b *SecurityScheme) Scheme(v string) *SecurityScheme {
	b.target.Scheme = v
	return b
}

func (b *SecurityScheme) BearerFormat(v string) *SecurityScheme {
	b.target.BearerFormat = v
	return b
}

func (b *SecurityScheme) Flows(v *entity.OAuthFlows) *SecurityScheme {
	b.target.Flows = v
	return b
}

func (b *SecurityScheme) OpenIDConnectURL(v string) *SecurityScheme {
	b.target.OpenIDConnectURL = v
	return b
}
