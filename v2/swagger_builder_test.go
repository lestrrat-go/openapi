package openapi_test

import (
	"testing"

	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/stretchr/testify/assert"
)

func TestSwaggerValidate(t *testing.T) {
	t.Run("invalid info", func(t *testing.T) {
		_, err := openapi.NewSwagger(nil, petsPaths).Do()
		if !assert.Error(t, err, `info = nil should result in error`) {
			return
		}
	})
	t.Run("valid host", func(t *testing.T) {
		for _, h := range []string{"example.com", "example.com:8080"} {
			_, err := openapi.NewSwagger(petsInfo, petsPaths).
				Host(h).
				Do()
			if !assert.NoError(t, err, `valid host %s should NOT result in error`, h) {
				return
			}
		}
	})
	t.Run("invalid host", func(t *testing.T) {
		for _, h := range []string{"https://example.com", "example.com/foo/bar"} {
			_, err := openapi.NewSwagger(petsInfo, petsPaths).
				Host(h).
				Do()
			if !assert.Error(t, err, `invalid host %s should result in error`, h) {
				return
			}
		}
	})
	t.Run("invalid basePath", func(t *testing.T) {
		_, err := openapi.NewSwagger(petsInfo, petsPaths).
			BasePath("foo").
			Do()
		if !assert.Error(t, err, `invalid basePath %s should result in error`, "foo") {
			return
		}
	})
	t.Run("invalid paths", func(t *testing.T) {
		_, err := openapi.NewSwagger(petsInfo, nil).
			Do()
		if !assert.Error(t, err, `paths = nil should result in error`) {
			return
		}
	})
}
