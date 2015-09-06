package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetKeyHandler(ctxt *Context, req *http.Request) Response {
	key := []byte(mux.Vars(req)["key"])
	data, err := (*ctxt.Store).Get(key)
	if err != nil {
		return ResponseFromError(err)
	}
	// TODO: return NotFound if data is nil.
	// TODO: Encoding/Decoding stored data, for now we only
	// do strings.
	return Response{http.StatusOK, string(data)}
}

func SetKeyHandler(ctxt *Context, req *http.Request) Response {
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
	// TODO: Return id !?
	return Response{http.StatusCreated, true}
}
