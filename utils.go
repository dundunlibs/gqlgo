package gqlgo

import (
	"github.com/graphql-go/graphql"
)

func isID(op graphql.Type) bool {
	switch op := op.(type) {
	case *graphql.Scalar:
		return op == graphql.ID
	case *graphql.NonNull:
		return isID(op.OfType)
	case *graphql.List:
		return isID(op.OfType)
	default:
		return false
	}
}
