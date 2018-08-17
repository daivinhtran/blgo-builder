package server

import (
	"os"
	"path"
	"testing"
)

func TestBuilder(t *testing.T) {
	cwd, _ := os.Getwd()
	outputPath := path.Join(cwd, "output")
	Serve(outputPath, "8080")
}
