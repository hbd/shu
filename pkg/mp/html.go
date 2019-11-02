package mp

import "fmt"

// ToHTML produces escaped HTML from Markdown.
func ToHTML(md Markdown) string {
	var html string

	for _, line := range md.Lines {
		if line.Indented {
			html += fmt.Sprintf("<pre><code>%s</code></pre>", line.Str)
		}
		html += fmt.Sprintf("\n%s", line.Str)
	}

	return html
}
