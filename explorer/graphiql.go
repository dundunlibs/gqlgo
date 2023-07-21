package explorer

import (
	_ "embed"
	"html/template"
)

//go:embed graphiql.html.tmpl
var htmlGraphiQL string

var ExplorerGraphiQL, _ = template.New("GraphiQL").Parse(htmlGraphiQL)
