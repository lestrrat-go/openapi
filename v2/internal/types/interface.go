package types

const defaultSwaggerVersion = "2.0"
const refOnlyTmpl = `{"$ref":%s}`

type CollectionFormat string

const (
	CSV   CollectionFormat = "csv"
	SSV   CollectionFormat = "ssv"
	TSV   CollectionFormat = "tsv"
	Pipes CollectionFormat = "pipes"
	Multi CollectionFormat = "multi"
)

type Scheme = string
type MediaType = string
type MIMEType = string

const (
	WWWFormURLEncoded MIMEType = "application/x-www-form-urlencoded"
	MultiparFormDaata MIMEType = "multipart/form-data"
)

type Location string

const (
	InPath   Location = "path"
	InQuery  Location = "query"
	InHeader Location = "header"
	InBody   Location = "body"
	InForm   Location = "formData"
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
	File    PrimitiveType = "file"
	Null    PrimitiveType = "null"
)

type Extensions map[string]interface{}

type OpenAPI = Swagger
type Swagger interface {
	QueryJSON(string) (interface{}, bool)
}

type swagger struct {
	version             string                  `json:"swagger" builder:"-" default:"defaultSwaggerVersion"`
	info                Info                    `json:"info" builder:"required"`
	host                string                  `json:"host,omitempty"`
	basePath            string                  `json:"basePath,omitempty"`
	schemes             SchemeList              `json:"schemes,omitempty"`
	consumes            MIMETypeList            `json:"consumes,omitempty"`
	produces            MIMETypeList            `json:"produces,omitempty"`
	paths               Paths                   `json:"paths" builder:"required"`
	definitions         SchemaMap               `json:"definitions,omitempty"`
	parameters          ParameterMap            `json:"parameters,omitempty"`
	responses           ResponseMap             `json:"responses,omitempty"`
	securityDefinitions SecuritySchemeMap       `json:"securityDefinitions,omitempty"`
	security            SecurityRequirementList `json:"security,omitempty"`
	tags                TagList                 `json:"tags,omitempty"`
	externalDocs        ExternalDocumentation   `json:"externalDocs,omitempty"`
}

type Info interface {
}

type info struct {
	title          string  `json:"title" builder:"required"`
	version        string  `json:"version" builder:"required"`
	description    string  `json:"description"`
	termsOfService string  `json:"termsOfService"`
	contact        Contact `json:"contact"`
	license        License `json:"license"`
}

type Contact interface {
}

type contact struct {
	name  string `json:"name"`
	uRL   string `json:"url"`
	email string `json:"email"`
}

type License interface {
}

type license struct {
	name string `json:"name" builder:"required"`
	uRL  string `json:"url"`
}

type Paths interface {
}

type paths struct {
	paths PathItemMap `json:"-" mutator:"-"`
}

type PathItem interface {
	//gen:lazy setName(string)
	//gen:lazy setPath(string)
}

type pathItem struct {
	name string `json:"-" builder:"-" mutator:"-" resolve:"-"`
	path string `json:"-" builder:"-" mutator:"-" resolve:"-"`
	// reference string `json:"$ref"`
	get        Operation     `json:"get,omitempty"`
	put        Operation     `json:"put,omitempty"`
	post       Operation     `json:"post,omitempty"`
	delete     Operation     `json:"delete,omitempty"`
	options    Operation     `json:"options,omitempty"`
	head       Operation     `json:"head,omitempty"`
	patch      Operation     `json:"patch,omitempty"`
	parameters ParameterList `json:"parameters,omitempty"`
}

type Operation interface {
	//gen:lazy setPathItem(PathItem)
	//gen:lazy setVerb(string)
}

type operation struct {
	verb         string                  `json:"-" builder:"-" mutator:"-" resolve:"-"`
	pathItem     PathItem                `json:"-" builder:"-" mutator:"-" resolve:"-"`
	tags         TagList                 `json:"tags,omitempty"`
	summary      string                  `json:"summary,omitempty"`
	description  string                  `json:"description,omitempty"`
	externalDocs ExternalDocumentation   `json:"externalDocs,omitempty"`
	operationID  string                  `json:"operationId,omitempty"`
	consumes     MIMETypeList            `json:"consumes,omitempty"`
	produces     MIMETypeList            `json:"produces,omitempty"`
	parameters   ParameterList           `json:"parameters,omitempty"`
	responses    Responses               `json:"responses" builder:"required"`
	schemes      SchemeList              `json:"schemes,omitempty"`
	deprecated   bool                    `json:"deprecated,omitempty"`
	security     SecurityRequirementList `json:"security,omitempty"`
}

type ExternalDocumentation interface {
}

type externalDocumentation struct {
	uRL         string `json:"url" builder:"required"`
	description string `json:"description,omitempty"`
}

type Parameter interface {
}

type parameter struct {
	name        string   `json:"name" builder:"required"`
	description string   `json:"description,omitempty"`
	required    bool     `json:"required,omitempty"`
	in          Location `json:"in" builder:"required"`
	// Only applicable if when in == body
	schema Schema `json:"schema,omitempty"`
	// Only applicable if when in != body
	typ              PrimitiveType    `json:"type,omitempty"`
	format           string           `json:"format,omitempty"`
	title            string           `json:"title,omitempty"`
	allowEmptyValue  bool             `json:"allowEmptyValue,omitempty"`
	items            Items            `json:"items,omitempty"`
	collectionFormat CollectionFormat `json:"collectionFormat,omitempty"`
	defaultValue     interface{}      `json:"default,omitempty"`
	maximum          float64          `json:"maximum,omitempty"`
	exclusiveMaximum float64          `json:"exclusiveMaximum,omitempty"`
	minimum          float64          `json:"minimum,omitempty"`
	exclusiveMinimum float64          `json:"exclusiveMinimum,omitempty"`
	maxLength        int              `json:"maxLength,omitempty"`
	minLength        int              `json:"minLength,omitempty"`
	pattern          string           `json:"pattern,omitempty"`
	maxItems         int              `json:"maxItems,omitempty"`
	minItems         int              `json:"minItems,omitempty"`
	uniqueItems      bool             `json:"uniqueItems,omitempty"`
	enum             InterfaceList    `json:"enum,omitempty"`
	multipleOf       float64          `json:"multipleOf,omitempty"`
}

type Items interface {
}

type items struct {
	typ              string           `json:"type,omitempty"`
	format           string           `json:"format,omitempty"`
	items            Items            `json:"items,omitempty"`
	collectionFormat CollectionFormat `json:"collectionFormat,omitempty"`
	defaultValue     interface{}      `json:"default,omitempty"`
	maximum          float64          `json:"maximum,omitempty"`
	exclusiveMaximum float64          `json:"exclusiveMaximum,omitempty"`
	minimum          float64          `json:"minimum,omitempty"`
	exclusiveMinimum float64          `json:"exclusiveMinimum,omitempty"`
	maxLength        int              `json:"maxLength,omitempty"`
	minLength        int              `json:"minLength,omitempty"`
	pattern          string           `json:"pattern,omitempty"`
	maxItems         int              `json:"maxItems,omitempty"`
	minItems         int              `json:"minItems,omitempty"`
	uniqueItems      bool             `json:"uniqueItems,omitempty"`
	enum             InterfaceList    `json:"enum,omitempty"`
	multipleOf       float64          `json:"multipleOf,omitempty"`
}

type Responses interface {
}

type responses struct {
	defaultValue Response    `json:"default,omitempty"`
	responses    ResponseMap `json:"-"`
}

type Response interface {
}

type response struct {
	name        string     `json:"-"`
	description string     `json:"description" builder:"required"`
	schema      Schema     `json:"schema,omitempty"`
	headers     HeaderMap  `json:"headers,omitempty"`
	examples    ExampleMap `json:"example,omitempty"`
}

type Header interface {
}

type header struct {
	name             string           `json:"-"`
	description      string           `json:"description,omitempty"`
	typ              string           `json:"type" builder:"required"`
	format           string           `json:"format,omitempty"`
	items            Items            `json:"items,omitempty"`
	collectionFormat CollectionFormat `json:"collectionFormat,omitempty"`
	defaultValue     interface{}      `json:"default,omitempty"`
	maximum          float64          `json:"maximum,omitempty"`
	exclusiveMaximum float64          `json:"exclusiveMaximum,omitempty"`
	minimum          float64          `json:"minimum,omitempty"`
	exclusiveMinimum float64          `json:"exclusiveMinimum,omitempty"`
	maxLength        int              `json:"maxLength,omitempty"`
	minLength        int              `json:"minLength,omitempty"`
	pattern          string           `json:"pattern,omitempty"`
	maxItems         int              `json:"maxItems,omitempty"`
	minItems         int              `json:"minItems,omitempty"`
	uniqueItems      bool             `json:"uniqueItems,omitempty"`
	enum             InterfaceList    `json:"enum,omitempty"`
	multipleOf       float64          `json:"multipleOf,omitempty"`
}

type Schema interface {
}

type schema struct {
	name                string                `json:"-" builder:"-"` // This is only populated when applicable
	typ                 PrimitiveType         `json:"type,omitempty"`
	format              string                `json:"format,omitempty"`
	title               string                `json:"title,omitempty"`
	multipleOf          float64               `json:"multipleOf,omitempty"`
	maximum             float64               `json:"maximum,omitempty"`
	exclusiveMaximum    float64               `json:"exclusiveMaximum,omitempty"`
	minimum             float64               `json:"minimum,omitempty"`
	exclusiveMinimum    float64               `json:"exclusiveMinimum,omitempty"`
	maxLength           int                   `json:"maxLength,omitempty"`
	minLength           int                   `json:"minLength,omitempty"`
	pattern             string                `json:"pattern,omitempty"`
	maxItems            int                   `json:"maxItems,omitempty"`
	minItems            int                   `json:"minItems,omitempty"`
	uniqueItems         bool                  `json:"uniqueItems,omitempty"`
	maxProperties       int                   `json:"maxProperties,omitempty"`
	minProperties       int                   `json:"minProperties,omitempty"`
	required            StringList            `json:"required,omitempty"`
	enum                InterfaceList         `json:"enum,omitempty"`
	allOf               SchemaList            `json:"allOf,omitempty"`
	items               Schema                `json:"items,omitempty"`
	properties          SchemaMap             `json:"properties,omitempty"`
	additionaProperties SchemaMap             `json:"additionalProperties,omitempty"`
	defaultValue        interface{}           `json:"default,omitempty"`
	discriminator       string                `json:"discriminator,omitempty"`
	readOnly            bool                  `json:"readOnly,omitempty"`
	externalDocs        ExternalDocumentation `json:"externalDocs,omitempty"`
	example             interface{}           `json:"example,omitempty"`
	deprecated          bool                  `json:"deprecated,omitempty"`
	xml                 XML                   `json:"xml,omitempty"`
}

type XML interface {
}

type xml struct {
	name      string `json:"name,omitempty"`
	namespace string `json:"namespace,omitempty"`
	prefix    string `json:"prefix,omitempty"`
	attribute bool   `json:"attribute,omitempty"`
	wrapped   bool   `json:"wrapped,omitempty"`
}

type SecurityScheme interface {
}

type securityScheme struct {
	typ              string    `json:"type" builder:"required"`
	description      string    `json:"description,omitempty"`
	name             string    `json:"name,omitempty"`
	in               string    `json:"in,omitempty"`
	flow             string    `json:"flow,omitempty"`
	authorizationURL string    `json:"authorizationUrl,omitempty"`
	tokenURL         string    `json:"tokenUrl,omitempty"`
	scopes           StringMap `json:"scopes,omitempty"`
}

type SecurityRequirement interface {
}

type securityRequirement struct {
	data map[string][]string
}

type Tag interface {
}

type tag struct {
	name         string                `json:"name" builder:"required"`
	description  string                `json:"description,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type ExampleMapKey = string
type ExampleMap map[ExampleMapKey]interface{}
type HeaderMapKey = string
type HeaderMap map[HeaderMapKey]Header
type InterfaceList []interface{}
type MIMETypeList []MIMEType
type ParameterList []Parameter
type ParameterMapKey = string
type ParameterMap map[ParameterMapKey]Parameter
type PathItemMapKey = string
type PathItemMap map[PathItemMapKey]PathItem
type ResponseMapKey = string
type ResponseMap map[ResponseMapKey]Response
type SchemeList []string
type SchemaList []Schema
type SchemaMapKey = string
type SchemaMap map[SchemaMapKey]Schema
type SecurityRequirementList []SecurityRequirement
type SecuritySchemeMapKey = string
type SecuritySchemeMap map[SecuritySchemeMapKey]SecurityScheme
type StringList []string
type StringMapKey = string
type StringMap map[StringMapKey]string
type TagList []Tag
