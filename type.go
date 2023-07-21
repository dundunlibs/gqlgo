package gqlgo

import (
	"github.com/graphql-go/graphql"
)

type Type struct {
	Name        string
	Description string
	Fields      Fields
}

func (t *Type) graphqlObject(s Schema) *graphql.Object {
	if s.graphqlObjects[t] == nil {
		s.graphqlObjects[t] = graphql.NewObject(graphql.ObjectConfig{
			Name:        t.Name,
			Description: t.Description,
			Fields:      graphql.Fields{},
		})
		// append fields later to avoid cycle initialization
		for k, f := range t.Fields.graphqlFields(s) {
			s.graphqlObjects[t].AddFieldConfig(k, f)
		}
	}

	return s.graphqlObjects[t]
}

func (t *Type) Output(s Schema) graphql.Output {
	return t.graphqlObject(s)
}
