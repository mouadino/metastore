package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/context"
)

type Middleware func(http.Handler) http.Handler

func WithContext(values map[int]interface{}) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for k, v := range values {
				context.Set(r, k, v)
			}
			defer context.Clear(r)
			handler.ServeHTTP(w, r)
		})
	}
}

func Accept(contentTypes ...string) Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ct := normalizeContentType(r.Header)
			if !contains(ct, contentTypes) {
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
			} else if !contains(r.Method, methods) {
				message := fmt.Sprintf("unknown method '%s' support %s", r.Method, methods)
				http.Error(w, message, http.StatusMethodNotAllowed)
			} else {
				handler.ServeHTTP(w, r)
			}
		})
	}
}

func Trace() Middleware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceId := strconv.FormatInt(rand.Int63()&0x001fffffffffffff, 10)
			context.Set(r, traceIdKey, traceId)
			w.Header().Set("X-Trace-Id", traceId)
			handler.ServeHTTP(w, r)
		})
	}
}
