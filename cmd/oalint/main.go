package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/lestrrat-go/openapi/openapi2"
	"github.com/lestrrat-go/openapi/openapi3"
	"github.com/pkg/errors"
)

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

func _main() error {
	var file string
	var format string
	var oaVersion string

	flag.StringVar(&oaVersion, "openapi-version", "2.0", "OpenAPI version (2.0 or 3.0.1)")
	flag.StringVar(&format, "format", "yaml", "format (yaml|json) of the input data")
	flag.StringVar(&file, "file", "", "file name to read from (use '-' to indicate stdin)")
	flag.Parse()

	var input io.Reader
	switch file {
	case "":
		return errors.New(`'-file' is required`)
	case "-":
		input = os.Stdin
	default:
		f, err := os.Open(file)
		if err != nil {
			return errors.Wrapf(err, `failed to open file %s`, file)
		}
		defer f.Close()
		input = f
	}

	var spec interface{}
	var parseErr error

	switch oaVersion {
	case "2.0":
		spec, parseErr = parseV2(format, input)
	case "3", "3.0.1":
		spec, parseErr = parseV3(format, input)
	default:
		return errors.Errorf(`unsupported OpenAPI version %s`, oaVersion)
	}

	if parseErr != nil {
		return errors.Wrap(parseErr, `failed to parse input`)
	}

	var output []byte
	var serializeErr error
	switch strings.ToLower(format) {
	case "yaml":
		output, serializeErr = yaml.Marshal(spec)
	case "json":
		// This is silly, but doing multiple marshaling here allows us to
		// encode the entire thing in a sorted key order
		output, serializeErr = json.Marshal(spec)
		if serializeErr == nil {
			var m map[string]interface{}
			if serializeErr = json.Unmarshal(output, &m); serializeErr == nil {
				output, serializeErr = json.MarshalIndent(m, "", "  ")
			}
		}
	default:
		return errors.Errorf(`invalid format %s`, format)
	}

	if serializeErr != nil {
		return errors.Wrap(serializeErr, `failed to serialize`)
	}

	os.Stdout.Write(output)

	return nil
}

func parseV2(format string, input io.Reader) (interface{}, error) {
	switch strings.ToLower(format) {
	case "yaml":
		return openapi2.ParseYAML(input, openapi2.WithValidate(true))
	case "json":
		return openapi2.ParseJSON(input, openapi2.WithValidate(true))
	default:
		return nil, errors.Errorf(`invalid format %s`, format)
	}
}

func parseV3(format string, input io.Reader) (interface{}, error) {
	switch strings.ToLower(format) {
	case "yaml":
		return openapi3.ParseYAML(input, openapi3.WithValidate(true))
	case "json":
		return openapi3.ParseJSON(input, openapi3.WithValidate(true))
	default:
		return nil, errors.Errorf(`invalid format %s`, format)
	}
}
