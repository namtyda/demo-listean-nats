package repository

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/namtyda/demo-listean-nats/internal/models"
)

var ErrNotFound = errors.New("not found")

type repository struct {
	pool *pgxpool.Pool
}

type Repository interface {
	AddOrder(OrderUID, rawJson string) (err error)
	ReadAll() (rowSlice []models.Cache, err error)
}

func New(pool *pgxpool.Pool) *repository {
	return &repository{pool: pool}
}
