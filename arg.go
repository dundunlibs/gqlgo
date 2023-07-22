package gqlgo

import "github.com/graphql-go/graphql"

type Arg struct {
	Type        Output
	Description string
	Default     any
}

func (arg *Arg) graphqlArg(s Schema) *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{
		Type:         arg.Type.Output(s),
		Description:  arg.Description,
		DefaultValue: arg.Default,
	}
}

func (arg *Arg) graphqlInputField(s Schema) *graphql.InputObjectFieldConfig {
	return &graphql.InputObjectFieldConfig{
		Type:         arg.Type.Output(s),
		Description:  arg.Description,
		DefaultValue: arg.Default,
	}
}
