package types

import "github.com/dundunlabs/gqlgo"

var User = &gqlgo.Type{
	Name:        "User",
	Description: "User type",
	Fields: gqlgo.Fields{
		"id": &gqlgo.Field{
			Type:        gqlgo.ID,
			Description: "The ID of User type",
		},
	},
}
