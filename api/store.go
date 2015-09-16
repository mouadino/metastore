package api

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mouadino/metastore/storage"
)

type NewEntry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func GetKeyHandler(req *http.Request) (int, interface{}) {
	store := context.Get(req, storeKey).(storage.DB)
	key := []byte(mux.Vars(req)["key"])
	data, err := store.Get(key)
	if err != nil {
		return http.StatusNotFound, err
	}
	return http.StatusOK, string(data)
}

func SetKeyHandler(req *http.Request) (int, interface{}) {
	var entry NewEntry
	err := getJSON(req, &entry)
	if err != nil {
		return http.StatusBadRequest, err
	}
	store := context.Get(req, storeKey).(storage.DB)
	err = store.Put([]byte(entry.Key), []byte(entry.Value))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, true
}
