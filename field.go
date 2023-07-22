package gqlgo

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
)

type ResolveParams graphql.ResolveParams
type FieldResolver func(p ResolveParams) (any, error)
type FieldMiddlewareFn func(next FieldResolver) FieldResolver

type Field struct {
	Type        Output
	Description string
	Args        Args
	Resolve     FieldResolver
	Middlewares []FieldMiddlewareFn
}

func (f *Field) graphqlField(s Schema) *graphql.Field {
	graphqlType, resolve, ml := f.Type.Output(s), f.Resolve, len(f.Middlewares)

	// apply global ID if needed
	if isID(graphqlType) && s.Config.IDFromObject != nil {
		resolve = func(p ResolveParams) (interface{}, error) {
			return s.Config.IDFromObject(p.Source, p.Info, p.Context)
		}
	}

	// apply middlewares
	for i := 1; i <= ml; i++ {
		resolve = f.Middlewares[ml-i](resolve)
	}

	field := &graphql.Field{
		Type:        graphqlType,
		Description: f.Description,
	}

	if f.Args != nil {
		field.Args = f.Args.graphqlArgs(s)
	}

	if resolve != nil {
		field.Resolve = func(p graphql.ResolveParams) (interface{}, error) {
			return resolve(ResolveParams(p))
		}
	}

	return field
}

func defaultFieldResolver(p graphql.ResolveParams) (interface{}, error) {
	var src map[string]any
	if s, ok := p.Source.(map[string]any); ok {
		src = s
	} else {
		data, err := json.Marshal(p.Source)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, &src); err != nil {
			return nil, err
		}
	}
	return src[p.Info.FieldName], nil
}
