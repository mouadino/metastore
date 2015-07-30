package context

import (
	"github.com/mouadino/metastore/storage"
)

import (
	"fmt"
	"log"
	"os"
)

type Status map[string]interface{}

type Context struct {
	Host   string
	Port   int
	Logger *log.Logger
	Store  *storage.DB
}

func Create(host string, port int, store storage.DB) *Context {
	return &Context{
		Host:   host,
		Port:   port,
		Logger: log.New(os.Stderr, "", log.LstdFlags),
		Store:  &store,
	}
}

func (ctxt *Context) Status() Status {
	return Status{
		"api":   ctxt.APIAddress(),
		"store": (*ctxt.Store).Status(),
	}
}

func (ctxt *Context) APIAddress() string {
	return fmt.Sprintf("%s:%d", ctxt.Host, ctxt.Port)
}
