package gqlgo

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/dundunlabs/gqlgo/explorer"
	"github.com/graphql-go/graphql"
)

type HandlerOption func(h *Handler)

func WithExplorer(explorer explorer.Explorer) HandlerOption {
	return func(h *Handler) {
		h.explorerTmpl = explorer
	}
}

type HandlerBody struct {
	OperationName string         `json:"operationName"`
	Query         string         `json:"query"`
	Variables     map[string]any `json:"variables"`
}

func NewHandler(schema Schema, opts ...HandlerOption) *Handler {
	h := new(Handler)
	s, err := schema.graphqlSchema()
	if err != nil {
		panic(err)
	}
	h.schema = s
	for _, opt := range opts {
		opt(h)
	}
	return h
}

type Handler struct {
	schema       graphql.Schema
	explorerTmpl *template.Template
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if h.explorerTmpl != nil {
			h.explorerTmpl.Execute(w, explorer.Data{
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
