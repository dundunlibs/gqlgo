package gqlgo

import "github.com/graphql-go/graphql"

type Input struct {
	Name        string
	Description string
	Args        Args
}

func (i *Input) graphqlInput(s Schema) *graphql.InputObject {
	if s.graphqlInputs[i] == nil {
		s.graphqlInputs[i] = graphql.NewInputObject(graphql.InputObjectConfig{
			Name:        i.Name,
			Description: i.Description,
			Fields:      make(graphql.InputObjectConfigFieldMap),
		})

		// append fields later to avoid cycle initialization
		for k, f := range i.Args.graphqlInputFields(s) {
			s.graphqlInputs[i].AddFieldConfig(k, f)
		}
	}

	return s.graphqlInputs[i]
}

func (i *Input) Output(s Schema) graphql.Output {
	return i.graphqlInput(s)
}
