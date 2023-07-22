package gqlgo

import "github.com/graphql-go/graphql"

type Input struct {
	Name        string
	Description string
	Args        Args
}

func (i *Input) graphqlInput(s Schema) *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name:        i.Name,
		Description: i.Description,
		Fields:      i.Args.graphqlInputFields(s),
	})
}

func (i *Input) Output(s Schema) graphql.Output {
	return i.graphqlInput(s)
}
