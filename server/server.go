package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	_ "net/http/pprof"
)

type Server struct {
	host string
	port int
}

type Status struct {
	Address string
}

func MakeServer(host string, port int) *Server {
	return &Server{
		host: host,
		port: port,
	}
}

func (s *Server) ListenAndServe() error {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", StatusHandler)
	http.Handle("/", muxRouter)

	addr := s.getAddress()
	return http.ListenAndServe(addr, nil)
}

func (s *Server) getAddress() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

func (s *Server) Status() Status {
	return Status{
		Address: s.getAddress(),
	}
}
