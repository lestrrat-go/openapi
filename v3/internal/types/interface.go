// This file serves as the base template of openapi/v3/interface_gen.go.
// This is separated from the main openapi/v3 files because if we
// rely on code that resides in the same directory as where we are
// generating code, we risk messing up the code and not being able to
// run code generation tools again because of compile problems
package types

const (
	DefaultSpecVersion   = "0.0.1"
	DefaultVersion       = "3.0.1"
	defaultHeaderInValue = "header"
)

type Location string

const (
	InHeader Location = "header"
	InQuery  Location = "query"
	InPath   Location = "path"
	InCookie Location = "cookie"
)

type PrimitiveType string

const (
	Invalid PrimitiveType = "invalid"
	Integer PrimitiveType = "integer"
	Number  PrimitiveType = "number"
	String  PrimitiveType = "string"
	Boolean PrimitiveType = "boolean"
	Object  PrimitiveType = "object"
	Array   PrimitiveType = "array"
	Null    PrimitiveType = "null"
)

type OpenAPI interface {
	QueryJSON(string) (interface{}, bool)
}

type openAPI struct {
	version      string                `json:"openapi" builder:"required" default:"DefaultVersion"`
	info         Info                  `json:"info" builder:"required"`
	servers      ServerList            `json:"servers,omitempty"`
	paths        Paths                 `json:"paths" builder:"required"`
	components   Components            `json:"components,omitempty"`
	security     SecurityRequirement   `json:"security,omitempty"`
	tags         TagList               `json:"tags,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Info interface {
}

type info struct {
	title          string  `json:"title" builder:"required"`
	description    string  `json:"description,omitempty"`
	termsOfService string  `json:"termsOfService,omitempty"`
	contact        Contact `json:"contact,omitempty"`
	license        License `json:"license,omitempty"`
	version        string  `json:"version" builder:"required" default:"DefaultSpecVersion"`
}

// Contact represents the contact object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.1.md#contactObject
type Contact interface {
}

type contact struct {
	name  string `json:"name,omitempty"`
	url   string `json:"url,omitempty"`
	email string `json:"email,omitempty"`
}

type License interface {
}

type license struct {
	name string `json:"name" builder:"required"`
	url  string `json:"url,omitempty"`
}

type Server interface {
}

type server struct {
	url         string            `json:"url" builder:"required"`
	description string            `json:"description,omitempty"`
	variables   ServerVariableMap `json:"variables,omitempty"`
}

type ServerVariable interface {
	setName(string)
}

type serverVariable struct {
	name         string     `json:"-" builder:"-"`
	enum         StringList `json:"enum"`
	defaultValue string     `json:"default" builder:"required"`
	description  string     `json:"description"`
}

type Components interface {
}

type components struct {
	schemas         SchemaMap         `json:"schemas,omitempty"`         // or Reference
	responses       ResponseMap       `json:"responses,omitempty"`       // or Reference
	parameters      ParameterMap      `json:"parameters,omitempty"`      // or Reference
	examples        ExampleMap        `json:"examples,omitempty"`        // or Reference
	requestBodies   RequestBodyMap    `json:"requestBodies,omitempty"`   // or Reference
	headers         HeaderMap         `json:"headers,omitempty"`         // or Reference
	securitySchemes SecuritySchemeMap `json:"securitySchemes,omitempty"` // or Reference
	links           LinkMap           `json:"links,omitempty"`           // or Reference
	callbacks       CallbackMap       `json:"callbacks,omitempty"`       // or Reference
}

type Paths interface {
}

type paths struct {
	paths PathItemMap `json:"-" mutator:"-"`
}

type PathItem interface {
	setPath(string)
	setName(string)
	//gen:lazy Operations() *OperationListIterator
}

type pathItem struct {
	name        string        `json:"-" builder:"-"` // This is a secret variable that gets reset when the item is added to a path
	path        string        `json:"-" resolve:"-"` // This is a secret variable that gets reset when the item is added to a path
	summary     string        `json:"summary,omitempty"`
	description string        `json:"description,omitempty"`
	get         Operation     `json:"get,omitempty" builder:"-" mutator:"-"`
	put         Operation     `json:"put,omitempty" builder:"-" mutator:"-"`
	post        Operation     `json:"post,omitempty" builder:"-" mutator:"-"`
	delete      Operation     `json:"delete,omitempty" builder:"-" mutator:"-"`
	options     Operation     `json:"options,omitempty" builder:"-" mutator:"-"`
	head        Operation     `json:"head,omitempty" builder:"-" mutator:"-"`
	patch       Operation     `json:"patch,omitempty" builder:"-" mutator:"-"`
	trace       Operation     `json:"trace,omitempty" builder:"-" mutator:"-"`
	servers     ServerList    `json:"servers,omitempty"`
	parameters  ParameterList `json:"parameters,omitempty"` // or Reference
}

type Operation interface {
	setVerb(string)
	setPathItem(PathItem)
	Path() string
	Detached() bool
}

type operation struct {
	verb         string                  `json:"-" builder:"-" mutator:"-" resolve:"-"`
	pathItem     PathItem                `json:"-" builder:"-" mutator:"-" resolve:"-"`
	tags         StringList              `json:"tags,omitempty"`
	summary      string                  `json:"summary,omitempty"`
	description  string                  `json:"description,omitempty"`
	externalDocs ExternalDocumentation   `json:"externalDocs,omitempty"`
	operationID  string                  `json:"operationId,omitempty"`
	parameters   ParameterList           `json:"parameters,omitempty"`  // or Reference
	requestBody  RequestBody             `json:"requestBody,omitempty"` // or Reference
	responses    Responses               `json:"responses" builder:"required"`
	callbacks    CallbackMap             `json:"callbacks,omitempty"` // or Reference
	deprecated   bool                    `json:"deprecated,omitempty"`
	security     SecurityRequirementList `json:"security,omitempty"`
	servers      ServerList              `json:"servers,omitempty"`
}

type ExternalDocumentation interface {
}

type externalDocumentation struct {
	description string `json:"description"`
	url         string `json:"url" builder:"required"` // REQUIRED
}

type RequestBody interface {
	setName(string)
}

type requestBody struct {
	name        string       `json:"-" builder:"-"`
	description string       `json:"description,omitempty"`
	content     MediaTypeMap `json:"content,omitempty" builder:"-" mutator:"-"`
	required    bool         `json:"required,omitempty"`
}

type MediaType interface {
	setMime(string)
	setName(string)
}

type mediaType struct {
	name     string      `json:"-" builder:"-"`    // This is a secret variable that gets reset when the  is added to the container
	mime     string      `json:"-" builder:"-"`    // This is a secret variable that gets reset when the  is added to the container
	schema   Schema      `json:"schema,omitempty"` // or Reference
	examples ExampleMap  `json:"examples,omitempty"`
	encoding EncodingMap `json:"encoding,omitempty"`
}

type Encoding interface {
	setName(string)
}

type encoding struct {
	name          string    `json:"-" builder:"-"`
	contentType   string    `json:"contentType"`
	headers       HeaderMap `json:"headers"`
	explode       bool      `json:"explode"`
	allowReserved bool      `json:"allowReserved"`
}

type Responses interface {
}

type responses struct {
	defaultValue Response    `json:"default,omitempty"` // or Reference
	responses    ResponseMap `json:"-" builder:"-"`     // or Reference
}

type Response interface {
	setName(string)
}

type response struct {
	name        string       `json:"-" builder:"-"`
	description string       `json:"description" builder:"required"`
	headers     HeaderMap    `json:"headers,omitempty"` // or Reference
	content     MediaTypeMap `json:"content,omitempty" builder:"-"`
	links       LinkMap      `json:"links,omitempty" builder:"-"` // or Reference
}

type Callback interface {
	setName(string)
}

type callback struct {
	name string `json:"-" builder:"-"`
	urls map[string]PathItem
}

type Example interface {
	setName(string)
}

type example struct {
	name          string      `json:"-" builder:"-"`
	description   string      `json:"description"`
	value         interface{} `json:"value"`
	externalValue string      `json:"externalValue"`
}

type Link interface {
	setName(string)
}

type link struct {
	name         string       `json:"-" builder:"-"` // This is only populated when applicable
	operationRef string       `json:"operationRef"`
	operationID  string       `json:"operationId"`
	parameters   InterfaceMap `json:"parameters"`
	requestBody  interface{}  `json:"requestBody"`
	description  string       `json:"description"`
	server       Server       `json:"server"`
}

type Tag interface {
}

type tag struct {
	name         string                `json:"name" builder:"required"`
	description  string                `json:"description,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Schema interface {
	Type() PrimitiveType
	setName(string)
}

type schema struct {
	name             string                `json:"-" builder:"-"` // This is only populated when applicable
	title            string                `json:"title,omitempty"`
	multipleOf       float64               `json:"multipleOf,omitempty"`
	maximum          float64               `json:"maximum,omitempty"`
	exclusiveMaximum float64               `json:"exclusiveMaximum,omitempty"`
	minimum          float64               `json:"minimum,omitempty"`
	exclusiveMinimum float64               `json:"exclusiveMinimum,omitempty"`
	maxLength        int                   `json:"maxLength,omitempty"`
	minLength        int                   `json:"minLength,omitempty"`
	pattern          string                `json:"pattern,omitempty"`
	maxItems         int                   `json:"maxItems,omitempty"`
	minItems         int                   `json:"minItems,omitempty"`
	uniqueItems      bool                  `json:"uniqueItems,omitempty"`
	maxProperties    int                   `json:"maxProperties,omitempty"`
	minProperties    int                   `json:"minProperties,omitempty"`
	required         StringList            `json:"required,omitempty"`
	enum             InterfaceList         `json:"enum,omitempty"`
	typ              PrimitiveType         `json:"type,omitempty" accessor:"-"`
	allOf            SchemaList            `json:"allOf,omitempty"`
	oneOf            SchemaList            `json:"oneOf,omitempty"`
	anyOf            SchemaList            `json:"anyOf,omitempty"`
	not              Schema                `json:"not,omitempty"`
	items            Schema                `json:"items,omitempty"`
	properties       SchemaMap             `json:"properties,omitempty"`
	format           string                `json:"format,omitempty"`
	defaultValue     interface{}           `json:"default,omitempty"`
	nullable         bool                  `json:"nullable,omitempty"`
	discriminator    Discriminator         `json:"discriminator,omitempty"`
	readOnly         bool                  `json:"readOnly,omitempty"`
	writeOnly        bool                  `json:"writeOnly,omitempty"`
	externalDocs     ExternalDocumentation `json:"externalDocs,omitempty"`
	example          interface{}           `json:"example,omitempty"`
	deprecated       bool                  `json:"deprecated,omitempty"`
}

type Discriminator interface {
}

type discriminator struct {
	propertyName string    `json:"propertyName" builder:"required"`
	mapping      StringMap `json:"mapping"`
}

type SecurityScheme interface {
}

type securityScheme struct {
	typ              string     `json:"type" builder:"required"` // REQUIRED
	description      string     `json:"description"`
	name             string     `json:"name" builder:"required"`   // REQUIRED
	in               string     `json:"in" builder:"required"`     // REQUIRED
	scheme           string     `json:"scheme" builder:"required"` // REQUIRED
	bearerFormat     string     `json:"bearerFormat"`
	flows            OAuthFlows `json:"flows" builder:"required"`            // REQUIRED
	openIDConnectURL string     `json:"openIdConnectUrl" builder:"required"` // REQUIRED
}

type OAuthFlows interface {
}

type oAuthFlows struct {
	implicit          OAuthFlow `json:"implicit"`
	password          OAuthFlow `json:"password"`
	clientCredentials OAuthFlow `json:"clientCredentials"`
	authorizationCode OAuthFlow `json:"authorizationCode"`
}

type OAuthFlow interface {
}

type oAuthFlow struct {
	authorizationURL string   `json:"authorizationUrl"`
	tokenURL         string   `json:"tokenUrl"`
	refreshURL       string   `json:"refreshUrl"`
	scopes           ScopeMap `json:"scopes"`
}

type SecurityRequirement interface {
}

type securityRequirement struct {
	schemes StringListMap
}

type Header interface {
	//gen:lazy setName(string)
}

type header struct {
	name            string       `json:"-" builder:"-" mutator:"-" resolve:"-"`
	in              Location     `json:"-" builder:"required" default:"InHeader"`
	required        bool         `json:"required,omitempty"`
	description     string       `json:"description,omitempty"`
	deprecated      bool         `json:"deprecated,omitempty"`
	allowEmptyValue bool         `json:"allowEmptyValue,omitempty"`
	explode         bool         `json:"explode,omitempty"`
	allowReserved   bool         `json:"allowReserved,omitempty"`
	schema          Schema       `json:"schema,omitempty"`
	examples        ExampleMap   `json:"examples,omitempty"`
	content         MediaTypeMap `json:"content,omitempty"`
}

type Parameter interface {
}

type parameter struct {
	name            string       `json:"name,omitempty" builder:"required" resolve:"-"`
	in              Location     `json:"in" builder:"required"`
	required        bool         `json:"required,omitempty" default:"defaultParameterRequiredFromLocation(in)"`
	description     string       `json:"description,omitempty"`
	deprecated      bool         `json:"deprecated,omitempty"`
	allowEmptyValue bool         `json:"allowEmptyValue,omitempty"`
	explode         bool         `json:"explode,omitempty"`
	allowReserved   bool         `json:"allowReserved,omitempty"`
	schema          Schema       `json:"schema,omitempty"`
	examples        ExampleMap   `json:"examples,omitempty"`
	content         MediaTypeMap `json:"content,omitempty"`
}

// note: using type aliases for key types effectively allow us to refer
// to every key type as "base type name" + "key", which is a win for
// code generation (no reflect trickeries). And yet since it is
// treated as the underlying key, casual users are affected by errors like
//
//    "cannot use s (type string) as type FooBarKey in map index"

type CallbackMapKey = string
type CallbackMap map[CallbackMapKey]Callback
type EncodingMapKey = string
type EncodingMap map[EncodingMapKey]Encoding
type ExampleMapKey = string
type ExampleMap map[ExampleMapKey]Example
type HeaderMapKey = string
type HeaderMap map[HeaderMapKey]Header
type InterfaceList []interface{}
type InterfaceMapKey = string
type InterfaceMap map[InterfaceMapKey]interface{}
type LinkMapKey = string
type LinkMap map[LinkMapKey]Link
type MediaTypeMapKey = string
type MediaTypeMap map[MediaTypeMapKey]MediaType
type ParameterList []Parameter
type ParameterMapKey = string
type ParameterMap map[ParameterMapKey]Parameter
type PathItemMapKey = string
type PathItemMap map[PathItemMapKey]PathItem
type RequestBodyMapKey = string
type RequestBodyMap map[RequestBodyMapKey]RequestBody
type ResponseMapKey = string
type ResponseMap map[ResponseMapKey]Response
type ServerList []Server
type ServerVariableMapKey = string
type ServerVariableMap map[ServerVariableMapKey]ServerVariable
type SchemaList []Schema
type SchemaMapKey = string
type SchemaMap map[SchemaMapKey]Schema
type ScopeMapKey = string
type ScopeMap map[ScopeMapKey]string
type SecurityRequirementList []SecurityRequirement
type SecuritySchemeMapKey = string
type SecuritySchemeMap map[SecuritySchemeMapKey]SecurityScheme
type StringList []string
type StringListMapKey = string
type StringListMap map[StringListMapKey][]string
type StringMapKey = string
type StringMap map[StringMapKey]string
type TagList []Tag
