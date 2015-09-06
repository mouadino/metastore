package api

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mouadino/metastore/storage"
)

func Init(opts *Options, db *storage.DB) {
	ctxt := CreateContext(*db)
	server := MakeServer(opts, ctxt)
	go server.ListenAndServe()

	waitForTermination()
}

func waitForTermination() {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Print("Received SIGTERM, exiting ...")
	}
}
