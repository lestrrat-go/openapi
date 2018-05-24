package openapi_test

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/lestrrat-go/openapi/v3"
	client "github.com/lestrrat-go/openapi/v3/generator/client"
)

func ExampleBuild() {
	errReference := openapi.NewSchema().
		Reference("#/components/schemas/Error").
		Do()
	errResponse := openapi.NewResponse("unexpected error").
		Content("application/json",
			openapi.NewMediaType().
				Schema(errReference).
				Do(),
		).
		Do()
	petSchema :=
		openapi.NewSchema().
			Required([]string{"id", "name"}).
			Property("id", openapi.NewSchema().
				Type(openapi.Integer).
				Format("int64").
				Do(),
			).
			Property("name", openapi.NewSchema().
				Type(openapi.String).
				Do(),
			).
			Property("tag", openapi.NewSchema().
				Type(openapi.String).
				Do(),
			).
			Do()
	petsSchema := openapi.NewSchema().
		Type(openapi.Array).
		Items(openapi.NewSchema().
			Reference("#/components/schemas/Pet").
			Do(),
		).
		Do()
	errSchema := openapi.NewSchema().
		Required([]string{"code", "message"}).
		Property("code", openapi.NewSchema().
			Type(openapi.Integer).
			Format("int32").
			Do(),
		).
		Property("message", openapi.NewSchema().
			Type(openapi.String).
			Do(),
		).
		Do()

	components := openapi.NewComponents().
		Schema("Pet", petSchema).
		Schema("Pets", petsSchema).
		Schema("Error", errSchema).
		Do()

	petsPostOperation := openapi.NewOperation(
		openapi.NewResponses().
			Response(
				"201",
				openapi.NewResponse("Null response").
					Do(),
			).
			Default(errResponse).
			Do(),
	).
		Summary("Create a pet").
		OperationID("createPets").
		Tag("pets").
		Do()

	petsGetOperation := openapi.NewOperation(
		openapi.NewResponses().
			Response(
				"200",
				openapi.NewResponse("An paged array of pets").
					Header("x-next",
						openapi.NewHeader().
							Description("A link to the next page of responses").
							Schema(
								openapi.NewSchema().
									Type("string").
									Do(),
							).
							Do(),
					).
					Content("application/json",
						openapi.NewMediaType().
							Schema(
								openapi.NewSchema().
									Reference("#/components/schemas/Pets").
									Do(),
							).
							Do(),
					).
					Do(),
			).
			Default(errResponse).
			Do(),
	).
		Summary("List all pets").
		OperationID("listPets").
		Tag("pets").
		Parameter(
			openapi.NewParameter("limit", openapi.InQuery).
				Description("How many items to return at one time (max 100)").
				Schema(openapi.NewSchema().
					Type(openapi.Integer).
					Format("int32").
					Do(),
				).
				Do(),
		).
		Do()

	petsPath := openapi.NewPathItem().
		Post(petsPostOperation).
		Get(petsGetOperation).
		Do()

	petsIDPath := openapi.NewPathItem().
		Get(
			openapi.NewOperation(
				openapi.NewResponses().
					Response(
						"200",
						openapi.NewResponse("Expected response to a valid request").
							Content("application/json",
								openapi.NewMediaType().
									Schema(
										openapi.NewSchema().
											Reference("#/components/schemas/Pets").
											Do(),
									).
									Do(),
							).
							Do(),
					).
					Default(errResponse).
					Do(),
			).
				Summary("Info for a specific pet").
				OperationID("showPetById").
				Tag("pets").
				Parameter(
					openapi.NewParameter("petId", openapi.InPath).
						Description("The id of the pet to retrieve").
						Schema(
							openapi.NewSchema().Type(openapi.String).Do(),
						).
						Do(),
				).
				Do(),
		).
		Do()

	o := openapi.NewOpenAPI(
		openapi.NewInfo("Swagger Petstore").
			Version("1.0.0").
			License(
				openapi.NewLicense("MIT").Do(),
			).Do(),
		openapi.NewPaths().
			Path("/pets", petsPath).
			Path("/pets/{petId}", petsIDPath).
			Do(),
	).
		Components(components).
		Do()

	buf, err := yaml.Marshal(o)
	if err != nil {
		log.Printf("%s", err)
		return
	}
	os.Stdout.Write(buf)
	// OUTPUT:
	// components:
	//   schemas:
	//     Error:
	//       properties:
	//         code:
	//           format: int32
	//           type: integer
	//         message:
	//           type: string
	//       required:
	//       - code
	//       - message
	//     Pet:
	//       properties:
	//         id:
	//           format: int64
	//           type: integer
	//         name:
	//           type: string
	//         tag:
	//           type: string
	//       required:
	//       - id
	//       - name
	//     Pets:
	//       items:
	//         $ref: '#/components/schemas/Pet'
	//       type: array
	// info:
	//   license:
	//     name: MIT
	//   title: Swagger Petstore
	//   version: 1.0.0
	// openapi: 3.0.1
	// paths:
	//   /pets:
	//     get:
	//       operationId: listPets
	//       parameters:
	//       - description: How many items to return at one time (max 100)
	//         in: query
	//         name: limit
	//         schema:
	//           format: int32
	//           type: integer
	//       responses:
	//         "200":
	//           content:
	//             application/json:
	//               schema:
	//                 $ref: '#/components/schemas/Pets'
	//           description: An paged array of pets
	//           headers:
	//             x-next:
	//               description: A link to the next page of responses
	//               schema:
	//                 type: string
	//         default:
	//           content:
	//             application/json:
	//               schema:
	//                 $ref: '#/components/schemas/Error'
	//           description: unexpected error
	//       summary: List all pets
	//       tags:
	//       - pets
	//     post:
	//       operationId: createPets
	//       responses:
	//         "201":
	//           description: Null response
	//         default:
	//           content:
	//             application/json:
	//               schema:
	//                 $ref: '#/components/schemas/Error'
	//           description: unexpected error
	//       summary: Create a pet
	//       tags:
	//       - pets
	//   /pets/{petId}:
	//     get:
	//       operationId: showPetById
	//       parameters:
	//       - description: The id of the pet to retrieve
	//         in: path
	//         name: petId
	//         required: true
	//         schema:
	//           type: string
	//       responses:
	//         "200":
	//           content:
	//             application/json:
	//               schema:
	//                 $ref: '#/components/schemas/Pets'
	//           description: Expected response to a valid request
	//         default:
	//           content:
	//             application/json:
	//               schema:
	//                 $ref: '#/components/schemas/Error'
	//           description: unexpected error
	//       summary: Info for a specific pet
	//       tags:
	//       - pets
}

func ExampleParse() {
	const src = `
info:
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
	if err != nil {
		log.Printf("%s", err)
		return
	}

	buf, err := yaml.Marshal(spec)
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
		return
	}
	os.Stdout.Write(buf)

	for pathIter := spec.Paths().Paths(); pathIter.Next(); {
		_, p := pathIter.Item()
		for operIter := p.Operations(); operIter.Next(); {
			oper := operIter.Item()
			openapi.MutateOperation(oper).
				OperationID("foo").
				Do()
		}
	}

	// OUTPUT:
	// info:
	//   license:
	//     name: MIT
	//   title: Swagger Petstore
	//   version: 1.0.0
	// openapi: 3.0.1
	// paths:
	//   /pets:
	//     get:
	//       operationId: listPets
	//       parameters:
	//       - description: How many items to return at one time (max 100)
	//         in: query
	//         name: limit
	//         schema:
	//           format: int32
	//           type: integer
	//       responses:
	//         "200":
	//           content:
	//             application/json:
	//               schema:
	//                 $ref: '#/components/schemas/Pets'
	//           description: An paged array of pets
	//           headers:
	//             x-next:
	//               description: A link to the next page of responses
	//               schema:
	//                 type: string
	//         default:
	//           content:
	//             application/json:
	//               schema:
	//                 $ref: '#/components/schemas/Error'
	//           description: unexpected error
	//       summary: List all pets
	//       tags:
	//       - pets
	//     post:
	//       operationId: createPets
	//       responses:
	//         "201":
	//           description: Null response
	//         default:
	//           content:
	//             application/json:
	//               schema:
	//                 $ref: '#/components/schemas/Error'
	//           description: unexpected error
	//       summary: Create a pet
	//       tags:
	//       - pets
	//   /pets/{petId}:
	//     get:
	//       operationId: showPetById
	//       parameters:
	//       - description: The id of the pet to retrieve
	//         in: path
	//         name: petId
	//         required: true
	//         schema:
	//           type: string
	//       responses:
	//         "200":
	//           content:
	//             application/json:
	//               schema:
	//                 $ref: '#/components/schemas/Pets'
	//           description: Expected response to a valid request
	//         default:
	//           content:
	//             application/json:
	//               schema:
	//                 $ref: '#/components/schemas/Error'
	//           description: unexpected error
	//       summary: Info for a specific pet
	//       tags:
	//       - pets
}

func ExampleGenerateClient() {
	f, err := os.Open(filepath.Join("..", "spec", "examples", "v3.0", "petstore-expanded.yaml"))
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
		return
	}
	defer f.Close()

	spec, err := openapi.ParseYAML(f)
	if err != nil {
		os.Stdout.Write([]byte(err.Error()))
		return
	}

	c := client.New()
	if err := c.Generate(context.Background(), spec); err != nil {
		return
	}
	// OUTPUT:
}
