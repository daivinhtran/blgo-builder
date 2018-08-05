package builder

import (
	"os"
	"path"
	"testing"
)

func TestBuilder(t *testing.T) {
	cwd, _ := os.Getwd()
	sourcePath := path.Join(cwd, "test_source")
	outputPath := path.Join(cwd, "output")
	Build(sourcePath, outputPath)
}
