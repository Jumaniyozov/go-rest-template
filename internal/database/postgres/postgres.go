package postgres

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	"github.com/Jumaniyozov/go-rest-template/internal/database/entities"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/auth"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresDB struct {
	entity *entities.Entities
}

func New(cfg *config.Config) (repository.RepositoryI, error) {
	dbpool, err := pgxpool.New(context.Background(), cfg.DbUrl)
	if err != nil {
		return nil, err
	}

	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	q := db.New(dbpool)

	e := entities.New(q)

	return &postgresDB{
		entity: e,
	}, nil
}

func (p *postgresDB) UserRepository() user.UserI { return user.New(p.entity) }
func (p *postgresDB) AuthRepository() auth.AuthI { return auth.New(p.entity) }
