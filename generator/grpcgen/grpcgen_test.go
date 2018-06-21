package grpcgen_test

import (
	"context"
	"os"
	"testing"

	"github.com/lestrrat-go/openapi/generator/grpcgen"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Skip()

	f, err := os.Open("cats.yaml")
	if !assert.NoError(t, err, "reading spec file should succeed") {
		return
	}
	defer f.Close()

	spec, err := openapi.ParseYAML(f)
	if !assert.NoError(t, err, "openapi.ParseYAML should succeed") {
		return
	}

	if !assert.NoError(t, grpcgen.Generate(context.Background(), spec), `generate should succeed`) {
		return
	}
}
