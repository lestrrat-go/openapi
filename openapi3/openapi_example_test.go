package openapi3_test

import (
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/lestrrat-go/openapi/openapi3"
)

func ExampleBuild() {
	errReference := openapi3.NewSchema().
		Reference("#/components/schemas/Error").
		MustBuild()

	errMediaType := openapi3.NewMediaType().
		Schema(errReference).
		MustBuild()
	errResponse := openapi3.NewResponse("unexpected error").
		Content("application/json", errMediaType).
		MustBuild()

	idProp := openapi3.NewSchema().
		Type(openapi3.Integer).
		Format("int64").
		MustBuild()
	strSchema := openapi3.NewSchema().
		Type(openapi3.String).
		MustBuild()
	petsRef := openapi3.NewSchema().
		Reference("#/components/schemas/Pet").
		MustBuild()
	petSchema :=
		openapi3.NewSchema().
			Required([]string{"id", "name"}).
			Property("id", idProp).
			Property("name", strSchema).
			Property("tag", strSchema).
			MustBuild()
	petsSchema := openapi3.NewSchema().
		Type(openapi3.Array).
		Items(petsRef).
		MustBuild()
	int32Schema := openapi3.NewSchema().
		Type(openapi3.Integer).
		Format("int32").
		MustBuild()
	errSchema := openapi3.NewSchema().
		Required([]string{"code", "message"}).
		Property("code", int32Schema).
		Property("message", strSchema).
		MustBuild()

	components := openapi3.NewComponents().
		Schema("Pet", petSchema).
		Schema("Pets", petsSchema).
		Schema("Error", errSchema).
		MustBuild()

	nullResponse := openapi3.NewResponse("Null response").
		MustBuild()

	petsSchemaRef := openapi3.NewSchema().
		Reference("#/components/schemas/Pets").
		MustBuild()

	o := openapi3.NewOpenAPI(
		openapi3.NewInfo("Swagger Petstore").
			Version("1.0.0").
			License(
				openapi3.NewLicense("MIT").MustBuild(),
			).MustBuild(),
		openapi3.NewPaths().
			Path("/pets", openapi3.NewPathItem().
				Post(openapi3.NewOperation(
					openapi3.NewResponses().
						Response("201", nullResponse).
						Default(errResponse).
						MustBuild(),
				).
					Summary("Create a pet").
					OperationID("createPets").
					Tag("pets").
					MustBuild(),
				).
				Get(openapi3.NewOperation(openapi3.NewResponses().
					Response("200", openapi3.NewResponse("An paged array of pets").
						Header("x-next", openapi3.NewHeader().
							Description("A link to the next page of responses").
							Schema(strSchema).
							MustBuild(),
						).
						Content("application/json", openapi3.NewMediaType().
							Schema(petsSchemaRef).
							MustBuild(),
						).
						MustBuild(),
					).
					Default(errResponse).
					MustBuild(),
				).
					Summary("List all pets").
					OperationID("listPets").
					Tag("pets").
					Parameter(
						openapi3.NewParameter("limit", openapi3.InQuery).
							Description("How many items to return at one time (max 100)").
							Schema(int32Schema).
							MustBuild(),
					).
					MustBuild(),
				).
				MustBuild(),
			).
			Path("/pets/{petId}", openapi3.NewPathItem().
				Get(openapi3.NewOperation(openapi3.NewResponses().
					Response("200", openapi3.NewResponse("Expected response to a valid request").
						Content("application/json", openapi3.NewMediaType().
							Schema(petsSchemaRef).
							MustBuild(),
						).
						MustBuild(),
					).
					Default(errResponse).
					MustBuild(),
				).
					Summary("Info for a specific pet").
					OperationID("showPetById").
					Tag("pets").
					Parameter(
						openapi3.NewParameter("petId", openapi3.InPath).
							Description("The id of the pet to retrieve").
							Schema(strSchema).
							MustBuild(),
					).
					MustBuild(),
				).
				MustBuild(),
			).
			MustBuild(),
	).
		Components(components).
		MustBuild()

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

	spec, err := openapi3.ParseYAML(strings.NewReader(src))
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
			openapi3.MutateOperation(oper).
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
