package openapi_test

import (
	"testing"

	openapi "github.com/lestrrat-go/openapi/v3"
	"github.com/stretchr/testify/assert"
)

func TestOperationValidate(t *testing.T) {
	t.Run("missing responses", func(t *testing.T) {
		const src = `{
"operationId": "HelloWorld",
"description": "A hello world operation"
}`
		var oper openapi.Operation
		if !assert.NoError(t, openapi.OperationFromJSON([]byte(src), &oper), "unmarshal should not fail") {
			return
		}

		if !assert.Error(t, oper.Validate(false), "oper.Validate should fail") {
			return
		}
	})
}
