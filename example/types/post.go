package types

import (
	"database/sql"

	"github.com/dundunlabs/gqlgo"
	"github.com/dundunlabs/gqlgo/example/models"
)

var Post = &gqlgo.Type{
	Name:        "Post",
	Description: "Post type",
}

func NewPostFields(db *sql.DB) gqlgo.Fields {
	return gqlgo.Fields{
		"id": &gqlgo.Field{
			Type: gqlgo.NotNull(gqlgo.ID),
		},
		"title": &gqlgo.Field{
			Type: gqlgo.String,
		},
		"body": &gqlgo.Field{
			Type: gqlgo.String,
		},
		"authorId": &gqlgo.Field{
			Type: gqlgo.Int,
		},
		"author": &gqlgo.Field{
			Type: gqlgo.NotNull(User),
			Resolve: func(p gqlgo.ResolveParams) (interface{}, error) {
				row := db.QueryRowContext(p.Context, "SELECT * FROM users WHERE id = ? LIMIT 1", p.Source.(models.Post).AuthorID)
				var user models.User
				err := row.Scan(&user.ID, &user.Name)
				return user, err
			},
		},
	}
}
