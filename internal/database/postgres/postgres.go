package postgres

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/database/entities"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Config) (*repository.Repository, error) {
	dbpool, err := pgxpool.New(context.Background(), cfg.DbUrl)
	if err != nil {
		return nil, err
	}

	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	q := db.New(dbpool)

	e := entities.New(q)

	return repository.New(e), nil
}
