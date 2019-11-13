package openapi2

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = context.Background
var _ = errors.Cause

var ErrVisitAbort = errors.New(`visit aborted (non-error)`)

type exampleMapKeyVisitorCtxKey struct{}

type headerMapKeyVisitorCtxKey struct{}

type interfaceMapKeyVisitorCtxKey struct{}

type parameterMapKeyVisitorCtxKey struct{}

type pathItemMapKeyVisitorCtxKey struct{}

type responseMapKeyVisitorCtxKey struct{}

type schemaMapKeyVisitorCtxKey struct{}

type scopesMapKeyVisitorCtxKey struct{}

type securitySchemeMapKeyVisitorCtxKey struct{}

type stringMapKeyVisitorCtxKey struct{}

type swaggerVisitorCtxKey struct{}

type infoVisitorCtxKey struct{}

type contactVisitorCtxKey struct{}

type licenseVisitorCtxKey struct{}

type pathsVisitorCtxKey struct{}

type pathItemVisitorCtxKey struct{}

type operationVisitorCtxKey struct{}

type externalDocumentationVisitorCtxKey struct{}

type parameterVisitorCtxKey struct{}

type itemsVisitorCtxKey struct{}

type responsesVisitorCtxKey struct{}

type responseVisitorCtxKey struct{}

type headerVisitorCtxKey struct{}

type schemaVisitorCtxKey struct{}

type xmlVisitorCtxKey struct{}

type securitySchemeVisitorCtxKey struct{}

type securityRequirementVisitorCtxKey struct{}

type tagVisitorCtxKey struct{}

// Visit allows you to traverse through the OpenAPI structure
func Visit(ctx context.Context, handler, elem interface{}) error {
	if v, ok := handler.(SwaggerVisitor); ok {
		ctx = context.WithValue(ctx, swaggerVisitorCtxKey{}, v)
	}

	if v, ok := handler.(InfoVisitor); ok {
		ctx = context.WithValue(ctx, infoVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ContactVisitor); ok {
		ctx = context.WithValue(ctx, contactVisitorCtxKey{}, v)
	}

	if v, ok := handler.(LicenseVisitor); ok {
		ctx = context.WithValue(ctx, licenseVisitorCtxKey{}, v)
	}

	if v, ok := handler.(PathsVisitor); ok {
		ctx = context.WithValue(ctx, pathsVisitorCtxKey{}, v)
	}

	if v, ok := handler.(PathItemVisitor); ok {
		ctx = context.WithValue(ctx, pathItemVisitorCtxKey{}, v)
	}

	if v, ok := handler.(OperationVisitor); ok {
		ctx = context.WithValue(ctx, operationVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ExternalDocumentationVisitor); ok {
		ctx = context.WithValue(ctx, externalDocumentationVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ParameterVisitor); ok {
		ctx = context.WithValue(ctx, parameterVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ItemsVisitor); ok {
		ctx = context.WithValue(ctx, itemsVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ResponsesVisitor); ok {
		ctx = context.WithValue(ctx, responsesVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ResponseVisitor); ok {
		ctx = context.WithValue(ctx, responseVisitorCtxKey{}, v)
	}

	if v, ok := handler.(HeaderVisitor); ok {
		ctx = context.WithValue(ctx, headerVisitorCtxKey{}, v)
	}

	if v, ok := handler.(SchemaVisitor); ok {
		ctx = context.WithValue(ctx, schemaVisitorCtxKey{}, v)
	}

	if v, ok := handler.(XMLVisitor); ok {
		ctx = context.WithValue(ctx, xmlVisitorCtxKey{}, v)
	}

	if v, ok := handler.(SecuritySchemeVisitor); ok {
		ctx = context.WithValue(ctx, securitySchemeVisitorCtxKey{}, v)
	}

	if v, ok := handler.(SecurityRequirementVisitor); ok {
		ctx = context.WithValue(ctx, securityRequirementVisitorCtxKey{}, v)
	}

	if v, ok := handler.(TagVisitor); ok {
		ctx = context.WithValue(ctx, tagVisitorCtxKey{}, v)
	}

	return visit(ctx, elem)
}

func visit(ctx context.Context, elem interface{}) error {
	switch elem := elem.(type) {
	case Swagger:
		return visitSwagger(ctx, elem)
	case Info:
		return visitInfo(ctx, elem)
	case Contact:
		return visitContact(ctx, elem)
	case License:
		return visitLicense(ctx, elem)
	case Paths:
		return visitPaths(ctx, elem)
	case PathItem:
		return visitPathItem(ctx, elem)
	case Operation:
		return visitOperation(ctx, elem)
	case ExternalDocumentation:
		return visitExternalDocumentation(ctx, elem)
	case Parameter:
		return visitParameter(ctx, elem)
	case Items:
		return visitItems(ctx, elem)
	case Responses:
		return visitResponses(ctx, elem)
	case Response:
		return visitResponse(ctx, elem)
	case Header:
		return visitHeader(ctx, elem)
	case Schema:
		return visitSchema(ctx, elem)
	case XML:
		return visitXML(ctx, elem)
	case SecurityScheme:
		return visitSecurityScheme(ctx, elem)
	case SecurityRequirement:
		return visitSecurityRequirement(ctx, elem)
	case Tag:
		return visitTag(ctx, elem)
	default:
		return errors.Errorf(`unknown element %T`, elem)
	}
}
