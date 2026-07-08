package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPool creates a *pgxpool.Pool which we'll use to query to database
func NewPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	// create config based on connection string
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	// setup our connection related config
	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetimeJitter = 5 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConnLifetime = 30 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	// create a pgx pool with the config that we setup
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	// failed to reach database so let's just return error before we face any issue later on
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	// finally return the valid pool with no error
	return pool, nil
}
