openapi
=======

[WIP] OpenAPI for Go

[![Build Status](https://travis-ci.org/lestrrat-go/openapi.svg?branch=master)](https://travis-ci.org/lestrrat-go/openapi)

[![GoDoc](https://godoc.org/github.com/lestrrat-go/openapi?status.svg)](https://godoc.org/github.com/lestrrat-go/openapi)

# Programmatically Build an OpenAPI spec

* API is geared towards regularity and to be easier for a program to generate this code
* Use the `builder` API to build a spec
* Each `builder` constructor takes in the required arguments for that component
* Finalize and build the OpenAPI nodes by calling Build() at the end

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

# TODO

* Disable direct access to composite types such as `map[string]Schema`