package secret

import (
	"encoding/json"

	"github.com/auth-api/cmd/auth/internal/logger"
	"github.com/auth-api/cmd/auth/internal/model"
	"github.com/auth-api/pkg/e"
)

type recurse interface {

	// метод возвращает список эветнов или ошибку встучии:
	// проблем с бд,
	// недействительного токена, "invalid token"
	Audit(token string) (auditlist []model.Audit, err error)

	// отчищает список эвентов
	// или ошибку встучии:
	// проблем с бд,
	// недействительного токена, "invalid token"
	ClearAudit(token string) (err error)
}

type serves struct{}

const (
	errAuthFailed = "auth failed"

	// ошибки recurse
	errInvalidToken  = "invalid token"
	errTokenNotFound = "token not found"
)

func (serves) audit(token string, r recurse) (jsonAudit []byte, err error) {
	defer func() { err = e.IsError(errAuthFailed, err) }()
	logger.App.Debug("work", "who", "audit")
	list, err := r.Audit(token)
	if err != nil {
		return nil, err
	}

	jsonAudit, err = json.Marshal(list)
	if err != nil {
		return nil, err
	}

	return jsonAudit, nil
}

func (serves) clearAudit(token string, r recurse) (err error) {
	defer func() { err = e.IsError(errAuthFailed, err) }()
	logger.App.Debug("work", "who", "clearAudit")
	err = r.ClearAudit(token)
	if err != nil {
		return err
	}

	return nil
}
