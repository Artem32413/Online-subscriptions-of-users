package logger

import (
	"log/slog"
	"os"
)

func InitSwagLog() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	})

	return slog.New(handler)
}