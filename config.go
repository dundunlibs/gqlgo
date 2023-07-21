package gqlgo

import (
	"context"

	"github.com/graphql-go/graphql"
)

type Config struct {
	IDFromObject IDFromObjectFn
}

type IDFromObjectFn func(object any, info graphql.ResolveInfo, ctx context.Context) (string, error)
