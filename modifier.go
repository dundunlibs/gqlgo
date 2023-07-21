package gqlgo

import (
	"github.com/graphql-go/graphql"
)

type Modifier func(s Schema) graphql.Output

func (m Modifier) Output(s Schema) graphql.Output {
	return m(s)
}

func NotNull(output Output) Modifier {
	return func(s Schema) graphql.Output {
		return graphql.NewNonNull(output.Output(s))
	}
}

func List(output Output) Modifier {
	return func(s Schema) graphql.Output {
		return graphql.NewList(output.Output(s))
	}
}
