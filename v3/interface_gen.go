package openapi

// This file was automatically generated.
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
	Invalid PrimitiveType = "invalid"
	Integer PrimitiveType = "integer"
	Number  PrimitiveType = "number"
	String  PrimitiveType = "string"
	Boolean PrimitiveType = "boolean"
	Object  PrimitiveType = "object"
	Array   PrimitiveType = "array"
	Null    PrimitiveType = "null"
)

// Validator objects can validate themselves.
type Validator interface {
	Validate(bool) error
}

type Extensions map[string]interface{}

type OpenAPI interface {
	Version() string
	Info() Info
	Servers() *ServerListIterator
	Paths() Paths
	Components() Components
	Security() SecurityRequirement
	Tags() *TagListIterator
	ExternalDocs() ExternalDocumentation
	MarshalJSON() ([]byte, error)
	Clone() OpenAPI
	Validator
	QueryJSON(string) (interface{}, bool)
}

type openAPI struct {
	reference    string                `json:"$ref,omitempty"`
	resolved     bool                  `json:"-"`
	extensions   Extensions            `json:"-"`
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
	Title() string
	Description() string
	TermsOfService() string
	Contact() Contact
	License() License
	Version() string
	MarshalJSON() ([]byte, error)
	Clone() Info
	Validator
}

type info struct {
	reference      string     `json:"$ref,omitempty"`
	resolved       bool       `json:"-"`
	extensions     Extensions `json:"-"`
	title          string     `json:"title" builder:"required"`
	description    string     `json:"description,omitempty"`
	termsOfService string     `json:"termsOfService,omitempty"`
	contact        Contact    `json:"contact,omitempty"`
	license        License    `json:"license,omitempty"`
	version        string     `json:"version" builder:"required" default:"DefaultSpecVersion"`
}

// Contact represents the contact object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.1.md#contactObject
type Contact interface {
	Name() string
	URL() string
	Email() string
	MarshalJSON() ([]byte, error)
	Clone() Contact
	Validator
}

type contact struct {
	reference  string     `json:"$ref,omitempty"`
	resolved   bool       `json:"-"`
	extensions Extensions `json:"-"`
	name       string     `json:"name,omitempty"`
	url        string     `json:"url,omitempty"`
	email      string     `json:"email,omitempty"`
}

type License interface {
	Name() string
	URL() string
	MarshalJSON() ([]byte, error)
	Clone() License
	Validator
}

type license struct {
	reference  string     `json:"$ref,omitempty"`
	resolved   bool       `json:"-"`
	extensions Extensions `json:"-"`
	name       string     `json:"name" builder:"required"`
	url        string     `json:"url,omitempty"`
}

type Server interface {
	URL() string
	Description() string
	Variables() *ServerVariableMapIterator
	MarshalJSON() ([]byte, error)
	Clone() Server
	Validator
}

type server struct {
	reference   string            `json:"$ref,omitempty"`
	resolved    bool              `json:"-"`
	extensions  Extensions        `json:"-"`
	url         string            `json:"url" builder:"required"`
	description string            `json:"description,omitempty"`
	variables   ServerVariableMap `json:"variables,omitempty"`
}

type ServerVariable interface {
	Name() string
	Enum() *StringListIterator
	Default() string
	Description() string
	MarshalJSON() ([]byte, error)
	Clone() ServerVariable
	Validator
	setName(string)
}

type serverVariable struct {
	reference    string     `json:"$ref,omitempty"`
	resolved     bool       `json:"-"`
	extensions   Extensions `json:"-"`
	name         string     `json:"-" builder:"-"`
	enum         StringList `json:"enum"`
	defaultValue string     `json:"default" builder:"required"`
	description  string     `json:"description"`
}

type Components interface {
	Schemas() *SchemaMapIterator
	Responses() *ResponseMapIterator
	Parameters() *ParameterMapIterator
	Examples() *ExampleMapIterator
	RequestBodies() *RequestBodyMapIterator
	Headers() *HeaderMapIterator
	SecuritySchemes() *SecuritySchemeMapIterator
	Links() *LinkMapIterator
	Callbacks() *CallbackMapIterator
	MarshalJSON() ([]byte, error)
	Clone() Components
	Validator
}

type components struct {
	reference       string            `json:"$ref,omitempty"`
	resolved        bool              `json:"-"`
	extensions      Extensions        `json:"-"`
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
	Paths() *PathItemMapIterator
	MarshalJSON() ([]byte, error)
	Clone() Paths
	Validator
}

type paths struct {
	reference  string      `json:"$ref,omitempty"`
	resolved   bool        `json:"-"`
	extensions Extensions  `json:"-"`
	paths      PathItemMap `json:"-" mutator:"-"`
}

type PathItem interface {
	Name() string
	Path() string
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
	Servers() *ServerListIterator
	Parameters() *ParameterListIterator
	MarshalJSON() ([]byte, error)
	Clone() PathItem
	Validator
	setPath(string)
	setName(string)
	Operations() *OperationListIterator
}

type pathItem struct {
	reference   string        `json:"$ref,omitempty"`
	resolved    bool          `json:"-"`
	extensions  Extensions    `json:"-"`
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
	Verb() string
	PathItem() PathItem
	Tags() *StringListIterator
	Summary() string
	Description() string
	ExternalDocs() ExternalDocumentation
	OperationID() string
	Parameters() *ParameterListIterator
	RequestBody() RequestBody
	Responses() Responses
	Callbacks() *CallbackMapIterator
	Deprecated() bool
	Security() *SecurityRequirementListIterator
	Servers() *ServerListIterator
	MarshalJSON() ([]byte, error)
	Clone() Operation
	Validator
	setVerb(string)
	setPathItem(PathItem)
	Path() string
	Detached() bool
}

type operation struct {
	reference    string                  `json:"$ref,omitempty"`
	resolved     bool                    `json:"-"`
	extensions   Extensions              `json:"-"`
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
	Description() string
	URL() string
	MarshalJSON() ([]byte, error)
	Clone() ExternalDocumentation
	Validator
}

type externalDocumentation struct {
	reference   string     `json:"$ref,omitempty"`
	resolved    bool       `json:"-"`
	extensions  Extensions `json:"-"`
	description string     `json:"description"`
	url         string     `json:"url" builder:"required"` // REQUIRED
}

type RequestBody interface {
	Name() string
	Description() string
	Content() *MediaTypeMapIterator
	Required() bool
	MarshalJSON() ([]byte, error)
	Clone() RequestBody
	Validator
	setName(string)
}

type requestBody struct {
	reference   string       `json:"$ref,omitempty"`
	resolved    bool         `json:"-"`
	extensions  Extensions   `json:"-"`
	name        string       `json:"-" builder:"-"`
	description string       `json:"description,omitempty"`
	content     MediaTypeMap `json:"content,omitempty" builder:"-" mutator:"-"`
	required    bool         `json:"required,omitempty"`
}

type MediaType interface {
	Name() string
	Mime() string
	Schema() Schema
	Examples() *ExampleMapIterator
	Encoding() *EncodingMapIterator
	MarshalJSON() ([]byte, error)
	Clone() MediaType
	Validator
	setMime(string)
	setName(string)
}

type mediaType struct {
	reference  string      `json:"$ref,omitempty"`
	resolved   bool        `json:"-"`
	extensions Extensions  `json:"-"`
	name       string      `json:"-" builder:"-"`    // This is a secret variable that gets reset when the  is added to the container
	mime       string      `json:"-" builder:"-"`    // This is a secret variable that gets reset when the  is added to the container
	schema     Schema      `json:"schema,omitempty"` // or Reference
	examples   ExampleMap  `json:"examples,omitempty"`
	encoding   EncodingMap `json:"encoding,omitempty"`
}

type Encoding interface {
	Name() string
	ContentType() string
	Headers() *HeaderMapIterator
	Explode() bool
	AllowReserved() bool
	MarshalJSON() ([]byte, error)
	Clone() Encoding
	Validator
	setName(string)
}

type encoding struct {
	reference     string     `json:"$ref,omitempty"`
	resolved      bool       `json:"-"`
	extensions    Extensions `json:"-"`
	name          string     `json:"-" builder:"-"`
	contentType   string     `json:"contentType"`
	headers       HeaderMap  `json:"headers"`
	explode       bool       `json:"explode"`
	allowReserved bool       `json:"allowReserved"`
}

type Responses interface {
	Default() Response
	Responses() *ResponseMapIterator
	MarshalJSON() ([]byte, error)
	Clone() Responses
	Validator
}

type responses struct {
	reference    string      `json:"$ref,omitempty"`
	resolved     bool        `json:"-"`
	extensions   Extensions  `json:"-"`
	defaultValue Response    `json:"default,omitempty"` // or Reference
	responses    ResponseMap `json:"-" builder:"-"`     // or Reference
}

type Response interface {
	Name() string
	Description() string
	Headers() *HeaderMapIterator
	Content() *MediaTypeMapIterator
	Links() *LinkMapIterator
	MarshalJSON() ([]byte, error)
	Clone() Response
	Validator
	setName(string)
}

type response struct {
	reference   string       `json:"$ref,omitempty"`
	resolved    bool         `json:"-"`
	extensions  Extensions   `json:"-"`
	name        string       `json:"-" builder:"-"`
	description string       `json:"description" builder:"required"`
	headers     HeaderMap    `json:"headers,omitempty"` // or Reference
	content     MediaTypeMap `json:"content,omitempty" builder:"-"`
	links       LinkMap      `json:"links,omitempty" builder:"-"` // or Reference
}

type Callback interface {
	Name() string
	URLs() map[string]PathItem
	MarshalJSON() ([]byte, error)
	Clone() Callback
	Validator
	setName(string)
}

type callback struct {
	reference  string     `json:"$ref,omitempty"`
	resolved   bool       `json:"-"`
	extensions Extensions `json:"-"`
	name       string     `json:"-" builder:"-"`
	urls       map[string]PathItem
}

type Example interface {
	Name() string
	Description() string
	Value() interface{}
	ExternalValue() string
	MarshalJSON() ([]byte, error)
	Clone() Example
	Validator
	setName(string)
}

type example struct {
	reference     string      `json:"$ref,omitempty"`
	resolved      bool        `json:"-"`
	extensions    Extensions  `json:"-"`
	name          string      `json:"-" builder:"-"`
	description   string      `json:"description"`
	value         interface{} `json:"value"`
	externalValue string      `json:"externalValue"`
}

type Link interface {
	Name() string
	OperationRef() string
	OperationID() string
	Parameters() *InterfaceMapIterator
	RequestBody() interface{}
	Description() string
	Server() Server
	MarshalJSON() ([]byte, error)
	Clone() Link
	Validator
	setName(string)
}

type link struct {
	reference    string       `json:"$ref,omitempty"`
	resolved     bool         `json:"-"`
	extensions   Extensions   `json:"-"`
	name         string       `json:"-" builder:"-"` // This is only populated when applicable
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
	MarshalJSON() ([]byte, error)
	Clone() Tag
	Validator
}

type tag struct {
	reference    string                `json:"$ref,omitempty"`
	resolved     bool                  `json:"-"`
	extensions   Extensions            `json:"-"`
	name         string                `json:"name" builder:"required"`
	description  string                `json:"description,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Schema interface {
	Name() string
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
	Required() *StringListIterator
	Enum() *InterfaceListIterator
	AllOf() *SchemaListIterator
	OneOf() *SchemaListIterator
	AnyOf() *SchemaListIterator
	Not() Schema
	Items() Schema
	Properties() *SchemaMapIterator
	Format() string
	Default() interface{}
	Nullable() bool
	Discriminator() Discriminator
	ReadOnly() bool
	WriteOnly() bool
	ExternalDocs() ExternalDocumentation
	Example() interface{}
	Deprecated() bool
	MarshalJSON() ([]byte, error)
	Clone() Schema
	Validator
	Type() PrimitiveType
	setName(string)
}

type schema struct {
	reference        string                `json:"$ref,omitempty"`
	resolved         bool                  `json:"-"`
	extensions       Extensions            `json:"-"`
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
	PropertyName() string
	Mapping() *StringMapIterator
	MarshalJSON() ([]byte, error)
	Clone() Discriminator
	Validator
}

type discriminator struct {
	reference    string     `json:"$ref,omitempty"`
	resolved     bool       `json:"-"`
	extensions   Extensions `json:"-"`
	propertyName string     `json:"propertyName" builder:"required"`
	mapping      StringMap  `json:"mapping"`
}

type SecurityScheme interface {
	Type() string
	Description() string
	Name() string
	In() string
	Scheme() string
	BearerFormat() string
	Flows() OAuthFlows
	OpenIDConnectURL() string
	MarshalJSON() ([]byte, error)
	Clone() SecurityScheme
	Validator
}

type securityScheme struct {
	reference        string     `json:"$ref,omitempty"`
	resolved         bool       `json:"-"`
	extensions       Extensions `json:"-"`
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
	MarshalJSON() ([]byte, error)
	Clone() OAuthFlows
	Validator
}

type oauthFlows struct {
	reference         string     `json:"$ref,omitempty"`
	resolved          bool       `json:"-"`
	extensions        Extensions `json:"-"`
	implicit          OAuthFlow  `json:"implicit"`
	password          OAuthFlow  `json:"password"`
	clientCredentials OAuthFlow  `json:"clientCredentials"`
	authorizationCode OAuthFlow  `json:"authorizationCode"`
}

type OAuthFlow interface {
	AuthorizationURL() string
	TokenURL() string
	RefreshURL() string
	Scopes() *ScopeMapIterator
	MarshalJSON() ([]byte, error)
	Clone() OAuthFlow
	Validator
}

type oauthFlow struct {
	reference        string     `json:"$ref,omitempty"`
	resolved         bool       `json:"-"`
	extensions       Extensions `json:"-"`
	authorizationURL string     `json:"authorizationUrl"`
	tokenURL         string     `json:"tokenUrl"`
	refreshURL       string     `json:"refreshUrl"`
	scopes           ScopeMap   `json:"scopes"`
}

type SecurityRequirement interface {
	Schemes() *StringListMapIterator
	MarshalJSON() ([]byte, error)
	Clone() SecurityRequirement
	Validator
}

type securityRequirement struct {
	reference  string     `json:"$ref,omitempty"`
	resolved   bool       `json:"-"`
	extensions Extensions `json:"-"`
	schemes    StringListMap
}

type Header interface {
	Name() string
	In() Location
	Required() bool
	Description() string
	Deprecated() bool
	AllowEmptyValue() bool
	Explode() bool
	AllowReserved() bool
	Schema() Schema
	Examples() *ExampleMapIterator
	Content() *MediaTypeMapIterator
	MarshalJSON() ([]byte, error)
	Clone() Header
	Validator
	setName(string)
}

type header struct {
	reference       string       `json:"$ref,omitempty"`
	resolved        bool         `json:"-"`
	extensions      Extensions   `json:"-"`
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
	Name() string
	In() Location
	Required() bool
	Description() string
	Deprecated() bool
	AllowEmptyValue() bool
	Explode() bool
	AllowReserved() bool
	Schema() Schema
	Examples() *ExampleMapIterator
	Content() *MediaTypeMapIterator
	MarshalJSON() ([]byte, error)
	Clone() Parameter
	Validator
}

type parameter struct {
	reference       string       `json:"$ref,omitempty"`
	resolved        bool         `json:"-"`
	extensions      Extensions   `json:"-"`
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
