package grpc_test

import (
	"testing"

	"github.com/lestrrat-go/openapi/internal/codegen/grpc"
	"github.com/stretchr/testify/assert"
)

func TestFieldName(t *testing.T) {
	var data = []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "foo-bar-baz",
			Expected: "foo_bar_baz",
		},
	}

	for _, c := range data {
		if !assert.Equal(t, c.Expected, grpc.FieldName(c.Input), "expected %s", c.Expected) {
			return
		}
	}
}
