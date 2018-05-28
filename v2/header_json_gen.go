package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T19:20:54+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var _ = log.Printf
var _ = json.Unmarshal
var _ = errors.Cause

type headerMarshalProxy struct {
	Reference        string           `json:"$ref,omitempty"`
	Description      string           `json:"description,omitempty"`
	Type             string           `json:"type"`
	Format           string           `json:"format,omitempty"`
	Items            Items            `json:"items,omitempty"`
	CollectionFormat CollectionFormat `json:"collectionFormat,omitempty"`
	DefaultValue     interface{}      `json:"default,omitempty"`
	Maximum          float64          `json:"maximum,omitempty"`
	ExclusiveMaximum float64          `json:"exclusiveMaximum,omitempty"`
	Minimum          float64          `json:"minimum,omitempty"`
	ExclusiveMinimum float64          `json:"exclusiveMinimum,omitempty"`
	MaxLength        int              `json:"maxLength,omitempty"`
	MinLength        int              `json:"minLength,omitempty"`
	Pattern          string           `json:"pattern,omitempty"`
	MaxItems         int              `json:"maxItems,omitempty"`
	MinItems         int              `json:"minItems,omitempty"`
	UniqueItems      bool             `json:"uniqueItems,omitempty"`
	Enum             InterfaceList    `json:"enum,omitempty"`
	MultipleOf       float64          `json:"multipleOf,omitempty"`
}

type headerUnmarshalProxy struct {
	Reference        string           `json:"$ref,omitempty"`
	Description      string           `json:"description,omitempty"`
	Type             string           `json:"type"`
	Format           string           `json:"format,omitempty"`
	Items            json.RawMessage  `json:"items,omitempty"`
	CollectionFormat CollectionFormat `json:"collectionFormat,omitempty"`
	DefaultValue     interface{}      `json:"default,omitempty"`
	Maximum          float64          `json:"maximum,omitempty"`
	ExclusiveMaximum float64          `json:"exclusiveMaximum,omitempty"`
	Minimum          float64          `json:"minimum,omitempty"`
	ExclusiveMinimum float64          `json:"exclusiveMinimum,omitempty"`
	MaxLength        int              `json:"maxLength,omitempty"`
	MinLength        int              `json:"minLength,omitempty"`
	Pattern          string           `json:"pattern,omitempty"`
	MaxItems         int              `json:"maxItems,omitempty"`
	MinItems         int              `json:"minItems,omitempty"`
	UniqueItems      bool             `json:"uniqueItems,omitempty"`
	Enum             InterfaceList    `json:"enum,omitempty"`
	MultipleOf       float64          `json:"multipleOf,omitempty"`
}

func (v *header) MarshalJSON() ([]byte, error) {
	var proxy headerMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.Description = v.description
	proxy.Type = v.typ
	proxy.Format = v.format
	proxy.Items = v.items
	proxy.CollectionFormat = v.collectionFormat
	proxy.DefaultValue = v.defaultValue
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
	return json.Marshal(proxy)
}

func (v *header) UnmarshalJSON(data []byte) error {
	var proxy headerUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.description = proxy.Description
	v.typ = proxy.Type
	v.format = proxy.Format

	if len(proxy.Items) > 0 {
		var decoded items
		if err := json.Unmarshal(proxy.Items, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Items`)
		}

		v.items = &decoded
	}
	v.collectionFormat = proxy.CollectionFormat
	v.defaultValue = proxy.DefaultValue
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
	v.enum = proxy.Enum
	v.multipleOf = proxy.MultipleOf
	return nil
}

func (v *header) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*header)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Header, but got %T`, resolved)
		}
		mutator := MutateHeader(v)
		mutator.Name(asserted.Name())
		mutator.Description(asserted.Description())
		mutator.Type(asserted.Type())
		mutator.Format(asserted.Format())
		mutator.Items(asserted.Items())
		mutator.CollectionFormat(asserted.CollectionFormat())
		mutator.DefaultValue(asserted.DefaultValue())
		mutator.Maximum(asserted.Maximum())
		mutator.ExclusiveMaximum(asserted.ExclusiveMaximum())
		mutator.Minimum(asserted.Minimum())
		mutator.ExclusiveMinimum(asserted.ExclusiveMinimum())
		mutator.MaxLength(asserted.MaxLength())
		mutator.MinLength(asserted.MinLength())
		mutator.Pattern(asserted.Pattern())
		mutator.MaxItems(asserted.MaxItems())
		mutator.MinItems(asserted.MinItems())
		mutator.UniqueItems(asserted.UniqueItems())
		for iter := asserted.Enum(); iter.Next(); {
			item := iter.Item()
			mutator.Enum(item)
		}
		mutator.MultipleOf(asserted.MultipleOf())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.items != nil {
		if err := v.items.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Items`)
		}
	}
	if v.enum != nil {
		if err := v.enum.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Enum`)
		}
	}
	return nil
}

func (v *header) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "description":
		target = v.description
	case "type":
		target = v.typ
	case "format":
		target = v.format
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
