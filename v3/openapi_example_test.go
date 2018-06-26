package openapi_test

import (
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/lestrrat-go/openapi/v3"
)

func ExampleBuild() {
	errReference := openapi.NewSchema().
		Reference("#/components/schemas/Error").
		MustBuild()

	errMediaType := openapi.NewMediaType().
		Schema(errReference).
		MustBuild()
	errResponse := openapi.NewResponse("unexpected error").
		Content("application/json", errMediaType).
		MustBuild()

	idProp := openapi.NewSchema().
		Type(openapi.Integer).
		Format("int64").
		MustBuild()
	strSchema := openapi.NewSchema().
		Type(openapi.String).
		MustBuild()
	petsRef := openapi.NewSchema().
		Reference("#/components/schemas/Pet").
		MustBuild()
	petSchema :=
		openapi.NewSchema().
			Required([]string{"id", "name"}).
			Property("id", idProp).
			Property("name", strSchema).
			Property("tag", strSchema).
			MustBuild()
	petsSchema := openapi.NewSchema().
		Type(openapi.Array).
		Items(petsRef).
		MustBuild()
	int32Schema := openapi.NewSchema().
		Type(openapi.Integer).
		Format("int32").
		MustBuild()
	errSchema := openapi.NewSchema().
		Required([]string{"code", "message"}).
		Property("code", int32Schema).
		Property("message", strSchema).
		MustBuild()

	components := openapi.NewComponents().
		Schema("Pet", petSchema).
		Schema("Pets", petsSchema).
		Schema("Error", errSchema).
		MustBuild()

	nullResponse := openapi.NewResponse("Null response").
		MustBuild()

	petsSchemaRef := openapi.NewSchema().
		Reference("#/components/schemas/Pets").
		MustBuild()

	o := openapi.NewOpenAPI(
		openapi.NewInfo("Swagger Petstore").
			Version("1.0.0").
			License(
				openapi.NewLicense("MIT").MustBuild(),
			).MustBuild(),
		openapi.NewPaths().
			Path("/pets", openapi.NewPathItem().
				Post(openapi.NewOperation(
					openapi.NewResponses().
						Response("201", nullResponse).
						Default(errResponse).
						MustBuild(),
				).
					Summary("Create a pet").
					OperationID("createPets").
					Tag("pets").
					MustBuild(),
				).
				Get(openapi.NewOperation(openapi.NewResponses().
					Response("200", openapi.NewResponse("An paged array of pets").
						Header("x-next", openapi.NewHeader().
							Description("A link to the next page of responses").
							Schema(strSchema).
							MustBuild(),
						).
						Content("application/json", openapi.NewMediaType().
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
						openapi.NewParameter("limit", openapi.InQuery).
							Description("How many items to return at one time (max 100)").
							Schema(int32Schema).
							MustBuild(),
					).
					MustBuild(),
				).
				MustBuild(),
			).
			Path("/pets/{petId}", openapi.NewPathItem().
				Get(openapi.NewOperation(openapi.NewResponses().
					Response("200", openapi.NewResponse("Expected response to a valid request").
						Content("application/json", openapi.NewMediaType().
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
						openapi.NewParameter("petId", openapi.InPath).
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
