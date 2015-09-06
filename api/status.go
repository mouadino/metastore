package api

import (
	"net/http"
)

func StatusHandler(ctxt *Context, r *http.Request) Response {
	return Response{http.StatusOK, ctxt.Status()}
}
