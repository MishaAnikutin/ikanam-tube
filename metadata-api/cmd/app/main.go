package main

import (
	"log/slog"

	"github.com/MishaAnikutin/metadata-api/internal/config"
	"github.com/MishaAnikutin/metadata-api/pkg/logger"
)

const configPath = "config/local.yaml"

func main() {
	cfg := config.Load(configPath)

	logger := logger.Init(cfg.Env)

	logger.Info("Запускаем сервис metadata-api", slog.String("env", cfg.Env))
	logger.Debug("Режим Debug включен")

}
