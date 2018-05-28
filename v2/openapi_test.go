package openapi_test

import openapi "github.com/lestrrat-go/openapi/v2"

// These objects exist so that we can reuse them in tests later
var apiSupport openapi.Contact
var petsInfo openapi.Info
var petsLicense openapi.License
var petSchema openapi.Schema
var petListSchema openapi.Schema
var petsGetResponse openapi.Response
var petsResponses openapi.Responses
var petsGetOper openapi.Operation
var petsPathItem openapi.PathItem
var petsPaths openapi.Paths

func init() {
	apiSupport, _ = openapi.NewContact().
		Name("API Support").
		URL("http://www.swagger.io/support").
		Email("support@swagger.io").
		Do()

	petsLicense, _ = openapi.NewLicense("Apache 2.0").
		URL("http://www.apache.org/licenses/LICENSE-2.0.html").
		Do()

	petsInfo, _ = openapi.NewInfo("Swagger Sample App", "1.0.1").
		Description("This is a sample server Petstore server.").
		TermsOfService("http://swagger.io/terms/").
		Contact(apiSupport).
		License(petsLicense).
		Do()

	petSchema, _ = openapi.NewSchema().
		Reference("#/definitions/pet").
		Do()

	petListSchema, _ = openapi.NewSchema().
		Type(openapi.Array).
		Items(petSchema).
		Do()

	petsGetResponse, _ = openapi.NewResponse("A list of pets.").
		Schema(petListSchema).
		Do()

	petsResponses, _ = openapi.NewResponses().
		Response("200", petsGetResponse).
		Do()

	petsGetOper, _ = openapi.NewOperation(petsResponses).
		Description("Returns all pets from the system that the user has access to").
		Produces("application/json").
		Do()

	petsPathItem, _ = openapi.NewPathItem().
		Get(petsGetOper).
		Do()

	petsPaths, _ = openapi.NewPaths().
		Path("/pets", petsPathItem).
		Do()
}
