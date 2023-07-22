package gqlgo

import (
	"github.com/graphql-go/graphql"
)

type Schema struct {
	// used to cache all posible types
	graphqlObjects map[*Type]*graphql.Object
	graphqlInputs  map[*Input]*graphql.InputObject

	Query    *Type
	Mutation *Type
	Config   Config
}

func (s Schema) graphqlSchema() (graphql.Schema, error) {
	// initialize cache data
	s.graphqlObjects = make(map[*Type]*graphql.Object)
	s.graphqlInputs = make(map[*Input]*graphql.InputObject)

	// initialize graphql's schema config
	schemaConfig := graphql.SchemaConfig{}
	if s.Query != nil {
		schemaConfig.Query = s.Query.graphqlObject(s)
	}
	if s.Mutation != nil {
		schemaConfig.Mutation = s.Mutation.graphqlObject(s)
	}

	return graphql.NewSchema(schemaConfig)
}
