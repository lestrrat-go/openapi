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
		Build()
	errResponse := openapi.NewResponse("unexpected error").
		Content("application/json",
			openapi.NewMediaType().
				Schema(errReference).
				Build(),
		).
		Build()
	petSchema :=
		openapi.NewSchema().
			Required([]string{"id", "name"}).
			Property("id", openapi.NewSchema().
				Type(openapi.Integer).
				Format("int64").
				Build(),
			).
			Property("name", openapi.NewSchema().
				Type(openapi.String).
				Build(),
			).
			Property("tag", openapi.NewSchema().
				Type(openapi.String).
				Build(),
			).
			Build()
	petsSchema := openapi.NewSchema().
		Type(openapi.Array).
		Items(openapi.NewSchema().
			Reference("#/components/schemas/Pet").
			Build(),
		).
		Build()
	errSchema := openapi.NewSchema().
		Required([]string{"code", "message"}).
		Property("code", openapi.NewSchema().
			Type(openapi.Integer).
			Format("int32").
			Build(),
		).
		Property("message", openapi.NewSchema().
			Type(openapi.String).
			Build(),
		).
		Build()

	components := openapi.NewComponents().
		Schema("Pet", petSchema).
		Schema("Pets", petsSchema).
		Schema("Error", errSchema).
		Build()

	petsPostOperation := openapi.NewOperation(
		openapi.NewResponses().
			Response(
				"201",
				openapi.NewResponse("Null response").
					Build(),
			).
			Default(errResponse).
			Build(),
	).
		Summary("Create a pet").
		OperationID("createPets").
		Tag("pets").
		Build()

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
									Build(),
							).
							Build(),
					).
					Content("application/json",
						openapi.NewMediaType().
							Schema(
								openapi.NewSchema().
									Reference("#/components/schemas/Pets").
									Build(),
							).
							Build(),
					).
					Build(),
			).
			Default(errResponse).
			Build(),
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
					Build(),
				).
				Build(),
		).
		Build()

	petsPath := openapi.NewPathItem().
		Post(petsPostOperation).
		Get(petsGetOperation).
		Build()

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
											Build(),
									).
									Build(),
							).
							Build(),
					).
					Default(errResponse).
					Build(),
			).
				Summary("Info for a specific pet").
				OperationID("showPetById").
				Tag("pets").
				Parameter(
					openapi.NewParameter("petId", openapi.InPath).
						Description("The id of the pet to retrieve").
						Schema(
							openapi.NewSchema().Type(openapi.String).Build(),
						).
						Build(),
				).
				Build(),
		).
		Build()

	o := openapi.NewOpenAPI(
		openapi.NewInfo("Swagger Petstore").
			Version("1.0.0").
			License(
				openapi.NewLicense("MIT").Build(),
			).Build(),
		openapi.NewPaths().
			Path("/pets", petsPath).
			Path("/pets/{petId}", petsIDPath).
			Build(),
	).
		Components(components).
		Build()

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

	buf, err := yaml.Marshal(spec.Info())
	if err != nil {
		log.Printf("%s", err)
		return
	}

	for pathIter := spec.Paths().Items(); pathIter.Next(); {
		p := pathIter.Item()
		for operIter := p.Operations(); operIter.Next(); {
			log.Printf("%s", operIter.Operation().Verb())
		}
	}

	os.Stdout.Write(buf)

	// OUTPUT:
}
