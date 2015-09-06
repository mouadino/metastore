package api

import (
	"net/http"
)

type Handler struct {
	*Context
	handler func(*Context, *http.Request) Response
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := h.handler(h.Context, r)
	response.WriteTo(w)
}
