package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/pkg/errors"
)

// CallObjectName generates a Call object type name. It is NOT
// normalized, so the caller needs to normalize it
func CallObjectName(oper openapi.Operation) string {
	if operID := oper.OperationID(); operID != "" {
		return operID + "_Call"
	}

	pi := oper.PathItem()
	if pi == nil {
		buf, err := yaml.Marshal(oper)
		if err != nil {
			fmt.Fprintf(os.Stdout, err.Error())
		} else {
			os.Stdout.Write(buf)
		}
		panic("PathItem for operation is nil")
	}

	verb := strings.ToLower(oper.Verb())
	return verb + "_" + oper.PathItem().Path() + "_Call"
}

func DumpCode(dst io.Writer, src io.Reader) {
	scanner := bufio.NewScanner(src)
	lineno := 1
	for scanner.Scan() {
		fmt.Fprintf(dst, "%5d: %s\n", lineno, scanner.Text())
		lineno++
	}
}

func CallMethodName(oper openapi.Operation) string {
	rawMethodName, ok := oper.Extension(`x-call-method-name`)
	if !ok {
		return ""
	}

	if s, ok := rawMethodName.(string); ok {
		return s
	}
	return ""
}

func WriteToFile(fn string, data []byte) error {
	dir := filepath.Dir(fn)
	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.Wrapf(err, `failed to create directory %s`, dir)
		}
	}

	f, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return errors.Wrap(err, `failed to open file for writing`)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return errors.Wrap(err, `failed to write to file`)
	}

	return nil
}
