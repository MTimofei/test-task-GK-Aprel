package postgres

import (
	"errors"
	"fmt"

	"github.com/auth-api/cmd/auth/internal/config"
	"github.com/auth-api/cmd/auth/internal/logger"
	"github.com/auth-api/cmd/auth/internal/model"
	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

type Postgres struct {
	db *pg.DB
}

const (
	errInvalidToken = "invalid token"
)

func New() (*Postgres, error) {
	logger.App.Debug("work", "who", "New")
	db := pg.Connect(
		&pg.Options{
			Addr:     config.Cng.Postgres.Addr.Host + ":" + config.Cng.Postgres.Addr.Port,
			User:     config.Cng.Postgres.User,
			Password: config.Cng.Postgres.Password,
			Database: config.Cng.Postgres.Database,
		},
	)
	_, err := db.Exec(queryCheck)
	if err != nil {
		logger.App.Debug(err.Error())
		return nil, err
	}
	return &Postgres{db: db}, nil

}

func (p *Postgres) Close() {
	logger.App.Debug("work", "who", "Close")
	p.db.Close()
}

func (p *Postgres) Auth(login, password string) (id uuid.UUID, msg string, err error) {
	logger.App.Debug("work", "who", "Auth")
	var result struct {
		User_id    uuid.UUID
		Event_type string
	}
	_, err = p.db.Query(&result, fmt.Sprintf(queryAuth, password, login, password))
	if err != nil {
		logger.App.Debug(err.Error())
		return uuid.Nil, "", err
	}
	logger.App.Debug("result", "id", result.User_id, "msg", result.Event_type)
	return result.User_id, result.Event_type, nil
}

func (p *Postgres) Transaction(login string) error {
	logger.App.Debug("work", "who", "Transaction")
	_, err := p.db.Exec(fmt.Sprintf(queryTransaction, login))
	if err != nil {
		logger.App.Debug(err.Error())
		return err
	}
	return nil
}

func (p *Postgres) Commit() error {
	logger.App.Debug("work", "who", "Commit")
	_, err := p.db.Exec(queryCommit)
	if err != nil {
		logger.App.Debug(err.Error())
		return err
	}
	return nil
}

func (p *Postgres) Rollback() error {
	logger.App.Debug("work", "who", "Rollback")
	_, err := p.db.Exec(queryROLLBACK)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) SteToken(id uuid.UUID, token string) error {
	logger.App.Debug("work", "who", "SteToken")
	q := fmt.Sprintf(querySetToken, id, token)
	logger.App.Debug("work", "who", "SteToken", "q", q)
	_, err := p.db.Exec(q)
	if err != nil {
		logger.App.Debug(err.Error())
		return err
	}
	return nil
}

func (p *Postgres) Audit(token string) (audits []model.Audit, err error) {
	logger.App.Debug("work", "who", "Audit")
	q := fmt.Sprintf(queryAudit, token)
	logger.App.Debug("work", "who", "Audit", "q", q)
	_, err = p.db.Query(&audits, fmt.Sprintf(queryAudit, token))
	if err != nil {
		logger.App.Debug(err.Error())
		return nil, err
	}
	if audits == nil {
		logger.App.Debug(err.Error())
		return nil, errors.New(errInvalidToken)
	}
	return audits, nil
}

func (p *Postgres) ClearAudit(token string) (err error) {
	logger.App.Debug("work", "who", "ClearAudit")
	var ok []model.Audit
	q := fmt.Sprintf(queryClearAudit, token)
	logger.App.Debug("work", "who", "ClearAudit", "q", q)
	_, err = p.db.Query(&ok, q)
	if err != nil {
		logger.App.Debug(err.Error())
		return err
	}

	if ok == nil {
		logger.App.Debug(err.Error())
		return errors.New(errInvalidToken)
	}

	return nil
}
