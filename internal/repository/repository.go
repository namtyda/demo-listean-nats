package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrNotFound = errors.New("not found")

type repository struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func New(pool *pgxpool.Pool, ctx context.Context) *repository {
	return &repository{pool: pool, ctx: ctx}
}
