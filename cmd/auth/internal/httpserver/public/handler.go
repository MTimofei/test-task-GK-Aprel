package public

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/auth-api/cmd/auth/internal/util"
	"github.com/auth-api/pkg/e"
)

var s server

func HandlerAuth(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		util.Response(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}

	token, err := s.auth(
		query.Get("login"),
		query.Get("password"),
		nil,
	)

	switch {
	case err == nil:
		w.Header().Set("X-Token", token)
		util.Response(w, http.StatusOK, nil)
	case err.Error() == e.IsError(errAuthFailed, errors.New(blockedLogin)).Error():
		util.Response(w, http.StatusForbidden, []byte(err.Error()))
	case err.Error() == e.IsError(errAuthFailed, errors.New(unknownResponse)).Error():
		util.Response(w, http.StatusBadRequest, []byte(err.Error()))
	case err.Error() == e.IsError(errAuthFailed, errors.New(invalidPassword)).Error():
		util.Response(w, http.StatusBadRequest, []byte(err.Error()))
	}
}
