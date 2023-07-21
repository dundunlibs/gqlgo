package main

import (
	"github.com/dundunlabs/gqlgo"
)

var schema = gqlgo.Schema{
	Query: Query,
	Config: gqlgo.Config{
		IDFromObject: gqlgo.RelayIDFromObject,
	},
}
