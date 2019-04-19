package postgres_test

import (
	"database/sql"

	. "github.com/swensonhe/fanatick-backend/fanatick/postgres"
)

// NewDB opens and returns a new database connection.
func NewDB(dataSourceName string) *DB {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	return &DB{DB: db}
}
