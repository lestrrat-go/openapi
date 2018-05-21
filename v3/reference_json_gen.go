package openapi

// This file was automatically generated by genbuilders.go on 2018-05-21T19:54:19+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
)

var _ = errors.Cause

type referenceMarshalProxy struct {
}

type referenceUnmarshalProxy struct {
}

func (v *reference) MarshalJSON() ([]byte, error) {
	var proxy referenceMarshalProxy
	return json.Marshal(proxy)
}

func (v *reference) UnmarshalJSON(data []byte) error {
	var proxy referenceUnmarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	return nil
}

func (v *reference) QueryJSON(path string) (ret interface{}, ok bool) {
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
