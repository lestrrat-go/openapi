package openapi

// This file was automatically generated by genbuilders.go on 2018-05-24T12:53:28+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"strings"
)

var _ = log.Printf
var _ = errors.Cause

type discriminatorMarshalProxy struct {
	Reference    string            `json:"$ref,omitempty"`
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping"`
}

type discriminatorUnmarshalProxy struct {
	Reference    string `json:"$ref,omitempty"`
	PropertyName string `json:"propertyName"`
}

func (v *discriminator) MarshalJSON() ([]byte, error) {
	var proxy discriminatorMarshalProxy
	if len(v.reference) > 0 {
		proxy.Reference = v.reference
		return json.Marshal(proxy)
	}
	proxy.PropertyName = v.propertyName
	proxy.Mapping = v.mapping
	return json.Marshal(proxy)
}

func (v *discriminator) UnmarshalJSON(data []byte) error {
	var proxy discriminatorUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if len(proxy.Reference) > 0 {
		v.reference = proxy.Reference
		return nil
	}
	v.propertyName = proxy.PropertyName
	return nil
}

func (v *discriminator) Resolve(resolver *Resolver) error {
	log.Printf(`discriminator.Resolve`)
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*discriminator)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Discriminator, but got %T`, resolved)
		}
		log.Printf(`Setting fields for discriminator to those of resolved object`)
		mutator := MutateDiscriminator(v)
		mutator.PropertyName(asserted.PropertyName())
		for iter := asserted.Mapping(); iter.Next(); {
			key, item := iter.Item()
			mutator.Mapping(key, item)
		}
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.mapping != nil {
		if err := v.mapping.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Mapping`)
		}
	}
	return nil
}

func (v *discriminator) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "propertyName":
		target = v.propertyName
	case "mapping":
		target = v.mapping
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
