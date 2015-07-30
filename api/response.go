package api

import (
	"encoding/json"
	"net/http"
)

const (
	JSONContentType = "application/json; charset=UTF-8"
)

type Response struct {
	Status int
	Body   interface{}
}

func ResponseFromError(err error) Response {
	return Response{
		Status: http.StatusInternalServerError,
		Body: map[string]string{
			"message": err.Error(),
		},
	}
}

func (r *Response) WriteTo(w http.ResponseWriter) {
	// TODO: Add X-Request-ID.
	w.Header().Set("Content-Type", JSONContentType)
	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(r.Body)
}
