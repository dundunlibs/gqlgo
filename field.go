package gqlgo

import (
	"github.com/graphql-go/graphql"
)

type Field struct {
	Type        Output
	Description string
	Middlewares []FieldMiddlewareFn
	Resolve     graphql.FieldResolveFn
}

func (f *Field) graphqlField(s Schema) *graphql.Field {
	graphqlType, resolve, ml := f.Type.Output(s), f.Resolve, len(f.Middlewares)

	// apply global ID if needed
	if isID(graphqlType) && s.Config.IDFromObject != nil {
		resolve = func(p graphql.ResolveParams) (interface{}, error) {
			return s.Config.IDFromObject(p.Source, p.Info, p.Context)
		}
	}

	// apply middlewares
	for i := 1; i <= ml; i++ {
		resolve = f.Middlewares[ml-i](resolve)
	}

	return &graphql.Field{
		Type:        graphqlType,
		Description: f.Description,
		Resolve:     resolve,
	}
}

type FieldMiddlewareFn func(next graphql.FieldResolveFn) graphql.FieldResolveFn
