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
					"attributes": &gqlgo.Arg{
						Type:        gqlgo.NotNull(inputs.PostAttributes),
						Description: "Post's attributes",
					},
				},
				Resolve: func(p gqlgo.ResolveParams) (interface{}, error) {
					attrs := p.Args["attributes"].(map[string]any)
					aid, err := strconv.Atoi(attrs["authorId"].(string))
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
			"updatePost": &gqlgo.Field{
				Type: types.Post,
				Args: gqlgo.Args{
					"id": &gqlgo.Arg{
						Type:        gqlgo.NotNull(gqlgo.ID),
						Description: "Post's ID",
					},
					"attributes": &gqlgo.Arg{
						Type:        gqlgo.NotNull(inputs.PostAttributes),
						Description: "Post's attributes",
					},
				},
				Resolve: func(p gqlgo.ResolveParams) (interface{}, error) {
					attrs := p.Args["attributes"].(map[string]any)
					aid, err := strconv.Atoi(attrs["authorId"].(string))
					if err != nil {
						return nil, err
					}

					id, err := strconv.Atoi(p.Args["id"].(string))
					if err != nil {
						return nil, err
					}

					post := models.Post{
						ID:       id,
						Title:    attrs["title"].(string),
						Body:     attrs["body"].(string),
						AuthorID: aid,
					}

					if _, err := db.ExecContext(
						p.Context,
						"UPDATE posts SET author_id = ?, title = ?, body = ? WHERE id = ?;",
						post.AuthorID, post.Title, post.Body, post.ID,
					); err != nil {
						return nil, err
					}

					return post, nil
				},
			},
		},
	}
}
