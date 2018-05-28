//go:generate go run internal/cmd/gentypes/gentypes.go

package openapi

import "strings"

func validLocation(l Location) bool {
	switch l {
	case InPath, InQuery, InHeader, InBody, InForm:
		return true
	}
	return false
}

func extractFragFromPath(path string) (string, string) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	var frag string
	if i := strings.Index(path, `/`); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ``
	}
	return frag, path
}
