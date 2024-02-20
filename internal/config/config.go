package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path"
	"path/filepath"
	"time"
)

type Config struct {
	ServerHost                  string        `env:"SERVER_HOST" env-required`
	ServerPort                  int           `env:"SERVER_PORT" env-required`
	ProdEnv                     string        `env:"PROD_ENV" env-required`
	LogLevel                    string        `env:"LOG_LEVEL" env-required`
	PostgresHost                string        `env:"POSTGRES_HOST" env-required`
	PostgresPort                int           `env:"POSTGRES_PORT" env-required`
	PostgresDbName              string        `env:"POSTGRES_DB_NAME" env-required`
	PostgresPassword            string        `env:"POSTGRES_PASSWORD" env-required`
	PostgresUser                string        `env:"POSTGRES_USER" env-required`
	DbPoolMaxConnection         int           `env:"DB_POOL_MAX_CONNECTION" env-required`
	DockerPostgresContainerName string        `env:"DOCKER_POSTGRES_CONTAINER_NAME" env-required`
	DockerVolumeName            string        `env:"DOCKER_VOLUME_NAME" env-required`
	DbUrl                       string        `env:"DB_URL" env-required`
	TokenSymmetricKey           string        `env:"TOKEN_SYMMETRIC_KEY" env-required`
	AccessTokenDuration         time.Duration `env:"ACCESS_TOKEN_DURATION" env-required`
	RefreshTokenDuration        time.Duration `env:"REFRESH_TOKEN_DURATION" env-required`
	AccessTokenSecretKey        string        `env:"ACCESS_TOKEN_SECRET_KEY" env-required`
	RefreshTokenSecretKey       string        `env:"REFRESH_TOKEN_SECRET_KEY" env-required`
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
