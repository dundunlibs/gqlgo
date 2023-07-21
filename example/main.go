package main

import (
	"database/sql"
	_ "embed"
	"log"
	"net/http"

	"github.com/dundunlabs/gqlgo"
	"github.com/dundunlabs/gqlgo/explorer"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed scripts.sql
var scripts string

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	checkErr(err)

	defer db.Close()

	_, err = db.Exec(scripts)
	checkErr(err)

	handler := gqlgo.NewHandler(newSchema(db), gqlgo.WithExplorer(explorer.ExplorerGraphiQL))
	http.Handle("/graphql", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
