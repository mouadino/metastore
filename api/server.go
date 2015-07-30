package api

import (
	"net/http"
	_ "net/http/pprof" // Expose runtime profiling data.

	"github.com/gorilla/mux"

	"github.com/mouadino/metastore/context"
)

type Server struct {
	Ctxt *context.Context
}

func MakeServer(ctxt *context.Context) *Server {
	return &Server{ctxt}
}

func (s *Server) ListenAndServe() error {
	muxRouter := mux.NewRouter()
	// FIXME: Interface segregation principle
	muxRouter.Handle("/", Handler{s.Ctxt, StatusHandler})
	muxRouter.Handle("/store/{key}", Handler{s.Ctxt, GetKeyHandler}).
		Methods("GET")
	muxRouter.Handle("/store/", Handler{s.Ctxt, SetKeyHandler}).
		Methods("POST").
		Headers("content-type", "application/json")

	http.Handle("/", muxRouter)

	addr := s.Ctxt.APIAddress()
	s.Ctxt.Logger.Printf("Listening on %s ...", addr)
	return http.ListenAndServe(addr, nil)
}
