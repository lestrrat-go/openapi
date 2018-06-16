package golang

import (
	"github.com/lestrrat-go/openapi/internal/codegen/golang"
	"github.com/lestrrat-go/openapi/internal/codegen/restclient"
	openapi "github.com/lestrrat-go/openapi/v2"
)

func CallObjectName(oper openapi.Operation) string {
	return golang.ExportedName(restclient.CallObjectName(oper))
}

func CallMethodName(oper openapi.Operation) string {
	if name := restclient.CallMethodName(oper); name != "" {
		return golang.ExportedName(name)
	}
	return ""
}
