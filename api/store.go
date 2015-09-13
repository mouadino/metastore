package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mouadino/metastore/storage"
)

func GetKeyHandler(req *http.Request) (int, interface{}) {
	key := []byte(mux.Vars(req)["key"])
	store := context.Get(req, storeKey).(storage.DB)
	data, err := store.Get(key)
	if err != nil {
		return http.StatusNotFound, err
	}
	// TODO: Encoding/Decoding stored data, for now we only
	// do strings.
	return http.StatusOK, string(data)
}

func SetKeyHandler(req *http.Request) (int, interface{}) {
	// TODO: All this doesn't belong here.
	var newEntry map[string]string
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newEntry)
	if err != nil {
		return http.StatusBadRequest, err
	}
	key, ok := newEntry["key"]
	if !ok {
		return http.StatusBadRequest, errors.New("malformed body")
	}
	value, ok := newEntry["value"]
	if !ok {
		return http.StatusBadRequest, errors.New("malformed body")
	}
	//
	store := context.Get(req, storeKey).(storage.DB)
	err = store.Put([]byte(key), []byte(value))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, true
}
