package main

import "github.com/graphql-go/graphql"

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Description: "Root Query",
		Fields: graphql.Fields{
			"test": &graphql.Field{
				Type:        graphql.String,
				Description: "Test field",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "test", nil
				},
			},
		},
	}),
})
