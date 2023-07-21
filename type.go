package gqlgo

import (
	"github.com/graphql-go/graphql"
)

type Type struct {
	Name        string
	Description string
	Fields      Fields
}

func (t *Type) graphqlObject(config Config) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        t.Name,
		Description: t.Description,
		Fields:      t.Fields.graphqlFields(config),
	})
}

func (t *Type) Output(config Config) graphql.Output {
	return t.graphqlObject(config)
}
