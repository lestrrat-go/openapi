package codegen_test

import (
	"testing"

	"github.com/lestrrat-go/openapi/internal/codegen"
	"github.com/stretchr/testify/assert"
)

func TestExportedName(t *testing.T) {
	var tests = []struct {
		Expected string
		Input    string
	}{
		{Expected: "FooXmlBlah", Input: "fooXMLBlah"},
	}

	for _, test := range tests {
		t.Run(test.Expected, func(t *testing.T) {
			if !assert.Equal(t, test.Expected, codegen.ExportedName(test.Input), `names should match`) {
				return
			}
		})
	}
}
