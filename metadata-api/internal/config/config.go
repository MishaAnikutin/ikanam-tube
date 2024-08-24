package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	AppConfig struct {
		Env          string `yaml:"env" env-required:"true"`
		ServerConfig `yaml:"http"`
		MongoConfig  `yaml:"mongo"`
	}

	MongoConfig struct {
		Username     string
		Password     string
		DatabaseName string `yaml:"database_name" env-required:"true"`
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
		log.Fatalf("Не удалось загрузить конфиг: %s", err)
	}

	return &cfg
}
