package openapi

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var rxHostPortOnly = regexp.MustCompile(`^[^:/]+(:\d+)?$`)

func (v *swagger) Validate(recurse bool) error {
	if v.version != defaultSwaggerVersion {
		return errors.Errorf(`swagger field must be %s (got %s)`, strconv.Quote(defaultSwaggerVersion), strconv.Quote(v.version))
	}

	if v.info == nil {
		return errors.New(`info is required`)
	}

	if v.paths == nil {
		return errors.New(`paths is required`)
	}

	// The host (name or ip) serving the API. This MUST be the host
	// only and does not include the scheme nor sub-paths. It MAY
	// include a port. If the host is not included, the host serving
	// the documentation is to be used (including the port). The
	// host does not support path templating.
	if s := v.host; len(s) > 0 {
		if !rxHostPortOnly.MatchString(s) {
			return errors.New(`host field must be either "host" or "host:port"`)
		}
	}

	if s := v.basePath; len(s) > 0 {
		if !strings.HasPrefix(s, "/") {
			return errors.New(`basePath must start with a slash (/)`)
		}
	}

	if recurse {
		return v.recurseValidate()
	}

	return nil
}
