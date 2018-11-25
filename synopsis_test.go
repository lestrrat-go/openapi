package openapi_test

import (
	"log"
	"os"

	v2 "github.com/lestrrat-go/openapi/v2"
)

func ExampleV2() {
	const src = `spec/examples/v2.0/petstore-extended.yaml`
	f, err := os.Open(src)
	if err != nil {
		log.Printf(`failed to open file %s: %s`, src, err)
		return
	}

	spec, err := v2.ParseYAML(f)
	if err != nil {
		log.Printf(`failed to parse file %s: %s`, src, err)
		return
	}

	for pathItemIter := spec.Paths().Paths(); pathItemIter.Next(); {
		path, pathItem := pathItemIter.Item()
		for operIter := pathItem.Operations(); operIter.Next(); {
			oper := operIter.Item()
			log.Printf("%s %s", oper.Verb(), path)
		}
	}
}
