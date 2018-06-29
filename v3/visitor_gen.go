package openapi

// This file was automatically generated.
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"context"

	"github.com/pkg/errors"
)

var _ = errors.Cause

var ErrVisitAbort = errors.New(`visit aborted (non-error)`)

type callbackVisitorCtxKey struct{}

type componentsVisitorCtxKey struct{}

type contactVisitorCtxKey struct{}

type discriminatorVisitorCtxKey struct{}

type encodingVisitorCtxKey struct{}

type exampleVisitorCtxKey struct{}

type externalDocumentationVisitorCtxKey struct{}

type headerVisitorCtxKey struct{}

type infoVisitorCtxKey struct{}

type licenseVisitorCtxKey struct{}

type linkVisitorCtxKey struct{}

type mediaTypeVisitorCtxKey struct{}

type oauthFlowVisitorCtxKey struct{}

type oauthFlowsVisitorCtxKey struct{}

type openapiVisitorCtxKey struct{}

type operationVisitorCtxKey struct{}

type parameterVisitorCtxKey struct{}

type pathItemVisitorCtxKey struct{}

type pathsVisitorCtxKey struct{}

type requestBodyVisitorCtxKey struct{}

type responseVisitorCtxKey struct{}

type responsesVisitorCtxKey struct{}

type schemaVisitorCtxKey struct{}

type securityRequirementVisitorCtxKey struct{}

type securitySchemeVisitorCtxKey struct{}

type serverVisitorCtxKey struct{}

type serverVariableVisitorCtxKey struct{}

type tagVisitorCtxKey struct{}

// Visit allows you to traverse through the OpenAPI structure
func Visit(ctx context.Context, handler, elem interface{}) error {
	if v, ok := handler.(CallbackVisitor); ok {
		ctx = context.WithValue(ctx, callbackVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ComponentsVisitor); ok {
		ctx = context.WithValue(ctx, componentsVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ContactVisitor); ok {
		ctx = context.WithValue(ctx, contactVisitorCtxKey{}, v)
	}

	if v, ok := handler.(DiscriminatorVisitor); ok {
		ctx = context.WithValue(ctx, discriminatorVisitorCtxKey{}, v)
	}

	if v, ok := handler.(EncodingVisitor); ok {
		ctx = context.WithValue(ctx, encodingVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ExampleVisitor); ok {
		ctx = context.WithValue(ctx, exampleVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ExternalDocumentationVisitor); ok {
		ctx = context.WithValue(ctx, externalDocumentationVisitorCtxKey{}, v)
	}

	if v, ok := handler.(HeaderVisitor); ok {
		ctx = context.WithValue(ctx, headerVisitorCtxKey{}, v)
	}

	if v, ok := handler.(InfoVisitor); ok {
		ctx = context.WithValue(ctx, infoVisitorCtxKey{}, v)
	}

	if v, ok := handler.(LicenseVisitor); ok {
		ctx = context.WithValue(ctx, licenseVisitorCtxKey{}, v)
	}

	if v, ok := handler.(LinkVisitor); ok {
		ctx = context.WithValue(ctx, linkVisitorCtxKey{}, v)
	}

	if v, ok := handler.(MediaTypeVisitor); ok {
		ctx = context.WithValue(ctx, mediaTypeVisitorCtxKey{}, v)
	}

	if v, ok := handler.(OAuthFlowVisitor); ok {
		ctx = context.WithValue(ctx, oauthFlowVisitorCtxKey{}, v)
	}

	if v, ok := handler.(OAuthFlowsVisitor); ok {
		ctx = context.WithValue(ctx, oauthFlowsVisitorCtxKey{}, v)
	}

	if v, ok := handler.(OpenAPIVisitor); ok {
		ctx = context.WithValue(ctx, openapiVisitorCtxKey{}, v)
	}

	if v, ok := handler.(OperationVisitor); ok {
		ctx = context.WithValue(ctx, operationVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ParameterVisitor); ok {
		ctx = context.WithValue(ctx, parameterVisitorCtxKey{}, v)
	}

	if v, ok := handler.(PathItemVisitor); ok {
		ctx = context.WithValue(ctx, pathItemVisitorCtxKey{}, v)
	}

	if v, ok := handler.(PathsVisitor); ok {
		ctx = context.WithValue(ctx, pathsVisitorCtxKey{}, v)
	}

	if v, ok := handler.(RequestBodyVisitor); ok {
		ctx = context.WithValue(ctx, requestBodyVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ResponseVisitor); ok {
		ctx = context.WithValue(ctx, responseVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ResponsesVisitor); ok {
		ctx = context.WithValue(ctx, responsesVisitorCtxKey{}, v)
	}

	if v, ok := handler.(SchemaVisitor); ok {
		ctx = context.WithValue(ctx, schemaVisitorCtxKey{}, v)
	}

	if v, ok := handler.(SecurityRequirementVisitor); ok {
		ctx = context.WithValue(ctx, securityRequirementVisitorCtxKey{}, v)
	}

	if v, ok := handler.(SecuritySchemeVisitor); ok {
		ctx = context.WithValue(ctx, securitySchemeVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ServerVisitor); ok {
		ctx = context.WithValue(ctx, serverVisitorCtxKey{}, v)
	}

	if v, ok := handler.(ServerVariableVisitor); ok {
		ctx = context.WithValue(ctx, serverVariableVisitorCtxKey{}, v)
	}

	if v, ok := handler.(TagVisitor); ok {
		ctx = context.WithValue(ctx, tagVisitorCtxKey{}, v)
	}

	return visit(ctx, elem)
}

func visit(ctx context.Context, elem interface{}) error {
	switch elem := elem.(type) {
	case Callback:
		return visitCallback(ctx, elem)
	case Components:
		return visitComponents(ctx, elem)
	case Contact:
		return visitContact(ctx, elem)
	case Discriminator:
		return visitDiscriminator(ctx, elem)
	case Encoding:
		return visitEncoding(ctx, elem)
	case Example:
		return visitExample(ctx, elem)
	case ExternalDocumentation:
		return visitExternalDocumentation(ctx, elem)
	case Header:
		return visitHeader(ctx, elem)
	case Info:
		return visitInfo(ctx, elem)
	case License:
		return visitLicense(ctx, elem)
	case Link:
		return visitLink(ctx, elem)
	case MediaType:
		return visitMediaType(ctx, elem)
	case OAuthFlow:
		return visitOAuthFlow(ctx, elem)
	case OAuthFlows:
		return visitOAuthFlows(ctx, elem)
	case OpenAPI:
		return visitOpenAPI(ctx, elem)
	case Operation:
		return visitOperation(ctx, elem)
	case Parameter:
		return visitParameter(ctx, elem)
	case PathItem:
		return visitPathItem(ctx, elem)
	case Paths:
		return visitPaths(ctx, elem)
	case RequestBody:
		return visitRequestBody(ctx, elem)
	case Response:
		return visitResponse(ctx, elem)
	case Responses:
		return visitResponses(ctx, elem)
	case Schema:
		return visitSchema(ctx, elem)
	case SecurityRequirement:
		return visitSecurityRequirement(ctx, elem)
	case SecurityScheme:
		return visitSecurityScheme(ctx, elem)
	case Server:
		return visitServer(ctx, elem)
	case ServerVariable:
		return visitServerVariable(ctx, elem)
	case Tag:
		return visitTag(ctx, elem)
	default:
		return errors.Errorf(`unknown element %T`, elem)
	}
}
