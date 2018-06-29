package openapi

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

func (v *responses) QueryJSON(path string) (ret interface{}, ok bool) {
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
	case "default":
		target = v.defaultValue
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

// ResponsesFromJSON constructs a Responses from JSON buffer. `dst` must
// be a pointer to `Responses`
func ResponsesFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Responses)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Responses, but got %T`, dst)
	}
	var tmp responses
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Responses`)
	}
	*v = &tmp
	return nil
}
