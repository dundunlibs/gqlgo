package gqlgo

import "github.com/graphql-go/graphql"

type Output interface {
	Output(Config) graphql.Output
}
