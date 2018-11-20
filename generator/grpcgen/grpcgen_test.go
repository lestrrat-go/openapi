package grpcgen_test

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/lestrrat-go/openapi/generator/grpcgen"
	openapi "github.com/lestrrat-go/openapi/v2"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Skip()

	f, err := os.Open("cats.yaml")
	if !assert.NoError(t, err, "reading spec file should succeed") {
		return
	}
	defer f.Close()

	spec, err := openapi.ParseYAML(f)
	if !assert.NoError(t, err, "openapi.ParseYAML should succeed") {
		return
	}

	if !assert.NoError(t, grpcgen.Generate(context.Background(), spec), `generate should succeed`) {
		return
	}
}

func TestGenerateAllOf(t *testing.T) {
	const src = `
swagger: "2.0"
info:
  version: "1.0.0"
  title: "definition using allOf"
  description: "blah"
definitions:
  foo:
    type: object
    properties:
      one:
        type: string
      two:
        type: string
  bar:
    type: object
    properties:
      three:
        type: integer
      four:
        type: integer
paths:
  "/":
    get:
      operationId: "root"
      description: "root"
      parameters:
        - name: param
          in: body
          schema:
            allOf:
              - $ref: "#/definitions/foo"
              - $ref: "#/definitions/bar"
      responses:
        default:
          description: success
`

	spec, err := openapi.ParseYAML(strings.NewReader(src))
	if !assert.NoError(t, err, "openapi.ParseYAML should succeed") {
		return
	}

	var buf bytes.Buffer
	if !assert.NoError(t, grpcgen.Generate(context.Background(), spec, grpcgen.WithDestination(&buf)), `generate should succeed`) {
		return
	}

	const expected = `syntax = "proto3";

package myapp;

import "google/protobuf/empty.proto";

message Bar {
  int32 four = 1;
  int32 three = 2;
}

message Foo {
  string one = 1;
  string two = 2;
}

message RootRequest {
  message Param {
    int32 four = 1;
    string one = 2;
    int32 three = 3;
    string two = 4;
  }
  Param param = 1;
}

service Myapp {
  // root
  rpc Root(RootRequest) returns (google.protobuf.Empty) {
  }
}`

	if !assert.Equal(t, expected, buf.String(), "should match") {
		return
	}
}
