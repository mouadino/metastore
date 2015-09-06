package api

import (
	"log"
	"os"

	"github.com/mouadino/metastore/storage"
)

type Context struct {
	Store  *storage.DB
	Logger *log.Logger
}

func CreateContext(store storage.DB) *Context {
	// TODO: Configure logger.
	return &Context{
		Store:  &store,
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (ctxt *Context) Status() map[string]interface{} {
	return map[string]interface{}{
		"store": (*ctxt.Store).Status(),
	}
}
