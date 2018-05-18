package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type OAuthFlow struct {
	target *entity.OAuthFlow
}

func (b *OAuthFlow) Build() *entity.OAuthFlow {
	return b.target
}

func (b *Builder) NewOAuthFlow() *OAuthFlow {
	return &OAuthFlow{
		target: &entity.OAuthFlow{},
	}
}

func (b *OAuthFlow) AuthorizationURL(v string) *OAuthFlow {
	b.target.AuthorizationURL = v
	return b
}

func (b *OAuthFlow) TokenURL(v string) *OAuthFlow {
	b.target.TokenURL = v
	return b
}

func (b *OAuthFlow) RefreshURL(v string) *OAuthFlow {
	b.target.RefreshURL = v
	return b
}

func (b *OAuthFlow) Scopes(v map[string]string) *OAuthFlow {
	b.target.Scopes = v
	return b
}
