package db

import (
	"ai-agent-manager/internal/platform/config"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(cfg.PGConn)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, err
	}
	if err := ping(ctx, pool); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}

func ping(ctx context.Context, pool *pgxpool.Pool) error {
	return pool.Ping(ctx)
}
