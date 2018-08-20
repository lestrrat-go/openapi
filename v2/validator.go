package openapi

import (
	"context"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type validator struct {
	recurse bool
}

func newValidator(recurse bool) *validator {
	return &validator{recurse: recurse}
}

var rxHostPortOnly = regexp.MustCompile(`^[^:/]+(:\d+)?$`)

func (val *validator) VisitSwagger(ctx context.Context, v Swagger) error {
	if v.Version() != defaultSwaggerVersion {
		return errors.Errorf(`swagger field must be %s (got %s)`, strconv.Quote(defaultSwaggerVersion), strconv.Quote(v.Version()))
	}

	if v.Info() == nil {
		return errors.New(`info is required`)
	}

	if v.Paths() == nil {
		return errors.New(`paths is required`)
	}

	// The host (name or ip) serving the API. This MUST be the host
	// only and does not include the scheme nor sub-paths. It MAY
	// include a port. If the host is not included, the host serving
	// the documentation is to be used (including the port). The
	// host does not support path templating.
	if s := v.Host(); len(s) > 0 {
		if !rxHostPortOnly.MatchString(s) {
			return errors.New(`host field must be either "host" or "host:port"`)
		}
	}

	if s := v.BasePath(); len(s) > 0 {
		if !strings.HasPrefix(s, "/") {
			return errors.New(`basePath must start with a slash (/)`)
		}
	}

	return nil
}

func (val *validator) VisitInfo(ctx context.Context, v Info) error {
	// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#infoObject
	if v.Title() == "" {
		return errors.New(`missing required field "title"`)
	}

	if v.Version() == "" {
		return errors.New(`missing required field "version"`)
	}

	return nil
}

func (val *validator) VisitLicense(ctx context.Context, v License) error {
	// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#licenseObject
	if v.Name() == "" {
		return errors.New(`missing required field "name"`)
	}

	if uv := v.URL(); uv != "" {
		if _, err := url.Parse(uv); err != nil {
			return errors.Wrap(err, `field "url" must be a valid URL`)
		}
	}

	return nil
}

func (val *validator) VisitPaths(ctx context.Context, v Paths) error {
	for iter := v.Paths(); iter.Next(); {
		key, _ := iter.Item()
		if !strings.HasPrefix(key, "/") {
			return errors.Errorf(`invalid path item key "%s" (paths must start with a slash)`, key)
		}
	}
	return nil
}

func (val *validator) VisitPathItem(ctx context.Context, v PathItem) error {
	// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#pathItemObject
	seen := make(map[string]struct{})
	for iter := v.Parameters(); iter.Next(); {
		param := iter.Item()
		key := param.Name() + "\000" + string(param.In())
		if _, ok := seen[key]; ok {
			return errors.Errorf(`duplicate path item name = "%s", location = %s"`, param.Name(), param.In())
		}
		seen[key] = struct{}{}
	}
	return nil
}

func (val *validator) VisitOperation(ctx context.Context, v Operation) error {
	if v.Responses() == nil {
		return errors.New(`missing required field "responses"`)
	}

	inMap := make(map[Location][]string) // map of parameter location to param name
	for piter := v.Parameters(); piter.Next(); {
		param := piter.Item()
		inMap[param.In()] = append(inMap[param.In()], param.Name())
	}

	if names, ok := inMap[InBody]; ok {
		if len(names) > 1 {
			return errors.Errorf(`there can only be 1 body parameter got %v`, names)
		}

		// XXX this check is transitive. this case alone will suffice
		if formNames, ok := inMap[InForm]; ok {
			return errors.Errorf(`both "body" and "formData" parameters are present (can only have either) body = %v, formData = %v`, names, formNames)
		}
	}

	return nil
}

func (val *validator) VisitParameter(ctx context.Context, v Parameter) error {
	// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#parameterObject
	if len(v.Name()) == 0 {
		return errors.New(`invalid parameter: "name" field is required`)
	}

	if !validLocation(v.In()) {
		return errors.Errorf(`invalid parameter: invalid value in "in" field: %s`, strconv.Quote(string(v.In())))
	}

	if v.In() == InBody {
		if v.Schema() == nil {
			return errors.New(`invalid parameter: for parameters with {in: body}, you must specify the "schema" field`)
		}
		return nil
	}

	if v.AllowEmptyValue() {
		if v.In() != InQuery && v.In() != InForm {
			return errors.Errorf(`invalid parameter: {allowEmptyValue: true} is only applicable for "query" or "formData" parameters: got %s`, v.In())
		}
	}

	switch v.Type() {
	case String, Number, Integer, Boolean:
	case Array:
		if v.Items() == nil {
			return errors.Errorf(`invalid paramter: for {type: array}, "items" field must be specified`)
		}
	case File:
		if v.In() != InForm {
			return errors.Errorf(`invalid parameter: for {type: file}, "in" field must be "formData" (got %s)`, v.In())
		}
	default:
		return errors.Errorf(`invalid parameter: type must be one of "string", "number", "integer", "boolean", "array" or "file" (got %s)`, v.Type())
	}
	return nil
}

func (val *validator) VisitResponse(ctx context.Context, v Response) error {
	if len(v.Description()) == 0 {
		return errors.New(`response description is required`)
	}
	return nil
}

func (val *validator) VisitSchema(ctx context.Context, v Schema) error {
	if v.Reference() != "" {
		return nil
	}

	if typ := v.Type(); typ != "" && !typ.IsValid() {
		return errors.Errorf(`invalid type in schema: %s`, typ)
	}

	return nil
}
