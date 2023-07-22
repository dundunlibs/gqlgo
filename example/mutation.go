package main

import (
	"database/sql"
	"strconv"

	"github.com/dundunlabs/gqlgo"
	"github.com/dundunlabs/gqlgo/example/inputs"
	"github.com/dundunlabs/gqlgo/example/models"
	"github.com/dundunlabs/gqlgo/example/types"
)

func newMutation(db *sql.DB) *gqlgo.Type {
	return &gqlgo.Type{
		Name: "Mutation",
		Fields: gqlgo.Fields{
			"createPost": &gqlgo.Field{
				Type: types.Post,
				Args: gqlgo.Args{
					"authorId": &gqlgo.Arg{
						Type:        gqlgo.NotNull(gqlgo.ID),
						Description: "Author's ID",
					},
					"attributes": &gqlgo.Arg{
						Type:        gqlgo.NotNull(inputs.PostAttributes),
						Description: "Post's attributes",
					},
				},
				Resolve: func(p gqlgo.ResolveParams) (interface{}, error) {
					attrs := p.Args["attributes"].(map[string]any)
					aid, err := strconv.Atoi(p.Args["authorId"].(string))
					if err != nil {
						return nil, err
					}

					post := models.Post{
						AuthorID: aid,
						Title:    attrs["title"].(string),
						Body:     attrs["body"].(string),
					}

					result, err := db.ExecContext(p.Context, "INSERT INTO posts (author_id, title, body) VALUES (?, ?, ?);", post.AuthorID, post.Title, post.Body)
					if err != nil {
						return nil, err
					}

					if id, err := result.LastInsertId(); err != nil {
						return nil, err
					} else {
						post.ID = int(id)
					}

					return post, nil
				},
			},
		},
	}
}
