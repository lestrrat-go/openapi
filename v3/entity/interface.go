package entity

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

type OpenAPI struct {
	OpenAPI      string                 `json:"openapi" builder:"required" default:"DefaultVersion"`
	Info         *Info                  `json:"info" builder:"required"`
	Servers      []Server               `json:"servers,omitempty"`
	Paths        *Paths                 `json:"paths" builder:"required"`
	Components   *Components            `json:"components,omitempty"`
	Security     *SecurityRequirement   `json:"security,omitempty"`
	Tags         []*Tag                 `json:"tags,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Info struct {
	Title          string   `json:"title" builder:"required"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version" builder:"required" default:"DefaultSpecVersion"`
}

// Contact represents the contact object
// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.1.md#contactObject
type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

type License struct {
	Name string `json:"name" builder:"required"`
	URL  string `json:"url,omitempty"`
}

type Server struct {
	URL         string                     `json:"url" builder:"required"`
	Description string                     `json:"description,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty"`
}

type ServerVariable struct {
	Enum        []string `json:"enum"`
	Default     string   `json:"default" builder:"required"`
	Description string   `json:"description"`
}

type Components struct {
	Schemas         map[string]*Schema         `json:"schemas,omitempty"`         // or Reference
	Responses       map[string]*Response       `json:"responses,omitempty"`       // or Reference
	Parameters      map[string]*Parameter      `json:"parameters,omitempty"`      // or Reference
	Examples        map[string]*Example        `json:"examples,omitempty"`        // or Reference
	RequestBodies   map[string]*RequestBody    `json:"requestBodies,omitempty"`   // or Reference
	Headers         map[string]*Header         `json:"headers,omitempty"`         // or Reference
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty"` // or Reference
	Links           map[string]*Link           `json:"links,omitempty"`           // or Reference
	Callbacks       map[string]*Callback       `json:"callbacks,omitempty"`       // or Reference
}

type Paths struct {
	Paths map[string]*PathItem `json:"-"`
}

type PathItem struct {
	Reference   string       `json:"$ref,omitempty"`
	Summary     string       `json:"summary,omitempty"`
	Description string       `json:"description,omitempty"`
	Get         *Operation   `json:"get,omitempty"`
	Put         *Operation   `json:"put,omitempty"`
	Post        *Operation   `json:"post,omitempty"`
	Delete      *Operation   `json:"delete,omitempty"`
	Options     *Operation   `json:"options,omitempty"`
	Head        *Operation   `json:"head,omitempty"`
	Patch       *Operation   `json:"patch,omitempty"`
	Trace       *Operation   `json:"trace,omitempty"`
	Servers     []*Server    `json:"servers,omitempty"`
	Parameters  []*Parameter `json:"parameters,omitempty"` // or Reference
}

type Operation struct {
	Tags         []string               `json:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty"`
	Description  string                 `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	OperationID  string                 `json:"operationId,omitempty"`
	Parameters   []*Parameter           `json:"parameters,omitempty"`  // or Reference
	RequestBody  *RequestBody           `json:"requestBody,omitempty"` // or Reference
	Responses    *Responses             `json:"responses" builder:"required"`
	Callbacks    map[string]*Callback   `json:"callbacks,omitempty"` // or Reference
	Deprecated   bool                   `json:"deprecated,omitempty"`
	Security     []*SecurityRequirement `json:"security,omitempty"`
	Servers      []*Server              `json:"servers,omitempty"`
}

type ExternalDocumentation struct {
	Description string `json:"description" yaml:"description"`
	URL         string `json:"url" yaml:"url"` // REQUIRED
}

type RequestBody struct {
	Description string                `json:"description,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty"`
	Required    bool                  `json:"required,omitempty"`
}

type MediaType struct {
	Schema   *Schema              `json:"schema,omitempty"` // or Reference
	Example  interface{}          `json:"example,omitempty"`
	Examples map[string]*Example  `json:"examples,omitempty"`
	Encoding map[string]*Encoding `json:"encoding,omitempty"`
}

type Encoding struct {
	ContentType   string             `json:"contentType" yaml:"contentType"`
	Headers       map[string]*Header `json:"headers" yaml:"headers"`
	Explode       bool               `json:"explode" yaml:"explode"`
	AllowReserved bool               `json:"allowReserved" yaml:"allowReserved"`
}

type Responses struct {
	Default     *Response            `json:"default,omitempty" yaml:"default,omitempty"` // or Reference
	StatusCodes map[string]*Response `json:"-" yaml:"inline" builder:"-"`                // or Reference
}

type Response struct {
	Description string                `json:"description" builder:"required"`
	Headers     map[string]*Header    `json:"headers,omitempty" builder:"-"` // or Reference
	Content     map[string]*MediaType `json:"content,omitempty" builder:"-"`
	Links       map[string]*Link      `json:"links,omitempty" builder:"-"` // or Reference
}

type Callback struct {
	URLs map[string]*PathItem
}

type Example struct {
	Description   string      `json:"description" yaml:"description"`
	Value         interface{} `json:"value" yaml:"value"`
	ExternalValue string      `json:"external_value" yaml:"external_value"`
}

type Link struct {
	OperationRef string                 `json:"operationRef" yaml:"operationRef"`
	OperationID  string                 `json:"operationId" yaml:"operationId"`
	Parameters   map[string]interface{} `json:"parameters" yaml:"parameters"`
	RequestBody  interface{}            `json:"requestBody" yaml:"requestBody"`
	Description  string                 `json:"description" yaml:"description"`
	Server       *Server                `json:"server" yaml:"server"`
}

type Tag struct {
	Name         string                 `json:"name" builder:"required"`
	Description  string                 `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Reference struct {
	URL string // REQUIRED
}

type Schema struct {
	Reference        string                 `json:"$ref,omitempty"`
	Title            string                 `json:"title,omitempty"`
	MultipleOf       float64                `json:"multipleOf,omitempty"`
	Maximum          float64                `json:"maximum,omitempty"`
	ExclusiveMaximum float64                `json:"exclusiveMaximum,omitempty"`
	Minimum          float64                `json:"minimum,omitempty"`
	ExclusiveMinimum float64                `json:"exclusiveMinimum,omitempty"`
	MaxLength        int                    `json:"maxLength,omitempty"`
	MinLength        int                    `json:"minLength,omitempty"`
	Pattern          string                 `json:"pattern,omitempty"`
	MaxItems         int                    `json:"maxItems,omitempty"`
	MinItems         int                    `json:"minItems,omitempty"`
	UniqueItems      bool                   `json:"uniqueItems,omitempty"`
	MaxProperties    int                    `json:"maxProperties,omitempty"`
	MinProperties    int                    `json:"minProperties,omitempty"`
	Required         []string               `json:"required,omitempty"`
	Enum             []interface{}          `json:"enum,omitempty"`
	Type             PrimitiveType          `json:"type,omitempty"`
	AllOf            []*Schema              `json:"allOf,omitempty"`
	OneOf            []*Schema              `json:"oneOf,omitempty"`
	AnyOf            []*Schema              `json:"anyOf,omitempty"`
	Not              *Schema                `json:"not,omitempty"`
	Items            *Schema                `json:"items,omitempty"`
	Properties       map[string]*Schema     `json:"properties,omitempty"`
	Format           string                 `json:"format,omitempty"`
	Default          interface{}            `json:"default,omitempty"`
	Nullable         bool                   `json:"nullable,omitempty"`
	Discriminator    *Discriminator         `json:"discriminator,omitempty"`
	ReadOnly         bool                   `json:"read_only,omitempty"`
	WriteOnly        bool                   `json:"write_only,omitempty"`
	ExternalDocs     *ExternalDocumentation `json:"externalDocs,omitempty"`
	Example          interface{}            `json:"example,omitempty"`
	Deprecated       bool                   `json:"deprecated,omitempty"`
}

type Discriminator struct {
	PropertyName string            `json:"property_name" yaml:"property_name"` // REQUIRED
	Mapping      map[string]string `json:"mapping" yaml:"mapping"`
}

type SecurityScheme struct {
	Type             string      `json:"type" yaml:"type"` // REQUIRED
	Description      string      `json:"description" yaml:"description"`
	Name             string      `json:"name" yaml:"name"`     // REQUIRED
	In               string      `json:"in" yaml:"in"`         // REQUIRED
	Scheme           string      `json:"scheme" yaml:"scheme"` // REQUIRED
	BearerFormat     string      `json:"bearerFormat" yaml:"bearerFormat"`
	Flows            *OAuthFlows `json:"flows" yaml:"flows"`                       // REQUIRED
	OpenIDConnectURL string      `json:"openIdConnectUrl" yaml:"openIdConnectUrl"` // REQUIRED
}

type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit" yaml:"implicit"`
	Password          *OAuthFlow `json:"password" yaml:"password"`
	ClientCredentials *OAuthFlow `json:"clientCredentials" yaml:"clientCredentials"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode" yaml:"authorizationCode"`
}

type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl" yaml:"authorizationUrl"`
	TokenURL         string            `json:"tokenUrl" yaml:"tokenUrl"`
	RefreshURL       string            `json:"refreshUrl" yaml:"refreshUrl"`
	Scopes           map[string]string `json:"scopes" yaml:"scopes"`
}

type SecurityRequirement struct {
	Schemes map[string][]string
}
