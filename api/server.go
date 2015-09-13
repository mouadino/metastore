package api

import (
	"net/http"
	_ "net/http/pprof" // Expose runtime profiling data.

	"github.com/gorilla/mux"
)

type Server struct {
	Opts        *Options
	Mux         *mux.Router
	Middlewares []Middleware
}

func NewServer(opts *Options, middlewares ...Middleware) *Server {
	return &Server{
		Opts:        opts,
		Mux:         mux.NewRouter(),
		Middlewares: middlewares,
	}
}

func (s *Server) Handle(route string, handler http.Handler, middlewares ...Middleware) {
	middlewares = append(s.Middlewares, middlewares...)
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	s.Mux.Handle(route, handler)
}

func (s *Server) ListenAndServe() error {
	http.Handle("/", s.Mux)

	addr := s.Opts.Address()
	//FIXME: s.Ctxt.Logger.Printf("Listening on %s ...", addr)
	return http.ListenAndServe(addr, nil)
}
