package api

import (
	"github.com/mouadino/metastore/testhelpers"
)

// TODO: Reset state/tearDown !?
var store = testhelpers.InMemoryStore{
	"abc": "foobar",
}
var ctxt = CreateContext(&store)

var testServer = MakeServer(
	&Options{},
	ctxt,
)
