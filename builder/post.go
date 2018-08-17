package builder

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/russross/blackfriday"
)

// Post represents a single blog post
type Post struct {
	Body          template.HTML
	Date          time.Time
	RelativeLink  string
	CanonicalLink string
	Title         string
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
	p.RelativeLink = "/posts/" + strings.TrimSuffix(filepath.Base(filename), ".md") + ".html"

	if canonicalLink, ok := frontMatter["canonical"]; ok {
		p.CanonicalLink = canonicalLink.(string)
	}

	return nil
}
