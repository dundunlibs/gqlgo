package gqlgo

import (
	"github.com/graphql-go/graphql"
)

type Schema struct {
	// used to cache all posible types
	graphqlObjects map[*Type]*graphql.Object

	Query  *Type
	Config Config
}

func (s Schema) graphqlSchema() (graphql.Schema, error) {
	// initialize graphqlObjects
	s.graphqlObjects = make(map[*Type]*graphql.Object)

	// initialize graphql's schema config
	schemaConfig := graphql.SchemaConfig{}
	if s.Query != nil {
		schemaConfig.Query = s.Query.graphqlObject(s)
	}

	return graphql.NewSchema(schemaConfig)
}
