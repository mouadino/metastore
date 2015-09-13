package api

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/context"
	"github.com/mouadino/metastore/storage"
)

type Middleware func(http.Handler) http.Handler

func WithContext(store storage.DB) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, storeKey, store)
			defer context.Clear(r)
			handler.ServeHTTP(w, r)
		})
	}
}

func Accept(contentTypes ...string) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !in(r.Header.Get("content-type"), contentTypes) {
				http.Error(w, "", http.StatusUnsupportedMediaType)
			}
			handler.ServeHTTP(w, r)
		})
	}
}

func Method(methods ...string) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sort.Strings(methods)
			if r.Method == "OPTIONS" {
				w.Header().Set("Allow", strings.Join(methods, ", "))
				w.WriteHeader(http.StatusOK)
			} else if !in(r.Method, methods) {
				message := fmt.Sprintf("unknown method '%s' support %s", r.Method, methods)
				http.Error(w, message, http.StatusMethodNotAllowed)
			} else {
				handler.ServeHTTP(w, r)
			}
		})
	}
}
