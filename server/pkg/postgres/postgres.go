package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultMaxPoolSize  = 2
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

func NewPGXPool(connString string) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPGXPool - pgxpool.ParseConfig: %w", err)
	}

	conf.MaxConns = int32(defaultMaxPoolSize)

	var pool *pgxpool.Pool
	for i := 1; i <= defaultConnAttempts; i++ {
		pool, err = pgxpool.NewWithConfig(context.Background(), conf)
		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", defaultConnAttempts-i)
		time.Sleep(defaultConnTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPGXPool - pgxpool.NewWithConfig: %w", err)
	}

	return pool, nil
}
