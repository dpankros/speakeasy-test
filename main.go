package main

import (
	"github.com/speakeasy-api/easytemplate"
	"log"
)

func main() {
	// Create a new easytemplate engine.
	engine := easytemplate.New()
	data := 0 //context.Background()
	// Start the engine from a javascript entrypoint.
	err := engine.RunScript("main.js", data)
	if err != nil {
		log.Fatal(err)
	}
}
