package builder

import (
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Build builds html files from markdown files
func Build(sourcePath string, outputPath string) {
	// Create directory to store genereated posts
	createDir(outputPath)
	createDir(path.Join(outputPath, "posts"))

	postsPath := path.Join(sourcePath, "posts")

	files, err := filesInDir(postsPath, "*.md")
	if err != nil {
		log.Fatal("Unable to get posts from source")
	}

	tmpl := template.Must(template.ParseFiles(
		path.Join(sourcePath, "templates", "post.tmpl.html"),
	))
	for _, filename := range files {
		// ignore index file
		if filepath.Base(filename) == "_index.md" {
			continue
		}

		post := &Post{}
		if err := post.ReadMarkdown(filename); err != nil {
			log.Fatalln(err)
			return
		}

		outputFilename := strings.TrimSuffix(filepath.Base(filename), ".md") + ".html"
		var outputFile *os.File
		if outputFile, err = os.Create(path.Join(outputPath, "posts", outputFilename)); err != nil {
			log.Fatalln(err)
			return
		}
		defer outputFile.Close()

		if tmpl.ExecuteTemplate(outputFile, "post.tmpl.html", post); err != nil {
			log.Fatalln("Unable to execute template", err)
		}
	}
}

func filesInDir(dirName string, pattern string) ([]string, error) {
	filenames, err := filepath.Glob(path.Join(dirName, pattern))
	return filenames, err
}

func createDir(dirPath string) error {
	if stat, err := os.Stat(dirPath); err != nil && !os.IsExist(err) || !stat.IsDir() {
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
