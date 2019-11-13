package openapi2

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

func (v *paths) QueryJSON(path string) (ret interface{}, ok bool) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	if path == "" {
		return v, true
	}
	return nil, false
}

// PathsFromJSON constructs a Paths from JSON buffer. `dst` must
// be a pointer to `Paths`
func PathsFromJSON(buf []byte, dst interface{}) error {
	v, ok := dst.(*Paths)
	if !ok {
		return errors.Errorf(`dst needs to be a pointer to Paths, but got %T`, dst)
	}
	var tmp paths
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to unmarshal Paths`)
	}
	*v = &tmp
	return nil
}