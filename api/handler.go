package api

import (
	"encoding/json"
	"net/http"
)

type SimpleHandler func(*http.Request) (status int, body interface{})

func (h SimpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, body := h(r)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
	w.Header().Set("Content-Type", JSONContentType)
}
