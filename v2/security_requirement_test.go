package openapi_test

import (
	"testing"

	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/stretchr/testify/assert"
)

func TestSecurityRequirementJSON(t *testing.T) {
	const src = `{
  "foo": ["AAA", "BBB", "CCC"],
	"bar": ["DDD"],
	"baz": []
}`

	var sr openapi.SecurityRequirement
	if !assert.NoError(t, openapi.SecurityRequirementFromJSON([]byte(src), &sr), `unmarshaling JSON should work`) {
		return
	}

	t.Logf("%#v", sr)
}
