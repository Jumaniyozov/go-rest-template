package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	ProdEnv               string
	ServerHost            string
	ServerPort            int
	LogLevel              string
	PostgresHost          string
	PostgresPort          int
	PostgresDBName        string
	PostgresUser          string
	PostgresPassword      string
	DbUrl                 string
	DBPoolMaxConnections  int
	TokenSymmetricKey     string
	AccessTokenDuration   string
	RefreshTokenDuration  string
	AccessTokenSecretKey  string
	RefreshTokenSecretKey string
}

func SetupConfigs(logger *zerolog.Logger) (*Config, error) {
	if err := godotenv.Load("./.env"); err != nil {
		logger.Error().Msgf("Error loading .env file: %v", err)
		return nil, err
	}

	cfg := new(Config)

	cfg.ProdEnv = cast.ToString(getFromEnvOrDefault("PROD_ENV", "dev"))
	cfg.LogLevel = cast.ToString(getFromEnvOrDefault("LOG_LEVEL", "debug"))
	cfg.ServerPort = cast.ToInt(getFromEnvOrDefault("SERVER_PORT", 1324))
	cfg.ServerHost = cast.ToString(getFromEnvOrDefault("SERVER_HOST", "127.0.0.1"))

	cfg.PostgresPassword = cast.ToString(getFromEnvOrDefault("POSTGRES_PASSWORD", "password"))
	cfg.PostgresDBName = cast.ToString(getFromEnvOrDefault("POSTGRES_DB_NAME", "postgres"))
	cfg.PostgresPort = cast.ToInt(getFromEnvOrDefault("POSTGRES_PORT", 5434))
	cfg.PostgresUser = cast.ToString(getFromEnvOrDefault("POSTGRES_USER", "postgres"))
	cfg.PostgresHost = cast.ToString(getFromEnvOrDefault("POSTGRES_HOST", "localhost"))
	cfg.DbUrl = cast.ToString(getFromEnvOrDefault("DB_URL", "postgres://postgres:password@127.0.0.1:5434/postgres?sslmode=disable"))
	cfg.DBPoolMaxConnections = cast.ToInt(getFromEnvOrDefault("DB_POOL_MAX_CONNECTION", 30))

	cfg.TokenSymmetricKey = cast.ToString(getFromEnvOrDefault("TOKEN_SYMMETRIC_KEY", ""))
	cfg.AccessTokenDuration = cast.ToString(getFromEnvOrDefault("ACCESS_TOKEN_DURATION", ""))
	cfg.RefreshTokenDuration = cast.ToString(getFromEnvOrDefault("REFRESH_TOKEN_DURATION", ""))
	cfg.AccessTokenSecretKey = cast.ToString(getFromEnvOrDefault("ACCESS_TOKEN_SECRET_KEY", ""))
	cfg.RefreshTokenSecretKey = cast.ToString(getFromEnvOrDefault("REFRESH_TOKEN_SECRET_KEY", ""))

	if cfg.TokenSymmetricKey == "" || cfg.AccessTokenDuration == "" || cfg.RefreshTokenDuration == "" || cfg.AccessTokenSecretKey == "" || cfg.RefreshTokenSecretKey == "" {
		return nil, fmt.Errorf("token symmetric key, access token duration, refresh token duration, access token secret key, refresh token secret key must be set")
	}

	return cfg, nil
}

func getFromEnvOrDefault(name string, defaultValue any) any {
	value, exists := os.LookupEnv(name)
	if !exists {
		return defaultValue
	}

	return value
}
