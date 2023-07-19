package gqlgo

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/dundunlabs/gqlgo/playground"
	"github.com/graphql-go/graphql"
)

type HandlerOption func(h *Handler)

func WithPlayground(v bool) HandlerOption {
	return func(h *Handler) {
		if v {
			h.playgroundTmpl = playground.TemplateHTML
		}
	}
}

type HandlerBody struct {
	OperationName string         `json:"operationName"`
	Query         string         `json:"query"`
	Variables     map[string]any `json:"variables"`
}

func NewHandler(schema graphql.Schema, opts ...HandlerOption) *Handler {
	h := new(Handler)
	h.schema = schema
	for _, opt := range opts {
		opt(h)
	}
	return h
}

type Handler struct {
	schema         graphql.Schema
	playgroundTmpl *template.Template
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if h.playgroundTmpl != nil {
			h.playgroundTmpl.Execute(w, playground.Data{
				Endpoint: r.URL.Path,
			})
			return
		}
	case http.MethodPost:
		var body HandlerBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := graphql.Do(graphql.Params{
			Context:        r.Context(),
			Schema:         h.schema,
			OperationName:  body.OperationName,
			RequestString:  body.Query,
			VariableValues: body.Variables,
		})

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}
