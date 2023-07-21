package gqlgo

import (
	"encoding/json"

	"github.com/graphql-go/graphql"
)

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
