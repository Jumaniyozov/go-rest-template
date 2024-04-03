package postgres

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/database/entities"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBClient struct {
	Entity *entities.Entities
}

type EntityClient struct {
	entity *entities.Entities
}

func New(cfg *config.Config) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), cfg.DbUrl)
	if err != nil {
		return nil, err
	}

	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return dbpool, nil
}

func NewDBEntities(dbpool *pgxpool.Pool) *DBClient {
	return &DBClient{
		Entity: entities.New(db.New(dbpool)),
	}
}
