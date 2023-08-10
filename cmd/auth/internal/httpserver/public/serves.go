package public

import (
	"errors"

	"github.com/auth-api/cmd/auth/internal/logger"
	"github.com/auth-api/cmd/auth/internal/util"
	"github.com/auth-api/pkg/e"
	"github.com/google/uuid"
)

type resource interface {
	Auth(login, password string) (id uuid.UUID, msg string, err error)
	Transaction(login string) error
	Commit() error
	Rollback() error
	SteToken(id uuid.UUID, token string) error
}

const (
	errAuthFailed = "auth failed"
)

// ответы resource
const (
	invalidPassword = "invalid password"
	successfulLogin = "successful login"
	blockedLogin    = "block"
	unknownResponse = "unknown response"
)

type server struct{}

func (s *server) auth(login, password string, r resource) (token string, err error) {
	defer func() { err = e.IsError(errAuthFailed, err) }()
	logger.App.Debug("work", "who", "auth")
	err = r.Transaction(login)
	if err != nil {
		r.Rollback()
		return "", err
	}

	id, msg, err := r.Auth(login, password)
	if err != nil {
		r.Rollback()
		return "", err
	}

	err = handlerMsg(msg, r)
	if err != nil {
		return "", err
	}

	token = util.GenerateToken()

	err = r.SteToken(id, token)
	if err != nil {
		return "", err
	}

	return token, nil
}

// обработка сообщений от resource
func handlerMsg(msg string, r resource) (err error) {
	logger.App.Debug("work", "who", "handlerMsg", "msg", msg)

	switch msg {
	case successfulLogin:
		err = r.Commit()
		if err != nil {
			r.Rollback()
			return err
		}

		return nil

	case invalidPassword:
		err = r.Commit()
		if err != nil {
			r.Rollback()
			return err
		}

		return errors.New(invalidPassword)

	case blockedLogin:
		err = r.Commit()
		if err != nil {
			r.Rollback()
			return err
		}

		return errors.New(blockedLogin)

	default:
		r.Rollback()

		return errors.New(unknownResponse)
	}
}
