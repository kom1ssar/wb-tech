package storage

import (
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"os"
	"sync"
	"tech-wb/internal/config"
)

type Postgres struct {
	db *pgx.Conn
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func (db *Postgres) connect() (err error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"

	db.db, err = pgx.Connect(pgx.ConnConfig{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return nil
}

func NewConnection(config config.DBConfig) *Postgres {

	connection, err := pgx.Connect(pgx.ConnConfig{
		Host:     config.GetHost(),
		Port:     config.GetPort(),
		User:     config.GetUser(),
		Password: config.GetPassword(),
		Database: config.GetName(),
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Println("DB connected")

	return &Postgres{db: connection}

}
