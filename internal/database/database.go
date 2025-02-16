package database

import (
	"context"
	"database/sql"

	sqlc "github.com/epuerta9/openchef/internal/database/sqlc"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type DB struct {
	*sql.DB
	*sqlc.Queries
}

func New(dbURL string) (*DB, error) {
	db, err := sql.Open("libsql", dbURL)
	if err != nil {
		return nil, err
	}

	q := sqlc.New(db)
	return &DB{
		DB:      db,
		Queries: q,
	}, nil
}

func (d *DB) Close() error {
	return d.DB.Close()
}

// Example query method
func (d *DB) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.DB.QueryContext(ctx, query, args...)
}
