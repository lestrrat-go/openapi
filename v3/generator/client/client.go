package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
	openapi "github.com/lestrrat-go/openapi/v3"
	"github.com/pkg/errors"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) Generate(ctx context.Context, spec openapi.OpenAPI) error {
	resolver := openapi.NewResolver(spec)

	spec.Resolve(resolver)

	var buf bytes.Buffer
	var dst io.Writer = &buf

	// We need a base server address. If we don't have one, complaing
	var serverURL string
	for siter := spec.Servers(); siter.Next(); {
		server := siter.Item()
		if v := server.URL(); v != "" {
			serverURL = v
			break
		}
	}

	if serverURL == "" {
		return errors.New(`could not find a usable server URL`)
	}

	fmt.Fprintf(dst, "\n\nfunc New() *Client {")
	fmt.Fprintf(dst, "\nreturn &Client{")
	fmt.Fprintf(dst, "\nserver: %s,", strconv.Quote(serverURL))
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (c *Client) formatURL(path string) string {")
	fmt.Fprintf(dst, "\nreturn c.server + path")
	fmt.Fprintf(dst, "\n}")

	for piter := spec.Paths().Paths(); piter.Next(); {
		_, item := piter.Item()
		for opiter := item.Operations(); opiter.Next(); {
			oper := opiter.Item()
			if err := c.GenerateCall(dst, oper); err != nil {
				return errors.Wrapf(err, `failed to generate code for operation %s (%s)`, oper.Path(), oper.Verb())
			}
		}
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.String())
		return errors.Wrap(err, `failed to format source`)
	}

	_ = formatted
	// os.Stdout.Write(formatted)
	return nil
}

func methodName(oper openapi.Operation) string {
	if id := oper.OperationID(); id != "" {
		return strcase.ToCamel(id)
	}

	// otherwise, guess from the given path
	return strcase.ToCamel(strings.ToLower(oper.Verb()) + " " + oper.Path())
}

func paramType(param openapi.Parameter) string {
	switch param.Schema().Type() {
	case openapi.Integer:
		return integerType(param.Schema().Format())
	case openapi.String:
		return "string"
	default:
		return "unimplemented"
	}
}

func integerType(format string) string {
	switch format {
	case "int32":
		return format
	default:
		return "int"
	}
}

func (c *Client) GenerateCall(dst io.Writer, oper openapi.Operation) error {
	mname := methodName(oper)

	// For simplicity, we're only going to support ONE requestBody
	// element per operation. Patches welcome

	var requestBodyType openapi.MediaType
	if rb := oper.RequestBody(); rb != nil {
		for citer := rb.Content(); citer.Next(); {
			_, mt := citer.Item()
			if requestBodyType != nil {
				return errors.New(`currently this tool does not support multiple requestBody types`)
			}
			requestBodyType = mt
		}
	}
	if requestBodyType != nil {
		if s := requestBodyType.Schema(); s.Type() != openapi.Object {
			log.Printf("%+v", requestBodyType)
			json.NewEncoder(os.Stderr).Encode(requestBodyType.Schema())
			return errors.Errorf(`currently this tool only supports requestBody types where type = object (got %s)`, s.Type())
		}
	}

	fmt.Fprintf(dst, "\n\ntype %sCall struct {", mname)
	for piter := oper.Parameters(); piter.Next(); {
		param := piter.Item()
		typ := paramType(param)
		fmt.Fprintf(dst, "\n%s %s", strcase.ToLowerCamel(param.Name()), typ)
	}
	if requestBodyType != nil {
		for piter := requestBodyType.Schema().Properties(); piter.Next(); {
			_, prop := piter.Item()
			fmt.Fprintf(dst, "\n%s %s", prop.Name(), "string")
		}
	}
	fmt.Fprintf(dst, "\n}")

	fmt.Fprintf(dst, "\n\nfunc (c *Client) %s(", mname)
	fmt.Fprintf(dst, ") %sCall {", mname)
	fmt.Fprintf(dst, "\nreturn &%sCall{", mname)
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\n}")

	var hasQuery bool
	for piter := oper.Parameters(); piter.Next(); {
		param := piter.Item()
		if param.In() == openapi.InQuery {
			hasQuery = true
		}
		typ := paramType(param)
		fmt.Fprintf(dst, "\n\nfunc (call *%sCall) %s(%s %s) *%sCall {", mname, strcase.ToCamel(param.Name()), param.Name(), typ, mname)
		fmt.Fprintf(dst, "\ncall.%s = %s", strcase.ToLowerCamel(param.Name()), param.Name())
		fmt.Fprintf(dst, "\nreturn call")
		fmt.Fprintf(dst, "\n}")
	}

	fmt.Fprintf(dst, "\n\nfunc (c *%sCall) Do(ctx context.Context) (%sResponse, error) {", mname, mname)
	fmt.Fprintf(dst, "\nconst basepath = %s", strconv.Quote(oper.Path()))
	fmt.Fprintf(dst, "\nvar path = basepath")
	if hasQuery {
		fmt.Fprintf(dst, "\nquery := url.Values{}")
	}

	for piter := oper.Parameters(); piter.Next(); {
		param := piter.Item()
		switch param.In() {
		case openapi.InPath:
			fmt.Fprintf(dst, "\npath = strings.Replace(path, %s, c.%s, -1)", strconv.Quote("{"+param.Name()+"}"), strcase.ToCamel(param.Name()))
		case openapi.InQuery:
			fmt.Fprintf(dst, "\nquery.Add(%s, c.%s)", strconv.Quote(param.Name()), strcase.ToLowerCamel(param.Name()))
		}
	}

	fmt.Fprintf(dst, "\nu := c.formatURL(path)")
	if hasQuery {
		fmt.Fprintf(dst, "+ `?` + query.Encode()")
	}

	switch oper.Verb() {
	case http.MethodGet:
		fmt.Fprintf(dst, "\nreq, err := http.NewRequest(http.MethodGet, u, nil)")
	case http.MethodPost:
		fmt.Fprintf(dst, "\nvar body io.Reader")
		if rb := oper.RequestBody(); rb != nil {
			fmt.Fprintf(dst, "\nvar buf bytes.Buffer")
			fmt.Fprintf(dst, "\nbody = &buf")

			fmt.Fprintf(dst, "\nif strings.Contains(call.requestBodyMediaType, `application/json`) {")
			fmt.Fprintf(dst, "\nif err := json.NewEncoder(&buf).Encode(call.requestBody); err != nil {")
			fmt.Fprintf(dst, "\nreturn errors.Wrap(err, `failed to encode request body into JSON`)")
			fmt.Fprintf(dst, "\n}")
			fmt.Fprintf(dst, "\n}")
		}
		fmt.Fprintf(dst, "\nreq, err := http.NewRequest(http.MethodPost, u, body)")

	}

	fmt.Fprintf(dst, "\nif err != nil {")
	fmt.Fprintf(dst, "\nreturn nil, errors.Wrap(err, `failed to construct HTTP request`)")
	fmt.Fprintf(dst, "\n}")
	fmt.Fprintf(dst, "\nhttp.DefaultClient.Do(req)")
	fmt.Fprintf(dst, "\n}")

	return nil
}
