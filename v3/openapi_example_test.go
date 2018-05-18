package openapi_test

import (
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/lestrrat-go/openapi/v3/builder"
	"github.com/lestrrat-go/openapi/v3/entity"
)

func ExampleBuild() {
	b := builder.New()

	errSchema := b.NewSchema().
		Reference("#/components/schemas/Error").
		Build()
	errResponse := b.NewResponse("unexpected error").
		Content("application/json",
			b.NewMediaType().
				Schema(errSchema).
				Build(),
		).
		Build()

	components := b.NewComponents().
		Schema("Pet",
			b.NewSchema().
				Required([]string{"id", "name"}).
				Property("id", b.NewSchema().
					Type(entity.Integer).
					Format("int64").
					Build(),
				).
				Property("name", b.NewSchema().
					Type(entity.String).
					Build(),
				).
				Property("tag", b.NewSchema().
					Type(entity.String).
					Build(),
				).
				Build(),
		).
		Schema("Pets",
			b.NewSchema().
				Type(entity.Array).
				Items(b.NewSchema().
					Reference("#/components/schemas/Pet").
					Build(),
				).
				Build(),
		).
		Schema("Error",
			b.NewSchema().
				Required([]string{"code", "message"}).
				Property("code", b.NewSchema().
					Type(entity.Integer).
					Format("int32").
					Build(),
				).
				Property("message", b.NewSchema().
					Type(entity.String).
					Build(),
				).
				Build(),
		).
		Build()

	petsPath := b.NewPathItem().
		Post(
			b.NewOperation(
				b.NewResponses().
					StatusCode(
						"201",
						b.NewResponse("Null response").
							Build(),
					).
					Default(errResponse).
					Build(),
			).
				Summary("Create a pet").
				OperationID("createPets").
				Tag("pets").
				Build(),
		).
		Get(
			b.NewOperation(
				b.NewResponses().
					StatusCode(
						"200",
						b.NewResponse("An paged array of pets").
							Header("x-next",
								b.NewHeader().
									Description("A link to the next page of responses").
									Schema(
										b.NewSchema().
											Type("string").
											Build(),
									).
									Build(),
							).
							Content("application/json",
								b.NewMediaType().
									Schema(
										b.NewSchema().
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
					b.NewParameter("limit", entity.InQuery).
						Description("How many items to return at one time (max 100)").
						Schema(b.NewSchema().
							Type(entity.Integer).
							Format("int32").
							Build(),
						).
						Build(),
				).
				Build(),
		).Build()

	petsIDPath := b.NewPathItem().
		Get(
			b.NewOperation(
				b.NewResponses().
					StatusCode(
						"200",
						b.NewResponse("Expected response to a valid request").
							Content("application/json",
								b.NewMediaType().
									Schema(
										b.NewSchema().
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
					b.NewParameter("petId", entity.InPath).
						Description("The id of the pet to retrieve").
						Schema(
							b.NewSchema().Type(entity.String).Build(),
						).
						Build(),
				).
				Build(),
		).
		Build()

	o := b.NewOpenAPI(
		b.NewInfo("Swagger Petstore").
			Version("1.0.0").
			License(
				b.NewLicense("MIT").Build(),
			).Build(),
		b.NewPaths().
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
