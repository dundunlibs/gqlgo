package explorer

import "html/template"

type Explorer *template.Template

type Data struct {
	Endpoint string
}
