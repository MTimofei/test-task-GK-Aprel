package dbms

import (
	"github.com/auth-api/cmd/auth/internal/dbms/postgres"
	"github.com/auth-api/cmd/auth/internal/logger"
)

var DBMS *postgres.Postgres

func init() {
	var err error
	DBMS, err = postgres.New()
	if err != nil {
		logger.App.Error(err.Error())
	}
}
