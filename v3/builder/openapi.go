package builder

import "github.com/lestrrat-go/openapi/v3/entity"

func (b *OpenAPI) Tag(tag *entity.Tag) *OpenAPI {
	b.target.Tags = append(b.target.Tags, tag)
	return b
}
