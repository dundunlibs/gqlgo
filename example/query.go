package main

import (
	"github.com/dundunlabs/gqlgo"
	"github.com/dundunlabs/gqlgo/example/types"
	"github.com/graphql-go/graphql"
)

var Query = &gqlgo.Type{
	Name:        "Query",
	Description: "Root query",
	Fields: gqlgo.Fields{
		"test": &gqlgo.Field{
			Type:        gqlgo.String,
			Description: "test field",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "test", nil
			},
		},
		"user": &gqlgo.Field{
			Type:        types.User,
			Description: "Current user",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return map[string]any{
					"id": 1,
				}, nil
			},
		},
	},
}
