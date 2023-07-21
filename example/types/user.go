package types

import (
	"database/sql"

	"github.com/dundunlabs/gqlgo"
	"github.com/dundunlabs/gqlgo/example/models"
	"github.com/graphql-go/graphql"
)

var User = &gqlgo.Type{
	Name:        "User",
	Description: "User type",
}

func NewUserFields(db *sql.DB) gqlgo.Fields {
	return gqlgo.Fields{
		"id": &gqlgo.Field{
			Type:        gqlgo.NotNull(gqlgo.ID),
			Description: "ID of user",
		},
		"name": &gqlgo.Field{
			Type:        gqlgo.String,
			Description: "Name of user",
		},
		"posts": &gqlgo.Field{
			Type: gqlgo.NotNull(gqlgo.List(Post)),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				rows, err := db.QueryContext(p.Context, "SELECT * FROM posts WHERE author_id = ?", p.Source.(models.User).ID)
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
	}
}
