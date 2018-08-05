package builder

import "time"

// Index represents global settings/variables and the index of the posts
// the index.html will generated from Index
type Index struct {
	Title     string
	Posts     []*Post
	URL       string
	XMLURL    string
	UpdatedAt time.Time
}
