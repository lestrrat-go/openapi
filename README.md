openapi
=======

[WIP] OpenAPI for Go

[![Build Status](https://travis-ci.org/lestrrat-go/openapi.svg?branch=master)](https://travis-ci.org/lestrrat-go/openapi)

[![GoDoc](https://godoc.org/github.com/lestrrat-go/openapi?status.svg)](https://godoc.org/github.com/lestrrat-go/openapi)

# Status

* v2/v3 generally works
* can generate simple gRPC protobuf definition
* can generate simple Go REST client
* can generate simple ES6/flow REST client
* can lint v2/v3 documents (note: validation rules are still not 100% implemented)

# Programmatically Build an OpenAPI spec

* API is geared towards regularity and to be easier for a program to generate this code
* Use the `builder` API to build a spec
* Each `builder` constructor takes in the required arguments for that component
* Finalize and build the OpenAPI nodes by calling Build() at the end
* Properly validate structures

# Easier Navigation

OpenAPI by design requires you to remember context as you navigate through the structure.
For example, when parsing `paths` component, when you eventually get to the `operation`
component you need to have remembered 

1. The name of the path you are working on
2. The HTTP verb that the operation is for

This is because at each level the structure does not contain these information. It's
controlled and catalogued by the component one scope above, and that creates really
hard-to-work with method signatures like

```go
func processOperation(path, verb string, oper Operation) { ... }
```

Where you really want to just do

```go
func processOperation(oper Operation) { ... }
```

This library gives you tools to work with such annoyances by making sure that the
structure is immutable once the tree is created, and in doing so we can
automatically memoizing some of the information in the components towards the
inner scope.

# Validation / Lint

`oalint` command is available to lint your OpenAPI files. Compile using `go build ./cmd/oalint/main.go`

```
oalint -file=/path/to/v2spec.yaml -format=yaml
```

```
oalint -file=/path/to/v3spec.yaml -format=yaml -openapi-version=3
```

# Code Generation

`oagen` command is available to generate code. Compile using `go build ./cmd/oagen/main.go`

## Protobuf (gRPC)

Generate a protobuf (gRPC) definition file from OpenAPI spec. This feature has not been extensively tested, but in most simple cases, it should just work. Look into openapi2proto for prior arts.

```
oagen protobuf \
    -output=/path/to/file.proto \
    -global-option=go_package=myapi \
    spec.yaml
```

## Go/ES6(+Flow) REST Client

Currently `oagen restclient` command will generate a Go REST client loosely based on
Google-ish api (see https://google.golang.org/api)

```
oagen restclient \
    -target=go \ # or es6flow
    -directory=path/to/dir \
    -package=packageName \
    spec.yaml
```

# CAVEATS

*JSON Reference* (those pesky `$ref`s) are supported, but only if it's internal to the document itself.
I plan on incorporating external references, but in a separate level. I believe external references
should be resolved before we do anything with this library.