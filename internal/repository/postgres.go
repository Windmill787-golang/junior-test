package repository

import (
	"database/sql"
	"fmt"
	"github.com/Windmill787-golang/junior-test/internal/config"

	_ "github.com/lib/pq"
)

const (
	booksTable = "books"
	usersTable = "users"
)

func NewPostgres(config *config.Postgres) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", config.Username, config.Password, config.Host, config.Port, config.Database, config.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
