package logger

import (
	"log/slog"

	"github.com/auth-api/cmd/auth/internal/config"
	"github.com/auth-api/pkg/logger"
)

var App *slog.Logger

func init() {
	App = logger.New(config.Cng.Evn)
}
