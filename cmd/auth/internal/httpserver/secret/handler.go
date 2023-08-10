package secret

import (
	"fmt"
	"net/http"

	"github.com/auth-api/cmd/auth/internal/util"
)

var s serves

func HandlerAudit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	token := r.Header.Get("X-Token")
	if token == "" {
		util.Response(w, http.StatusBadRequest, []byte("token is empty"))
		return
	}

	jsonAudit, err := s.audit(
		token,
		nil,
	)

	if err != nil {
		switch err.Error() {
		case errInvalidToken:
			util.Response(w, http.StatusForbidden, []byte(errInvalidToken))
		case errTokenNotFound:
			util.Response(w, http.StatusNotFound, []byte(errTokenNotFound))
		default:
			util.Response(w, http.StatusInternalServerError, []byte("internal error"))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Contest-Length", fmt.Sprintf("%d", len(jsonAudit)))
	util.Response(w, http.StatusOK, jsonAudit)
}

func HandlerClearAudit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	token := r.Header.Get("X-Token")
	if token == "" {
		util.Response(w, http.StatusBadRequest, []byte("token is empty"))
		return
	}

	err := s.clearAudit(
		token,
		nil,
	)

	if err != nil {
		switch err.Error() {
		case errInvalidToken:
			util.Response(w, http.StatusForbidden, []byte(errInvalidToken))
		case errTokenNotFound:
			util.Response(w, http.StatusNotFound, []byte(errTokenNotFound))
		default:
			util.Response(w, http.StatusInternalServerError, []byte("internal error"))
		}
	}

	util.Response(w, http.StatusOK, nil)
}
