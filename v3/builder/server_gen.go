package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Server struct {
	target *entity.Server
}

func (b *Server) Build() *entity.Server {
	return b.target
}

func (b *Builder) NewServer(url string) *Server {
	return &Server{
		target: &entity.Server{
			URL: url,
		},
	}
}

func (b *Server) Description(v string) *Server {
	b.target.Description = v
	return b
}

func (b *Server) Variables(v map[string]*entity.ServerVariable) *Server {
	b.target.Variables = v
	return b
}
