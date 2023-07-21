package gqlgo

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/graphql-go/graphql"
)

func RelayIDFromObject(object any, info graphql.ResolveInfo, ctx context.Context) (string, error) {
	id, err := defaultFieldResolver(graphql.ResolveParams{
		Source:  object,
		Info:    info,
		Context: ctx,
	})
	idData := fmt.Sprintf("%v:%v", info.ParentType, id)
	return base64.StdEncoding.EncodeToString([]byte(idData)), err
}
