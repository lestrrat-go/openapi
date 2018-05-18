package builder

import "github.com/lestrrat-go/openapi/v3/entity"

func (b *Operation) Tag(s string) *Operation {
	b.target.Tags = append(b.target.Tags, s)
	return b
}

func (b *Operation) Parameter(p *entity.Parameter) *Operation {
	b.target.Parameters = append(b.target.Parameters, p)
	return b
}
