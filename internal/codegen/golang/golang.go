package golang

// Package golang contains tools to work with Go code

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/lestrrat-go/openapi/internal/codegen/common"
	"github.com/lestrrat-go/openapi/internal/stringutil"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

func CallObjectName(oper openapi.Operation) string {
	return ExportedName(common.CallObjectName(oper))
}

func CallMethodName(oper openapi.Operation) string {
	if name := common.CallMethodName(oper); name != "" {
		return ExportedName(name)
	}
	return ""
}

func DumpCode(dst io.Writer, src io.Reader) {
	common.DumpCode(dst, src)
}

func ExportedName(s string) string {
	switch s {
	case "defaultValue":
		return "Default"
	case "url":
		return "URL"
	case "xml":
		return "XML"
	case "typ":
		return "Type"
	}

	s = stringutil.Camel(stringutil.Snake(s))
	s = strings.Replace(s, "Id", "ID", -1)
	s = strings.Replace(s, "Url", "URL", -1)
	return s
}

func UnexportedName(s string) string {
	switch s {
	case "Default":
		return "defaultValue"
	case "URL":
		return "url"
	case "XML":
		return "xml"
	case "Type":
		return "typ"
	}

	s = stringutil.LcFirst(stringutil.Camel(stringutil.Snake(s)))
	s = strings.Replace(s, "Id", "ID", -1)
	s = strings.Replace(s, "Url", "URL", -1)
	return s
}

func WritePreamble(dst io.Writer, pkg string) error {
	fmt.Fprintf(dst, "\n\npackage %s", pkg)
	fmt.Fprintf(dst, "\n\n// This file was automatically generated.")
	fmt.Fprintf(dst, "\n// DO NOT EDIT MANUALLY. All changes will be lost\n")
	return nil
}

var importDummies = map[string]string{
	"github.com/pkg/errors": "errors.Cause",
	"encoding/json":         "json.Unmarshal",
	"log":                   "log.Printf",
	"net/url":               "url.Parse",
	"sort":                  "sort.Strings",
	"strconv":               "strconv.Quote",
}

func WriteImports(dst io.Writer, libs ...string) error {
	l := len(libs)
	if l == 0 {
		return nil
	}

	fmt.Fprintf(dst, "\n\nimport ")
	if l == 1 {
		fmt.Fprintf(dst, "%s", strconv.Quote(libs[0]))
		return nil
	}

	// first, separate out stdlib and everything else
	var stdlibs []string
	var extlibs []string
	for _, lib := range libs {
		i := strings.IndexByte(lib, '/')
		if i == -1 {
			stdlibs = append(stdlibs, lib)
			continue
		}

		i = strings.IndexByte(lib[:i], '.')
		if i == -1 {
			stdlibs = append(stdlibs, lib)
		} else {
			extlibs = append(extlibs, lib)
		}
	}

	sort.Strings(stdlibs)
	sort.Strings(extlibs)

	fmt.Fprintf(dst, "(")

	// Start with stdlibs
	for _, lib := range stdlibs {
		fmt.Fprintf(dst, "\n%s", strconv.Quote(lib))
	}

	if len(extlibs) > 0 {
		if len(stdlibs) > 0 {
			fmt.Fprintf(dst, "\n")
		}

		for _, lib := range extlibs {
			fmt.Fprintf(dst, "\n%s", strconv.Quote(lib))
		}
	}

	fmt.Fprintf(dst, "\n)")

	// check to see if we need dummies
	var buf bytes.Buffer
	for _, lib := range libs {
		if v, ok := importDummies[lib]; ok {
			fmt.Fprintf(&buf, "\nvar _ = %s", v)
		}
	}

	if buf.Len() > 0 {
		fmt.Fprintf(dst, "\n")
		buf.WriteTo(dst)
	}
	return nil
}

func WriteFormattedToFile(fn string, code []byte) error {
	formatted, err := format.Source(code)
	if err != nil {
		return errors.Wrap(err, `failed to format source code`)
	}

	f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, `failed to open file for writing`)
	}
	defer f.Close()

	if _, err := f.Write(formatted); err != nil {
		return errors.Wrap(err, `failed to write to file`)
	}
	return nil
}
