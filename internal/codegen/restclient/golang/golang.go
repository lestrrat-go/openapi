package golang

import (
	"github.com/lestrrat-go/openapi/internal/codegen/golang"
	"github.com/lestrrat-go/openapi/internal/codegen/restclient"
	"github.com/lestrrat-go/openapi/openapi2"
)

func CallObjectName(oper openapi2.Operation) string {
	return golang.ExportedName(restclient.CallObjectName(oper))
}

func CallMethodName(oper openapi2.Operation) string {
	if name := restclient.CallMethodName(oper); name != "" {
		return golang.ExportedName(name)
	}
	return ""
}
