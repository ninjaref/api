package api

import (
	"database/sql"
	"os"

	"github.com/go-gorp/gorp"
	"github.com/lib/pq"
)

// NewDatabase creates a new connection to our PostgreSQL database.
func NewDatabase() (*gorp.DbMap, error) {
	connection, _ := pq.ParseURL(os.Getenv("DATABASE_URL"))
	connection += " sslmode=require" // NOTE: Required by Heroku.

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}, nil
}
