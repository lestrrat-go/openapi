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

type tagMarshalProxy struct {
	Reference    string                `json:"$ref,omitempty"`
	Name         string                `json:"name"`
	Description  string                `json:"description,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

func (v *tag) MarshalJSON() ([]byte, error) {
	var proxy tagMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Name = v.name
	proxy.Description = v.description
	proxy.ExternalDocs = v.externalDocs
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

// UnmarshalJSON defines how tag is deserialized from JSON
func (v *tag) UnmarshalJSON(data []byte) error {
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

	mutator := MutateTag(v)

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

	const externalDocsMapKey = "externalDocs"
	if raw, ok := proxy[externalDocsMapKey]; ok {
		var decoded externalDocumentation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		mutator.ExternalDocs(&decoded)
		delete(proxy, externalDocsMapKey)
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

func (v *tag) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "externalDocs":
		target = v.externalDocs
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

// TagFromJSON constructs a Tag from JSON buffer. `dst` must
// be a pointer to `Tag`
func TagFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Tag)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Tag, but got %T`, dst)
	}
	var tmp tag
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Tag`)
	}
	*v = &tmp
	return nil
}
