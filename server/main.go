package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var currentServer *Server

func Main(host string, port int) {
	currentServer = MakeServer(host, port)
	go currentServer.ListenAndServe()

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
