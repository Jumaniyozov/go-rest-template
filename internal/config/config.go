package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path"
	"path/filepath"
	"time"
)

type Config struct {
	ServerHost                  string        `env:"SERVER_HOST"`
	ServerPort                  int           `env:"SERVER_PORT"`
	ProdEnv                     string        `env:"PROD_ENV"`
	LogLevel                    string        `env:"LOG_LEVEL"`
	PostgresHost                string        `env:"POSTGRES_HOST"`
	PostgresPort                int           `env:"POSTGRES_PORT"`
	PostgresDbName              string        `env:"POSTGRES_DB_NAME"`
	PostgresPassword            string        `env:"POSTGRES_PASSWORD"`
	PostgresUser                string        `env:"POSTGRES_USER"`
	DbPoolMaxConnection         int           `env:"DB_POOL_MAX_CONNECTION"`
	DockerPostgresContainerName string        `env:"DOCKER_POSTGRES_CONTAINER_NAME"`
	DockerVolumeName            string        `env:"DOCKER_VOLUME_NAME"`
	DbUrl                       string        `env:"DB_URL"`
	TokenSymmetricKey           string        `env:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration         time.Duration `env:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration        time.Duration `env:"REFRESH_TOKEN_DURATION"`
	AccessTokenSecretKey        string        `env:"ACCESS_TOKEN_SECRET_KEY"`
	RefreshTokenSecretKey       string        `env:"REFRESH_TOKEN_SECRET_KEY"`
}

func New() (*Config, error) {
	var cfg Config

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	err = cleanenv.ReadConfig(path.Join(exPath, ".env"), &cfg)
	if err != nil {
		return nil, err

	}

	return &cfg, nil
}
