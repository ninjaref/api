package api

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// NewDatabase creates a new connection to our PostgreSQL database.
func NewDatabase() (*gorm.DB, error) {
	connection, _ := pq.ParseURL(os.Getenv("DATABASE_URL"))
	connection += " sslmode=require" // NOTE: Required by Heroku.

	db, err := gorm.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	return db, nil
}
