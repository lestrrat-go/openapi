package builder

import "github.com/lestrrat-go/openapi/v3/entity"

func (b *Paths) Path(path string, item *entity.PathItem) *Paths {
	if b.target.Paths == nil {
		b.target.Paths = make(map[string]*entity.PathItem)
	}

	b.target.Paths[path] = item
	return b
}
