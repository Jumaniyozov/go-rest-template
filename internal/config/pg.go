package config

import (
	"errors"
)

const (
	dsnEnvName = "PG_DSN"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func (cfg *Config) NewPGConfig() (PGConfig, error) {
	dsn := cfg.DbUrl
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
