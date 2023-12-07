package postgresql

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"tech-wb/internal/config"
	"tech-wb/pkg/utils"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, config config.DBConfig) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error

	delay := 3 * time.Second

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, delay)

		defer cancel()

		pool, err = pgxpool.Connect(ctx, config.GetDSN())
		if err != nil {
			return err
		}
		return nil

	}, config.GetMaxAttempts(), delay)

	if err != nil {
		return nil, err
	}

	return pool, nil
}
