package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type OAuthFlows struct {
	target *entity.OAuthFlows
}

func (b *OAuthFlows) Build() *entity.OAuthFlows {
	return b.target
}

func (b *Builder) NewOAuthFlows() *OAuthFlows {
	return &OAuthFlows{
		target: &entity.OAuthFlows{},
	}
}

func (b *OAuthFlows) Implicit(v *entity.OAuthFlow) *OAuthFlows {
	b.target.Implicit = v
	return b
}

func (b *OAuthFlows) Password(v *entity.OAuthFlow) *OAuthFlows {
	b.target.Password = v
	return b
}

func (b *OAuthFlows) ClientCredentials(v *entity.OAuthFlow) *OAuthFlows {
	b.target.ClientCredentials = v
	return b
}

func (b *OAuthFlows) AuthorizationCode(v *entity.OAuthFlow) *OAuthFlows {
	b.target.AuthorizationCode = v
	return b
}
