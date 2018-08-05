package builder

import (
	"os"
	"path"
	"testing"
)

func TestStuff(t *testing.T) {
	cwd, _ := os.Getwd()
	sourcePath := path.Join(cwd, "source")
	outputPath := path.Join(cwd, "output")
	Build(sourcePath, outputPath)
}
