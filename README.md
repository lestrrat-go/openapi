openapi
=======

[WIP] OpenAPI for Go

# Marshaling to JSON/YAML

* Use `json/encoding` on `entity.*` objects to marshal objets into JSON (note: no validation)
* Use `github.com/ghodss/yaml` to marshal objects into YAML (not `gopkg.in/yaml.v2`)

# Programmatically Build an OpenAPI spec

* API is geared towards regularity and to be easier for a program to generate this code
* Use the `builder` API to build a spec
* Each `builder` constructor takes in the required arguments for that component
* Finalize and build the OpenAPI nodes (`entity.*`) by calling Build() at the end