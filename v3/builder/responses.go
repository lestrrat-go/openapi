package builder

import "github.com/lestrrat-go/openapi/v3/entity"

func (b *Responses) StatusCode(code string, v *entity.Response) *Responses {
	if b.target.StatusCodes == nil {
		b.target.StatusCodes = make(map[string]*entity.Response)
	}
	b.target.StatusCodes[code] = v
	return b
}
