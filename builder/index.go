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

// ReadMarkdown reads markdown file to update the receiver index
func (i *Index) ReadMarkdown(filename string) error {
	markdownBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	frontMatter, err := parseFrontmatter(&markdownBytes)
	if err != nil {
		return err
	}

	if name, ok := frontMatter["name"]; ok {
		i.Name = name.(string)
	} else {
		return fmt.Errorf("could not read the name from %v", filename)
	}
	return nil
}
