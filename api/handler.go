package api

import (
	"net/http"

	"github.com/mouadino/metastore/context"
)

type Handler struct {
	*context.Context
	handler func(*context.Context, *http.Request) Response
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := h.handler(h.Context, r)
	response.WriteTo(w)
}
