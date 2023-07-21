package gqlgo

import (
	"github.com/graphql-go/graphql"
)

type Schema struct {
	Query  *Type
	Config Config
}

func (s Schema) graphqlSchema() (graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{}
	if s.Query != nil {
		schemaConfig.Query = s.Query.graphqlObject(s.Config)
	}

	return graphql.NewSchema(schemaConfig)
}
