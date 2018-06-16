package restclientgen

import (
	"github.com/lestrrat-go/openapi/generator/restclientgen/es6flow"
	"github.com/lestrrat-go/openapi/generator/restclientgen/golang"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

func Generate(spec openapi.Swagger, options ...Option) error {
	if err := spec.Validate(true); err != nil {
		return errors.Wrap(err, `failed to validate spec`)
	}

	target := "go"
	for _, option := range options {
		switch option.Name() {
		case optkeyTarget:
			target = option.Value().(string)
		}
	}

	switch target {
	case "go":
		return golang.Generate(spec, options...)
	case "es6flow":
		return es6flow.Generate(spec, options...)
	default:
		return errors.Errorf(`invalid generation target %s`, target)
	}
	return errors.New(`unreachable`)
}
