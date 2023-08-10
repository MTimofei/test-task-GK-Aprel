package main

import (
	"os"

	"github.com/auth-api/cmd/auth/internal/config"
	"github.com/auth-api/cmd/auth/internal/dbms"
	"github.com/auth-api/cmd/auth/internal/httpserver"
	"github.com/auth-api/cmd/auth/internal/logger"
)

func main() {
	defer dbms.DBMS.Close()

	logger.App.Debug("config app",
		"env", config.Cng.Evn,
		"server host", config.Cng.Server.Host,
		"server port", config.Cng.Server.Port,
		"postgres host", config.Cng.Postgres.Addr.Host,
		"postgres port", config.Cng.Postgres.Addr.Port,
		"postgres database", config.Cng.Postgres.Database,
		"postgres user", config.Cng.Postgres.User,
		"postgres password", config.Cng.Postgres.Password)

	s := httpserver.New(config.Cng.Server.Host, config.Cng.Server.Port)

	go func() {
		err := s.Run()
		if err != nil {
			logger.App.Error(err.Error())
			os.Exit(1)
		}
	}()

	err := s.Shutdown()
	if err != nil {
		logger.App.Error(err.Error())
		os.Exit(1)
	}
}
