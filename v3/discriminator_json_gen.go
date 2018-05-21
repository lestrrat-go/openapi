package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
)

var _ = errors.Cause

type discriminatorMarshalProxy struct {
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping"`
}

type discriminatorUnmarshalProxy struct {
	PropertyName string `json:"propertyName"`
}

func (v *discriminator) MarshalJSON() ([]byte, error) {
	var proxy discriminatorMarshalProxy
	proxy.PropertyName = v.propertyName
	proxy.Mapping = v.mapping
	return json.Marshal(proxy)
}

func (v *discriminator) UnmarshalJSON(data []byte) error {
	var proxy discriminatorUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	v.propertyName = proxy.PropertyName
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
