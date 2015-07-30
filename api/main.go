package api

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mouadino/metastore/context"
)

func Main(ctxt *context.Context) {
	server := MakeServer(ctxt)
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
