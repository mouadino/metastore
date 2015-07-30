package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mouadino/metastore/context"
)

func GetKeyHandler(ctxt *context.Context, req *http.Request) Response {
	key := []byte(mux.Vars(req)["key"])
	data, err := (*ctxt.Store).Get(key)
	if err != nil {
		return ResponseFromError(err)
	}
	// TODO: return NotFound if data is empty.
	return Response{http.StatusOK, data}
}

func SetKeyHandler(ctxt *context.Context, req *http.Request) Response {
	var newEntry map[string]string
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newEntry)
	if err != nil {
		return ResponseFromError(err)
	}
	key, ok := newEntry["key"]
	if !ok {
		return ResponseFromError(fmt.Errorf("malformed body"))
	}
	value, ok := newEntry["value"]
	if !ok {
		return ResponseFromError(fmt.Errorf("malformed body"))
	}
	err = (*ctxt.Store).Put([]byte(key), []byte(value))
	if err != nil {
		return ResponseFromError(err)
	}
	return Response{http.StatusCreated, true}
}
