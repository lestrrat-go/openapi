package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T19:20:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

const defaultSwaggerVersion = "2.0"

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

type Swagger interface {
	Version() string
	Info() Info
	Host() string
	BasePath() string
	Schemes() *SchemeListIterator
	Consumes() *MIMETypeListIterator
	Produces() *MIMETypeListIterator
	Paths() Paths
	Definitions() *SchemaMapIterator
	Parameters() *ParameterMapIterator
	Responses() *ResponseMapIterator
	SecurityDefinitions() *SecuritySchemeMapIterator
	Security() *SecurityRequirementListIterator
	Tags() *TagListIterator
	ExternalDocs() ExternalDocumentation
	Clone() Swagger
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
	QueryJSON(string) (interface{}, bool)
}

type swagger struct {
	reference           string                  `json:"$ref,omitempty"`
	resolved            bool                    `json:"-"`
	version             string                  `json:"swagger" builder:"-" default:"defaultSwaggerVersion"`
	info                Info                    `json:"info" builder:"required"`
	host                string                  `json:"host"`
	basePath            string                  `json:"basePath"`
	schemes             SchemeList              `json:"schemes"`
	consumes            MIMETypeList            `json:"consumes,omitempty"`
	produces            MIMETypeList            `json:"produces,omitempty"`
	paths               Paths                   `json:"paths" builder:"required"`
	definitions         SchemaMap               `json:"definitions"`
	parameters          ParameterMap            `json:"parameters"`
	responses           ResponseMap             `json:"responses"`
	securityDefinitions SecuritySchemeMap       `json:"securityDefinitions"`
	security            SecurityRequirementList `json:"security"`
	tags                TagList                 `json:"tags"`
	externalDocs        ExternalDocumentation   `json:"externalDocs"`
}

type Info interface {
	Title() string
	Version() string
	Description() string
	TermsOfService() string
	Contact() Contact
	License() License
	Clone() Info
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type info struct {
	reference      string  `json:"$ref,omitempty"`
	resolved       bool    `json:"-"`
	title          string  `json:"title" builder:"required"`
	version        string  `json:"version" builder:"required"`
	description    string  `json:"description"`
	termsOfService string  `json:"termsOfService"`
	contact        Contact `json:"contact"`
	license        License `json:"license"`
}

type Contact interface {
	Name() string
	URL() string
	Email() string
	Clone() Contact
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type contact struct {
	reference string `json:"$ref,omitempty"`
	resolved  bool   `json:"-"`
	name      string `json:"name"`
	uRL       string `json:"url"`
	email     string `json:"email"`
}

type License interface {
	Name() string
	URL() string
	Clone() License
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type license struct {
	reference string `json:"$ref,omitempty"`
	resolved  bool   `json:"-"`
	name      string `json:"name" builder:"required"`
	uRL       string `json:"url"`
}

type Paths interface {
	Paths() *PathItemMapIterator
	Clone() Paths
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type paths struct {
	reference string      `json:"$ref,omitempty"`
	resolved  bool        `json:"-"`
	paths     PathItemMap `json:"-"`
}

type PathItem interface {
	Name() string
	Path() string
	Get() Operation
	Put() Operation
	Post() Operation
	Delete() Operation
	Options() Operation
	Head() Operation
	Patch() Operation
	Parameters() *ParameterListIterator
	Clone() PathItem
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
	setName(string)
	setPath(string)
}

type pathItem struct {
	reference string `json:"$ref,omitempty"`
	resolved  bool   `json:"-"`
	name      string `json:"-"`
	path      string `json:"-"`
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
	Tags() *TagListIterator
	Summary() string
	Description() string
	ExternalDocs() ExternalDocumentation
	OperationID() string
	Consumes() *MIMETypeListIterator
	Produces() *MIMETypeListIterator
	Parameters() *ParameterListIterator
	Responses() Responses
	Schemes() *SchemeListIterator
	Deprecated() bool
	Security() *SecurityRequirementListIterator
	Clone() Operation
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type operation struct {
	reference    string                  `json:"$ref,omitempty"`
	resolved     bool                    `json:"-"`
	tags         TagList                 `json:"tags,omitempty"`
	summary      string                  `json:"summary,omitempty"`
	description  string                  `json:"description,omitempty"`
	externalDocs ExternalDocumentation   `json:"externalDocs,omitempty"`
	operationID  string                  `json:"operationId,omitempty"`
	consumes     MIMETypeList            `json:"consumes,omitempty"`
	produces     MIMETypeList            `json:"produces,omitempty"`
	parameters   ParameterList           `json:"parameters,omitempty"`
	responses    Responses               `json:"responses" builder:"required"`
	schemes      SchemeList              `json:"schemes"`
	deprecated   bool                    `json:"deprecated,omitempty"`
	security     SecurityRequirementList `json:"security,omitempty"`
}

type ExternalDocumentation interface {
	URL() string
	Description() string
	Clone() ExternalDocumentation
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type externalDocumentation struct {
	reference   string `json:"$ref,omitempty"`
	resolved    bool   `json:"-"`
	uRL         string `json:"url" builder:"required"`
	description string `json:"description,omitempty"`
}

type Parameter interface {
	Name() string
	Description() string
	Required() bool
	In() Location
	Schema() Schema
	Type() PrimitiveType
	Format() string
	Title() string
	AllowEmptyValue() bool
	Items() Items
	CollectionFormat() CollectionFormat
	DefaultValue() interface{}
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
	Enum() *InterfaceListIterator
	MultipleOf() float64
	Clone() Parameter
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type parameter struct {
	reference   string   `json:"$ref,omitempty"`
	resolved    bool     `json:"-"`
	name        string   `json:"name" builder:"required"`
	description string   `json:"description,omitempty"`
	required    bool     `json:"required,omitempty"`
	in          Location `jsonm:"in" builder:"required"`
	// Only applicable if when in == body
	schema Schema `json:"schema"`
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
	Type() string
	Format() string
	Items() Items
	CollectionFormat() CollectionFormat
	DefaultValue() interface{}
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
	Enum() *InterfaceListIterator
	MultipleOf() float64
	Clone() Items
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type items struct {
	reference        string           `json:"$ref,omitempty"`
	resolved         bool             `json:"-"`
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
	DefaultValue() Response
	Responses() *ResponseMapIterator
	Clone() Responses
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type responses struct {
	reference    string      `json:"$ref,omitempty"`
	resolved     bool        `json:"-"`
	defaultValue Response    `json:"default,omitempty"`
	responses    ResponseMap `json:"-"`
}

type Response interface {
	Name() string
	Description() string
	Schema() Schema
	Headers() *HeaderMapIterator
	Examples() *ExampleMapIterator
	Clone() Response
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type response struct {
	reference   string     `json:"$ref,omitempty"`
	resolved    bool       `json:"-"`
	name        string     `json:"-"`
	description string     `json:"description" builder:"required"`
	schema      Schema     `json:"schema,omitempty"`
	headers     HeaderMap  `json:"headers,omitempty"`
	examples    ExampleMap `json:"example,omitempty"`
}

type Header interface {
	Name() string
	Description() string
	Type() string
	Format() string
	Items() Items
	CollectionFormat() CollectionFormat
	DefaultValue() interface{}
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
	Enum() *InterfaceListIterator
	MultipleOf() float64
	Clone() Header
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type header struct {
	reference        string           `json:"$ref,omitempty"`
	resolved         bool             `json:"-"`
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
	Name() string
	Type() PrimitiveType
	Format() string
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
	Items() Schema
	Properties() *SchemaMapIterator
	AdditionaProperties() *SchemaMapIterator
	DefaultValue() interface{}
	Discriminator() string
	ReadOnly() bool
	ExternalDocs() ExternalDocumentation
	Example() interface{}
	Deprecated() bool
	XML() XML
	Clone() Schema
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type schema struct {
	reference           string                `json:"$ref,omitempty"`
	resolved            bool                  `json:"-"`
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
	Name() string
	Namespace() string
	Prefix() string
	Attribute() bool
	Wrapped() bool
	Clone() XML
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type xml struct {
	reference string `json:"$ref,omitempty"`
	resolved  bool   `json:"-"`
	name      string `json:"name,omitempty"`
	namespace string `json:"namespace,omitempty"`
	prefix    string `json:"prefix,omitempty"`
	attribute bool   `json:"attribute,omitempty"`
	wrapped   bool   `json:"wrapped,omitempty"`
}

type SecurityScheme interface {
	Type() string
	Description() string
	Name() string
	In() string
	Flow() string
	AuthorizationURL() string
	TokenURL() string
	Scopes() *StringMapIterator
	Clone() SecurityScheme
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type securityScheme struct {
	reference        string    `json:"$ref,omitempty"`
	resolved         bool      `json:"-"`
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
	Data() map[string][]string
	Clone() SecurityRequirement
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type securityRequirement struct {
	reference string `json:"$ref,omitempty"`
	resolved  bool   `json:"-"`
	data      map[string][]string
}

type Tag interface {
	Name() string
	Description() string
	ExternalDocs() ExternalDocumentation
	Clone() Tag
	IsUnresolved() bool
	MarshalJSON() ([]byte, error)
	Resolve(*Resolver) error
	Validate() error
}

type tag struct {
	reference    string                `json:"$ref,omitempty"`
	resolved     bool                  `json:"-"`
	name         string                `json:"name" builder:"required"`
	description  string                `json:"description,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type ExampleMapKey = string
type ExampleMap map[ExampleMapKey]interface{}
type HeaderMapKey = string
type HeaderMap map[HeaderMapKey]Header
type InterfaceList []interface{}
type MIMETypeList []string
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
