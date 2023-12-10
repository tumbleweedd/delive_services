package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

func NewPostgresDB(ctx context.Context, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 4)

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, err
}

// TODO: Close DB connections
