package main

import (
	"log"
	"net/http"

	"github.com/dundunlabs/gqlgo"
	"github.com/dundunlabs/gqlgo/explorer"
)

func main() {
	handler := gqlgo.NewHandler(schema, gqlgo.WithExplorer(explorer.ExplorerGraphiQL))
	http.Handle("/graphql", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
