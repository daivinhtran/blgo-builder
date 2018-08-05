package builder

import (
	"bytes"

	yaml "gopkg.in/yaml.v2"
)

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
