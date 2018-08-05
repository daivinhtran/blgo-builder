package builder

import (
	"fmt"
	"io/ioutil"
)

// Index represents global settings
type Index struct {
	Name  string
	Posts []*Post
}

func (index *Index) Len() int           { return len(index.Posts) }
func (index *Index) Swap(i, j int)      { index.Posts[i], index.Posts[j] = index.Posts[j], index.Posts[i] }
func (index *Index) Less(i, j int) bool { return index.Posts[i].Date.Before(index.Posts[j].Date) }

// ReadMarkdown reads markdown file to update the receiver index
func (index *Index) ReadMarkdown(filename string) error {
	markdownBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	frontMatter, err := parseFrontmatter(&markdownBytes)
	if err != nil {
		return err
	}

	if name, ok := frontMatter["name"]; ok {
		index.Name = name.(string)
	} else {
		return fmt.Errorf("could not read the name from %v", filename)
	}
	return nil
}
