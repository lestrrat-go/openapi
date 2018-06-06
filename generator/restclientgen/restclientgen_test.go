package restclientgen_test

import (
	"os"
	"testing"

	"github.com/lestrrat-go/openapi/generator/restclientgen"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	if _, err := os.Stat("cats.yaml"); err != nil {
		t.Skip("missing definition.")
	}

	var gen restclientgen.Generator
	f, err := os.Open("cats.yaml")
	if !assert.NoError(t, err, "reading spec file should succeed") {
		return
	}
	defer f.Close()

	spec, err := openapi.ParseYAML(f)
	if !assert.NoError(t, err, "openapi.ParseYAML should succeed") {
		return
	}

	if !assert.NoError(t, gen.Generate(os.Stdout, spec), `generate should succeed`) {
		return
	}
}
