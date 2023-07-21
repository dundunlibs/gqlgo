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

func (f *Field) graphqlField(config Config) *graphql.Field {
	resolve, ml := f.Resolve, len(f.Middlewares)

	// apply global ID if needed
	if f.Type == ID && config.IDFromObject != nil {
		resolve = func(p graphql.ResolveParams) (interface{}, error) {
			return config.IDFromObject(p.Source, p.Info, p.Context)
		}
	}

	// apply middlewares
	for i := 1; i <= ml; i++ {
		resolve = f.Middlewares[ml-i](resolve)
	}

	return &graphql.Field{
		Type:        f.Type.Output(config),
		Description: f.Description,
		Resolve:     resolve,
	}
}

type FieldMiddlewareFn func(next graphql.FieldResolveFn) graphql.FieldResolveFn
