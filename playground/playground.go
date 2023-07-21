package playground

import (
	_ "embed"
	"html/template"
)

//go:embed playground.html.tmpl
var tmplHTML string

var TemplateHTML, _ = template.New("Playground").Parse(tmplHTML)

type Data struct {
	Endpoint string
}
