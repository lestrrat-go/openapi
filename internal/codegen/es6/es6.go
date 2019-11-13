package es6

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/lestrrat-go/openapi/internal/codegen/common"
	"github.com/lestrrat-go/openapi/internal/codegen/restclient"
	"github.com/lestrrat-go/openapi/internal/stringutil"
	"github.com/lestrrat-go/openapi/openapi2"
	"github.com/pkg/errors"
)

func FileName(s string) string {
	return stringutil.Snake(s)
}

func ClassName(s string) string {
	return stringutil.Camel(stringutil.Snake(s))
}

func MethodName(s string) string {
	return stringutil.LowerCamel(stringutil.Snake(s))
}

func FieldName(s string) string {
	return stringutil.LowerCamel(stringutil.Snake(s))
}

func CallObjectName(oper openapi2.Operation) string {
	return ClassName(restclient.CallObjectName(oper))
}

func CallMethodName(oper openapi2.Operation) string {
	if name := restclient.CallMethodName(oper); name != "" {
		return MethodName(name)
	}
	return ""
}

func DumpCode(dst io.Writer, src io.Reader) {
	common.DumpCode(dst, src)
}

func WritePreamble(dst io.Writer, pkg string) error {
	fmt.Fprintf(dst, "// @flow")
	fmt.Fprintf(dst, "\n\n// This file was automatically generated.")
	fmt.Fprintf(dst, "\n// DO NOT EDIT MANUALLY. All changes will be lost\n")
	return nil
}

func WriteFormattedToFile(fn string, code []byte) error {
	if err := common.WriteToFile(fn, code); err != nil {
		return errors.Wrap(err, `failed to write to file`)
	}

	cmd := exec.Command("prettier", "--write", fn)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, `failed to execute "prettier" tool`)
	}

	return nil
}
