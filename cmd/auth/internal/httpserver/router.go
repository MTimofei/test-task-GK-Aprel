package httpserver

import (
	"net/http"

	"github.com/auth-api/cmd/auth/internal/httpserver/middleware"
	"github.com/auth-api/cmd/auth/internal/httpserver/public"
	"github.com/auth-api/cmd/auth/internal/httpserver/secret"
)

const (
	endpointClearAudit = "/audit/clear"
	endpointAudit      = "/audit"
	endpointAuth       = "/auth"
)

var mux = http.NewServeMux()

func init() {
	mux.HandleFunc(endpointClearAudit, middleware.CheckMethod(secret.HandlerClearAudit))
	mux.HandleFunc(endpointAudit, middleware.CheckMethod(secret.HandlerAudit))
	mux.HandleFunc(endpointAuth, middleware.CheckMethod(public.HandlerAuth))
}

func router() *http.ServeMux {
	return mux
}
