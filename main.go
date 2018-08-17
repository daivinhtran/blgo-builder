package main

import (
	"blgo-builder/builder"
	"blgo-builder/server"
	"os"
	"path"
)

func main() {
	cwd, _ := os.Getwd()
	sourcePath := path.Join(cwd, "source")
	outputPath := path.Join(cwd, "output")
	builder.Build(sourcePath, outputPath)
	server.Serve(outputPath, "8080")
}
