package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"path"
	"path/filepath"
	"time"
)

type Config struct {
	ServerPort                  int           `env:"SERVER_PORT" env-required:"true"`
	PostgresPort                int           `env:"POSTGRES_PORT" env-required:"true"`
	DbPoolMaxConnection         int           `env:"DB_POOL_MAX_CONNECTION" env-required:"true"`
	RefreshTokenDuration        time.Duration `env:"REFRESH_TOKEN_DURATION" env-required:"true"`
	AccessTokenDuration         time.Duration `env:"ACCESS_TOKEN_DURATION" env-required:"true"`
	ProdEnv                     string        `env:"PROD_ENV" env-required:"true"`
	ServerHost                  string        `env:"SERVER_HOST" env-required:"true"`
	LogLevel                    string        `env:"LOG_LEVEL" env-required:"true"`
	PostgresHost                string        `env:"POSTGRES_HOST" env-required:"true"`
	PostgresDbName              string        `env:"POSTGRES_DB_NAME" env-required:"true"`
	PostgresPassword            string        `env:"POSTGRES_PASSWORD" env-required:"true"`
	PostgresUser                string        `env:"POSTGRES_USER" env-required:"true"`
	DockerPostgresContainerName string        `env:"DOCKER_POSTGRES_CONTAINER_NAME" env-required:"true"`
	DockerVolumeName            string        `env:"DOCKER_VOLUME_NAME" env-required:"true"`
	DbUrl                       string        `env:"DB_URL" env-required:"true"`
	TokenSymmetricKey           string        `env:"TOKEN_SYMMETRIC_KEY" env-required:"true"`
	AccessTokenSecretKey        string        `env:"ACCESS_TOKEN_SECRET_KEY" env-required:"true"`
	RefreshTokenSecretKey       string        `env:"REFRESH_TOKEN_SECRET_KEY" env-required:"true"`
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
