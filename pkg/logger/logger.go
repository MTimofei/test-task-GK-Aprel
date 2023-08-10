package logger

import (
	"log"
	"os"

	"log/slog"
)

const (
	evnLocal = "local"
	evnDev   = "dev"
	evnProd  = "prod"
)

func New(evn string) *slog.Logger {
	var logger *slog.Logger

	switch evn {
	case evnLocal:

		logger = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: slog.LevelDebug,
				},
			),
		)
	case evnDev:
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: slog.LevelDebug,
				},
			),
		)
	case evnProd:
		logger = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			),
		)
	default:
		log.Panic("logger ", "unknown environment ", evn)
	}

	return logger
}
