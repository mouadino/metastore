package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	_ "net/http/pprof"
)

func ListenAndServe(host string, port int) error {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", HomeHandler)
	http.Handle("/", muxRouter)
	addr := fmt.Sprintf("%s:%d", host, port)
	return http.ListenAndServe(addr, nil)
}
