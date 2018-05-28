package main

import (
	"log"
	"os"

	"github.com/lestrrat-go/openapi/v2/internal/types"
)

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

func _main() error {
	return types.GenerateCode()
}
