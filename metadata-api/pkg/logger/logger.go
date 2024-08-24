package logger

import (
	"log/slog"
	"os"
)

const (
	Local      string = "local"
	Production string = "production"
)

func Init(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	// Выводим в консоль текстом, начиная с уровня Debug
	case Local:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Выводим JSON, начиная с уровня Info
	case Production:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
