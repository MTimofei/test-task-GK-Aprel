package httpserver

import (
	"net/http"

	"github.com/auth-api/cmd/auth/internal/httpserver/public"
	"github.com/auth-api/cmd/auth/internal/httpserver/secret"
)

const (
	endpointClearAudit = "/audit/clear"
	endpointAudit      = "/audit"
	endpointAuth       = "/auth"
)

func router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(endpointClearAudit, secret.HandlerClearAudit)
	mux.HandleFunc(endpointAudit, secret.HandlerAudit)
	mux.HandleFunc(endpointAuth, public.HandlerAuth)

	return mux
}
