package openapi_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ghodss/yaml"
	openapi "github.com/lestrrat-go/openapi/v3"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/assert"
)

func TestOpenAPI(t *testing.T) {
	root := filepath.Join("..", "spec", "examples", "v3.0")
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if !assert.NoError(t, err, `reading from %s should succeed`, path) {
			return err
		}
		//		t.Logf("%s", data)
		buf := bytes.NewBuffer(data)
		rdr := bytes.NewReader(buf.Bytes())

		t.Run(path, func(t *testing.T) {
			spec, err := openapi.ParseYAML(rdr)
			if !assert.NoError(t, err, `Decoding spec should succeed`) {
				rdr.Seek(0, 0)
				scanner := bufio.NewScanner(rdr)
				lineno := 1
				for scanner.Scan() {
					t.Logf("%4d: %s", lineno, scanner.Text())
					lineno++
				}
				return
			}
			_ = spec
		})

		return nil
	})
}

func TestQuery(t *testing.T) {
	file := filepath.Join("..", "spec", "examples", "v3.0", "petstore-expanded.yaml")
	src, err := ioutil.ReadFile(file)
	if !assert.NoError(t, err, `Reading from file should succeed`) {
		return
	}

	spec, err := openapi.ParseYAML(bytes.NewReader(src))
	if !assert.NoError(t, err, `Parsing spec should succeed`) {
		return
	}

	query := "#/components/schemas/NewPet"
	v, ok := spec.QueryJSON(query)
	if !assert.True(t, ok, `Querying should succeed`) {
		return
	}

	if !assert.Implements(t, (*openapi.Schema)(nil), v, "The result should be a schema") {
		return
	}
}

func TestParse(t *testing.T) {
	const src = `info:
  license:
    name: MIT
  title: Swagger Petstore
  version: 1.0.0
openapi: 3.0.1
paths:
  /pets:
    get:
      operationId: listPets
      parameters:
      - description: How many items to return at one time (max 100)
        in: query
        name: limit
        schema:
          format: int32
          type: integer
      responses:
        "200":
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Pets'
          description: An paged array of pets
          headers:
            x-next:
              description: A link to the next page of responses
              schema:
                type: string
        default:
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
          description: unexpected error
      summary: List all pets
      tags:
      - pets
    post:
      operationId: createPets
      responses:
        "201":
          description: Null response
        default:
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
          description: unexpected error
      summary: Create a pet
      tags:
      - pets
  /pets/{petId}:
    get:
      operationId: showPetById
      parameters:
      - description: The id of the pet to retrieve
        in: path
        name: petId
        required: true
        schema:
          type: string
      responses:
        "200":
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Pets'
          description: Expected response to a valid request
        default:
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Error'
          description: unexpected error
      summary: Info for a specific pet
      tags:
      - pets
`

	spec, err := openapi.ParseYAML(strings.NewReader(src))
	if !assert.NoError(t, err, `openapi.ParseYAML should succeed`) {
		return
	}

	buf, err := yaml.Marshal(spec)
	if !assert.NoError(t, err, `yaml.Marshal should succeed`) {
		os.Stdout.Write([]byte(err.Error()))
		return
	}

	if !assert.Equal(t, string(buf), src, "output should be the same") {
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(src),
			B:        difflib.SplitLines(string(buf)),
			FromFile: "Expected",
			ToFile:   "Generated",
			Context:  3,
		}
		text, _ := difflib.GetUnifiedDiffString(diff)
		t.Logf("%s", text)
		return
	}

	for pathIter := spec.Paths().Paths(); pathIter.Next(); {
		_, p := pathIter.Item()
		for operIter := p.Operations(); operIter.Next(); {
			oper := operIter.Item()
			openapi.MutateOperation(oper).
				OperationID("foo").
				Do()
		}
	}
}
