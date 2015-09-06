package api

import (
	"net/http"
	_ "net/http/pprof" // Expose runtime profiling data.

	"github.com/gorilla/mux"
)

// TODO: I don't like Context.
type Server struct {
	opts *Options
	Ctxt *Context
}

func MakeServer(opts *Options, ctxt *Context) *Server {
	return &Server{opts, ctxt}
}

func (s *Server) ListenAndServe() error {
	router := s.getRouter()
	http.Handle("/", router)

	addr := s.opts.Address()
	s.Ctxt.Logger.Printf("Listening on %s ...", addr)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) getRouter() http.Handler {
	muxRouter := mux.NewRouter()
	// FIXME: Interface segregation principle
	muxRouter.Handle("/", Handler{s.Ctxt, StatusHandler})
	muxRouter.Handle("/store/{key}/", Handler{s.Ctxt, GetKeyHandler}).
		Methods("GET")
	muxRouter.Handle("/store/", Handler{s.Ctxt, SetKeyHandler}).
		Methods("POST").
		Headers("content-type", "application/json")
	return muxRouter
}
