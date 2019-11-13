package openapi3_test

import (
	"testing"

	"github.com/lestrrat-go/openapi/openapi3"
	"github.com/stretchr/testify/assert"
)

func TestOperationValidate(t *testing.T) {
	t.Run("missing responses", func(t *testing.T) {
		const src = `{
"operationId": "HelloWorld",
"description": "A hello world operation"
}`
		var oper openapi3.Operation
		if !assert.NoError(t, openapi3.OperationFromJSON([]byte(src), &oper), "unmarshal should not fail") {
			return
		}

		if !assert.Error(t, oper.Validate(false), "oper.Validate should fail") {
			return
		}
	})
}
