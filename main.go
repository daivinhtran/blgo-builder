package main

import (
	"blgo-builder/builder"
	"os"
	"path"
)

func main() {
	cwd, _ := os.Getwd()
	sourcePath := path.Join(cwd, "source")
	outputPath := path.Join(cwd, "output")
	builder.Build(sourcePath, outputPath)
}
