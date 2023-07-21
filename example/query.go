package main

import (
	"database/sql"

	"github.com/dundunlabs/gqlgo"
	"github.com/dundunlabs/gqlgo/example/models"
	"github.com/dundunlabs/gqlgo/example/types"
	"github.com/graphql-go/graphql"
)

func newQuery(db *sql.DB) *gqlgo.Type {
	return &gqlgo.Type{
		Name:        "Query",
		Description: "Root query",
		Fields: gqlgo.Fields{
			"users": &gqlgo.Field{
				Type:        gqlgo.List(gqlgo.NotNull(types.User)),
				Description: "All users",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.QueryContext(p.Context, "SELECT * FROM users")
					users := []models.User{}
					if err != nil {
						return users, err
					}

					for rows.Next() {
						var user models.User
						if err := rows.Scan(&user.ID, &user.Name); err != nil {
							return users, err
						}
						users = append(users, user)
					}

					return users, rows.Err()
				},
			},
			"posts": &gqlgo.Field{
				Type:        gqlgo.List(gqlgo.NotNull(types.Post)),
				Description: "All posts",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.QueryContext(p.Context, "SELECT * FROM posts")
					posts := []models.Post{}
					if err != nil {
						return posts, err
					}

					for rows.Next() {
						var post models.Post
						if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.AuthorID); err != nil {
							return posts, err
						}
						posts = append(posts, post)
					}

					return posts, rows.Err()
				},
			},
		},
	}
}
