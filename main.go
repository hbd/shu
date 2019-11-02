package main

import (
	"html/template"
	"os"
	"shu/pkg/mp"
	"time"

	"github.com/pkg/errors"
)

// HTMLDoc represents an HTML document.
type HTMLDoc struct {
	Title string
	Date  time.Time
	Body  []HTMLBody // Slice of body elements.
}

type HTMLBody struct {
	Subtitle string
	Contents string
}

const tmpl = `
<head>
  <link rel="stylesheet" type="text/css" href="./index.css">
</head>
<body>
  <p><a class="post_title">{{ .Title }}</a></p>
  <div>
    {{range .Body}}
    {{ if .Subtitle }}<b>{{ .Subtitle }}</b>{{end}}<p>{{ .Contents }}</p>{{end}}
  </div>
</body>
`

const mainpage = `

`

func main() {
	mp.Parse(" hdd\t\tdd\ne\n    l\t l\nl\ts\no\n")

	htmlDoc := HTMLDoc{
		Title: "read me",
		Date:  time.Now(),
		Body: []HTMLBody{
			{Subtitle: "first title", Contents: "contents"},
			{Contents: "contents"},
			{Contents: "contents"},
		},
	}
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		panic(errors.Wrap(err, "template.Parse template"))
	}

	if err := t.Execute(os.Stdout, htmlDoc); err != nil {
		panic(errors.Wrap(err, "template.Execture html"))
	}
}
