package middleware

import (
	"net/http"

	"github.com/auth-api/cmd/auth/internal/logger"
)

const (
	endpointClearAudit = "/audit/clear"
	endpointAudit      = "/audit"
	endpointAuth       = "/auth"
)

func CheckMethod(handler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.App.Debug("work", "who", "MiddlewareCheckMethod", "path", r.URL.Path)
		switch r.URL.Path {
		case endpointAuth:
			if r.Method != http.MethodGet {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}

		case endpointAudit:
			if r.Method != http.MethodGet {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}

		case endpointClearAudit:
			if r.Method != http.MethodDelete {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}

		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return

		}

		handler(w, r)
	}
}
