package gqlgo

import "github.com/graphql-go/graphql"

func FromGraphqlScalar(scalar *graphql.Scalar) *Scalar {
	return &Scalar{
		Scalar: scalar,
	}
}

type Scalar struct {
	*graphql.Scalar
}

func (s *Scalar) Output(Config) graphql.Output {
	return s.Scalar
}

var ID = FromGraphqlScalar(graphql.ID)
var String = FromGraphqlScalar(graphql.String)
var Int = FromGraphqlScalar(graphql.Int)
var Float = FromGraphqlScalar(graphql.Float)
var Boolean = FromGraphqlScalar(graphql.Boolean)
var DateTime = FromGraphqlScalar(graphql.DateTime)
