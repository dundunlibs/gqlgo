package main

import (
	"log"
	"net/http"

	"github.com/dundunlabs/gqlgo"
)

func main() {
	handler := gqlgo.NewHandler(schema, gqlgo.WithPlayground(true))
	http.Handle("/graphql", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
