package main

import (
	"database/sql"

	"github.com/dundunlabs/gqlgo"
	"github.com/dundunlabs/gqlgo/example/types"
)

func newSchema(db *sql.DB) gqlgo.Schema {
	userFields := types.NewUserFields(db)
	postFields := types.NewPostFields(db)

	types.User.Fields = userFields
	types.Post.Fields = postFields

	return gqlgo.Schema{
		Query:    newQuery(db),
		Mutation: newMutation(db),
		Config: gqlgo.Config{
			IDFromObject: gqlgo.RelayIDFromObject,
		},
	}
}
