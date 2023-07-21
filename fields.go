package gqlgo

import (
	"github.com/graphql-go/graphql"
)

type Fields map[string]*Field

func (f Fields) graphqlFields(s Schema) graphql.Fields {
	fields := make(graphql.Fields)
	for k, v := range f {
		fields[k] = v.graphqlField(s)
	}
	return fields
}
