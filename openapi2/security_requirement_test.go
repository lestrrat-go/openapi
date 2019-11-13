package openapi2_test

import (
	"testing"

	"github.com/lestrrat-go/openapi/openapi2"
	"github.com/stretchr/testify/assert"
)

func TestSecurityRequirementJSON(t *testing.T) {
	const src = `{
  "foo": ["AAA", "BBB", "CCC"],
	"bar": ["DDD"],
	"baz": []
}`

	var sr openapi2.SecurityRequirement
	if !assert.NoError(t, openapi2.SecurityRequirementFromJSON([]byte(src), &sr), `unmarshaling JSON should work`) {
		return
	}

	t.Logf("%#v", sr)
}
