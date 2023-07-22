package gqlgo

import "github.com/graphql-go/graphql"

type Args map[string]*Arg

func (args Args) graphqlArgs(s Schema) graphql.FieldConfigArgument {
	config := make(graphql.FieldConfigArgument)
	for k, v := range args {
		config[k] = v.graphqlArg(s)
	}
	return config
}

func (args Args) graphqlInputFields(s Schema) graphql.InputObjectConfigFieldMap {
	config := make(graphql.InputObjectConfigFieldMap)
	for k, v := range args {
		config[k] = v.graphqlInputField(s)
	}
	return config
}
