package explorer

import (
	_ "embed"
	"html/template"
)

//go:embed playground.html.tmpl
var htmlPlayground string

var ExplorerPlayground, _ = template.New("Playground").Parse(htmlPlayground)
