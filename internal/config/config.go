package config

import (
	"github.com/1makarov/go-dater/internal/mongodb"
	"os"
)

type Config struct {
	Port string
	DB   mongodb.Config
}

func Init() *Config {
	var cfg Config

	setFromEnv(&cfg)

	return &cfg
}

func setFromEnv(cfg *Config) {
	cfg.Port = os.Getenv("GRPC_PORT")

	cfg.DB.Name = os.Getenv("DB_NAME")
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Host = os.Getenv("DB_HOST")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.Port = os.Getenv("DB_PORT")
}
