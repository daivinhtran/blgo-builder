package builder

import (
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/otiai10/copy"
)

// Build builds html files from markdown files
func Build(sourcePath string, outputPath string) {
	index := &Index{}
	if err := index.ReadMarkdown(filepath.Join(sourcePath, "index.md")); err != nil {
		log.Fatalln(err)
		return
	}

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
		path.Join(sourcePath, "templates", "index.tmpl.html"),
	))
	var outputFile *os.File

	for _, filename := range files {

		post := &Post{}
		if err := post.ReadMarkdown(filename); err != nil {
			log.Fatalln(err)
			return
		}

		outputFilename := strings.TrimSuffix(filepath.Base(filename), ".md") + ".html"
		if outputFile, err = os.Create(path.Join(outputPath, "posts", outputFilename)); err != nil {
			log.Fatalln(err)
			return
		}
		defer outputFile.Close()

		if tmpl.ExecuteTemplate(outputFile, "post.tmpl.html", post); err != nil {
			log.Fatalln("Unable to execute template", err)
		}

		index.Posts = append(index.Posts, post)
	}

	sort.Sort(sort.Reverse(index))

	if outputFile, err = os.Create(path.Join(outputPath, "index.html")); err != nil {
		log.Fatalln(err)
	}
	if err := tmpl.ExecuteTemplate(outputFile, "index.tmpl.html", index); err != nil {
		log.Fatalln(err)
	}

	// copy assets
	copy.Copy(path.Join(sourcePath, "assets"), path.Join(outputPath, "assets"))
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
