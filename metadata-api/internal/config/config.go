package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	AppConfig struct {
		Env            string `yaml:"env" env-required:"true"`
		ServerConfig   `yaml:"http"`
		PostgresConfig `yaml:"postgres"`
	}

	PostgresConfig struct {
		DatabaseHost string `yaml:"host" env-required:"true"`
		DatabasePort string `yaml:"port" env-default:"5432"`
		Username     string
		Password     string
		DatabaseName string `yaml:"database" env-required:"true"`
	}

	ServerConfig struct {
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
		IdleTimeout  time.Duration `yaml:"idle_timeout"`
	}
)

func Load(configPath string) *AppConfig {

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Путь не найдет: %s", configPath)
	}

	var cfg AppConfig

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(fmt.Sprintf("Не удалось загрузить конфиг: %s", err))
	}

	cfg.Username = os.Getenv("POSTGRES_USER")
	cfg.Password = os.Getenv("POSTGRES_PASSWORD")

	if cfg.Username == "" || cfg.Password == "" {
		panic("Не удалось загрузить пользователя или пароль от базы данных")
	}

	return &cfg
}

func PostgresURL(cfg *AppConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
	)

}
