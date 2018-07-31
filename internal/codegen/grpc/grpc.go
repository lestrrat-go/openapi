package grpc

import (
	"strings"

	"github.com/lestrrat-go/openapi/internal/codegen/golang"
)

var fieldNameReplacer = strings.NewReplacer(
	"-", "_",
	" ", "_",
)

func FieldName(s string) string {
	// All non-alphanumeric characters except for "_" are
	// converted to "_"
	// (note: too lazy to make list right now...
	return fieldNameReplacer.Replace(s)
}

func MessageName(s string) string {
	return golang.ExportedName(s)
}
