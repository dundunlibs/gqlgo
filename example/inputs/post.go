package inputs

import (
	"github.com/dundunlabs/gqlgo"
)

var PostAttributes = &gqlgo.Input{
	Name:        "PostAttributes",
	Description: "Post's attributes",
	Args: gqlgo.Args{
		"title": &gqlgo.Arg{
			Type:    gqlgo.String,
			Default: "",
		},
		"body": &gqlgo.Arg{
			Type:    gqlgo.String,
			Default: "",
		},
		"authorId": &gqlgo.Arg{
			Type:        gqlgo.NotNull(gqlgo.ID),
			Description: "Author's ID",
		},
	},
}
