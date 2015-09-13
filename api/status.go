package api

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/mouadino/metastore/storage"
)

func StatusHandler(req *http.Request) (int, interface{}) {
	store := context.Get(req, storeKey).(storage.DB)
	return http.StatusOK, map[string]interface{}{
		"store": store.Status(),
	}
}
