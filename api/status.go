package api

import (
	"net/http"

	"github.com/mouadino/metastore/context"
)

func StatusHandler(ctxt *context.Context, r *http.Request) Response {
	return Response{http.StatusOK, ctxt.Status()}
}
