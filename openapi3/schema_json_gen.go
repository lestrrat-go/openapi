package openapi3

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/pkg/errors"
)

var _ = log.Printf
var _ = json.Unmarshal
var _ = errors.Cause

type schemaMarshalProxy struct {
	Reference        string                `json:"$ref,omitempty"`
	Title            string                `json:"title,omitempty"`
	MultipleOf       float64               `json:"multipleOf,omitempty"`
	Maximum          float64               `json:"maximum,omitempty"`
	ExclusiveMaximum float64               `json:"exclusiveMaximum,omitempty"`
	Minimum          float64               `json:"minimum,omitempty"`
	ExclusiveMinimum float64               `json:"exclusiveMinimum,omitempty"`
	MaxLength        int                   `json:"maxLength,omitempty"`
	MinLength        int                   `json:"minLength,omitempty"`
	Pattern          string                `json:"pattern,omitempty"`
	MaxItems         int                   `json:"maxItems,omitempty"`
	MinItems         int                   `json:"minItems,omitempty"`
	UniqueItems      bool                  `json:"uniqueItems,omitempty"`
	MaxProperties    int                   `json:"maxProperties,omitempty"`
	MinProperties    int                   `json:"minProperties,omitempty"`
	Required         StringList            `json:"required,omitempty"`
	Enum             InterfaceList         `json:"enum,omitempty"`
	Type             PrimitiveType         `json:"type,omitempty"`
	AllOf            SchemaList            `json:"allOf,omitempty"`
	OneOf            SchemaList            `json:"oneOf,omitempty"`
	AnyOf            SchemaList            `json:"anyOf,omitempty"`
	Not              Schema                `json:"not,omitempty"`
	Items            Schema                `json:"items,omitempty"`
	Properties       SchemaMap             `json:"properties,omitempty"`
	Format           string                `json:"format,omitempty"`
	Default          interface{}           `json:"default,omitempty"`
	Nullable         bool                  `json:"nullable,omitempty"`
	Discriminator    Discriminator         `json:"discriminator,omitempty"`
	ReadOnly         bool                  `json:"readOnly,omitempty"`
	WriteOnly        bool                  `json:"writeOnly,omitempty"`
	ExternalDocs     ExternalDocumentation `json:"externalDocs,omitempty"`
	Example          interface{}           `json:"example,omitempty"`
	Deprecated       bool                  `json:"deprecated,omitempty"`
}

type schemaUnmarshalProxy struct {
	Reference        string          `json:"$ref,omitempty"`
	Title            string          `json:"title,omitempty"`
	MultipleOf       float64         `json:"multipleOf,omitempty"`
	Maximum          float64         `json:"maximum,omitempty"`
	ExclusiveMaximum float64         `json:"exclusiveMaximum,omitempty"`
	Minimum          float64         `json:"minimum,omitempty"`
	ExclusiveMinimum float64         `json:"exclusiveMinimum,omitempty"`
	MaxLength        int             `json:"maxLength,omitempty"`
	MinLength        int             `json:"minLength,omitempty"`
	Pattern          string          `json:"pattern,omitempty"`
	MaxItems         int             `json:"maxItems,omitempty"`
	MinItems         int             `json:"minItems,omitempty"`
	UniqueItems      bool            `json:"uniqueItems,omitempty"`
	MaxProperties    int             `json:"maxProperties,omitempty"`
	MinProperties    int             `json:"minProperties,omitempty"`
	Required         StringList      `json:"required,omitempty"`
	Enum             InterfaceList   `json:"enum,omitempty"`
	Type             PrimitiveType   `json:"type,omitempty"`
	AllOf            SchemaList      `json:"allOf,omitempty"`
	OneOf            SchemaList      `json:"oneOf,omitempty"`
	AnyOf            SchemaList      `json:"anyOf,omitempty"`
	Not              json.RawMessage `json:"not,omitempty"`
	Items            json.RawMessage `json:"items,omitempty"`
	Properties       SchemaMap       `json:"properties,omitempty"`
	Format           string          `json:"format,omitempty"`
	Default          interface{}     `json:"default,omitempty"`
	Nullable         bool            `json:"nullable,omitempty"`
	Discriminator    json.RawMessage `json:"discriminator,omitempty"`
	ReadOnly         bool            `json:"readOnly,omitempty"`
	WriteOnly        bool            `json:"writeOnly,omitempty"`
	ExternalDocs     json.RawMessage `json:"externalDocs,omitempty"`
	Example          interface{}     `json:"example,omitempty"`
	Deprecated       bool            `json:"deprecated,omitempty"`
}

func (v *schema) MarshalJSON() ([]byte, error) {
	var proxy schemaMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Title = v.title
	proxy.MultipleOf = v.multipleOf
	proxy.Maximum = v.maximum
	proxy.ExclusiveMaximum = v.exclusiveMaximum
	proxy.Minimum = v.minimum
	proxy.ExclusiveMinimum = v.exclusiveMinimum
	proxy.MaxLength = v.maxLength
	proxy.MinLength = v.minLength
	proxy.Pattern = v.pattern
	proxy.MaxItems = v.maxItems
	proxy.MinItems = v.minItems
	proxy.UniqueItems = v.uniqueItems
	proxy.MaxProperties = v.maxProperties
	proxy.MinProperties = v.minProperties
	proxy.Required = v.required
	proxy.Enum = v.enum
	proxy.Type = v.typ
	proxy.AllOf = v.allOf
	proxy.OneOf = v.oneOf
	proxy.AnyOf = v.anyOf
	proxy.Not = v.not
	proxy.Items = v.items
	proxy.Properties = v.properties
	proxy.Format = v.format
	proxy.Default = v.defaultValue
	proxy.Nullable = v.nullable
	proxy.Discriminator = v.discriminator
	proxy.ReadOnly = v.readOnly
	proxy.WriteOnly = v.writeOnly
	proxy.ExternalDocs = v.externalDocs
	proxy.Example = v.example
	proxy.Deprecated = v.deprecated
	return json.Marshal(proxy)
}

func (v *schema) UnmarshalJSON(data []byte) error {
	var proxy schemaUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return errors.Wrapf(err, `failed to unmarshal schema`)
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.title = proxy.Title
	v.multipleOf = proxy.MultipleOf
	v.maximum = proxy.Maximum
	v.exclusiveMaximum = proxy.ExclusiveMaximum
	v.minimum = proxy.Minimum
	v.exclusiveMinimum = proxy.ExclusiveMinimum
	v.maxLength = proxy.MaxLength
	v.minLength = proxy.MinLength
	v.pattern = proxy.Pattern
	v.maxItems = proxy.MaxItems
	v.minItems = proxy.MinItems
	v.uniqueItems = proxy.UniqueItems
	v.maxProperties = proxy.MaxProperties
	v.minProperties = proxy.MinProperties
	v.required = proxy.Required
	v.enum = proxy.Enum
	v.typ = proxy.Type
	v.allOf = proxy.AllOf
	v.oneOf = proxy.OneOf
	v.anyOf = proxy.AnyOf

	if len(proxy.Not) > 0 {
		var decoded schema
		if err := json.Unmarshal(proxy.Not, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Not`)
		}

		v.not = &decoded
	}

	if len(proxy.Items) > 0 {
		var decoded schema
		if err := json.Unmarshal(proxy.Items, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Items`)
		}

		v.items = &decoded
	}
	v.properties = proxy.Properties
	v.format = proxy.Format
	v.defaultValue = proxy.Default
	v.nullable = proxy.Nullable

	if len(proxy.Discriminator) > 0 {
		var decoded discriminator
		if err := json.Unmarshal(proxy.Discriminator, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Discriminator`)
		}

		v.discriminator = &decoded
	}
	v.readOnly = proxy.ReadOnly
	v.writeOnly = proxy.WriteOnly

	if len(proxy.ExternalDocs) > 0 {
		var decoded externalDocumentation
		if err := json.Unmarshal(proxy.ExternalDocs, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		v.externalDocs = &decoded
	}
	v.example = proxy.Example
	v.deprecated = proxy.Deprecated
	return nil
}

func (v *schema) QueryJSON(path string) (ret interface{}, ok bool) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	if path == "" {
		return v, true
	}

	var frag string
	if i := strings.Index(path, "/"); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ""
	}

	var target interface{}

	switch frag {
	case "title":
		target = v.title
	case "multipleOf":
		target = v.multipleOf
	case "maximum":
		target = v.maximum
	case "exclusiveMaximum":
		target = v.exclusiveMaximum
	case "minimum":
		target = v.minimum
	case "exclusiveMinimum":
		target = v.exclusiveMinimum
	case "maxLength":
		target = v.maxLength
	case "minLength":
		target = v.minLength
	case "pattern":
		target = v.pattern
	case "maxItems":
		target = v.maxItems
	case "minItems":
		target = v.minItems
	case "uniqueItems":
		target = v.uniqueItems
	case "maxProperties":
		target = v.maxProperties
	case "minProperties":
		target = v.minProperties
	case "required":
		target = v.required
	case "enum":
		target = v.enum
	case "type":
		target = v.typ
	case "allOf":
		target = v.allOf
	case "oneOf":
		target = v.oneOf
	case "anyOf":
		target = v.anyOf
	case "not":
		target = v.not
	case "items":
		target = v.items
	case "properties":
		target = v.properties
	case "format":
		target = v.format
	case "default":
		target = v.defaultValue
	case "nullable":
		target = v.nullable
	case "discriminator":
		target = v.discriminator
	case "readOnly":
		target = v.readOnly
	case "writeOnly":
		target = v.writeOnly
	case "externalDocs":
		target = v.externalDocs
	case "example":
		target = v.example
	case "deprecated":
		target = v.deprecated
	default:
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}
	if path == "" {
		return target, true
	}
	return nil, false
}

// SchemaFromJSON constructs a Schema from JSON buffer. `dst` must
// be a pointer to `Schema`
func SchemaFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Schema)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Schema, but got %T`, dst)
	}
	var tmp schema
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Schema`)
	}
	*v = &tmp
	return nil
}
