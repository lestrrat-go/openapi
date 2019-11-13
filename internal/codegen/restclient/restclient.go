package restclient

import (
	"fmt"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/lestrrat-go/openapi/openapi2"
)

// CallObjectName generates a Call object type name. It is NOT
// normalized, so the caller needs to normalize it
func CallObjectName(oper openapi2.Operation) string {
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

func CallMethodName(oper openapi2.Operation) string {
	rawMethodName, ok := oper.Extension(`x-call-method-name`)
	if !ok {
		return ""
	}

	if s, ok := rawMethodName.(string); ok {
		return s
	}
	return ""
}
