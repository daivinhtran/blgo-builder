package builder

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"time"

	"github.com/russross/blackfriday"
	yaml "gopkg.in/yaml.v2"
)

// Post represents a single blog post
type Post struct {
	// Index          *Index
	// Slug           string
	// OutputFilename string
	Body template.HTML
	Date time.Time
	// Description    string
	// GUID           string
	// Link           string
	// RelativeLink   string
	Title string
	// XMLDesc        string
	// XMLTitle       string
	// Draft          bool
}

// ReadMarkdown reads markdown file to update the receiver post
func (p *Post) ReadMarkdown(filename string) error {
	markdownBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	frontMatter, err := parseFrontmatter(&markdownBytes)
	if err != nil {
		return err
	}

	if title, ok := frontMatter["title"]; ok {
		p.Title = title.(string)
	} else {
		return fmt.Errorf("could not read the title from post")
	}

	if date, ok := frontMatter["date"]; ok {
		if p.Date, err = time.Parse("2006-01-02", date.(string)); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("could not read the date from post")
	}

	p.Body = template.HTML(blackfriday.MarkdownCommon([]byte(markdownBytes)))
	return nil
}

func parseFrontmatter(body *[]byte) (map[string]interface{}, error) {
	var frontmatterBuf bytes.Buffer
	buf := bytes.NewBuffer(*body)
	started := false
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return nil, err
		}

		if line == "---\n" {
			if started {
				break
			}
			started = true
		}
		if started {
			frontmatterBuf.Write([]byte(line))
		}
	}

	*body = buf.Bytes()
	frontmatter := make(map[string]interface{})
	return frontmatter, yaml.Unmarshal(frontmatterBuf.Bytes(), &frontmatter)
}

// func (p *Post) Read(filename string, body []byte) error {

// }
