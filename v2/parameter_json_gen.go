package openapi

// This file was automatically generated by gentypes.go
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

var _ = json.Unmarshal
var _ = fmt.Fprintf
var _ = log.Printf
var _ = strconv.ParseInt
var _ = errors.Cause

type parameterMarshalProxy struct {
	Reference        string           `json:"$ref,omitempty"`
	Name             string           `json:"name"`
	Description      string           `json:"description,omitempty"`
	Required         bool             `json:"required,omitempty"`
	In               Location         `json:"in"`
	Schema           Schema           `json:"schema,omitempty"`
	Type             PrimitiveType    `json:"type,omitempty"`
	Format           string           `json:"format,omitempty"`
	Title            string           `json:"title,omitempty"`
	AllowEmptyValue  bool             `json:"allowEmptyValue,omitempty"`
	Items            Items            `json:"items,omitempty"`
	CollectionFormat CollectionFormat `json:"collectionFormat,omitempty"`
	Default          interface{}      `json:"default,omitempty"`
	Maximum          *float64         `json:"maximum,omitempty"`
	ExclusiveMaximum *float64         `json:"exclusiveMaximum,omitempty"`
	Minimum          *float64         `json:"minimum,omitempty"`
	ExclusiveMinimum *float64         `json:"exclusiveMinimum,omitempty"`
	MaxLength        *int             `json:"maxLength,omitempty"`
	MinLength        *int             `json:"minLength,omitempty"`
	Pattern          string           `json:"pattern,omitempty"`
	MaxItems         *int             `json:"maxItems,omitempty"`
	MinItems         *int             `json:"minItems,omitempty"`
	UniqueItems      bool             `json:"uniqueItems,omitempty"`
	Enum             InterfaceList    `json:"enum,omitempty"`
	MultipleOf       *float64         `json:"multipleOf,omitempty"`
}

func (v *parameter) MarshalJSON() ([]byte, error) {
	var proxy parameterMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Name = v.name
	proxy.Description = v.description
	proxy.Required = v.required
	proxy.In = v.in
	proxy.Schema = v.schema
	proxy.Type = v.typ
	proxy.Format = v.format
	proxy.Title = v.title
	proxy.AllowEmptyValue = v.allowEmptyValue
	proxy.Items = v.items
	proxy.CollectionFormat = v.collectionFormat
	proxy.Default = v.defaultValue
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
	proxy.Enum = v.enum
	proxy.MultipleOf = v.multipleOf
	buf, err := json.Marshal(proxy)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal struct`)
	}
	if len(v.extensions) > 0 {
		extBuf, err := json.Marshal(v.extensions)
		if err != nil || len(extBuf) <= 2 {
			return nil, errors.Wrap(err, `failed to marshal struct (extensions)`)
		}
		buf = append(append(buf[:len(buf)-1], ','), extBuf[1:]...)
	}
	return buf, nil
}

// UnmarshalJSON defines how parameter is deserialized from JSON
func (v *parameter) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if raw, ok := proxy["$ref"]; ok {
		if err := json.Unmarshal(raw, &v.reference); err != nil {
			return errors.Wrap(err, `failed to unmarshal $ref`)
		}
		return nil
	}

	mutator := MutateParameter(v)

	const nameMapKey = "name"
	if raw, ok := proxy[nameMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field name`)
		}
		mutator.Name(decoded)
		delete(proxy, nameMapKey)
	}

	const descriptionMapKey = "description"
	if raw, ok := proxy[descriptionMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field description`)
		}
		mutator.Description(decoded)
		delete(proxy, descriptionMapKey)
	}

	const requiredMapKey = "required"
	if raw, ok := proxy[requiredMapKey]; ok {
		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field required`)
		}
		mutator.Required(decoded)
		delete(proxy, requiredMapKey)
	}

	const inMapKey = "in"
	if raw, ok := proxy[inMapKey]; ok {
		var decoded Location
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field in`)
		}
		mutator.In(decoded)
		delete(proxy, inMapKey)
	}

	const schemaMapKey = "schema"
	if raw, ok := proxy[schemaMapKey]; ok {
		var decoded schema
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Schema`)
		}

		mutator.Schema(&decoded)
		delete(proxy, schemaMapKey)
	}

	const typMapKey = "type"
	if raw, ok := proxy[typMapKey]; ok {
		var decoded PrimitiveType
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field type`)
		}
		mutator.Type(decoded)
		delete(proxy, typMapKey)
	}

	const formatMapKey = "format"
	if raw, ok := proxy[formatMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field format`)
		}
		mutator.Format(decoded)
		delete(proxy, formatMapKey)
	}

	const titleMapKey = "title"
	if raw, ok := proxy[titleMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field title`)
		}
		mutator.Title(decoded)
		delete(proxy, titleMapKey)
	}

	const allowEmptyValueMapKey = "allowEmptyValue"
	if raw, ok := proxy[allowEmptyValueMapKey]; ok {
		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field allowEmptyValue`)
		}
		mutator.AllowEmptyValue(decoded)
		delete(proxy, allowEmptyValueMapKey)
	}

	const itemsMapKey = "items"
	if raw, ok := proxy[itemsMapKey]; ok {
		var decoded items
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Items`)
		}

		mutator.Items(&decoded)
		delete(proxy, itemsMapKey)
	}

	const collectionFormatMapKey = "collectionFormat"
	if raw, ok := proxy[collectionFormatMapKey]; ok {
		var decoded CollectionFormat
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field collectionFormat`)
		}
		mutator.CollectionFormat(decoded)
		delete(proxy, collectionFormatMapKey)
	}

	const defaultValueMapKey = "default"
	if raw, ok := proxy[defaultValueMapKey]; ok {
		var decoded interface{}
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field default`)
		}
		mutator.Default(decoded)
		delete(proxy, defaultValueMapKey)
	}

	const maximumMapKey = "maximum"
	if raw, ok := proxy[maximumMapKey]; ok {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field maximum`)
		}
		mutator.Maximum(decoded)
		delete(proxy, maximumMapKey)
	}

	const exclusiveMaximumMapKey = "exclusiveMaximum"
	if raw, ok := proxy[exclusiveMaximumMapKey]; ok {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field exclusiveMaximum`)
		}
		mutator.ExclusiveMaximum(decoded)
		delete(proxy, exclusiveMaximumMapKey)
	}

	const minimumMapKey = "minimum"
	if raw, ok := proxy[minimumMapKey]; ok {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field minimum`)
		}
		mutator.Minimum(decoded)
		delete(proxy, minimumMapKey)
	}

	const exclusiveMinimumMapKey = "exclusiveMinimum"
	if raw, ok := proxy[exclusiveMinimumMapKey]; ok {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field exclusiveMinimum`)
		}
		mutator.ExclusiveMinimum(decoded)
		delete(proxy, exclusiveMinimumMapKey)
	}

	const maxLengthMapKey = "maxLength"
	if raw, ok := proxy[maxLengthMapKey]; ok {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field maxLength`)
		}
		mutator.MaxLength(decoded)
		delete(proxy, maxLengthMapKey)
	}

	const minLengthMapKey = "minLength"
	if raw, ok := proxy[minLengthMapKey]; ok {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field minLength`)
		}
		mutator.MinLength(decoded)
		delete(proxy, minLengthMapKey)
	}

	const patternMapKey = "pattern"
	if raw, ok := proxy[patternMapKey]; ok {
		var decoded string
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field pattern`)
		}
		mutator.Pattern(decoded)
		delete(proxy, patternMapKey)
	}

	const maxItemsMapKey = "maxItems"
	if raw, ok := proxy[maxItemsMapKey]; ok {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field maxItems`)
		}
		mutator.MaxItems(decoded)
		delete(proxy, maxItemsMapKey)
	}

	const minItemsMapKey = "minItems"
	if raw, ok := proxy[minItemsMapKey]; ok {
		var decoded int
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field minItems`)
		}
		mutator.MinItems(decoded)
		delete(proxy, minItemsMapKey)
	}

	const uniqueItemsMapKey = "uniqueItems"
	if raw, ok := proxy[uniqueItemsMapKey]; ok {
		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field uniqueItems`)
		}
		mutator.UniqueItems(decoded)
		delete(proxy, uniqueItemsMapKey)
	}

	const enumMapKey = "enum"
	if raw, ok := proxy[enumMapKey]; ok {
		var decoded InterfaceList
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Enum`)
		}
		for _, elem := range decoded {
			mutator.Enum(elem)
		}
		delete(proxy, enumMapKey)
	}

	const multipleOfMapKey = "multipleOf"
	if raw, ok := proxy[multipleOfMapKey]; ok {
		var decoded float64
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field multipleOf`)
		}
		mutator.MultipleOf(decoded)
		delete(proxy, multipleOfMapKey)
	}

	for name, raw := range proxy {
		if strings.HasPrefix(name, `x-`) {
			var ext interface{}
			if err := json.Unmarshal(raw, &ext); err != nil {
				return errors.Wrapf(err, `failed to unmarshal field %s`, name)
			}
			mutator.Extension(name, ext)
		}
	}

	if err := mutator.Apply(); err != nil {
		return errors.Wrap(err, `failed to  unmarshal JSON`)
	}
	return nil
}

func (v *parameter) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "name":
		target = v.name
	case "description":
		target = v.description
	case "required":
		target = v.required
	case "in":
		target = v.in
	case "schema":
		target = v.schema
	case "type":
		target = v.typ
	case "format":
		target = v.format
	case "title":
		target = v.title
	case "allowEmptyValue":
		target = v.allowEmptyValue
	case "items":
		target = v.items
	case "collectionFormat":
		target = v.collectionFormat
	case "default":
		target = v.defaultValue
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
	case "enum":
		target = v.enum
	case "multipleOf":
		target = v.multipleOf
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

// ParameterFromJSON constructs a Parameter from JSON buffer. `dst` must
// be a pointer to `Parameter`
func ParameterFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Parameter)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Parameter, but got %T`, dst)
	}
	var tmp parameter
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Parameter`)
	}
	*v = &tmp
	return nil
}
