package openapi_test

import (
	"testing"

	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/stretchr/testify/assert"
)

func TestValidateInfo(t *testing.T) {
	t.Run("Complete info", func(t *testing.T) {
		const src = `{
"title": "Swagger Sample App",
"description": "This is a sample server Petstore server.",
"termsOfService": "http://swagger.io/terms/",
"contact": {
  "name": "API Support",
  "url": "http://www.swagger.io/support",
  "email": "support@swagger.io"
},
"license": {
  "name": "Apache 2.0",
  "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
},
"version": "1.0.1"
}`
		var info openapi.Info
		if !assert.NoError(t, openapi.InfoFromJSON([]byte(src), &info), "reading from JSON should succeed") {
			return
		}

		if !assert.NoError(t, info.Validate(true), "validation should succeed") {
			return
		}
	})
	t.Run("Missing title", func(t *testing.T) {
		const src = `{
"description": "This is a sample server Petstore server.",
"termsOfService": "http://swagger.io/terms/",
"contact": {
  "name": "API Support",
  "url": "http://www.swagger.io/support",
  "email": "support@swagger.io"
},
"license": {
  "name": "Apache 2.0",
  "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
},
"version": "1.0.1"
}`
		var info openapi.Info
		if !assert.NoError(t, openapi.InfoFromJSON([]byte(src), &info), "reading from JSON should succeed") {
			return
		}
		if !assert.Error(t, info.Validate(true), "validation should fail") {
			return
		}
	})
	t.Run("Missing version", func(t *testing.T) {
		const src = `{
"title": "Swagger Sample App",
"description": "This is a sample server Petstore server.",
"termsOfService": "http://swagger.io/terms/",
"contact": {
  "name": "API Support",
  "url": "http://www.swagger.io/support",
  "email": "support@swagger.io"
},
"license": {
  "name": "Apache 2.0",
  "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
}
}`
		var info openapi.Info
		if !assert.NoError(t, openapi.InfoFromJSON([]byte(src), &info), "reading from JSON should succeed") {
			return
		}
		if !assert.Error(t, info.Validate(true), "validation should fail") {
			return
		}
	})
}

func TestValidateLicense(t *testing.T) {
	t.Run("Complete license", func(t *testing.T) {
		const src = `{
  "name": "Apache 2.0",
  "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
}`
		var v openapi.License
		if !assert.NoError(t, openapi.LicenseFromJSON([]byte(src), &v), "reading from JSON should succeed") {
			return
		}

		if !assert.NoError(t, v.Validate(true), "validation should succeed") {
			return
		}
	})
	t.Run("Missing name", func(t *testing.T) {
		const src = `{
  "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
}`
		var v openapi.License
		if !assert.NoError(t, openapi.LicenseFromJSON([]byte(src), &v), "reading from JSON should succeed") {
			return
		}
		if !assert.Error(t, v.Validate(true), "validation should fail") {
			return
		}
	})
	t.Run("Invalid url", func(t *testing.T) {
		// Boy, it's incredibly hard to find a URL pattern that actually fails to
		// parse. Here we just grab something from https://golang.org/src/net/url/url_test.go
		const src = `{
  "name": "API Support",
  "url": "http://us\ner:pass\nword@foo.com/"
}`
		var v openapi.License
		if !assert.NoError(t, openapi.LicenseFromJSON([]byte(src), &v), "reading from JSON should succeed") {
			return
		}
		if !assert.Error(t, v.Validate(true), "validation should fail") {
			return
		}
	})
}

func TestValidatePaths(t *testing.T) {
	t.Run("Invalid path pattern", func(t *testing.T) {
		const src = `{
	"relative/path": {
	}
}`
		var v openapi.Paths
		if !assert.NoError(t, openapi.PathsFromJSON([]byte(src), &v), "reading from JSON should succeed") {
			return
		}
		if !assert.Error(t, v.Validate(true), "validation should fail") {
			return
		}
	})
}

