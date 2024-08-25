package main

import (
	"context"
	"log/slog"

	"github.com/MishaAnikutin/metadata-api/internal/app"
	"github.com/MishaAnikutin/metadata-api/internal/config"
	"github.com/MishaAnikutin/metadata-api/internal/di"
	"github.com/MishaAnikutin/metadata-api/internal/migrations"
	"github.com/MishaAnikutin/metadata-api/internal/server"
	"github.com/MishaAnikutin/metadata-api/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

const configPath = "config/local.yaml"

func main() {
	cfg := config.Load(configPath)

	logger := logger.Init(cfg.Env)

	logger.Info("Запускаем сервис metadata-api", slog.String("env", cfg.Env))
	logger.Debug("Режим Debug включен")

	logger.Info("Подключаемся к базе данных", slog.String("env", cfg.Env))

	url, err := pgxpool.ParseConfig(config.PostgresURL(cfg))

	if err != nil {
		logger.Error("Ошибка при парсинге строки подключения: %s\n", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), url)
	if err != nil {
		logger.Error("Ошибка при подключении к БД: %s\n", err)
	}
	defer pool.Close()

	logger.Info("Проводим миграции")
	migrations.UpgradeHead(context.Background(), pool)

	logger.Info("Внедряем зависимости")
	router := di.Inject(context.Background(), pool)

	logger.Info("Поднимаем сервер")
	Server := server.New(cfg.Port, router)

	App := app.New(Server)

	logger.Info("Запускаем приложение...")
	App.Run()
}
