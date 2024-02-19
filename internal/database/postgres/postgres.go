package postgres

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	db "github.com/Jumaniyozov/go-rest-template/internal/database/sqlc"
	"github.com/Jumaniyozov/go-rest-template/internal/repository"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/auth"
	"github.com/Jumaniyozov/go-rest-template/internal/repository/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresDB struct {
	q *db.Queries
}

func NewPostgresDB(cfg *config.Config) (repository.RepositoryI, error) {
	dbpool, err := pgxpool.New(context.Background(), cfg.DbUrl)
	if err != nil {
		return nil, err
	}

	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	q := db.New(dbpool)

	return &postgresDB{
		q: q,
	}, nil
}

func (p *postgresDB) UserRepository() user.UserI {
	return user.NewRepository(p.q)
}
func (p *postgresDB) AuthRepository() auth.AuthI {
	return auth.NewRepository(p.q)
}
