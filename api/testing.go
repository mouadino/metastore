package api

import "github.com/mouadino/metastore/testing"

var testServer = MakeServer(
	&Options{},
	&testing.InMemoryStore{
		"abc": "foobar",
	},
)
