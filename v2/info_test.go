package openapi_test

import (
	"testing"

	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/stretchr/testify/assert"
)

func TestInfoValidate(t *testing.T) {
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

