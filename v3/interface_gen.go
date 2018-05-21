package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

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
	Integer PrimitiveType = "integer"
	Number  PrimitiveType = "number"
	String  PrimitiveType = "string"
	Boolean PrimitiveType = "boolean"
	Object  PrimitiveType = "object"
	Array   PrimitiveType = "array"
	Null    PrimitiveType = "null"
)

type OpenAPI interface {
	Version() string
	Info() Info
	Paths() Paths
	Components() Components
	Security() SecurityRequirement
	Tags() []Tag
	ExternalDocs() ExternalDocumentation
	Clone() OpenAPI
	QueryJSON(string) (interface{}, bool)
	Servers() *ServerListIterator
}

type openAPI struct {
	version      string                `json:"openapi" builder:"required" default:"DefaultVersion"`
	info         Info                  `json:"info" builder:"required"`
	servers      []Server              `json:"servers,omitempty" accessor:"-"`
	paths        Paths                 `json:"paths" builder:"required"`
	components   Components            `json:"components,omitempty"`
	security     SecurityRequirement   `json:"security,omitempty"`
	tags         []Tag                 `json:"tags,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Info interface {
	Title() string
	Description() string
	TermsOfService() string
	Contact() Contact
	License() License
	Version() string
	Clone() Info
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
	Name() string
	URL() string
	Email() string
	Clone() Contact
}

type contact struct {
	name  string `json:"name,omitempty"`
	uRL   string `json:"url,omitempty"`
	email string `json:"email,omitempty"`
}

type License interface {
	Name() string
	URL() string
	Clone() License
}

type license struct {
	name string `json:"name" builder:"required"`
	uRL  string `json:"url,omitempty"`
}

type Server interface {
	URL() string
	Description() string
	Variables() map[string]ServerVariable
	Clone() Server
}

type server struct {
	uRL         string            `json:"url" builder:"required"`
	description string            `json:"description,omitempty"`
	variables   ServerVariableMap `json:"variables,omitempty"`
}

type ServerVariableMap map[string]ServerVariable

type ServerVariable interface {
	Enum() []string
	Default() string
	Description() string
	Clone() ServerVariable
}

type serverVariable struct {
	enum         []string `json:"enum"`
	defaultValue string   `json:"default" builder:"required"`
	description  string   `json:"description"`
}

type Components interface {
	Schemas() map[string]Schema
	Responses() map[string]Response
	Parameters() map[string]Parameter
	Examples() map[string]Example
	RequestBodies() map[string]RequestBody
	Headers() map[string]Header
	SecuritySchemes() map[string]SecurityScheme
	Links() map[string]Link
	Callbacks() map[string]Callback
	Clone() Components
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
	Paths() map[string]PathItem
	Clone() Paths
	Items() *PathItemListIterator
}

type paths struct {
	paths PathItemMap `json:"-"`
}

type PathItemMap map[string]PathItem

type PathItem interface {
	Path() string
	Reference() string
	Summary() string
	Description() string
	Get() Operation
	Put() Operation
	Post() Operation
	Delete() Operation
	Options() Operation
	Head() Operation
	Patch() Operation
	Trace() Operation
	Servers() []Server
	Parameters() []Parameter
	Clone() PathItem
	Operations() *OperationListIterator
	setPath(string)
}

type pathItem struct {
	path        string      `json:"-"` // This is a secret variable that gets reset when the item is added to a path
	reference   string      `json:"$ref,omitempty"`
	summary     string      `json:"summary,omitempty"`
	description string      `json:"description,omitempty"`
	get         Operation   `json:"get,omitempty"`
	put         Operation   `json:"put,omitempty"`
	post        Operation   `json:"post,omitempty"`
	delete      Operation   `json:"delete,omitempty"`
	options     Operation   `json:"options,omitempty"`
	head        Operation   `json:"head,omitempty"`
	patch       Operation   `json:"patch,omitempty"`
	trace       Operation   `json:"trace,omitempty"`
	servers     []Server    `json:"servers,omitempty"`
	parameters  []Parameter `json:"parameters,omitempty"` // or Reference
}

type Operation interface {
	PathItem() PathItem
	Verb() string
	Tags() []string
	Summary() string
	Description() string
	ExternalDocs() ExternalDocumentation
	OperationID() string
	RequestBody() RequestBody
	Responses() Responses
	Callbacks() map[string]Callback
	Deprecated() bool
	Security() []SecurityRequirement
	Servers() []Server
	Clone() Operation
	setVerb(string)
	setPathItem(PathItem)
	Path() string
	Detached() bool
	Parameters() *ParameterListIterator
}

type operation struct {
	pathItem     PathItem              `json:"-" builder:"-"` // This is a secreate variable that gets reset when the operation is added to a pathItem
	verb         string                `json:"-" builder:"-"` // This is a secreate variable that gets reset when the operation is added to a pathItem
	tags         []string              `json:"tags,omitempty"`
	summary      string                `json:"summary,omitempty"`
	description  string                `json:"description,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
	operationID  string                `json:"operationId,omitempty"`
	parameters   []Parameter           `json:"parameters,omitempty" accessor:"-"` // or Reference
	requestBody  RequestBody           `json:"requestBody,omitempty"`             // or Reference
	responses    Responses             `json:"responses" builder:"required"`
	callbacks    CallbackMap           `json:"callbacks,omitempty"` // or Reference
	deprecated   bool                  `json:"deprecated,omitempty"`
	security     []SecurityRequirement `json:"security,omitempty"`
	servers      []Server              `json:"servers,omitempty"`
}

type ExternalDocumentation interface {
	Description() string
	URL() string
	Clone() ExternalDocumentation
}

type externalDocumentation struct {
	description string `json:"description"`
	uRL         string `json:"url" builder:"required"` // REQUIRED
}

type RequestBodyMap map[string]RequestBody
type RequestBody interface {
	Description() string
	Required() bool
	Clone() RequestBody
	Contents() *MediaTypeListIterator
}

type requestBody struct {
	description string       `json:"description,omitempty"`
	content     MediaTypeMap `json:"content,omitempty" builder:"-" accessor:"-"`
	required    bool         `json:"required,omitempty"`
}

type MediaTypeMap map[string]MediaType

type MediaType interface {
	Mime() string
	Schema() Schema
	Example() interface{}
	Examples() map[string]Example
	Encoding() map[string]Encoding
	Clone() MediaType
	setMime(string)
}

type mediaType struct {
	mime     string      `json:"-" builder:"-"`    // This is a secret variable that gets reset when the  is added to the container
	schema   Schema      `json:"schema,omitempty"` // or Reference
	example  interface{} `json:"example,omitempty"`
	examples ExampleMap  `json:"examples,omitempty"`
	encoding EncodingMap `json:"encoding,omitempty"`
}

type EncodingMap map[string]Encoding
type Encoding interface {
	ContentType() string
	Headers() map[string]Header
	Explode() bool
	AllowReserved() bool
	Clone() Encoding
}

type encoding struct {
	contentType   string    `json:"contentType"`
	headers       HeaderMap `json:"headers"`
	explode       bool      `json:"explode"`
	allowReserved bool      `json:"allowReserved"`
}

type Responses interface {
	Default() Response
	Responses() map[string]Response
	Clone() Responses
}

type responses struct {
	defaultValue Response    `json:"default,omitempty"` // or Reference
	responses    ResponseMap `json:"-" builder:"-"`     // or Reference
}

type ResponseMap map[string]Response

type Response interface {
	Description() string
	Headers() map[string]Header
	Content() map[string]MediaType
	Links() map[string]Link
	Clone() Response
}

type response struct {
	description string       `json:"description" builder:"required"`
	headers     HeaderMap    `json:"headers,omitempty" builder:"-"` // or Reference
	content     MediaTypeMap `json:"content,omitempty" builder:"-"`
	links       LinkMap      `json:"links,omitempty" builder:"-"` // or Reference
}

type CallbackMap map[string]Callback

type Callback interface {
	URLs() map[string]PathItem
	Clone() Callback
}

type callback struct {
	uRLs map[string]PathItem
}

type ExampleMap map[string]Example

type Example interface {
	Description() string
	Value() interface{}
	ExternalValue() string
	Clone() Example
}

type example struct {
	description   string      `json:"description"`
	value         interface{} `json:"value"`
	externalValue string      `json:"externalValue"`
}

type LinkMap map[string]Link

type Link interface {
	OperationRef() string
	OperationID() string
	Parameters() map[string]interface{}
	RequestBody() interface{}
	Description() string
	Server() Server
	Clone() Link
}

type InterfaceMap map[string]interface{}

type link struct {
	operationRef string       `json:"operationRef"`
	operationID  string       `json:"operationId"`
	parameters   InterfaceMap `json:"parameters"`
	requestBody  interface{}  `json:"requestBody"`
	description  string       `json:"description"`
	server       Server       `json:"server"`
}

type Tag interface {
	Name() string
	Description() string
	ExternalDocs() ExternalDocumentation
	Clone() Tag
}

type tag struct {
	name         string                `json:"name" builder:"required"`
	description  string                `json:"description,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Reference interface {
	URL() string
	Clone() Reference
}

type reference struct {
	uRL string `json:"-"` // REQUIRED
}

type SchemaMap map[string]Schema

type Schema interface {
	Name() string
	Reference() string
	Title() string
	MultipleOf() float64
	Maximum() float64
	ExclusiveMaximum() float64
	Minimum() float64
	ExclusiveMinimum() float64
	MaxLength() int
	MinLength() int
	Pattern() string
	MaxItems() int
	MinItems() int
	UniqueItems() bool
	MaxProperties() int
	MinProperties() int
	Required() []string
	Enum() []interface{}
	Type() PrimitiveType
	AllOf() []Schema
	OneOf() []Schema
	AnyOf() []Schema
	Not() Schema
	Items() Schema
	Format() string
	Default() interface{}
	Nullable() bool
	Discriminator() Discriminator
	ReadOnly() bool
	WriteOnly() bool
	ExternalDocs() ExternalDocumentation
	Example() interface{}
	Deprecated() bool
	Clone() Schema
	Properties() *SchemaListIterator
	setName(string)
}

type schema struct {
	name             string                `json:"-" builder:"-"`
	reference        string                `json:"$ref,omitempty"`
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
	required         []string              `json:"required,omitempty"`
	enum             []interface{}         `json:"enum,omitempty"`
	typ              PrimitiveType         `json:"type,omitempty"`
	allOf            []Schema              `json:"allOf,omitempty"`
	oneOf            []Schema              `json:"oneOf,omitempty"`
	anyOf            []Schema              `json:"anyOf,omitempty"`
	not              Schema                `json:"not,omitempty"`
	items            Schema                `json:"items,omitempty"`
	properties       SchemaMap             `json:"properties,omitempty" accessor:"-" builder:"-"`
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
	PropertyName() string
	Mapping() map[string]string
	Clone() Discriminator
}

type StringMap map[string]string
type StringListMap map[string][]string

type discriminator struct {
	propertyName string    `json:"propertyName" builder:"required"`
	mapping      StringMap `json:"mapping"`
}

type SecuritySchemeMap map[string]SecurityScheme

type SecurityScheme interface {
	Type() string
	Description() string
	Name() string
	In() string
	Scheme() string
	BearerFormat() string
	Flows() OAuthFlows
	OpenIDConnectURL() string
	Clone() SecurityScheme
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
	Implicit() OAuthFlow
	Password() OAuthFlow
	ClientCredentials() OAuthFlow
	AuthorizationCode() OAuthFlow
	Clone() OAuthFlows
}

type oAuthFlows struct {
	implicit          OAuthFlow `json:"implicit"`
	password          OAuthFlow `json:"password"`
	clientCredentials OAuthFlow `json:"clientCredentials"`
	authorizationCode OAuthFlow `json:"authorizationCode"`
}

type ScopeMap map[string]string

type OAuthFlow interface {
	AuthorizationURL() string
	TokenURL() string
	RefreshURL() string
	Scopes() map[string]string
	Clone() OAuthFlow
}

type oAuthFlow struct {
	authorizationURL string   `json:"authorizationUrl"`
	tokenURL         string   `json:"tokenUrl"`
	refreshURL       string   `json:"refreshUrl"`
	scopes           ScopeMap `json:"scopes"`
}

type SecurityRequirement interface {
	Schemes() map[string][]string
	Clone() SecurityRequirement
}

type securityRequirement struct {
	schemes StringListMap
}

type HeaderMap map[string]Header
type Header interface {
	In() Location
	Required() bool
	Description() string
	Deprecated() bool
	AllowEmptyValue() bool
	Explode() bool
	AllowReserved() bool
	Schema() Schema
	Example() interface{}
	Examples() map[string]Example
	Content() map[string]MediaType
	Clone() Header
}

type header struct {
	in              Location     `json:"-" builder:"required" default:"InHeader"`
	required        bool         `json:"required,omitempty"`
	description     string       `json:"description,omitempty"`
	deprecated      bool         `json:"deprecated,omitempty"`
	allowEmptyValue bool         `json:"allowEmptyValue,omitempty"`
	explode         bool         `json:"explode,omitempty"`
	allowReserved   bool         `json:"allowReserved,omitempty"`
	schema          Schema       `json:"schema,omitempty"`
	example         interface{}  `json:"example,omitempty"`
	examples        ExampleMap   `json:"examples,omitempty"`
	content         MediaTypeMap `json:"content,omitempty"`
}

type ParameterMap map[string]Parameter

type Parameter interface {
	Name() string
	In() Location
	Required() bool
	Description() string
	Deprecated() bool
	AllowEmptyValue() bool
	Explode() bool
	AllowReserved() bool
	Schema() Schema
	Example() interface{}
	Examples() map[string]Example
	Content() map[string]MediaType
	Clone() Parameter
}

type parameter struct {
	name            string       `json:"name,omitempty" builder:"required"`
	in              Location     `json:"in" builder:"required"`
	required        bool         `json:"required,omitempty" default:"defaultParameterRequiredFromLocation(in)"`
	description     string       `json:"description,omitempty"`
	deprecated      bool         `json:"deprecated,omitempty"`
	allowEmptyValue bool         `json:"allowEmptyValue,omitempty"`
	explode         bool         `json:"explode,omitempty"`
	allowReserved   bool         `json:"allowReserved,omitempty"`
	schema          Schema       `json:"schema,omitempty"`
	example         interface{}  `json:"example,omitempty"`
	examples        ExampleMap   `json:"examples,omitempty"`
	content         MediaTypeMap `json:"content,omitempty"`
}
