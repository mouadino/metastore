package api

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mouadino/metastore/storage"
)

func Init(opts *Options, db storage.DB) {
	server := MakeServer(opts, db)
	go server.ListenAndServe()

	waitForTermination()
}

func MakeServer(opts *Options, db storage.DB) *Server {
	server := NewServer(opts, WithContext(db))

	server.Handle(
		"/",
		SimpleHandler(StatusHandler),
		Method("GET"),
	)
	server.Handle(
		"/store/{key}/",
		SimpleHandler(GetKeyHandler),
		Method("GET"),
	)
	server.Handle(
		"/store/",
		SimpleHandler(SetKeyHandler),
		Method("POST"),
		Accept(JSONContentType),
	)

	return server
}

func waitForTermination() {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Print("Received SIGTERM, exiting ...")
	}
}
