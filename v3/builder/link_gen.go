package builder

import (
	"github.com/lestrrat-go/openapi/v3/entity"
)

type Link struct {
	target *entity.Link
}

func (b *Link) Build() *entity.Link {
	return b.target
}

func (b *Builder) NewLink() *Link {
	return &Link{
		target: &entity.Link{},
	}
}

func (b *Link) OperationRef(v string) *Link {
	b.target.OperationRef = v
	return b
}

func (b *Link) OperationID(v string) *Link {
	b.target.OperationID = v
	return b
}

func (b *Link) Parameters(v map[string]interface{}) *Link {
	b.target.Parameters = v
	return b
}

func (b *Link) RequestBody(v interface{}) *Link {
	b.target.RequestBody = v
	return b
}

func (b *Link) Description(v string) *Link {
	b.target.Description = v
	return b
}

func (b *Link) Server(v *entity.Server) *Link {
	b.target.Server = v
	return b
}
