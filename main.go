package main

import (
	"fmt"

	"blgo-builder/builder"
	"blgo-builder/server"
)

func main() {
	builder.Build()
	server.Serve()
	fmt.Println("Test")
}
